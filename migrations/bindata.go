// Code generated by go-bindata.
// sources:
// migrations_data/1_create_port_range_down.sql
// migrations_data/1_create_port_range_up.sql
// DO NOT EDIT!

package migration

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

var _migrations_data1_create_port_range_downSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\xc8\x2f\x2a\x89\x2f\x4a\xcc\x4b\x4f\x2d\x56\xb0\xe6\x02\x04\x00\x00\xff\xff\x56\x9a\x44\xc3\x23\x00\x00\x00")

func migrations_data1_create_port_range_downSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations_data1_create_port_range_downSql,
		"migrations_data/1_create_port_range_down.sql",
	)
}

func migrations_data1_create_port_range_downSql() (*asset, error) {
	bytes, err := migrations_data1_create_port_range_downSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations_data/1_create_port_range_down.sql", size: 35, mode: os.FileMode(420), modTime: time.Unix(1444042540, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations_data1_create_port_range_upSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\xcc\xb1\x8e\x82\x40\x14\x46\xe1\x9e\xa7\xf8\x4b\x48\xb6\xda\x2c\xd5\x66\x8b\x81\xbd\xea\x44\xb8\x98\xe1\x62\xa4\x22\x46\xd0\x4c\xe1\x0c\x19\x78\xff\x88\x14\x16\xc6\xfa\x7c\x39\xb9\x21\x25\x04\x51\x59\x41\xd0\x1b\x70\x25\xa0\x93\xae\xa5\xc6\xe8\xc3\xdc\x85\xb3\xbb\x0d\x13\xe2\x08\xb0\x3d\x34\xcb\x2a\xb8\x29\x0a\xa8\x46\xaa\x4e\xf3\x72\x28\x89\xe5\x6b\x11\xfd\x30\x5d\x82\x1d\x67\xeb\x1d\x8e\xca\xe4\x3b\x65\xe2\xef\x34\x4d\x56\xff\x04\xeb\xf2\x1a\xfc\xfd\x95\x7f\xde\xeb\xec\x3f\xb6\x83\xd1\xa5\x32\x2d\xf6\xd4\x22\xb6\x7d\x92\x44\xc4\x5b\xcd\x84\x3f\x68\xe7\xfc\x7f\xf6\x1b\x3d\x02\x00\x00\xff\xff\xb7\x33\x78\xb1\xcb\x00\x00\x00")

func migrations_data1_create_port_range_upSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations_data1_create_port_range_upSql,
		"migrations_data/1_create_port_range_up.sql",
	)
}

func migrations_data1_create_port_range_upSql() (*asset, error) {
	bytes, err := migrations_data1_create_port_range_upSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations_data/1_create_port_range_up.sql", size: 203, mode: os.FileMode(420), modTime: time.Unix(1444043292, 0)}
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
	"migrations_data/1_create_port_range_down.sql": migrations_data1_create_port_range_downSql,
	"migrations_data/1_create_port_range_up.sql": migrations_data1_create_port_range_upSql,
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
	"migrations_data": &bintree{nil, map[string]*bintree{
		"1_create_port_range_down.sql": &bintree{migrations_data1_create_port_range_downSql, map[string]*bintree{}},
		"1_create_port_range_up.sql": &bintree{migrations_data1_create_port_range_upSql, map[string]*bintree{}},
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

