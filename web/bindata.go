// Code generated by go-bindata.
// sources:
// pub/close.svg
// pub/index.css
// pub/index.html
// pub/index.js
// DO NOT EDIT!

package web

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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

var _closeSvg = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2c\x8e\xd1\xca\x84\x20\x14\x84\x5f\xe5\x70\xee\xff\xa3\x1e\xc9\xfc\x97\xec\x62\xaf\xdb\x87\x58\x28\x32\x68\xb7\xd8\x24\x63\x9f\x7e\xd5\x02\x19\xc6\x6f\x66\xc4\x66\xdb\x47\x38\x5e\xf3\x7b\x73\xe8\x43\x58\x6f\x42\xc4\x18\x29\x6a\x5a\x3e\xa3\x60\x29\xa5\x48\x0d\x84\x38\xf5\xc1\x3b\xd4\x06\xc1\x0f\xd3\xe8\xc3\xe9\xf7\x69\x88\xf7\xe5\x70\x28\x41\x82\x36\xe9\x60\xdb\xac\xcf\xe0\xa1\x77\xf8\x60\x4b\x15\xfc\x93\xe1\x8e\x0d\x69\x0b\x75\xba\x2a\x0b\xaa\x22\x6b\x0b\x2f\xa4\xbe\x4a\x17\x57\x67\xaf\x2c\x66\x26\xc5\x90\xa5\x4b\x98\x65\x32\xb3\xcd\x2f\x15\xc9\xfc\xaf\x84\x25\x49\xcb\x2f\x8a\xb6\xc9\x1f\x6e\x7f\x01\x00\x00\xff\xff\xe5\x79\xbd\x91\xd8\x00\x00\x00"

func closeSvgBytes() ([]byte, error) {
	return bindataRead(
		_closeSvg,
		"close.svg",
	)
}

