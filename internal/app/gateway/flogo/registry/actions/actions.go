// Code generated by go-bindata.
// sources:
// vendor/github.com/TIBCOSoftware/flogo-contrib/action/flow/action.json
// DO NOT EDIT!

package actions

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
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
	name    string
	size    int64
	mode    os.FileMode
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

var _vendorGithubComTibcosoftwareFlogoContribActionFlowActionJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8e\xcd\xae\x82\x30\x10\x85\xf7\x24\xbc\x43\xd3\xf5\xbd\x14\xb7\xdd\xa9\x89\x89\x2b\x17\xf8\x02\xa5\x0e\xd0\x84\xfe\x64\x3a\x48\x88\xf1\xdd\x0d\x15\x22\x0b\x17\x6e\xe7\xcc\x77\xce\xf7\xc8\x33\xc6\xb8\x53\x16\xb8\x64\xbc\xe9\xfd\xc8\xff\xd2\x89\xa6\xb0\x9e\x5a\x2f\x95\x26\xe3\xdd\x12\x21\x34\x73\xd2\x1a\xea\x86\xba\xd0\xde\x8a\xeb\xf9\x70\xbc\x54\xbe\xa1\x51\x21\x88\x84\xfc\x6b\xef\x08\x4d\x2d\xde\xa8\xd8\x54\xdf\x01\xe3\xdc\x26\x19\x2f\x8b\xb2\xd8\xad\x8b\x86\xfa\x34\x59\x19\x1b\x7a\x60\xa7\x0f\x71\x83\xa8\xd1\x04\x5a\xa8\xcd\x03\xdb\x6f\xcd\x3a\x6f\x21\xa8\x36\xb5\x74\x44\x21\x4a\x21\x7e\xd4\x24\x04\x10\x56\x45\x02\xfc\xa2\xac\xe2\xe4\x34\x97\x8c\x70\x80\x3c\x7b\xe6\xd9\x2b\x00\x00\xff\xff\xb5\x95\x8c\x25\x38\x01\x00\x00")

func vendorGithubComTibcosoftwareFlogoContribActionFlowActionJsonBytes() ([]byte, error) {
	return bindataRead(
		_vendorGithubComTibcosoftwareFlogoContribActionFlowActionJson,
		"vendor/github.com/TIBCOSoftware/flogo-contrib/action/flow/action.json",
	)
}

func vendorGithubComTibcosoftwareFlogoContribActionFlowActionJson() (*asset, error) {
	bytes, err := vendorGithubComTibcosoftwareFlogoContribActionFlowActionJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vendor/github.com/TIBCOSoftware/flogo-contrib/action/flow/action.json", size: 312, mode: os.FileMode(438), modTime: time.Unix(1531402965, 0)}
	a := &asset{bytes: bytes, info: info}
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
	if err != nil {
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
	"vendor/github.com/TIBCOSoftware/flogo-contrib/action/flow/action.json": vendorGithubComTibcosoftwareFlogoContribActionFlowActionJson,
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
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"vendor": &bintree{nil, map[string]*bintree{
		"github.com": &bintree{nil, map[string]*bintree{
			"TIBCOSoftware": &bintree{nil, map[string]*bintree{
				"flogo-contrib": &bintree{nil, map[string]*bintree{
					"action": &bintree{nil, map[string]*bintree{
						"flow": &bintree{nil, map[string]*bintree{
							"action.json": &bintree{vendorGithubComTibcosoftwareFlogoContribActionFlowActionJson, map[string]*bintree{}},
						}},
					}},
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
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
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
		err = RestoreAssets(dir, filepath.Join(name, child))
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

