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

var _vendorGithubComTibcosoftwareFlogoContribActionFlowActionJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8e\x3d\x8e\xc2\x30\x10\x85\xfb\x9c\xc2\x72\xbd\x1b\x67\x5b\x77\x0b\x12\x12\x15\x45\xb8\x80\x63\x26\x89\xa5\xd8\x63\x8d\x27\x44\x11\xe2\xee\x28\x3f\x40\x24\x1a\xda\x79\xf3\xbe\xf7\xdd\x32\x21\x64\x30\x1e\xa4\x16\xb2\xee\x70\x90\x3f\xd3\x85\xc7\xf8\xbc\x34\xa8\x8d\x65\x87\x61\x49\x08\xea\x29\x68\x1c\xb7\x7d\x95\x5b\xf4\xea\x7c\xdc\xed\x4f\x25\xd6\x3c\x18\x02\x35\x37\x7e\x2d\x06\x26\x57\xa9\xa5\xa9\xde\xe0\x2b\x50\x9a\x58\x5a\xc8\x22\x2f\xf2\xbf\x75\xce\x71\x37\xef\x95\xce\xc7\x0e\xc4\xe1\xf5\x7f\x81\x64\xc9\x45\x5e\x3b\x9b\x5c\xfc\x6f\xac\x5a\xf4\x10\x4d\x33\x33\x5a\xe6\x98\xb4\x52\x5f\x2a\x32\x01\x28\x6f\x12\x03\x7d\xea\x9a\x34\x06\x2b\xb5\x60\xea\x21\xbb\x67\x8f\x00\x00\x00\xff\xff\xdd\xfd\x73\xa2\x2e\x01\x00\x00")

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

	info := bindataFileInfo{name: "vendor/github.com/TIBCOSoftware/flogo-contrib/action/flow/action.json", size: 302, mode: os.FileMode(436), modTime: time.Unix(1535446427, 0)}
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
