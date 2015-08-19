package appfile

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform/config/module"
	"github.com/hashicorp/terraform/dag"
)

const (
	// CompileVersion is the current version that we're on for
	// compilation formats. This can be used in the future to change
	// the directory structure and on-disk format of compiled appfiles.
	CompileVersion = 1

	CompileFilename        = "Appfile.compiled"
	CompileDepsFolder      = "deps"
	CompileVersionFilename = "version"
)

// Compiled represents a "Compiled" Appfile. A compiled Appfile is one
// that has loaded all of its dependency Appfiles, completed its imports,
// verified it is valid, etc.
//
// Appfile compilation is a process that requires network activity and
// has to occur once. The idea is that after compilation, a fully compiled
// Appfile can then be loaded in the future without network connectivity.
// Additionally, since we can assume it is valid, we can load it very quickly.
type Compiled struct {
	// File is the raw Appfile
	File *File

	// Graph is the DAG that has all the dependencies. This is already
	// verified to have no cycles. Each vertex is a *CompiledGraphVertex.
	Graph *dag.AcyclicGraph
}

func (c *Compiled) Validate() error {
	var result error

	// First validate that there are no cycles in the dependency graph
	if cycles := c.Graph.Cycles(); len(cycles) > 0 {
		for _, cycle := range cycles {
			vertices := make([]string, len(cycle))
			for i, v := range cycle {
				vertices[i] = dag.VertexName(v)
			}

			result = multierror.Append(result, fmt.Errorf(
				"Dependency cycle: %s", strings.Join(vertices, ", ")))
		}
	}

	return result
}

func (c *Compiled) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("Compiled Appfile: %s\n\n", c.File.Path))
	buf.WriteString("Dep Graph:\n")
	buf.WriteString(c.Graph.String())
	buf.WriteString("\n")
	return buf.String()
}

// CompiledGraphVertex is the type of the vertex within the Graph of Compiled.
type CompiledGraphVertex struct {
	// File is the raw Appfile that this represents
	File *File

	// Dir is the directory of the data root for this dependency. This
	// is only non-empty for dependencies (the root vertex does not have
	// this value).
	Dir string

	// Don't use this outside of this package.
	NameValue string
}

func (v *CompiledGraphVertex) Name() string {
	return v.NameValue
}

// CompileOpts are the options for compilation.
type CompileOpts struct {
	// Dir is the directory where all the compiled data will be stored.
	// For use of Otto with a compiled Appfile, this directory must not
	// be deleted.
	Dir string

	// Callback is an optional way to receive notifications of events
	// during the compilation process. The CompileEvent argument should be
	// type switched to determine what it is.
	Callback func(CompileEvent)
}

// CompileEvent is a potential event that a Callback can receive during
// Compilation.
type CompileEvent interface{}

// CompileEventDep is the event that is called when a dependency is
// being loaded.
type CompileEventDep struct {
	Source string
}

// LoadCompiled loads and verifies a compiled Appfile (*Compiled) from
// disk.
func LoadCompiled(dir string) (*Compiled, error) {
	f, err := os.Open(filepath.Join(dir, CompileFilename))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var c Compiled
	dec := json.NewDecoder(f)
	if err := dec.Decode(&c); err != nil {
		return nil, err
	}

	return &c, nil
}

