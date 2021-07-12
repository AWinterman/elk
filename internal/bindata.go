// Package internal Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// template/http/create.tmpl
// template/http/handler.tmpl
// template/http/list.tmpl
// template/http/read.tmpl
// template/http/update.tmpl
package internal

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templateHttpCreateTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x57\x5f\x6f\x13\x39\x10\x7f\xef\xa7\x98\x5b\x45\x68\xd3\x4b\x1d\xee\xb5\x08\x9d\x80\x16\x88\x0e\x5a\x8e\xf4\xe0\x01\x21\xe1\xac\x67\x77\x4d\x1d\x7b\x6b\x7b\x13\x42\xd8\xef\x7e\xb2\xbd\xc9\x76\x37\x7f\x08\x41\x10\x09\x91\xb8\x33\xe3\x99\xdf\xcc\x6f\x66\xbc\x5c\x0e\x4f\x21\x53\x76\x51\xe0\x39\xa0\xb4\x99\x22\x5c\x0d\x51\x5a\xf7\x2f\x19\x66\x28\xc9\x0b\x4d\x8b\x1c\x4e\x87\x55\x75\x72\xb2\x5c\x02\xc3\x94\x4b\x84\x28\xb7\xb6\x18\x26\x1a\xa9\xc5\x08\xaa\xea\x04\x00\x60\xb9\x3c\x83\x39\xb7\x39\xe0\x17\x8b\x92\x41\x0f\xa2\x37\x34\xb9\xa5\x19\x46\x41\x21\x82\xb3\x5a\x34\x88\x83\xc5\x69\x21\xa8\x75\xf6\x90\x32\xd4\x11\x90\xc6\x18\x38\x1b\xf5\x2f\x3e\x2d\x94\xb6\x10\x65\xdc\xe6\xe5\x84\x24\x6a\x3a\xcc\xd4\x59\x21\xe8\x22\xd3\xaa\x94\x6c\x38\xa3\x82\x33\x6a\x95\x1e\xce\xfe\x7a\x18\x81\x0f\xec\x26\xe7\x06\xb8\x01\x89\xc8\x90\x41\xaa\x34\x18\x5b\x16\x9c\xc1\x78\xf4\x0e\x74\x29\xb0\x8e\xab\xbe\x4f\x53\x99\x21\xf4\x24\x9c\x3f\x86\x1e\xb9\x52\x0c\x0d\xdc\xf3\x77\x38\x84\x37\x74\x21\x14\x65\xa0\x52\xa0\x4e\xa3\x27\xc9\x15\x9d\x22\x54\x15\x04\x2c\x40\xe3\x5d\x89\xc6\x92\xb5\x96\xc3\xb6\x2d\xfa\xcc\x4b\xbe\x0d\x82\x60\xac\x2e\x13\x0b\xcb\xb5\x42\x70\xc6\xb9\x7f\x7d\x71\x7d\x0e\x2f\xe9\x8c\xcb\x0c\xa8\x10\x50\x28\x2e\x2d\x6a\x03\x39\x6a\x04\x83\x38\x35\x20\x95\x05\xcd\xb3\xdc\x12\x78\x4d\x17\x13\x04\xeb\x82\x4e\xa8\x84\x09\x02\x53\x12\x81\x4b\xa0\x52\xd9\x1c\x35\xcc\xe9\x02\x08\x21\x21\xea\xf6\x7d\xab\xe0\x53\x1f\xbc\x24\xcf\x39\x0a\x66\x5a\xf9\xba\x27\xdc\x4b\xc9\xd8\xfb\xed\xc5\x5c\xf8\xa7\xe1\xf4\x66\x51\xa0\xfb\x93\x73\xb9\xaa\xe0\xd3\x67\xa3\xe4\x79\xb4\x5c\x02\x97\x0c\xbf\x40\x6c\x0a\xc1\x2d\xc4\x96\x66\xaf\x94\xba\x2d\x8b\xc6\xd2\x0d\xcd\x20\x72\xe2\x51\x1f\xa2\x41\xd4\x87\x87\x50\x55\x4e\xd3\x57\x54\x2f\x25\x4f\xa4\x54\x96\x5a\xae\xa4\x21\x97\xe2\x16\xbe\x01\x8a\xdb\xe6\xd0\x5d\x57\x17\x01\xfa\x1b\xc9\xbb\xf0\x8b\x2b\x79\x43\x33\x03\x51\x53\xad\xd1\xba\xb8\x3e\x75\x61\x70\xc7\xdd\xa0\x1b\x74\xb0\x46\xe7\x92\x65\xb8\x1b\x1c\xec\x80\xe3\xc2\x4f\xdd\xf1\x7f\x92\xdf\x95\xae\x04\x1c\x5a\x28\x8c\xfb\xfa\xe1\xe3\xda\x99\xa0\xeb\x21\x1c\x5d\x1c\x01\x25\x1e\x06\x25\xfe\x46\x28\xef\x21\x54\xf3\xac\x66\x52\x20\x41\xcd\x1a\x03\x14\x24\xce\x3b\x8c\xa2\x92\x81\xb1\x4a\xa3\x01\x6e\x5d\x19\xdb\x1c\x81\x51\x4b\x27\xd4\x60\x43\xb0\xb4\x94\x09\xc4\x79\x5b\xf9\x25\x95\x4c\xa0\xee\xd7\xf7\xc4\x73\x70\xfd\x87\xbc\x45\x53\x28\x69\xf0\xbd\xe6\x16\xf5\x00\x34\x9c\xd6\xe7\x9e\x8c\xfd\x0e\x0d\x85\x4b\x77\x4e\x84\xca\xc8\x7b\x6e\xf3\xf8\x2b\x2d\xea\x8c\xc4\xd1\x14\x6d\xae\x58\x34\x80\x28\x5c\x11\xf5\xfb\x2d\xdd\xe1\x10\x5e\xa0\xf5\x3e\x17\xca\x58\xef\x38\x69\x49\xcc\xa8\x06\xb6\xa7\x35\xb4\x84\x79\x0a\xa8\xb5\xf3\xc7\x25\x96\x5c\xe1\xfc\x02\x13\xc5\x50\xc7\x9a\x3c\x55\x6c\xd1\x27\xe1\x77\xfc\x80\xf5\x1f\x79\xd1\x3f\x1e\x83\xe4\xa2\x13\x91\x8f\x8a\x5c\x6a\xad\x74\x1c\xa1\xfb\x0f\x98\xd3\x73\x45\xe6\x2b\x66\x00\x2e\xc8\x20\x80\x5a\x77\x82\x72\x1f\x8d\x92\xa1\x26\x4f\x29\xab\xfd\x8c\xe7\x03\xd0\x03\x88\xb8\xf4\x45\xe3\xed\xb8\xb6\xc6\x65\x16\x6d\x53\xb7\xa5\x96\xad\xe3\xaa\x0b\x5c\x5d\x6e\xb8\xce\x38\xd9\x81\x45\x4e\xd6\x7d\xbf\xae\xfd\xf8\xbb\xe1\x07\xed\x01\xa8\x5b\x67\x01\xb5\x26\xf1\x69\x63\x65\x14\x82\x68\x0a\xde\x43\xd1\x7f\xe4\xc4\x37\x6d\x6d\x81\xb3\x36\xe5\x00\xad\x67\x81\x8f\xe0\x00\x60\xef\x81\x3b\x72\x7d\x5e\x52\x31\x46\x3d\x43\x1d\x94\x02\xca\x92\x8b\x5d\x9a\x1b\xb8\x6e\x62\x1b\xfc\x1d\xc9\x54\xc5\xd1\x6c\x1d\x22\xa4\x94\x0b\x64\x3f\x91\x7b\x27\x7d\x54\xa6\xc7\x74\xb6\x2b\xcb\x93\x90\xe0\x44\x70\x94\x96\xb4\x68\x42\x6a\x56\x6f\x30\x2e\x8c\xcc\x79\x4e\x2d\xd0\x89\x2a\x2d\x18\xc1\x13\x84\x34\x0c\x33\xeb\xce\x73\x77\x63\x52\x1a\xab\xa6\x30\xa5\xda\xe4\x54\x08\xd4\xe6\xef\xe3\xe7\x21\x4f\x81\x91\xad\x43\x71\x67\x0d\xfa\xf8\xc8\x18\xed\x36\xb5\xf8\x74\xbb\xb9\x4d\x84\x37\xc6\xd4\xcf\x4e\xaf\x75\x28\xdd\x11\xb6\x3f\x94\xee\x80\xdb\x66\xba\x09\x3b\xd8\x7f\x5d\x86\x61\x33\x46\x7b\x2f\xe6\xee\xbd\xdb\x6b\x7d\x35\x3d\x7f\xe4\x9e\x27\xcc\x63\xbb\xfd\x1a\x42\xc8\xae\x9b\xce\xba\x73\x6c\xf5\x39\x08\x7c\x57\xe2\x6e\x84\xb9\xf1\xb5\x39\xba\xdc\x07\x07\xab\x66\x36\x21\x8e\x0d\xb1\x26\xcf\x94\xb4\xf8\xc5\xc6\x1d\x02\xd6\x6d\xef\xd0\xbe\x6e\xc2\xe2\x78\x8f\x38\xdf\xe0\x16\x27\x74\xe2\xa6\xf6\xc1\x54\xff\xc1\x4e\x74\x08\xe7\xdf\xa2\xdf\xa1\x51\x5a\xbd\x68\x63\x71\xb7\x87\xf2\xff\x96\xa8\x17\x71\x9f\xbc\x77\xcb\x6f\xdc\x8a\x4a\xa8\x39\x6a\x27\x33\xba\x88\xb1\x56\x1b\x5d\x74\x2b\xa9\x7f\xd2\x4d\x6c\x78\xab\x38\x32\xdc\xa8\x57\xce\xa5\x9e\xbc\xbf\xda\x6c\x04\x37\x1c\xc2\x25\xcd\x50\x43\xf0\xdf\xb3\xc8\xf7\x14\xaa\xc3\xe2\xcf\x35\x32\x50\x72\xf5\x16\x50\x05\x6a\x5f\x7b\x64\xdb\x9a\x48\x9e\x29\x86\x10\xdd\x6d\x5c\xb5\xa3\xe6\xea\x42\x79\x0c\x77\xe4\x5a\x8a\xc5\xf1\x75\x62\xe6\xdc\x26\x79\x98\x7c\xee\x75\xd2\xdd\x7a\x56\x9f\x84\x1a\x84\x53\x97\x89\x2b\x65\x9f\xbb\x57\x96\x4f\xff\xf9\x4e\xca\xad\x26\xcb\xd6\x92\xf3\x2f\x95\xd4\x59\xa9\x8b\x6f\x24\xed\x4a\x74\x74\xb1\x4a\x73\x34\x80\xdd\x19\x3c\x6c\x84\x42\x53\xbc\x2b\xbf\x57\x1b\xca\xf7\x1c\xdb\x67\x70\xeb\x74\xdd\x80\x69\xcc\x65\x56\x0a\xaa\xbf\x8b\x54\x4d\x55\x56\x16\x82\x27\xae\x58\x3c\x1d\xfc\x1b\x75\x1f\x63\x7f\x03\x68\x9b\x8b\x5d\xe3\xe4\x76\x00\x83\xeb\x9e\x4d\x9c\x41\x04\x7f\xba\xed\x2f\x51\x72\x46\x46\x56\xd1\x83\x29\x79\x20\xde\x0c\x53\x5a\x0a\x7b\x00\xb8\xa1\x0f\xa6\x68\x93\x7c\x67\x27\x84\x54\xab\x29\xb0\xc9\xef\xc3\xf7\x88\xdd\x6e\x0f\x22\xd5\x9e\x5e\xfb\x79\x3d\x5a\x4c\x8e\x9a\xa7\x29\x79\x1d\x16\x9e\xf8\xc1\xea\xe0\xba\xf0\xaf\xc0\xcd\x06\x30\x92\x89\x28\x19\x5e\x4e\x0b\xbb\xb8\xa1\xd9\x39\x58\x5d\xe2\x60\x43\xec\x85\x56\x65\x61\xce\xe1\xc3\xc7\xb0\xf0\x2f\x9b\xc6\xda\xeb\x3e\x33\xc7\x49\x8e\x53\x5a\xef\x6e\x41\xd1\xef\x72\x8a\x61\xf3\x1c\xde\xce\xd1\xa8\x19\xaf\x50\xb5\xbd\xa8\x06\x80\x47\xce\x49\x83\x9a\x53\xc1\xbf\x86\x25\xd8\x57\xcb\x2f\x2c\x83\x5f\x34\x50\xf7\x37\xdd\x70\x29\x1e\xdf\x73\xdb\x1e\xd5\x31\x5c\xff\x53\xbb\xfc\xb9\x7f\xd2\xf6\xaa\x79\xf6\xaf\xbf\xfd\x1f\x00\x00\xff\xff\xc7\x14\x34\x6d\x64\x14\x00\x00")

func templateHttpCreateTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateHttpCreateTmpl,
		"template/http/create.tmpl",
	)
}

func templateHttpCreateTmpl() (*asset, error) {
	bytes, err := templateHttpCreateTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/http/create.tmpl", size: 5220, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateHttpHandlerTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x92\x41\x8b\xdb\x3e\x10\xc5\xef\xf9\x14\x0f\x11\xfe\x24\x26\x91\xfe\xbd\x06\x7a\xea\xa1\x2d\x94\x50\xd8\xb2\x3d\xab\xd6\x44\x16\x95\x25\x23\x8f\xb3\xdd\x1a\x7f\xf7\xa2\xd8\xeb\xc4\xac\x05\x82\x27\xf4\xd3\x9b\x79\x83\xfa\x5e\x15\xb0\x91\x5f\x1b\x3a\x81\x02\xdb\x28\x5d\x54\x14\x38\xef\x52\x59\x0a\xf2\x73\xd2\x4d\x85\x42\x0d\xc3\x66\xd3\xf7\x30\x74\x71\x81\x20\x2a\xe6\x46\x55\x3a\x18\x4f\x49\x60\x18\x36\x00\xd0\xf7\x47\xbc\x38\xae\x40\x7f\x98\x82\xc1\x16\xe2\xbb\x2e\x7f\x6b\x4b\x62\x7c\x21\x70\x9c\xd0\x11\x07\x53\xdd\x78\xcd\xd9\x90\xb4\xc9\x56\xf2\x6e\x86\xec\x31\x9d\x5c\xdd\xc4\xc4\x10\xd6\x71\xd5\xfd\x92\x65\xac\x95\x8d\xc7\xc6\xeb\x57\x9b\x62\x17\x8c\xba\x6a\xef\x8c\xe6\x98\xd4\xf5\xc3\xff\x02\xb7\x64\x3f\x2a\xd7\xc2\xb5\x08\x44\x86\x0c\x2e\x31\xa1\xe5\xae\x71\x06\x4f\x5f\x9f\x91\x3a\x4f\x53\xb0\xa9\x5e\xd2\xc1\x12\xb6\x01\xa7\x8f\xd8\xca\x73\x34\xd4\xe2\xa1\x5f\xa5\x32\xb4\x0d\xf2\xac\x6b\xc2\x30\x7c\x19\xe3\x63\x1c\x43\x8b\x9c\x10\x65\xea\x0c\x62\x43\x49\xb3\x8b\xa1\x45\x0c\xe0\x8a\x96\x0f\x51\x47\x43\x5e\xce\xc6\x79\xfe\xeb\xd6\x2d\xa7\xae\x64\xf4\x33\x9a\x57\xe9\x1d\x05\xce\xaa\xa0\xc0\xf2\xd3\xed\xb8\x20\x7c\xb4\x93\x2a\xfe\xea\x46\x7e\x8b\xd6\x52\x5a\x10\xf3\xb8\x50\xcc\x52\x3e\x8f\x8a\x66\x72\x1a\x4d\x5e\x97\x2e\x94\x38\xd3\xcb\x5a\x9b\xbb\xf2\xb1\x93\x03\xfc\x63\xd9\x03\xae\x6b\x35\xf6\x28\x56\x13\x2f\xa3\x26\xe2\x2e\x05\xfc\xb7\x86\x2e\xc9\xfb\x60\x4e\x37\x79\x78\x77\xeb\xa3\x3d\xbd\x49\xf9\xd3\x71\xb5\xcb\x4d\x3e\x71\x72\xc1\xee\xc4\xdb\x5f\x3e\x40\xac\x15\x13\xfb\xfd\x7b\xc7\x39\xd5\x09\xd7\xe5\xed\xb0\x59\xaa\xfb\x67\x9e\xd5\xbf\x00\x00\x00\xff\xff\xee\x48\x80\xb8\x7c\x03\x00\x00")

func templateHttpHandlerTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateHttpHandlerTmpl,
		"template/http/handler.tmpl",
	)
}

func templateHttpHandlerTmpl() (*asset, error) {
	bytes, err := templateHttpHandlerTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/http/handler.tmpl", size: 892, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateHttpListTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x56\x51\x6f\xdc\x36\x0c\x7e\xef\xaf\x60\x85\x60\xb5\x03\x47\xb7\x62\x6f\x37\xdc\x43\x56\x04\x59\xb0\xac\xc9\x92\x0e\x7d\x18\xf6\xa0\x3b\xd1\xb6\x10\x59\xf2\x51\x74\xd2\xc4\xf5\x7f\x1f\x64\x3b\x89\x7d\x97\xb5\xc3\x30\xac\x02\x0c\xd8\x34\xf9\xf1\x13\x3f\x9a\x72\xdb\x2e\x0e\xa1\xf0\x7c\x5f\xe3\x12\xd0\x71\xe1\xa5\xf1\x0b\x74\x1c\xaf\xcd\xa2\x40\x27\x4f\x49\xd5\x25\x1c\x2e\xba\xee\xd5\xab\xb6\x05\x8d\xb9\x71\x08\xa2\x64\xae\x17\xd6\x04\x16\xd0\x75\xaf\x00\x00\xda\xf6\x08\xee\x0c\x97\x80\x9f\x18\x9d\x86\x03\x10\x97\x6a\x73\xa3\x0a\x14\x83\xbb\x80\xa3\xd1\x75\x70\x07\xc6\xaa\xb6\x8a\x23\x1a\x2a\x8d\x24\x40\x3e\x83\x41\xc4\x88\x49\xc7\x47\x52\xae\x40\x38\x70\xb0\x5c\xc1\x81\x7c\xef\x35\x06\x98\xc0\x2d\x16\x70\x85\x4a\x43\x8e\xbc\x29\x31\x00\x97\x18\xa3\x0e\x9c\x7c\xaf\x2a\x84\xae\x83\xca\x6b\xb4\x60\x34\x3a\x36\xb9\x41\x0d\xeb\x7b\x50\x50\x98\x5b\x74\xd0\x90\x3d\xaa\x15\xa9\x0a\x19\x09\x72\xf2\x55\x04\x98\x82\x6b\xc5\x6a\xad\x02\x82\x72\x1a\x08\xb9\x21\x17\xc0\x30\xb0\xef\x53\x6d\xac\x41\xc7\xf2\x29\x22\x6f\xdc\x06\x92\x12\x0e\x67\x1c\x7e\x56\x4e\x5b\xa4\x14\xce\x4d\xe0\xe4\x0e\x62\x55\xe4\x15\x86\xda\xbb\x80\x1f\xc9\x30\x52\x06\x04\x87\xa3\x7d\xdb\x60\xe0\x14\xda\x27\xd4\xb8\x6c\x2c\x40\x29\xad\x2f\xe4\x47\xc3\x65\xf2\xa0\x6a\x79\xcd\x64\x5c\x91\x88\x0a\xb9\xf4\x5a\x64\x20\x62\x02\x91\xa6\xb3\xc8\xed\x10\x39\x52\x9d\x11\x93\xbf\x35\x48\xf7\xc9\xdc\xff\x59\x50\x5d\x60\xf8\xe0\xcf\xbd\xd2\x51\x00\x31\x93\x7d\xba\x16\x0b\x38\x51\x05\x12\xd8\xe8\xda\x87\x01\x97\x8a\x41\x11\x02\xe1\xb6\x31\x84\x1a\xbc\x83\x88\x00\xbe\x46\x52\x6c\xbc\x93\x7b\x40\x6d\x0b\xf2\x9d\xd7\x08\x62\xbb\x97\x28\xd2\x1a\x7b\x63\x6a\xbe\x55\x04\x48\xfd\xe5\x69\xf6\xa6\x56\x05\xc6\xad\xbf\x9d\x59\x4d\x0e\x3a\x5a\x49\xfe\x7e\x75\xfe\xb8\x7f\x79\x8a\x9c\x88\x18\x20\xd2\x1f\x41\xc3\xeb\x15\x08\xb1\x23\xc0\x23\x64\xd6\xa7\x5b\x41\x60\xda\x78\x77\x2b\x8f\xd9\x9b\x44\xa7\x7b\xbe\x26\xef\x1d\x5f\xaf\xc0\x19\xfb\x02\x56\x2f\xaa\x3c\x73\xb9\x4f\x44\x4f\x1e\x6a\x45\xc1\xb8\x02\xb6\x91\x14\x3c\xf7\xe5\x9b\x98\xf6\x8d\xc8\x60\xaa\x79\x4f\x36\x03\x9d\x0e\xe6\x93\x88\x90\x20\x51\xba\xcf\x24\x2e\x42\xa7\x91\xe4\x4f\x4a\x8f\xfd\x95\xdc\x65\x40\x19\xf4\x38\x50\x35\x81\x61\x1d\x9b\x1c\x8c\x63\x8c\x52\x16\x84\x2a\x26\x7f\x40\xf2\xe2\xef\x30\xe3\xe7\xb0\xf7\x6a\xae\xcf\xfc\xc9\x30\x56\xe1\x12\xe9\x72\xd4\xe6\x87\xef\xff\xa9\x38\xd3\xc8\x2f\x8b\x34\xf5\xfc\x06\x62\x4d\xd3\xef\x8a\x36\xdb\xc4\x7f\x20\xde\xac\x9c\xff\x9b\x88\x18\x86\xb2\x2e\x57\xb0\x95\xe7\xa6\x32\x9c\x4c\x89\xa4\xf2\x22\xcf\x03\x72\x92\xf4\xbd\x75\x04\x6f\x53\x38\x84\xb9\xc7\xb1\xb5\x09\xc9\x77\xde\x31\x7e\xe2\x64\x67\xdb\x5f\x13\xc3\x8e\x05\x1b\x95\xe8\x47\x7f\x94\x62\x32\xda\x3e\xc3\x0d\xae\xd5\x1a\x3e\x43\x6d\x1b\x52\xb6\xeb\x86\xe9\xae\xd7\xe2\xeb\x25\x1f\xcb\x7d\xe6\x18\xc9\x29\x7b\x8d\x74\x8b\x34\x04\x0c\x75\x77\xc6\xbe\x14\xb5\x57\xc8\x79\xd9\xf4\x53\xd5\x42\x89\x64\xf2\x5c\xfe\xaa\x28\x94\xca\x26\xdf\x3d\x1a\x2e\xea\x38\x17\xc3\xfe\x96\xcf\xdc\xc6\x36\x1a\x4f\xaa\x9a\xef\x3f\xa8\x62\x09\x4c\x0d\x66\x7b\x6e\xa7\xe4\x9b\x3a\x2c\xe1\x8f\x3f\x43\xdf\x71\xed\xf3\x30\x3f\x90\xc7\xce\x79\xee\x07\x6f\x90\x27\xf6\xe6\x7a\x53\x62\xa5\x64\x3c\x3d\x87\xb0\x78\x1e\xc4\xe1\x1b\xcf\x60\x1b\xe2\xf1\x20\x5e\xa8\xe8\x60\x8d\x83\xf8\xa8\xeb\xa0\x9b\x73\xe8\x32\xc0\xf0\x2f\xc5\x0c\x48\x46\x59\xf3\xd0\x33\x1c\xc6\xf9\xb7\xd3\xea\xf1\x5b\xff\x42\x4b\xc5\x7f\x8b\x21\x3b\xea\x91\xe8\x99\xe3\x44\xa8\xca\x37\x8e\x45\x06\x16\x5d\x82\x21\xdd\x21\x3d\x12\xbe\xf8\x65\xe4\x37\x19\x46\xbb\xbf\x40\x4f\x77\x7f\x05\x00\x00\xff\xff\x94\xf1\xf1\xbd\xaf\x09\x00\x00")

func templateHttpListTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateHttpListTmpl,
		"template/http/list.tmpl",
	)
}

func templateHttpListTmpl() (*asset, error) {
	bytes, err := templateHttpListTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/http/list.tmpl", size: 2479, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateHttpReadTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\x4d\x4f\xe4\x46\x13\xbe\xf3\x2b\xea\xb5\xd0\x2b\x0f\x19\x7a\x4e\xb9\x20\xed\x61\x03\x84\xb5\x42\x60\x03\x6c\x38\x44\x39\xf4\xb8\xcb\x76\x0b\x4f\xb7\xa7\xba\x0c\x3b\x4c\xe6\xbf\x47\x65\x7b\x3e\xed\x04\x12\x65\x5b\xb2\xc0\xee\xaa\xea\xa7\x3e\x9e\xa7\x67\xb9\x9c\x9c\x40\xee\x79\x51\xe1\x19\xa0\xe3\xdc\x2b\xeb\x27\xe8\x58\x9e\x74\x92\xa3\x53\x57\xa4\xab\x02\x4e\x26\xab\xd5\xd1\xd1\x72\x09\x06\x33\xeb\x10\xa2\x82\xb9\x9a\x10\x6a\x13\xc1\x6a\x75\x04\x00\xb0\x5c\x9e\xc2\x8b\xe5\x02\xf0\x2b\xa3\x33\x70\x0c\xd1\x67\x9d\x3e\xe9\x1c\xa3\xd6\x3c\x82\xd3\xce\xb4\x35\x07\xc6\x59\x55\x6a\x96\x68\xa8\x0d\x52\x04\x6a\x1b\x0c\x24\x46\xf7\x66\x67\x95\x27\x86\x28\xb7\x5c\xd4\x53\x95\xfa\xd9\x24\xf7\xa7\x69\x61\x27\xf2\x3c\x7f\x1f\x41\x93\xc8\x43\x61\x03\xd8\x00\x0e\xd1\xa0\x81\xcc\x13\x04\xae\x2b\x6b\xe0\x3e\xf9\x15\xa8\x2e\xb1\xcb\xa3\x3b\x81\xb4\xcb\x11\x8e\x1d\x9c\x7d\x80\x63\x75\xe3\x0d\x06\xd8\x41\x38\x99\xc0\x1d\x6a\x03\x19\x72\x5a\x60\x00\x2e\x50\xbc\x8e\x9d\xba\xd1\x33\x84\xd5\x0a\x66\xde\x60\x09\xd6\xa0\x63\x9b\x59\x34\x30\x5d\x80\x86\xdc\x3e\xa3\x83\x9a\xca\xd3\x4a\x93\x9e\x21\x23\x41\x46\x7e\x26\x01\x76\x83\x1b\xcd\x7a\xaa\x03\x82\x76\x06\x08\xb9\x26\x17\xc0\x32\xb0\x6f\x8e\x4a\x4b\x8b\x8e\xd5\xc6\x23\xab\x5d\x0a\x71\x01\x27\x7b\x18\x3e\x69\x67\x4a\xa4\x51\x03\x35\x7e\x01\x29\xb4\xba\xc3\x50\x79\x17\xf0\x91\x2c\x23\x8d\x81\xe0\xa4\xfb\x3e\xaf\x31\xf0\x08\x96\x9b\xa8\xb2\x4a\x29\x40\xa1\x4a\x9f\xab\x47\xcb\x45\xfc\xaa\x2b\x75\xcf\x64\x5d\x1e\x47\x33\xe4\xc2\x9b\x68\x0c\x91\x1c\x10\x8d\x46\x7b\x9e\x93\x09\x24\x17\x52\xf2\x2f\x77\xd7\xb0\xc9\x56\xed\xd9\xc8\x5c\xd8\x4c\x20\x27\x17\x2a\x09\x89\xe3\xdd\x22\xaf\x97\x35\x63\x40\x22\x41\x12\x98\x52\xef\x9e\xd5\x47\xf6\x36\x4e\x0b\xab\xbe\xdc\x5d\x7f\x96\xd8\x31\x8d\x21\xb2\x3d\x10\x8d\x7b\xd6\x78\xff\xef\x03\x38\x5b\x1e\xa4\xb7\x49\x53\x5d\x12\x79\x8a\x23\x94\x3f\x90\x23\xb3\x75\x39\x58\xd3\x76\xa7\xa6\x72\x9b\x42\x34\x86\xdd\x2a\x58\xa9\xc0\x30\x94\xd6\xb0\x8d\x8c\x44\x03\xd8\x64\x11\x3a\x83\xa4\x7e\xd0\xa6\xeb\x41\xfc\x32\x86\x36\x06\xcc\xea\xc0\x30\x95\x31\x00\xeb\x18\x73\x24\xc8\x09\xb5\x8c\xcd\x2b\x92\x8f\xfe\x2a\xa2\x0c\x4c\x6f\x6b\x75\x50\x7b\xc0\x32\xe0\x70\xc1\xa5\xd6\x83\x39\xf5\x42\x38\xd3\x50\xf7\xb0\xf5\xe7\x0d\xca\x66\x58\xe7\x35\xd2\x42\x26\xb7\x21\x4b\x9f\x2a\x7b\xae\xf3\x76\xde\xba\x01\xdf\xb3\x53\xbf\x48\xa0\x78\xa4\x1e\x0b\x24\x8c\x77\xf6\xfe\x80\xd2\xbf\x20\x89\x4d\x72\x21\x1b\x9b\x99\xfa\xa4\xc3\x95\x7f\x58\x54\xe2\xdf\x3a\x24\x17\xaa\x7b\x8f\xad\x19\x6d\x8b\x60\xcd\x46\x57\x46\xa3\xa3\xde\x94\xb6\xea\x65\x72\x0c\x0f\xfe\xda\x6b\x23\xd2\x10\xed\x69\xdc\x41\x01\x2e\xb5\x34\xab\x14\xd3\xc6\x0d\xb8\xd0\x0c\x9a\x10\x08\xe7\xb5\x25\x34\xe0\x1d\x48\x04\xf0\x15\x92\x66\xeb\x9d\xea\x05\x5a\x2e\x41\x9d\x7b\x83\x10\xcd\x7b\x07\x09\xac\x1d\x21\x5c\x2f\xdc\x90\x65\xae\x6e\x5d\xb9\x88\x49\x9d\x7b\xc7\xf8\x95\xe3\x83\x01\x7c\x8b\x18\xe1\xc5\x4a\xc3\x90\x48\xc5\x72\x0d\x1c\x6a\xc3\x7a\xa5\x22\x54\x27\xd2\xaf\x1b\xcf\x3f\xfa\xda\x99\x66\xe4\xcf\x06\x8d\xa1\x61\x5b\xe2\x32\x1f\x47\x7b\x2d\x7c\xc2\xa9\x9e\x8a\x72\x3a\xcf\x90\x49\x94\x8e\x68\x89\xe3\xb5\x69\x72\xb1\x1e\x86\x68\x0c\xd6\xbc\x93\x60\xb0\x25\xd9\x1a\xe0\x9a\x62\x6f\x21\x18\x0e\xb8\x97\xf0\xbd\x75\x79\x5d\x6a\x7a\x33\xe7\x4e\x61\x4c\x5d\x95\x36\x15\x72\xa0\x63\x5a\x34\x57\xd1\x20\x8c\x6f\x90\x7e\x5f\x63\xb6\x68\x86\x4b\xd1\x62\x6c\xa6\xdf\x1a\x88\xe0\xbb\x8d\x04\x27\xec\xb5\x50\x68\xf8\x50\x83\x99\xae\x4b\x7e\x47\x3d\x5a\xc5\x6d\xb4\x41\x24\x77\x18\x46\xa3\xc2\x66\xfa\x0d\x4a\x92\x38\x46\x72\xba\xbc\x47\x7a\xc6\xb6\x89\x5d\x6d\x9c\x2d\xfb\xde\x7d\xae\x0f\x88\xed\xbe\xd1\xce\xed\x55\x20\xd9\x2c\x53\x3f\x6b\x0a\x85\x2e\xe3\xff\xaf\x3f\xdc\x56\x42\xff\xd0\xa7\x57\xe2\xd2\xb2\x36\x78\x39\xab\x78\xf1\xa0\xf3\x33\x60\xaa\x71\xdc\x33\xbb\x22\x5f\x57\xe1\x0c\x7e\xfb\x3d\x34\xb7\xd2\x72\xab\x59\xc7\xea\xa3\x73\x9e\x1b\x7d\x09\xea\xb2\x7c\xba\x4f\x0b\x9c\x69\x25\x57\x76\xeb\x26\xb2\x29\x1a\xb3\xd5\xc2\x61\x5a\x44\x5b\xb9\x87\xd5\x3e\x86\xd5\x18\xf0\x9f\x69\xcb\xa6\xfd\x01\xc9\xea\xd2\xbe\x36\x00\xa1\x19\x86\xff\xa2\xcb\xff\xae\xbb\x6f\xf6\xf2\xef\x95\xab\x3d\x14\xdf\x21\x5c\xfb\x47\x77\x60\x6f\x7f\xea\xb0\x99\xed\xf6\xe1\x4f\xde\xcd\x7f\x7f\x06\x00\x00\xff\xff\x81\xf1\x42\xb0\x9f\x0b\x00\x00")

func templateHttpReadTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateHttpReadTmpl,
		"template/http/read.tmpl",
	)
}

func templateHttpReadTmpl() (*asset, error) {
	bytes, err := templateHttpReadTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/http/read.tmpl", size: 2975, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateHttpUpdateTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x58\x51\x6f\xdb\x38\x12\x7e\xcf\xaf\x98\x13\x8c\x85\x9c\x73\xa8\xbd\x87\x7b\xc9\xa2\x38\x74\x9b\x6c\x6b\x5c\xdb\xf4\x92\x74\xfb\xb0\x58\x60\x69\x71\x24\xb1\xa1\x49\x85\xa4\x9c\xba\x59\xff\xf7\xc3\x90\xb2\x15\xcb\x72\x92\xa6\xd7\x02\x77\x27\x20\x88\x2d\x0f\x87\x33\xdf\xcc\x7c\x33\xe4\xed\x6d\x76\x08\xa5\xf1\xcb\x1a\x8f\x01\xb5\x2f\x0d\x93\x26\x43\xed\xe9\x2f\xcf\x4a\xd4\xec\xa5\xe5\x75\x05\x87\xd9\x6a\x75\x70\x70\x7b\x0b\x02\x0b\xa9\x11\x92\xca\xfb\x3a\x6b\x6a\xc1\x3d\x26\xb0\x5a\x1d\x00\x00\xdc\xde\x1e\xc1\x8d\xf4\x15\xe0\x27\x8f\x5a\xc0\x08\x92\x77\x3c\xbf\xe2\x25\x26\x71\x41\x02\x47\xad\x68\x14\x07\x8f\xf3\x5a\x71\x4f\xfa\x90\x0b\xb4\x09\xb0\x4e\x19\x90\x8e\xf6\x9b\x9c\xd7\xc6\x7a\x48\x21\x58\x7c\x59\x49\x07\xd2\x81\x46\x14\x28\xa0\x30\x16\x9c\x6f\x6a\x29\xe0\x62\xfa\x2b\xd8\x46\x61\x34\x78\xbd\x53\x52\x4a\x5f\x35\x33\x96\x9b\x79\x56\x9a\xa3\xbc\x92\x19\xfd\x2d\xfe\x9e\xec\x13\xa9\x15\x5f\x96\xd6\x34\x5a\x64\x0b\xae\xa4\xe0\xde\xd8\x6c\xf1\xb7\x1f\xe3\x82\xf1\xc1\xda\x44\xcb\x75\x89\x30\xd2\x70\xfc\x0c\x46\xec\xad\x11\xe8\xe0\xce\xc6\x59\x06\xef\xf8\x52\x19\x2e\xc0\x14\xc0\x69\xc5\x48\xb3\xb7\x7c\x8e\xb0\x5a\x41\x84\x0f\x2c\x5e\x37\xe8\x3c\xdb\xac\xa2\x70\x6c\x8b\xbe\x0f\x92\xe7\x51\x10\x9c\xb7\x4d\xee\xe1\x76\xb3\x20\x1a\x43\xc0\x9c\x9d\x9c\x1d\xc3\x2b\xbe\x90\xba\x04\xae\x14\xd4\x46\x6a\x8f\xd6\x41\x85\x16\xc1\x21\xce\x1d\x68\xe3\xc1\xca\xb2\xf2\x0c\xde\xf0\xe5\x0c\xc1\x13\x9c\x39\xd7\x30\x43\x10\x46\x23\x48\x0d\x5c\x1b\x5f\xa1\x85\x1b\xbe\x04\xc6\xd8\x36\x9e\xdb\xce\x17\xc1\x79\xcd\x7e\x91\xa8\x84\xdb\x0a\xf1\x1d\x61\x59\x84\x8d\x47\x05\x9b\xce\xe7\x8d\xe7\x33\x85\x83\xa2\xad\xf8\xa8\x60\x17\xc1\xcd\xa0\x95\xd0\x3a\x8c\x6f\x2f\x97\x35\xd2\x4f\xe4\xe1\x6a\x05\x7f\x7c\x74\x46\x1f\x27\xb4\x81\x16\xf8\x09\x52\x57\x2b\xe9\x21\xf5\xbc\x7c\x6d\xcc\x55\x53\x77\x9a\x2e\x79\x09\x09\x89\x27\x63\x48\x26\xc9\x18\x7e\x84\xd5\x8a\x56\x86\x9c\x1d\x15\xec\xb9\xd6\xc6\x73\x2f\x8d\x76\xec\x54\x5d\xc1\x9f\x80\xea\xaa\x7b\x49\xdb\xb5\xc9\x80\x61\x47\xf6\x6b\xfc\x26\x8d\xbe\xe4\xa5\x83\xa4\xab\x87\x64\x93\xbe\x7f\x0c\x80\x71\x74\x37\xb5\xef\x38\x4d\x6f\xfb\x98\x74\x38\x63\x8b\xf3\xa9\x28\x71\x2f\xcc\x23\xec\xe1\x16\xa1\x1f\x21\x7b\xaf\xe5\x75\x43\xc9\x44\x40\xa2\x72\xf4\xf1\xb7\xdf\x37\x76\xc6\xb5\x01\xdd\xe9\xc9\x13\x50\xc6\xc7\xa1\x8c\xdf\x03\xe5\x6d\xee\xa0\x67\x75\x70\xb7\x26\x63\x39\xb5\xf5\xe7\x80\x43\x29\x17\xa8\x7b\xd5\xc9\xb5\x00\xc7\x17\x28\xc0\x57\x08\x79\x45\x41\x70\xe0\x4d\xf8\x2a\xb8\xe7\x33\xee\xb0\x2b\xda\xa2\xd1\x39\xa4\xd5\xb6\x92\x57\x5c\x0b\x85\x76\xdc\xee\x98\xde\x00\xd1\x20\x3b\x47\x57\x1b\xed\xf0\x83\x95\x1e\xed\x04\x2c\x1c\xb6\xef\x43\x81\x8f\x7b\xa5\xad\x28\xf0\x15\x53\xa6\x64\x1f\xa4\xaf\xd2\xcf\xbc\x6e\x63\x93\x26\x73\xf4\x95\x11\xc9\x04\x92\xb8\x45\x32\x1e\x6f\xa7\x45\x96\xc1\xf4\x84\xc8\xf2\xfd\xf9\x6b\xa8\xb9\xe5\x73\xf4\x68\xd9\x41\x3f\x23\x29\x49\x34\x85\x7e\xea\xa6\xda\xf7\x93\x93\x1e\x29\x26\x80\xd6\x92\x31\xce\xdb\xdc\xe8\x05\x7b\xee\x8d\x4c\xf3\x4a\xb2\xf7\xe7\xaf\xdf\x91\xee\xd4\x4e\x20\x91\x62\xc7\x0a\x5a\x5e\x84\xd5\x7f\x79\x06\x5a\xaa\x9e\x87\x1b\x4f\xd9\xa9\xb5\xc6\xa6\x09\xd2\x3f\x28\xd1\x7b\xca\x40\x29\xa0\xb0\x66\x0e\x8d\x55\x9d\x0b\xc9\x04\xee\x02\x21\x09\x84\x61\x53\xa2\x60\xd4\x8c\xd6\x0e\xd8\x46\x8f\x45\x2d\xd0\xb2\x9f\xb9\x68\xc3\x90\xde\x4c\x20\xea\x80\x79\xe3\x3c\xb1\x23\xd7\x40\x7c\x5a\xa2\x85\xd2\x22\xf7\x68\xe1\x33\x5a\x93\xec\xd3\xe8\x1b\xab\x77\x7e\xda\x2d\xfb\x58\x8c\x03\x80\x13\xd6\x83\x3e\xed\x63\x8e\xad\xf7\x59\x06\x2f\xd1\x87\x7c\xad\x8d\xf3\x21\x69\xb7\x03\xbf\xe0\x16\xc4\x3d\xad\xe6\x60\x20\x80\xc7\xcf\x80\xca\x9b\xbd\xc5\x9b\x13\xcc\x8d\x40\x9b\x5a\xf6\xb3\x11\xcb\x31\x8b\xdf\xd3\x1f\xc4\xf8\xa7\xfb\x63\xdd\x8b\xb3\xa0\x75\x14\xe8\xc0\x1b\x0f\x87\x6b\x7f\xa8\x74\xa0\x8e\xa0\x87\x92\x54\xea\x72\x20\x36\x03\x71\x59\xf5\x81\x6b\x49\x07\x37\xd5\xce\xf6\x60\x51\xb1\xcd\x80\xd0\x32\x60\xfa\xa0\xfb\x71\xf5\x04\xcc\x15\x69\x40\x6b\x59\x7a\xd8\x69\x99\x46\x27\x3a\xda\x0b\x50\x8c\x7f\x22\xf1\x47\x95\x4d\xab\x8a\x00\x6d\x67\x8b\xe0\xc1\x23\x80\xbd\x03\xee\x94\xe6\x06\xcd\xd5\x05\xda\x05\xda\xb8\x28\xa2\xac\xa5\xfa\x8a\x7c\x8f\xf6\x4e\x75\x61\xd2\x64\xb1\x71\x11\x0a\x2e\x15\x8a\xaf\x88\x3d\x49\x3f\x29\xd2\x17\x7c\xb1\x2f\xca\xb3\x18\xe0\x5c\x49\xd4\x9e\x6d\x95\x09\x8b\x75\x72\xa6\x71\x7a\x92\xb6\x3d\x36\xd0\xe7\x2b\xee\x5e\x1a\x6a\x9f\x6d\x4f\xd5\xeb\x76\x0a\xab\x55\x2a\xc5\xb8\xab\x77\x29\x36\x7d\x6a\xdc\x37\x2a\x0e\x72\x37\x15\xf7\xc0\x67\xa6\xf1\xe0\x94\xcc\x11\x8a\x38\x62\x79\x7a\x5f\x91\xdd\x79\xe3\xbc\x99\xc3\x9c\x5b\x57\x71\xa5\xd0\xba\x7f\x7c\xb7\x29\x4d\x16\x20\xd8\xe0\xa8\x76\x2f\xc7\x07\x60\xd9\x05\xfa\xa1\xa5\xe9\xe1\xb0\xca\x71\x9c\x6f\xbd\x11\x06\x8e\x9e\x86\xcb\xee\x08\x3b\x9c\x12\x70\xef\x2c\xf6\xb5\x23\xda\x06\xb4\xfe\x9c\x76\x3f\x68\xfd\x29\x6e\x5f\x4c\x22\xb8\x51\xff\x9b\x26\x4e\x54\x17\xe8\xef\x20\xdb\xdf\x77\xb8\x94\xd7\x59\xfa\x25\xfb\xbc\x50\xc8\x2d\xed\x34\xee\xfd\xf0\x5c\x84\xd0\x0e\xef\xcf\x18\xdb\x67\xc2\xe0\x98\x0c\x83\x1d\x74\x20\x2a\x54\xda\xde\xd8\x70\x9a\xd9\x1d\xd7\xe8\xc1\xcd\x3c\x33\x63\xc4\x02\xa9\x65\x2f\x8c\xf6\xf8\xc9\xa7\x3d\xe2\x79\x68\x76\x71\x37\xd2\xe7\x55\x64\x72\x3a\xbd\xf5\x27\xb8\xf5\x93\x73\x87\x70\x48\x64\xf2\xd6\xf8\x5f\xe8\x78\x19\xa8\xee\x78\x2f\xc6\x6b\xa6\xbc\xc3\x3d\x7f\xc2\x15\xce\xf8\x8c\x52\x86\x4a\xb5\x20\x2d\x2d\x6f\x4e\xb5\x5f\x8b\x4e\x4f\xd6\x4c\x95\x4c\x40\x8a\x47\xce\x40\xd0\x11\xec\xda\xc0\x75\x6b\x7d\xc8\x82\x61\x85\x5b\x0e\x5f\x48\x5d\x36\x8a\xdb\x07\x7d\x6e\xbb\x99\x68\x6a\x25\x73\x6a\xc2\xa8\xbd\x5d\x86\x73\xfe\xa0\x19\xdf\xc0\xfd\xdd\xd9\xa2\xb3\x66\x18\x8a\x68\x63\x38\xe5\x48\x01\x09\xfc\x75\x33\x25\x4f\xbd\xe1\x29\xb2\x8d\x65\xbd\x02\x1c\xb6\x45\x60\xc1\x1b\xe5\x1f\x01\x53\x6c\xfa\x2e\x1e\xfa\xbf\x17\x3e\x5f\x38\x1e\xec\xd6\xf0\x63\x5a\xf3\x39\x86\xab\x93\x80\xec\x76\xe9\x5e\xdf\xd3\x99\xff\xd5\xa0\x5d\xa6\x63\xf6\xa1\x42\x8b\xe9\x16\x20\xca\xdc\x20\x51\x14\x9b\x9e\x3c\x3a\x20\xdd\xad\x16\x91\xfa\xa5\x79\x4d\x26\x8d\x34\x24\x79\x38\x01\x24\x43\x04\x95\x65\x70\xca\xe9\x94\x10\xed\x0f\xdd\x20\x34\x27\x6e\xe3\x7d\x8f\xb4\x28\xc0\x68\x88\x3a\xc0\xd4\x68\x03\x55\xb2\xa1\x56\xc4\x5e\x18\x81\x90\x5c\xef\x6c\xb5\x87\x22\x5b\x5e\x7b\x06\xd7\xec\x4c\xab\xe5\xff\x2c\xad\xed\x8f\xe0\xff\x33\xdb\xfd\x07\x51\xf9\xaf\x23\xc1\x02\x7d\x5e\xed\xa5\xc1\x78\x89\x20\x66\xdf\x0f\xc0\x6f\xcf\x92\x1f\xbb\x3b\x99\x0a\xad\x2c\x0a\xf6\x26\xce\xbc\xe9\x0f\xeb\x17\x67\x75\xb8\x6c\xdb\x2d\xdd\xa9\xce\x55\x23\xf0\x74\x5e\xfb\xe5\x25\x2f\x8f\xc1\xdb\x06\x27\x3b\x62\x2f\xad\x69\x6a\x77\x0c\xbf\xfd\x1e\x4f\xd4\xb7\x1d\x25\x8e\xfa\xb7\x79\x17\x79\x85\x73\xce\x5e\x04\x5a\x8b\x0b\x89\x6e\x89\xc0\xba\x83\xcf\x70\x75\x25\xdd\x1c\x07\xab\x6d\x2b\x56\x13\xc0\x2f\x63\xae\x4d\x5e\x38\xb4\x92\x2b\xf9\x39\x9e\x32\x43\x96\x7c\xc3\xf0\x3f\x2d\xec\x0f\x06\xf9\x7e\xba\x8c\x9b\xe2\xd3\xd9\x72\xdb\xa2\xd6\x87\xb3\x7f\xb6\x26\x7f\x1c\x1f\x6c\x5b\xd5\xdd\xae\x6e\x3e\xfd\x3b\x00\x00\xff\xff\x7e\xb8\x18\xd3\x48\x1a\x00\x00")

func templateHttpUpdateTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateHttpUpdateTmpl,
		"template/http/update.tmpl",
	)
}

func templateHttpUpdateTmpl() (*asset, error) {
	bytes, err := templateHttpUpdateTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/http/update.tmpl", size: 6728, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"template/http/create.tmpl":  templateHttpCreateTmpl,
	"template/http/handler.tmpl": templateHttpHandlerTmpl,
	"template/http/list.tmpl":    templateHttpListTmpl,
	"template/http/read.tmpl":    templateHttpReadTmpl,
	"template/http/update.tmpl":  templateHttpUpdateTmpl,
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
	"template": &bintree{nil, map[string]*bintree{
		"http": &bintree{nil, map[string]*bintree{
			"create.tmpl":  &bintree{templateHttpCreateTmpl, map[string]*bintree{}},
			"handler.tmpl": &bintree{templateHttpHandlerTmpl, map[string]*bintree{}},
			"list.tmpl":    &bintree{templateHttpListTmpl, map[string]*bintree{}},
			"read.tmpl":    &bintree{templateHttpReadTmpl, map[string]*bintree{}},
			"update.tmpl":  &bintree{templateHttpUpdateTmpl, map[string]*bintree{}},
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