func closeSvg() (*asset, error) {
	bytes, err := closeSvgBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "close.svg", size: 216, mode: os.FileMode(420), modTime: time.Unix(1435948337, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _indexCss = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x54\xdd\x6e\x1b\x2d\x10\xbd\xf7\x53\xa0\xe4\x22\xb1\x94\x75\x88\xfd\xd9\xfa\x4c\xa4\x3e\x44\xef\x7a\xc9\xc2\xec\x1a\x85\xdd\x59\x01\xeb\x9f\x54\x7d\xf7\x0e\xb0\x3f\xae\x9b\x48\xa9\x64\x5b\x30\x73\x06\xce\x9c\x39\xb8\x44\x7d\x61\x3f\x17\x8c\x95\x52\xbd\xd5\x0e\xfb\x56\x0b\x76\x5f\x55\xd5\x2b\xc5\x2a\x6c\x43\x51\xc9\xc6\xd8\x8b\x60\x0f\xdf\xa5\x85\x93\xbc\x3c\x3c\x31\x2f\x5b\x5f\x78\x70\x66\x46\x79\xf3\x0e\x82\xfd\xb7\xee\xce\x53\xe8\x04\xa6\x3e\x04\xc1\x36\x9c\xbf\x2e\x7e\x2d\x16\x15\xba\x26\x5d\x15\xe0\x1c\x0a\x69\x4d\xdd\x0a\xa6\xa0\x0d\xe0\x52\xfe\xbe\x94\x2e\xe5\x4f\x46\x87\x83\x60\xbb\xed\x70\x5a\x23\x5d\x6d\x08\xcb\x99\xec\x03\xc6\x48\x87\xde\x04\x83\x14\x73\x60\x65\x30\x47\xc8\x27\xf4\xce\xa6\x13\xfe\x99\xf8\xe6\x53\xe2\x33\x1d\xce\x33\xa6\x93\x5a\x9b\xb6\x16\x6c\xbd\xcd\x01\x85\x16\x1d\xa9\xb6\xdf\xef\xe3\xb6\x44\xa7\xc1\x15\x4e\x6a\xd3\x7b\xd2\x24\x83\x72\x54\xb0\x97\xee\xcc\x3c\x5a\xa3\xd9\xbd\x52\x2a\x66\xb0\x0f\xd6\xb4\x44\xa2\xc5\x16\x32\xf4\x5c\xf8\x83\xd4\x78\x8a\x3d\x13\x33\xb6\xa3\xaf\xab\x4b\xf9\xc8\x9f\xd8\xf0\x59\xad\x97\x53\xcf\xa2\x42\xd5\xfb\x3c\xc6\xbf\xaf\xe1\xfb\x6a\x46\x0a\xea\xae\x7c\x33\xa1\x30\x6d\xd7\x87\xa2\xb3\x52\xc1\x01\x2d\xd5\xa4\xf2\xa9\x17\xad\xf5\x75\x51\x83\xef\x5f\xc0\xaa\xa6\xbb\xf1\x52\x31\x62\x00\xe0\x43\xed\x06\x71\xb7\x3b\xfe\xd9\xac\xaf\xdd\x62\xa1\x0a\x37\x93\x5b\xbf\xe4\xc2\xe0\x68\xb6\x83\x29\xd2\x3a\xd9\x6d\xcd\x79\xe3\x19\x48\x0f\xd4\x70\x41\x4a\x4f\xd0\x98\x16\xcc\x2b\xf2\xc6\x8f\x47\xbe\xfc\x23\x5e\xa0\x33\x89\x45\xc0\x6e\x72\xe8\x34\xd8\x12\x43\xc0\xa6\x88\x5c\x3e\x1e\xf2\x88\x70\xd1\x46\x5f\xf5\xc1\xa8\xdf\x37\x26\x93\x86\xa3\x70\x69\x7a\x83\x0a\x1a\x14\x3a\x99\x7b\xcc\x66\x49\x55\x36\x8f\x7e\x7e\x14\xb2\xa4\x93\xfb\x90\x14\xa7\x16\x48\xcd\xb8\x72\xd9\xd5\x3c\xb3\x88\x0c\x87\xcd\x68\xf0\x91\xe1\x3c\x3d\xd3\xc8\x9a\x34\x26\x0f\x3c\xde\x3d\xfb\x67\x65\xd1\xc3\xca\x1f\xeb\xbb\xe5\x0d\x70\xbe\x3b\xcb\x75\xad\xda\x8c\x72\xd0\x81\x0c\x91\xfb\xb0\xbc\xc9\x0f\x7f\x22\xff\x93\x34\xf1\x27\xbd\xad\xde\xf9\xa8\x43\x87\x66\x3c\x10\x3b\xa9\x4c\xa0\x87\xcd\x57\x9b\xb8\xd7\xc6\x93\x35\x2f\x37\x92\x88\x03\x1e\x07\xa3\x5e\x15\xec\x62\xfe\x77\x00\x00\x00\xff\xff\x7e\x79\x1e\xec\xf4\x04\x00\x00"

func indexCssBytes() ([]byte, error) {
	return bindataRead(
		_indexCss,
		"index.css",
	)
}

func indexCss() (*asset, error) {
	bytes, err := indexCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.css", size: 1268, mode: os.FileMode(420), modTime: time.Unix(1435985561, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _indexHtml = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x52\x4d\x73\xd3\x30\x10\xbd\xf7\x57\x2c\x3a\x13\x6d\x42\x38\x90\x62\x87\x19\x92\x0c\x30\x53\xa0\x53\xc2\x00\x47\xc5\x5e\xc7\x0a\xb2\xe5\x4a\xeb\x26\xfe\xf7\x48\xb6\xdb\x86\xd2\x13\x3a\xed\xd7\x7b\x6f\x3f\x94\xbc\x58\x7f\x5d\x6d\x7f\x5d\x6f\xa0\xe4\xca\xc0\xf5\xf7\xf7\x57\x9f\x56\x20\x26\x88\x3f\xe6\x2b\xc4\xf5\x76\x0d\x3f\x3f\x6e\x3f\x5f\xc1\x4c\x4e\xe1\x1b\x3b\x9d\x31\xe2\xe6\x8b\x00\x51\x32\x37\x97\x88\xc7\xe3\x51\x1e\xe7\xd2\xba\x3d\x6e\x6f\xf0\x14\x59\x66\x11\x36\x9a\x13\xdf\x63\x64\xce\xb9\x58\x5e\x24\xbd\xc8\xa9\x32\xb5\x4f\x9f\x21\x98\x2d\x16\x8b\x01\x17\x6a\x01\x92\x92\x54\x1e\x8d\x60\xb2\x66\x43\xcb\x0f\x36\xc1\xc1\x1a\xa2\x15\xb1\x82\xc8\x33\xa1\xdb\x56\xdf\xa5\x62\x65\x6b\xa6\x9a\x27\xdb\xae\x21\x01\xd9\xe0\xa5\x82\xe9\xc4\x18\x79\xdf\x42\x56\x2a\xe7\x89\xd3\x96\x8b\xc9\x1b\x81\x23\x91\xd1\xf5\x6f\x28\x1d\x15\xa9\x40\x8f\xba\xce\xe9\x24\x33\xef\x45\x9f\x8d\xcf\x91\x49\x85\xe7\xce\x90\x2f\x89\xf8\x31\xc1\x41\x69\x14\x88\x80\x7f\xf9\xc6\x31\x8b\xd0\x8b\x97\x7b\x6b\xf7\x86\x54\xa3\xbd\xcc\x6c\x15\x11\xef\x0a\x55\x69\xd3\xa5\x37\xca\xd0\x51\x75\x97\xaf\xa7\xd3\x97\xf3\xe9\xf4\xff\x94\x13\xbc\x5f\x59\xb2\xb3\x79\x37\x36\x53\x58\x57\x81\x6a\xd9\x06\xc9\xc6\x10\x07\x94\x2d\x8a\xb1\xd5\x90\xcf\xf5\x1d\xe8\x3c\x15\x3b\xe5\x1e\x82\x67\xe1\xcc\x04\xf2\x04\x83\x7b\x96\xd4\x75\xd3\xf2\x59\x0b\xa2\x2f\x6d\x9d\x11\xd0\x18\x95\x51\x69\x4d\x4e\x2e\x15\x9b\x70\x01\x07\x5c\x12\x84\x1c\xb0\x05\x5f\x5a\x17\x8e\x12\x19\x7b\x8e\x87\x2e\xce\x05\x1e\xb5\xab\xe6\x2f\xed\x04\xe3\x30\xcb\x8b\xc1\xf1\x99\xd3\x0d\x83\x77\x59\x38\x1b\xaa\x83\x3a\x3d\x5d\x70\x8c\xa1\xd1\x3b\x8f\x87\xdb\x96\x5c\x87\xaf\xe4\x4c\xce\x47\x47\x56\xba\x96\x87\x7e\xb8\x81\x69\xf9\x0c\xed\xfd\x6f\x78\x5a\x97\xe0\xb0\xe1\xa4\xff\x57\xcb\x8b\x3f\x01\x00\x00\xff\xff\xfc\x61\x6a\x79\x49\x03\x00\x00"

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 841, mode: os.FileMode(420), modTime: time.Unix(1435948513, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _indexJs = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x55\x4d\x6f\x1b\x37\x10\xbd\xfb\x57\xb0\xa8\x11\x72\x51\x85\x92\xd3\x5b\x84\x1c\xda\xd8\x45\x5b\x04\x55\x51\xbb\x87\xa2\xe8\x81\xdd\x9d\x95\x18\x71\xc9\x05\xc9\xb5\xad\x3a\xfa\xef\x1d\x0e\xb9\xbb\x8a\xe0\x04\xfd\xd0\x41\x4b\xce\x0c\xc9\xc7\x37\x6f\x86\xa2\x1d\x6c\x1d\xb5\xb3\xa2\x62\x4f\x17\x17\xf7\xca\xb3\xcb\xd6\x77\xec\x0d\xbb\x14\xbc\x75\xbe\xe3\xd5\xe2\x82\xe1\xef\xb2\xee\xfa\x6c\xfd\x12\x47\xb3\xd5\x84\xd1\x6a\xc2\x64\x1d\xbc\x29\x56\x1c\x8d\x56\xa3\x42\xfc\xd5\x9b\x75\x3e\xc5\x43\xd0\x7f\x01\x46\x7d\x04\x80\xb1\xec\xab\x63\x5a\x8f\x40\xe4\x16\xa2\x58\x55\xe9\xf3\xad\x1b\x6c\xa3\xed\xf6\xad\xd1\x60\xe3\x2f\x18\x23\xaa\x35\xae\xa0\xb0\x3a\x04\xc1\x3b\xe5\xb7\xda\xbe\x8c\xae\xe7\x0b\xf6\xa0\x6d\xe3\x1e\xa4\xb6\x16\xfc\xf7\xa0\xb7\xbb\xb8\xfc\x9a\xbd\xa4\xbd\xe5\x2e\xcf\x5f\xe1\xfa\x63\xc1\x63\x55\x07\xdf\x79\xd7\x9d\x22\x1a\xbc\x9e\x41\xf5\xca\xc7\x74\x57\x34\xca\x30\xfc\x19\xa2\x47\x2c\xe2\xaa\x92\xa1\x37\x3a\x0a\xbe\xe4\x84\xc6\x43\x1c\xbc\xcd\xd1\xbf\x5f\xfd\x31\x1f\x60\x9c\x6a\x9e\xbf\x6e\x3a\x1a\x3d\x23\x02\x61\x5c\xad\x52\x88\xec\x55\xdc\x25\x6b\xbe\xa6\x54\xef\xd5\xa3\x78\x22\x2e\x91\xd6\xd7\x8c\x2f\x55\xaf\x97\x38\x5c\x72\xf6\x15\x2d\xcf\x44\x37\x2a\xaa\xbb\x43\x0f\x18\xf1\x3e\x38\xcb\xd1\x78\xac\xa4\x32\x0f\xea\x10\xe6\x7c\xa7\xa8\x8c\x81\x31\xdd\x32\xf1\x45\x32\x48\xb7\x1f\x6d\x8c\x2d\x97\xec\x6e\x73\xbd\x11\x7b\xeb\x7c\x74\xb6\x7a\xcd\x6e\xbc\x77\xbe\x78\xf3\x45\xd7\x34\x3b\x5e\xd0\x87\x72\xe7\x86\x98\x6e\x43\xbb\xd1\x64\x51\x16\x10\x6a\xf4\xe4\x88\x17\x2f\xf2\x40\x26\xe3\x87\x0f\x8c\xf3\xf5\xa4\x1d\x79\xaf\x0c\x92\x6f\x2a\xd9\xba\x7a\x08\x39\xcd\xb4\xfc\x5a\x37\x6f\x77\xca\x6e\x21\xdb\x8e\x29\x81\x99\xde\xb0\x73\x0f\xef\xb4\xdd\x9f\x52\x4c\xdc\x4d\x34\x1b\x72\x4e\xe4\x3a\xaf\x51\x2c\x48\x1c\x9f\xe8\x5b\xa7\x6b\x24\xa1\xcb\x16\xc5\x23\xb8\xe2\x95\xf4\xd0\xb9\x7b\x3a\xae\x6c\xa3\x48\xd9\x0d\x02\xeb\x50\x86\xb2\xf6\xa0\x22\xdc\x18\x48\x33\x5a\x52\x11\x58\xa9\x62\xf4\x82\xef\x3c\xb4\x28\x46\x3c\xbb\x98\x23\x3c\x46\x31\x4f\x55\xdf\x83\x6d\xee\x9c\xa0\x63\xc9\xc9\x71\x8b\x19\x09\x09\x3b\x7a\x65\x03\x55\xe3\x82\xf1\x50\x2b\x03\xbf\xa1\xf2\x78\x0e\xc3\xd2\xb8\x05\x03\x45\x55\x32\x60\xa5\xa8\x00\xdf\xd8\xe6\xe6\x31\x26\x4c\xaa\xd4\xd0\x82\xad\x16\x6c\x9e\x5c\x9d\x88\x7f\xa7\x1b\x38\x27\x2f\x13\xf7\x79\x0c\x2b\xc2\x30\xee\x72\x9a\x9f\xe7\x95\x5e\xfa\xc2\x98\xe2\x4a\x62\x11\x75\x39\x95\x49\x82\xe4\x7e\x33\x36\x8a\x51\x88\xb3\xd0\x48\x66\xc5\x4b\x95\x68\x88\x80\xb2\x74\x8c\x4f\x3d\x49\xb6\xaa\x81\x1f\xac\x78\xb5\x5a\x65\xa1\x30\x30\x01\xce\x03\x36\x43\x9c\x23\xe8\x1a\xd4\x4b\x10\x32\xc7\x0a\xef\x74\xc4\xab\x4e\xb7\x28\x4a\x02\xd9\x7b\xb8\x47\x5e\xaf\xa1\x55\x83\x29\x3d\xe8\x9f\x95\xf1\x58\x09\x9f\xa4\xe1\xbc\xcc\x63\xae\xe2\x9f\x37\xb7\x77\x7c\x31\xd6\x00\xfb\x7c\xe5\xa3\xfb\xc7\xdb\xcd\x4f\x32\x37\x28\xdd\x1e\xc4\x53\x59\x95\xfe\x8f\xd5\xc7\x3d\x82\xfd\xdf\x26\x31\x4a\x67\x2c\xd2\x7f\xd1\x18\xd6\xf3\xb6\x34\xff\x0f\x9b\x9e\x74\x94\xb9\x91\xcc\x0d\xa7\xa4\x24\xfb\x69\x72\xd2\x69\xce\x64\x93\x4e\x0d\xd1\xf9\x03\x96\x7c\x6f\x54\x0d\xb7\x11\x0b\x5b\x3c\x1d\x17\xcc\x0e\xc6\xa0\xe8\x97\xd0\xe8\x38\x31\x3e\x41\x1b\x3b\x8f\x38\xb1\x1e\xc7\xe6\x94\x92\x4a\x89\x4e\xa2\xda\xc3\x01\x9f\x24\x8b\xaa\x3a\x2d\x95\xd2\x09\x52\x40\x8f\xda\x86\x4f\xbb\x6b\x32\x9c\xfb\xd3\x09\x49\xd1\x14\x61\x74\xbd\x7f\x46\xb5\x93\xd6\x38\x9f\xdf\xcc\xac\xf1\xcc\xf1\x79\x6f\x25\xe0\xe5\x09\x55\x4d\x73\x93\x14\xff\x0e\xf9\x01\x7c\x4d\x05\xcf\x8f\x37\x9e\x93\x07\x18\x9b\x07\x69\xe5\xf9\x4e\xe9\xdd\x23\x6d\x1f\xab\xf4\xf9\x3b\x00\x00\xff\xff\xd8\xb9\xc7\xb1\x71\x08\x00\x00"

func indexJsBytes() ([]byte, error) {
	return bindataRead(
		_indexJs,
		"index.js",
	)
}

func indexJs() (*asset, error) {
	bytes, err := indexJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.js", size: 2161, mode: os.FileMode(420), modTime: time.Unix(1435986941, 0)}
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
	"close.svg": closeSvg,
	"index.css": indexCss,
	"index.html": indexHtml,
	"index.js": indexJs,
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
	"close.svg": &bintree{closeSvg, map[string]*bintree{
	}},
	"index.css": &bintree{indexCss, map[string]*bintree{
	}},
	"index.html": &bintree{indexHtml, map[string]*bintree{
	}},
	"index.js": &bintree{indexJs, map[string]*bintree{
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
