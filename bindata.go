package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
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

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _tmpl_about_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x56\x48\x49\x4d\xcb\xcc\x4b\x55\x50\x4a\x4c\xca\x2f\x2d\x51\x52\xa8\xad\xe5\x02\x0a\x96\xa4\xe6\x16\xe4\x24\x96\x00\x85\x33\x52\x13\x53\x52\x8b\x94\x14\x94\x1c\xe1\xf2\x0a\x0a\x36\x19\x86\x76\x60\xbe\x42\x78\x66\x76\x66\xb8\x63\x88\x8d\x3e\x50\x04\x24\x91\x92\x59\x06\xa2\x81\xac\x02\xbb\x90\x8c\x54\x05\xb0\xa9\x0a\x99\x79\x69\xf9\x36\xfa\x05\x60\x15\xfa\x60\x25\x28\x76\xa4\xe5\xe7\x97\x80\xec\x80\xd8\x9d\x9a\x97\x02\x64\x01\x02\x00\x00\xff\xff\x24\x88\x2c\xa7\x9a\x00\x00\x00")

func tmpl_about_html_bytes() ([]byte, error) {
	return bindata_read(
		_tmpl_about_html,
		"tmpl/about.html",
	)
}

func tmpl_about_html() (*asset, error) {
	bytes, err := tmpl_about_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "tmpl/about.html", size: 154, mode: os.FileMode(420), modTime: time.Unix(1418650467, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _tmpl_base_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x92\xc1\x8e\xd3\x30\x10\x86\xef\x79\x8a\x21\x88\xcb\x8a\xac\x2b\x21\x84\x14\xdc\x48\x2b\x40\xe2\x06\x87\x95\x56\x1c\xdd\x78\xd2\x8c\xd6\xb5\x2b\xc7\x59\x1a\x42\xde\x9d\xb1\xd3\x64\x5b\xd1\x93\x2d\xcf\xcc\xe7\x99\xf9\xff\x71\x04\x8d\x0d\x59\x84\xbc\x45\xa5\xd1\xe7\x30\x4d\x99\x7c\xf3\xf5\xc7\x97\xc7\x5f\x3f\xbf\x41\x1b\x0e\xa6\xca\xe4\x72\x70\x4a\x95\x01\xc8\x40\xc1\x60\x35\x8e\xf7\xd3\x04\x7f\xe1\x89\x9e\xe9\xe9\xe1\x51\x8a\xf9\x39\x26\x74\x61\x30\x08\x61\x38\xe2\x36\x0f\x78\x0a\xa2\xee\xba\x3c\x46\x20\x21\x61\x4c\x57\x80\x9d\x3b\x15\x1d\xfd\x21\xbb\x2f\xf9\xee\xb9\x81\x82\x9f\x3e\x2f\x51\x55\x3f\xef\xbd\xeb\xad\x2e\x6a\x67\x9c\x2f\xe1\x6d\xd3\xe0\x1c\x9d\xb2\x74\xdc\xbd\x87\xbb\x72\x87\x8d\xf3\x18\x6f\xaa\x09\xe8\x6f\xd2\xc9\xb6\xe8\x29\x5c\x15\x6b\x7a\xb9\xaf\x9d\x0d\x8a\x17\xf0\x5a\xf5\x9b\x74\x68\x4b\xf8\xf4\xf1\xdd\xd2\xc7\x41\xf9\x3d\xd9\x12\x36\xa0\xfa\xe0\x96\xd7\xa3\xd2\x3a\xa1\x3f\x6c\x8e\x27\xd8\x5c\x91\xad\x7a\x59\x79\x71\xfe\x42\x19\xda\x33\xa1\x46\xcb\x0d\xfe\x97\xda\xbf\x6e\x64\xa5\x6e\x96\x7f\x0c\x75\xa1\x48\x0b\x2d\xc1\x3a\x8b\xb7\xaa\x0d\xad\x00\x4d\xdd\xd1\xa8\x21\x4e\x6c\x68\xc9\xbe\x9c\x21\xb6\xbb\x30\x58\x2a\x91\xd0\xac\xae\x98\xe5\x95\x3b\xa7\x87\x24\x22\xb3\x67\xc9\x64\x6f\xaa\x33\x46\x1a\xaa\xa4\x82\xd6\x63\xb3\xcd\x45\x5e\x7d\x77\x07\x94\x42\x55\x52\x70\xe0\x66\x8e\xda\xb9\x3e\xe4\xd5\x43\x3c\xae\x33\xa5\x98\xb9\x52\x9c\x7f\x92\x2c\x08\xd4\x46\x75\xdd\x36\x5f\x75\x61\xdb\xb0\x4b\xd1\xea\xe8\xcc\xec\xc2\xb0\x8d\x73\xe1\x6c\xd8\xc8\xe0\xda\x38\xc4\xdc\x3d\x0f\x93\x2c\xbb\x56\xfe\x0b\x00\x00\xff\xff\x9c\xe3\xcd\x0e\xe9\x02\x00\x00")

func tmpl_base_html_bytes() ([]byte, error) {
	return bindata_read(
		_tmpl_base_html,
		"tmpl/base.html",
	)
}

func tmpl_base_html() (*asset, error) {
	bytes, err := tmpl_base_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "tmpl/base.html", size: 745, mode: os.FileMode(420), modTime: time.Unix(1418650826, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _tmpl_edit_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x50\x31\x6e\xc3\x30\x0c\xdc\xf3\x0a\x82\x0f\xa8\xd2\x4e\x1d\x64\x0d\x05\x3a\xb7\x40\xf2\x01\x26\xa2\x6b\x01\x96\x68\xd8\xb4\xdb\xc0\xf0\xdf\x4b\x21\x28\xdc\x6c\xd4\x1d\xef\x74\xc7\x75\x85\xc8\x6d\x2a\x0c\xc8\x31\x29\xc2\xb6\x1d\x0c\x53\xce\x43\x4f\x6a\x68\xc7\x14\x79\x44\x78\x3a\x27\xed\xb9\xd2\x00\xbe\x7b\x0e\xef\xb6\x9d\xca\x17\xac\xeb\x9d\xd9\x36\xef\x0c\x3e\x54\xba\x95\x31\x03\x5d\x35\x49\x69\xd0\x4d\xb4\xb0\xdb\xd7\x10\x32\x6b\x27\xb1\xc1\xcf\x8f\xd3\x19\x83\x09\x4c\x12\xd3\x12\xbc\xf2\x8f\xd2\xc8\x04\x85\x32\x37\x78\x91\x78\x43\x18\xe5\x7b\x6a\xf0\xe5\x88\x70\x95\xde\xa6\xd7\x23\x06\x73\x7b\x33\xb2\xfe\xf9\xa7\x09\xde\x55\x8f\x7f\x6e\xa9\x0c\xb3\x82\xde\x06\xb3\x9a\xe6\x4b\xae\xed\x16\xea\x67\x7b\x9e\x2c\x12\xee\x0a\xef\x6a\xe2\xf0\x58\xbc\x15\xd1\x5a\xfc\x7e\x10\x2e\xd1\xa6\xdf\x00\x00\x00\xff\xff\x8e\x73\x58\x65\x2e\x01\x00\x00")

func tmpl_edit_html_bytes() ([]byte, error) {
	return bindata_read(
		_tmpl_edit_html,
		"tmpl/edit.html",
	)
}

func tmpl_edit_html() (*asset, error) {
	bytes, err := tmpl_edit_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "tmpl/edit.html", size: 302, mode: os.FileMode(420), modTime: time.Unix(1418650785, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _tmpl_home_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x52\x4d\xcb\xdb\x30\x0c\xbe\xef\x57\xe8\xcd\x65\x97\xb2\xbc\xbb\x95\x12\x0a\x3b\x0c\x0a\x83\x9d\x0a\x3d\xab\x89\x12\x8b\xba\x56\x70\xd4\x98\x50\xf2\xdf\x27\xd7\xb4\xdd\xd8\x25\x51\x1e\xcb\x7a\x3e\xa2\xfb\x1d\x3a\xea\x39\x10\x54\x4e\xae\x54\xc1\xba\x7e\x31\x4c\xe9\x3a\x7a\xd4\x8c\x12\x76\x14\x2b\xa8\x0e\xcf\x63\x80\xc6\x7d\xdf\x9f\xc8\xb7\x86\x80\x0a\x9c\xf8\xc2\xa7\x1f\xc7\x8f\xa6\x36\x3c\x1f\x77\x3c\xe7\xb7\x55\xe3\xfe\xe8\x08\x92\x35\x40\x72\x14\x09\x82\xa8\xe3\x30\x00\x4f\xd0\x4b\xa4\x99\xe2\x06\xf2\x73\x79\xc1\xa1\x80\x18\xba\x52\x66\x0c\x7d\xc2\x65\xfa\xd6\xd4\xe3\x6b\xee\xc1\xa6\x7d\x9d\xc0\x49\x02\x56\x48\x12\x2f\xd3\xee\x7d\x2e\xbe\x14\x56\x7a\x36\x0d\x1c\x2e\x20\x3d\x4c\xa6\xb8\x10\x2d\x72\x83\x84\x41\xb3\xfe\x14\xd9\x9c\xe2\x59\x6e\x6a\x14\xd6\xff\xd7\xd5\xdf\x38\xf3\x90\x83\xb0\x3e\x2f\x2d\x7a\x27\x93\xee\xb6\x9f\xdb\xcf\x7a\x66\x4a\xf5\xd1\xa1\xe6\xf1\xc3\x7f\x37\x7f\xb2\x9a\x63\x98\xc8\xfc\x5b\x13\x48\xfe\x34\xc1\x38\xd3\x83\x51\x29\x3c\x4c\x52\x67\xfa\x59\x37\x20\x11\xda\x48\x99\x0b\xcd\x79\x2a\xa9\x8d\x38\x10\x44\x1e\x9c\x29\xcd\x01\x7e\xbc\x69\x9a\xfa\xe9\xd2\xe2\xf8\x45\x34\x02\x07\xb8\xb2\x8d\xd4\xcc\x87\x61\x91\x40\x39\xc8\xe5\x11\xfd\x06\x5a\x0c\x85\xcd\xcc\xc7\x32\x39\x99\x48\xfb\x27\x80\x6d\x2b\xb7\xa0\x78\x66\xcf\xba\x94\xa0\xff\x59\x83\x5e\x44\xf3\x1a\x94\xf5\x20\x23\x59\xd7\x3f\x01\x00\x00\xff\xff\x03\x75\x01\x54\x3c\x02\x00\x00")

func tmpl_home_html_bytes() ([]byte, error) {
	return bindata_read(
		_tmpl_home_html,
		"tmpl/home.html",
	)
}

func tmpl_home_html() (*asset, error) {
	bytes, err := tmpl_home_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "tmpl/home.html", size: 572, mode: os.FileMode(420), modTime: time.Unix(1418650369, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _tmpl_view_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x8e\xb1\x0a\x03\x21\x10\x44\xfb\xfb\x8a\xc5\x0f\x88\xa4\x37\x16\xf9\x86\x74\x21\x85\xb0\x23\x0a\x97\x53\x8e\xe5\x42\x10\xff\x3d\x2e\x16\xe1\xba\x9d\xb7\xc3\xcc\xb4\x46\x8c\x98\x37\x90\x39\x32\x3e\x86\x7a\x5f\x06\x13\xbc\xeb\x1a\x64\xd0\x84\xc0\xd8\x0d\x5d\x1e\x59\x56\xe8\x9b\xc8\xa5\xab\x6f\x6d\x92\xde\x9d\x1d\x72\x51\x5c\xfd\xd3\x05\x4a\x3b\xe2\xcd\x58\x70\x16\xfb\x77\x19\xaf\xc0\xd9\xe0\x5f\xce\xd6\xe9\xe7\x7c\x68\xce\xbd\xf0\x57\x63\x54\x9e\xcb\x63\x29\xa2\xe5\x73\x14\x36\x1e\xd7\x2f\x00\x00\xff\xff\xf6\xde\x88\x06\xb2\x00\x00\x00")

func tmpl_view_html_bytes() ([]byte, error) {
	return bindata_read(
		_tmpl_view_html,
		"tmpl/view.html",
	)
}

func tmpl_view_html() (*asset, error) {
	bytes, err := tmpl_view_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "tmpl/view.html", size: 178, mode: os.FileMode(420), modTime: time.Unix(1418650375, 0)}
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
	"tmpl/about.html": tmpl_about_html,
	"tmpl/base.html": tmpl_base_html,
	"tmpl/edit.html": tmpl_edit_html,
	"tmpl/home.html": tmpl_home_html,
	"tmpl/view.html": tmpl_view_html,
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
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"tmpl": &_bintree_t{nil, map[string]*_bintree_t{
		"about.html": &_bintree_t{tmpl_about_html, map[string]*_bintree_t{
		}},
		"base.html": &_bintree_t{tmpl_base_html, map[string]*_bintree_t{
		}},
		"edit.html": &_bintree_t{tmpl_edit_html, map[string]*_bintree_t{
		}},
		"home.html": &_bintree_t{tmpl_home_html, map[string]*_bintree_t{
		}},
		"view.html": &_bintree_t{tmpl_view_html, map[string]*_bintree_t{
		}},
	}},
}}

// Restore an asset under the given directory
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