// Compile compiles an Appfile.
//
// This may require network connectivity if there are imports or
// non-local dependencies. The repositories that dependencies point to
// will be fully loaded into the given directory, and the compiled Appfile
// will be saved there.
//
// LoadCompiled can be used to load a pre-compiled Appfile.
//
// If you have no interest in reloading a compiled Appfile, you can
// recursively delete the compilation directory after this is completed.
// Note that certain functions of Otto such as development environments
// will depend on those directories existing, however.
func Compile(f *File, opts *CompileOpts) (*Compiled, error) {
	// First clear the directory. In the future, we can keep it around
	// and do incremental compilations.
	if err := os.RemoveAll(opts.Dir); err != nil {
		return nil, err
	}
	if err := os.MkdirAll(opts.Dir, 0755); err != nil {
		return nil, err
	}

	// Write the version of the compilation that we'll be completing.
	if err := compileVersion(opts.Dir); err != nil {
		return nil, fmt.Errorf("Error writing compiled Appfile version: %s", err)
	}

	// Start building our compiled Appfile
	compiled := &Compiled{File: f, Graph: new(dag.AcyclicGraph)}

	// Add our root vertex for this Appfile
	vertex := &CompiledGraphVertex{File: f, NameValue: f.Application.Name}
	compiled.Graph.Add(vertex)

	// Build the storage we'll use for storing downloaded dependencies,
	// then use that to trigger the recursive call to download all our
	// dependencies.
	storage := &module.FolderStorage{
		StorageDir: filepath.Join(opts.Dir, CompileDepsFolder)}
	if err := compileDependencies(storage, compiled.Graph, opts, vertex); err != nil {
		return nil, err
	}

	// Validate the compiled file tree.
	if err := compiled.Validate(); err != nil {
		return nil, err
	}

	// Write the compiled Appfile data
	if err := compileWrite(opts.Dir, compiled); err != nil {
		return nil, err
	}

	return compiled, nil
}

func compileDependencies(
	storage module.Storage,
	graph *dag.AcyclicGraph,
	opts *CompileOpts,
	root *CompiledGraphVertex) error {
	// Make a map to keep track of the dep source to vertex mapping
	vertexMap := make(map[string]*CompiledGraphVertex)

	// Store ourselves in the map
	key, err := module.Detect(".", filepath.Dir(root.File.Path))
	if err != nil {
		return err
	}
	vertexMap[key] = root

	// Make a queue for the other vertices we need to still get
	// dependencies for. We arbitrarily make the cap for this slice
	// 30, since that is a ton of dependencies and we don't expect the
	// average case to have more than this.
	queue := make([]*CompiledGraphVertex, 1, 30)
	queue[0] = root

	// While we still have dependencies to get, continue loading them.
	// TODO: parallelize
	for len(queue) > 0 {
		var current *CompiledGraphVertex
		current, queue = queue[len(queue)-1], queue[:len(queue)-1]

		log.Printf("[DEBUG] compiling dependencies for: %s", current.Name())
		for _, dep := range current.File.Application.Dependencies {
			key, err := module.Detect(dep.Source, filepath.Dir(current.File.Path))
			if err != nil {
				return fmt.Errorf(
					"Error loading source: %s", err)
			}

			vertex := vertexMap[key]
			if vertex == nil {
				log.Printf("[DEBUG] loading dependency: %s", key)

				// Call the callback if we have one
				if opts.Callback != nil {
					opts.Callback(&CompileEventDep{
						Source: key,
					})
				}

				// Download the dependency
				if err := storage.Get(key, key, true); err != nil {
					return err
				}
				dir, _, err := storage.Dir(key)
				if err != nil {
					return err
				}

				// Parse the Appfile
				f, err := ParseFile(filepath.Join(dir, "Appfile"))
				if err != nil {
					return fmt.Errorf(
						"Error parsing Appfile in %s: %s", key, err)
				}

				// Build the vertex for this
				vertex = &CompiledGraphVertex{
					File:      f,
					Dir:       dir,
					NameValue: f.Application.Name,
				}

				// Add the vertex since it is new, store the mapping, and
				// queue it to be loaded later.
				graph.Add(vertex)
				vertexMap[key] = vertex
				queue = append(queue, vertex)
			}

			// Connect the dependencies
			graph.Connect(dag.BasicEdge(current, vertex))
		}
	}

	return nil
}

func compileVersion(dir string) error {
	f, err := os.Create(filepath.Join(dir, CompileVersionFilename))
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprintf(f, "%d", CompileVersion)
	return err
}

func compileWrite(dir string, compiled *Compiled) error {
	// Pretty-print the JSON data so that it can be more easily inspected
	data, err := json.MarshalIndent(compiled, "", "    ")
	if err != nil {
		return err
	}

	// Write it out
	f, err := os.Create(filepath.Join(dir, CompileFilename))
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, bytes.NewReader(data))
	return err
}