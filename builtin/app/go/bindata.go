// Code generated by go-bindata.
// sources:
// data/aws-vpc-public-private/build/template.json
// data/common/dev/Vagrantfile.tpl
// DO NOT EDIT!

package goapp

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"reflect"
	"strings"
	"unsafe"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindataRead(data, name string) ([]byte, error) {
	var empty [0]byte
	sx := (*reflect.StringHeader)(unsafe.Pointer(&data))
	b := empty[:]
	bx := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bx.Data = sx.Data
	bx.Len = len(data)
	bx.Cap = bx.Len

	gz, err := gzip.NewReader(bytes.NewBuffer(b))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _dataAwsVpcPublicPrivateBuildTemplateJson = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xe6\x52\x00\x02\xa5\xdc\xcc\xbc\xf8\x82\xc4\xe4\xec\xd4\xa2\xf8\xb2\xd4\xa2\xe2\xcc\xfc\x3c\x25\x2b\x05\x25\x03\x3d\x0b\x3d\x03\x25\xae\x5a\x2e\x40\x00\x00\x00\xff\xff\x1b\x82\xd0\x08\x26\x00\x00\x00"

func dataAwsVpcPublicPrivateBuildTemplateJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataAwsVpcPublicPrivateBuildTemplateJson,
		"data/aws-vpc-public-private/build/template.json",
	)
}

func dataAwsVpcPublicPrivateBuildTemplateJson() (*asset, error) {
	bytes, err := dataAwsVpcPublicPrivateBuildTemplateJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws-vpc-public-private/build/template.json", size: 38, mode: os.FileMode(420), modTime: time.Unix(1435862031, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataCommonDevVagrantfileTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x90\x41\x8b\x14\x31\x10\x85\xef\xfd\x2b\x9e\x19\x04\x85\xb1\x1b\x44\x3c\x08\x2e\x78\x12\x0f\x22\xe2\xe2\x45\xc4\xcd\x24\xd5\x9d\x62\xd3\xa9\x90\xa4\x5b\x87\xdd\xfd\xef\x26\x33\x3d\xe2\x2e\xf4\xa1\x79\xa9\x7a\x5f\xbd\xb7\xc3\x47\x0a\x94\x74\x21\x8b\xc3\x11\x5f\x4a\x91\x3d\xac\x20\x48\x01\x59\x2e\xcf\xba\x5d\xb7\xc3\xb5\xe3\x8c\xfa\x15\x47\xf8\xae\xa7\xa4\x43\x19\xd9\x13\xa6\xa7\xbb\x18\x25\x9d\xa6\x2c\xad\xe4\x25\xce\x14\x0a\x64\xac\x16\xa5\x59\xe8\x18\x3d\x1b\x5d\x58\xc2\x90\x29\xad\x6c\xa8\xc7\xa7\x82\xec\x64\xf1\xf6\x04\x3d\x10\x9c\x0e\xf6\x55\x83\x93\xed\x71\x2d\x98\xc5\xf2\x78\x6c\xb6\xd5\xe7\x3f\xfc\x1e\x4b\xa6\x13\xed\x43\x8c\x4d\xe8\xbb\x6e\x7b\xee\x8d\x84\x91\xa7\x25\xd1\x0b\xf5\x5a\xbd\x6c\x89\xee\xcf\xd2\x7d\x07\x9c\xff\xfa\x75\xee\x0f\xf2\x07\xef\xa1\x9c\xce\x8e\x8d\xa4\x38\xc4\x44\x86\x33\xbd\x7d\xa3\xba\x3a\xb8\xc3\x37\x2a\x4b\x84\x46\x3e\x06\x53\x63\x8e\xe2\x2d\x25\x8c\x49\x66\xc8\x92\xf0\x5b\xd2\x2d\x87\x09\x96\xeb\x5e\x91\x54\xaf\x14\x0c\xeb\xf9\x88\x47\xa4\xb3\xc1\xaf\xcd\x40\xdd\xdd\x21\xea\xe2\xfa\x8b\xc1\xc3\x83\xda\x43\x5d\x36\x37\xf8\x67\x7d\x4b\xe0\x5a\x8f\xd4\x94\xba\xe0\x66\x7b\x46\xce\xee\x06\x93\x50\xde\xc0\xfe\xc4\x6d\x4d\xd4\x10\x4d\x68\xfa\x23\x7c\x4c\xb2\x72\xae\xbd\x43\x65\x47\xde\x57\x1a\x07\xcf\x81\xde\xd5\x31\xe0\xf9\xd7\x1f\x64\x9c\x40\x19\xfb\xef\x7c\x85\xab\x2b\x0c\x4e\x66\xba\x28\x43\x7f\xa8\x45\x25\xf3\xb3\xa3\x60\xbb\xbf\x01\x00\x00\xff\xff\xdf\x53\x5d\xfd\x3b\x02\x00\x00"

func dataCommonDevVagrantfileTplBytes() ([]byte, error) {
	return bindataRead(
		_dataCommonDevVagrantfileTpl,
		"data/common/dev/Vagrantfile.tpl",
	)
}

func dataCommonDevVagrantfileTpl() (*asset, error) {
	bytes, err := dataCommonDevVagrantfileTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/common/dev/Vagrantfile.tpl", size: 571, mode: os.FileMode(420), modTime: time.Unix(1439845158, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if (err != nil) {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"data/aws-vpc-public-private/build/template.json": dataAwsVpcPublicPrivateBuildTemplateJson,
	"data/common/dev/Vagrantfile.tpl": dataCommonDevVagrantfileTpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"data": &bintree{nil, map[string]*bintree{
		"aws-vpc-public-private": &bintree{nil, map[string]*bintree{
			"build": &bintree{nil, map[string]*bintree{
				"template.json": &bintree{dataAwsVpcPublicPrivateBuildTemplateJson, map[string]*bintree{
				}},
			}},
		}},
		"common": &bintree{nil, map[string]*bintree{
			"dev": &bintree{nil, map[string]*bintree{
				"Vagrantfile.tpl": &bintree{dataCommonDevVagrantfileTpl, map[string]*bintree{
				}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
        if err != nil {
                return err
        }
        err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
        if err != nil {
                return err
        }
        err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
        if err != nil {
                return err
        }
        return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        // File
        if err != nil {
                return RestoreAsset(dir, name)
        }
        // Dir
        for _, child := range children {
                err = RestoreAssets(dir, path.Join(name, child))
                if err != nil {
                        return err
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

