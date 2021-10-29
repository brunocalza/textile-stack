package migrations

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var __001_init_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\x48\xcd\x2f\xc8\x49\xb5\xe6\x02\x04\x00\x00\xff\xff\xe9\xc7\x3b\x06\x13\x00\x00\x00")

func _001_init_down_sql() ([]byte, error) {
	return bindata_read(
		__001_init_down_sql,
		"001_init.down.sql",
	)
}

var __001_init_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xcb\xc1\x0a\x82\x40\x14\x85\xe1\xbd\x4f\x71\x96\x05\xbd\x41\xab\xd1\xae\x30\x34\x4e\x32\x1e\x41\x57\x32\xe5\x5d\x08\x5a\x12\xbe\x3f\xc1\x08\x2e\x5a\xde\xff\x7e\xa7\x08\x62\x28\xa0\xc9\x9d\xc0\x96\xf0\x0f\x42\x3a\xdb\xb0\xc1\xaa\x9f\x75\x56\x9c\x32\x00\x98\x46\x58\x4f\xd4\xc1\x56\x26\xf4\xb8\x4b\x7f\x49\xfd\x1d\x17\x05\xa5\x63\x5a\xfa\xd6\xb9\xbd\xeb\x12\xa7\x39\x3d\xf6\x7b\x7d\x0e\x63\xdc\x22\xf2\x9e\x62\xfe\xec\xeb\xab\x71\xd3\x71\x88\x1b\x68\x2b\x69\x68\xaa\xfa\x20\xb8\x49\x69\x5a\x47\x14\x6d\x08\xe2\x39\x1c\x24\x3b\x5f\xb3\x5f\x00\x00\x00\xff\xff\xd9\x16\x9a\xf6\xbf\x00\x00\x00")

func _001_init_up_sql() ([]byte, error) {
	return bindata_read(
		__001_init_up_sql,
		"001_init.up.sql",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"001_init.down.sql": _001_init_down_sql,
	"001_init.up.sql": _001_init_up_sql,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"001_init.down.sql": &_bintree_t{_001_init_down_sql, map[string]*_bintree_t{
	}},
	"001_init.up.sql": &_bintree_t{_001_init_up_sql, map[string]*_bintree_t{
	}},
}}
