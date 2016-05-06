// Code generated by go-bindata.
// sources:
// data/awskms-key.template
// data/kmseditkeypolicy.txt
// data/kmsgrantcreate.txt
// data/kmsgrantslist.txt
// data/kmsgrantsretire.txt
// data/kmsinit.txt
// data/usage.txt
// DO NOT EDIT!

package main

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

var _awskmsKeyTemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5a\xdf\x6f\xdb\xbe\x11\x7f\xcf\x5f\x41\xf8\x25\x5b\xe1\x78\x6d\x5f\x86\xf9\x4d\x4d\xda\xc2\x6b\x93\x05\x71\xda\x62\x18\x82\x81\x91\xce\x31\x11\x49\x54\x49\x2a\x99\x57\xe4\x7f\xff\xde\x91\x92\x4c\x49\x94\x9c\xa4\x69\xfb\x6d\xd1\xa0\x48\x6a\xf1\xc8\xfb\xf5\xb9\x1f\x3c\xf9\xcb\x1e\x63\x93\xe8\xd3\xf2\x1c\xb2\x22\xe5\x06\xde\x48\x95\x71\xf3\x11\x94\x16\x32\x9f\xcc\xd9\xe4\xe5\xf3\x17\xcf\x0f\x9e\xff\x03\xff\x4d\xa6\x44\x7b\x04\x3a\x56\xa2\x30\xd5\xf2\x19\x68\x59\xaa\x18\x34\x5b\x49\xc5\x4a\x0d\xec\x56\x98\x35\x7b\x25\x74\x5c\x0a\x33\x73\x7b\x8e\xc1\xf0\x84\x1b\x8e\x1b\x88\x1f\x3e\xf9\xb4\xe6\x26\x52\x70\xbe\x86\x33\x99\x82\x46\xae\xb8\xf6\x1f\xbb\x86\xab\xaf\xf3\x58\x6d\x0a\xf3\x09\x4f\x7a\x07\x1b\xa2\x60\x3c\x4f\xd8\x11\xf4\x1e\xdf\x8a\x34\x65\x96\x89\xdd\xc8\xd3\x54\xde\x32\xb3\x06\x56\x28\x91\xc7\xa2\xe0\xa9\x66\x22\xc7\x27\x12\x25\x53\xc4\x8a\x19\xc9\x32\x7e\x4d\x14\x3c\x36\x22\xe6\xa9\x95\x5a\xae\xec\xb6\x77\xc7\x4b\x76\x0d\x1b\xcb\xce\x1e\xe6\xf4\xda\x72\xd0\x50\x70\xc5\x49\x7b\xda\xa2\x40\x17\x32\xd7\xe2\x52\xa4\xc2\x08\xd0\x33\x86\x9a\x30\xf8\x1f\x47\x63\xc2\x94\x71\x46\x5a\xa3\x00\x57\xa0\xed\x16\x0d\xea\x06\x14\xcb\xc4\xd5\xda\x78\x87\xca\x3c\xdd\xb0\x4b\x70\x1c\x21\x21\x11\xc1\x99\x80\xdd\xae\x41\x01\xd7\x78\x54\xa1\x24\x9a\x59\xe3\x61\x2c\x97\x09\x30\xa1\x7d\xfa\xc4\xd9\x66\xe6\x9d\x7a\x22\x0d\xa0\x52\xdc\x30\x61\x88\xfa\x0a\x72\x50\xb8\x85\x58\x19\x83\x62\xe0\x36\x52\x9d\x74\x7e\xab\x78\x6e\x9c\xae\xb1\xcc\x0d\x5a\x2a\x25\x46\x3c\x26\x96\x44\xe8\x09\x5b\x80\x33\x80\x66\x32\xb7\x46\x43\x83\xe9\x29\xbb\x2c\x0d\xd3\x32\x03\x23\x32\xb2\xf2\x1a\xf0\x68\x73\x2b\x2b\xab\x73\x65\xd5\x63\x1b\x59\xb2\x1c\x20\x99\x22\xad\x77\x28\x92\x6f\x2c\x0d\xa9\x6b\xc5\x40\x42\x2b\xcb\x0d\xe4\x02\x8d\x01\xb3\x89\xa5\xbd\xc0\xdf\x77\x16\x54\xa7\xe8\x07\xe4\x86\x48\xdd\xc2\x2a\x22\x83\x2c\xa2\xe3\x53\x99\x8a\x18\xfd\x71\x2e\x0f\x9d\x36\x88\x97\xc8\xea\xd2\xd0\xf6\xa1\xbc\x58\xb1\xfd\x15\x02\x06\xf6\xa7\x56\x2d\xeb\x95\xa2\x3a\xc9\x19\x92\x24\x84\x1b\x9e\x96\x18\x29\x95\xd9\x51\x80\x4c\xe4\xe0\x50\xe3\xcc\x45\x54\x0e\x71\x09\xac\x70\x2d\x71\x10\x44\x4b\x23\x8d\x15\x6d\x33\xb5\x08\x43\x49\xb7\x0c\x2c\x94\x73\x69\x08\x08\x31\x81\x2a\x41\x53\x24\xec\x2f\x78\x52\x79\xb5\x66\xfb\xce\x45\xfb\xc4\x95\xf6\x39\xb3\xda\x4d\xb8\x61\x2d\x73\x89\xd4\x7f\x9d\x31\xd2\xc2\xa8\xb2\x56\xa2\xc5\x22\xe6\x79\xed\x5e\xcf\xb5\x1e\xee\xf5\x6c\xeb\x92\xf3\x4d\x01\x64\x96\xa5\xc1\x50\xba\xda\x3e\x8f\x1c\xe8\x3e\xa2\x15\x40\x7b\x51\x4b\x5e\x44\xbe\x0d\x21\x7e\xb6\xd6\x9c\x54\x9f\x2f\xec\xdf\xbb\x69\xe5\xa9\x04\xad\x26\xb4\x41\x28\x49\x75\xda\x04\xeb\x88\x7b\xde\x94\xa8\x6a\x74\x76\xa2\x29\xf0\x10\xb7\xca\xf3\x89\x17\x08\x88\x4e\x84\x4f\x66\xc1\xc6\xb7\x4c\x28\xfa\x56\x65\x1e\x7b\xc0\xc5\x90\x40\x95\x03\x1a\x1f\xca\x2c\xe3\x47\x90\x8a\x4c\xa0\x9b\xdf\xe3\x01\x93\x96\xec\x1f\x90\xf9\xb7\x12\x39\xe1\x9b\x03\x23\x0f\xf0\x0f\xf3\x02\xcd\xe6\xd4\xaf\x90\x18\x81\xd7\x16\xed\xcb\x2e\x2f\x77\x34\xc1\x44\x8d\x60\x6e\x1e\x31\x6d\xc9\x49\x78\x6e\x0c\x8f\xd7\x35\x8e\x62\x4c\x56\x14\x1a\x56\xcc\x96\x08\x87\x76\x65\x29\x28\x33\xda\x9c\x7f\x1f\x29\x1e\x8d\xb5\x21\x3d\x96\x60\x48\x54\x17\x21\x4c\xb8\xa4\x7f\xff\x72\xa3\x31\x16\xd3\xc4\x46\xa8\x53\x74\xc6\xce\xc9\x2b\x0a\x3e\x97\x02\x8b\x01\x3b\x8c\x4e\xa3\x57\xd1\xe2\xfd\xe2\xfc\xdf\xff\xc5\xd0\x9b\xf9\x06\x5d\xf1\x32\x35\x24\x84\x15\xdd\x19\xa7\xc9\x66\x98\xa7\x12\x61\x9d\xbd\xcd\x66\x88\xb3\x43\x5e\x70\x5b\x5c\x36\x0b\x9e\x35\x34\xbe\xe9\xde\xe4\xf3\xf9\xeb\xcf\xa5\x03\xe2\xd6\x3e\x5f\x9a\xff\x31\x2a\xd2\x2b\x8b\x92\x9e\x0b\x1a\xa2\xbb\x69\xc7\xb2\xc1\xa0\x7d\x9d\xf3\xcb\x14\x50\x92\x3a\xbf\x7e\xad\x44\xf7\xc8\xd7\xf7\x15\xb1\xb1\x64\xd3\x91\x6c\x0d\x59\x75\x23\x78\x68\x08\x74\xd8\xfd\xcc\xe7\x98\x00\xf1\x17\x12\x34\x0e\x3b\x55\x14\x80\x54\xd5\xbd\x4d\x3d\x44\x85\xb4\xea\x44\x5b\x50\x03\x67\x4a\x8b\x2b\xc3\x6b\x74\x76\x31\x8d\xcb\xae\x5c\x74\x19\x2d\x12\x22\xaf\xd5\xca\xf4\x96\x70\xea\x93\xb5\xdb\xb8\x97\x07\xd8\xc9\xbd\xf8\x7b\x9b\x64\x89\xdc\x21\x83\xdc\xb4\x7c\xd5\xf5\x57\xe3\xd7\xc5\xaa\x47\xe7\xa9\x13\x44\xc6\xb4\x47\xdd\x3d\xd9\x09\x22\xac\x4e\xee\xa0\x76\xd9\xc2\x80\xbd\xa2\xe2\xe7\xd5\x2c\xd7\x72\x34\x5d\x1a\x56\x43\x69\x97\x65\x49\x64\x94\xfa\x5d\xa6\x9f\x05\xd8\x93\xb8\xab\x15\xc4\xa6\x41\x60\x98\xa8\x49\xf1\x1d\xeb\x7b\x24\x08\x9d\xa0\x3d\x86\xf5\xac\x36\x92\x2d\xff\x29\x45\x3e\xb2\xdb\xd2\xcd\x83\xa2\xd5\x3f\x63\x5b\xa9\x1b\x56\xf9\x9c\xdf\xea\xb9\xe0\xd9\xf8\x39\x63\xa2\x56\x67\xd5\x01\x4b\xb1\x12\x39\x3b\x23\x08\x47\x77\xdd\x8d\xb3\x9c\x90\xcf\xc6\x4e\xb8\x18\x5c\x1b\x5a\xb9\x0b\x3e\x0f\x51\x07\x65\x9b\x44\x71\x1d\x89\xd7\x99\x9e\x3f\x0b\xc3\xa2\xce\x2f\x44\xf6\xac\x2f\x7f\xe0\xe4\x31\xbc\x47\x7d\xf8\x22\xbc\x15\xe0\x65\x2c\x76\x4d\xa5\x0d\x83\xcd\x6f\x20\xff\x06\xf2\x23\x80\x1c\x36\xac\x85\xf7\x5b\x30\xe1\xb2\xd1\xa1\xa3\x6e\xb2\x26\xa4\x62\x38\x42\x7a\x5a\x7a\x47\x06\xc8\x2e\x1e\x15\x50\x9d\x27\x6d\x2b\x74\xf4\xef\x15\xad\x56\x98\x55\xf5\x83\xae\x97\x74\x15\x6b\xdd\x41\xfa\x8a\xed\x0c\xae\x1d\x81\x55\x05\x55\x30\xfa\x6b\x14\x0e\xdc\x82\x76\x1a\xa1\xe7\xf6\x31\x97\x5b\xdf\xb8\xe6\x2f\x94\xd3\xec\xb2\x6b\x58\x2e\x87\x09\x5c\x61\x1e\x5c\x26\x94\x0c\x2e\x22\x2e\x06\xd7\x3e\x14\xc9\x98\x5c\x67\x70\x23\xaf\x47\xc4\x16\x7a\x54\x2c\x04\xf9\x88\xca\x29\x8c\x70\x5e\xc6\x6b\x48\x4a\xdb\xa3\x59\xca\x70\x2f\xe3\x6c\xcb\xf3\x18\x52\x9f\xb0\x8b\xd9\x9e\xbb\x46\x40\xff\x20\x48\x7b\xb3\xab\x6b\xe8\xc7\xf1\x77\xc0\x70\xe7\x36\xfc\x0d\xa0\x5b\xdd\xcf\x86\xfd\x38\xba\x7c\x06\xd5\xfe\x11\x90\xd0\x60\xcc\xc0\x11\x37\x1c\x9d\xb8\x33\x46\xe8\x9a\xf0\xdd\x1c\xec\xee\xd7\xd4\xa1\x93\x9f\x0b\xea\xe8\xb5\xa1\x4f\xaa\xb9\xe7\xfc\x92\x4e\x77\xf9\xca\x0e\xbd\xc6\x72\x8e\x9b\x8a\x8d\x67\x0f\x77\xc8\x03\x3d\xd6\x5b\x0e\xdd\x75\xbd\xe5\x57\x52\x0e\xf5\x57\x0e\x64\x24\xc4\x82\x26\xdd\x68\x5d\x8f\x15\xdd\xf9\x76\x9b\xef\x41\xf8\x19\xbb\xa8\x0d\xcf\x14\x1e\xd3\xb7\x56\x53\x67\x1a\x05\x21\x38\x79\xbe\x61\x34\x1c\x85\x7a\x7a\x0d\x09\x2b\xf3\x84\x26\xcd\xbd\xd1\x55\xeb\xcc\x3f\x45\x17\x8b\xa5\x22\x32\xfd\x6b\x70\x87\xb2\x3f\x14\x1a\x6d\x48\x27\x91\xea\x15\x83\xed\xcf\x8f\x6c\xff\xea\xac\xf9\x64\x6d\xda\x3d\xf1\xe3\xf7\xdf\x27\xd2\x4e\xf6\x9e\xba\xe7\xfb\x7e\xf8\xaf\x70\x4e\xf8\xff\xd5\x80\xde\x1f\x8a\xfe\xa4\x40\x1f\xee\x1e\x1a\x92\x4e\x03\xf0\x53\xc6\xc4\x5e\x68\xa5\x7e\xea\xfe\xd6\x63\xdc\x40\x0e\x1b\x1a\x8b\x2e\xa2\xe3\xf9\xbc\xe5\xfc\x56\x1d\xbc\x47\x38\xf5\x5f\xfe\x56\xa7\x64\xd5\xd4\x31\x20\x0e\xb5\x39\x99\x30\xba\xa9\x2f\xee\x6d\xab\xad\x2d\x3a\x5c\x5c\xea\x17\xb7\xf4\xa6\x44\x96\xc6\xd6\x22\xec\x90\x8c\x12\xf5\x20\x70\xaf\xe3\x93\xc1\x01\x6f\xa4\x75\x99\xd9\x01\xb9\xbb\xc9\x1e\xc9\xb8\xac\x64\x6d\x0d\x61\x9f\x76\xba\xfa\xb5\x3d\xdb\x12\xd4\x8d\xb0\x48\x0c\xc5\xc2\x04\xe2\x97\x33\x9e\xf1\xff\xcb\x9c\xdf\xea\x59\x2c\xb3\x3e\xc0\xba\x61\xf7\xb0\x9e\x4d\x1b\x3d\xdf\x5a\xae\xd7\x6d\x3d\x1a\xac\x81\x3c\xf4\x23\xc1\x1a\x78\x57\x54\x83\xb5\x2e\x06\x2d\xb0\xfe\x86\xa8\xc7\xea\x57\x83\xe8\x5e\xfd\xce\xe9\x5f\xa5\x29\x4a\xe3\xbd\x71\x42\x6c\x2c\x8e\x46\xde\x0c\xd3\xf8\x09\x09\x1a\x47\xbb\x9c\xdf\xf2\x71\x55\x14\xbc\x77\x57\xc1\xf8\xa0\x37\x65\x2a\xdf\xc1\x2a\x3a\x3b\x19\xe5\x35\x30\xc5\x9d\xb4\x5c\xdc\xb6\x68\x33\x94\xa5\x02\x3a\xd9\xd1\x91\xf9\x05\xee\x0c\xae\x7a\x13\x92\x8e\x1f\x7b\xd3\xe2\xd1\x03\x07\xa6\xb8\xbd\x33\x31\xf6\xfe\x76\xcf\x73\xfb\x46\xaf\xce\x0c\xa2\xe2\xe2\x11\x89\xab\xe3\xa4\x13\x9e\x35\x73\x1c\xfa\x66\x47\xf5\x3d\x01\xf7\x0d\xa4\x76\x62\xc1\xb4\x81\xe9\x64\xa4\xe3\x7c\x70\xd2\x1b\x06\x5f\x40\x87\x47\x74\x14\x03\xba\xf6\xf4\x6c\x57\xfb\x46\xcf\xdd\x97\xc9\x27\xd4\x38\xa0\x49\x3f\xe6\xf7\xee\xf6\xfe\x08\x00\x00\xff\xff\xf2\xcb\xef\xa1\x38\x27\x00\x00")

func awskmsKeyTemplateBytes() ([]byte, error) {
	return bindataRead(
		_awskmsKeyTemplate,
		"awskms-key.template",
	)
}

func awskmsKeyTemplate() (*asset, error) {
	bytes, err := awskmsKeyTemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "awskms-key.template", size: 10040, mode: os.FileMode(436), modTime: time.Unix(1462467644, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _kmseditkeypolicyTxt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x91\x4b\x0e\xdc\x20\x0c\x86\xf7\x9c\xe2\x3f\xc0\x24\x97\xa8\xba\x68\x47\x7d\x48\x73\x02\x87\x31\x89\x35\x04\x22\x20\x13\xe5\xf6\x35\x90\x56\x6d\xd5\x2e\x11\xff\xe3\xb3\xfd\xf1\x29\x05\x65\x61\xdc\xbf\x3c\x70\xe7\x13\xdf\xa3\x17\x7b\xc2\xc5\x04\x82\xa7\x89\x3d\xc8\xa6\x98\x33\x12\xcf\x12\x43\x1e\x8d\x61\x35\x0d\x2f\x3e\x87\xad\x8b\xe3\xc6\x21\x63\xdf\x70\xc6\x3d\xc1\xd1\x3b\x26\x29\x8c\x2a\xd3\x98\x43\xca\xd2\x2a\x3e\x3f\xbe\x7d\xd5\x94\x2d\x71\xe6\x50\xa8\x68\x9a\x89\xee\x5f\xed\x65\xa1\x02\xc9\xb0\x7b\x4a\x2a\xf5\x27\x32\x17\xc4\xf0\x4b\xab\xe5\x19\x94\x73\xb4\x42\x85\x9f\xad\xc3\x10\xf2\xc6\x56\x9c\xd8\x0e\x3e\xe2\xc3\x42\x61\xe6\x5c\xb9\xb0\xd2\x8b\x51\xa2\x46\x68\xb0\x13\xcf\xa0\xc4\x38\x94\xb4\x70\xc0\x44\xf6\xd5\x7f\xb9\x85\x1b\x09\x20\xef\x71\xf1\xd1\xb6\x29\x18\x4d\xea\xfa\xff\x1a\x0e\x51\xc3\x9b\x93\xb8\x6b\x82\xea\x6c\x7f\xa2\x0c\x7f\x06\x5e\x29\x95\xc1\xd4\x77\xa6\x95\x31\xb1\xae\x9d\xab\x2c\x1e\x12\xe6\x86\xad\x4c\xfc\xf3\x46\xbd\x67\xc4\xa7\x96\x71\xe2\x29\xce\x71\xba\x35\x5d\x2d\x37\x81\x75\x17\xea\xd0\x0d\x47\xff\xe6\x66\xea\x22\x0e\x56\x19\x56\x0a\xbb\xa6\xf7\x88\xee\xca\x4b\x35\x44\xc5\x3e\xda\xd1\x3a\xa3\xf9\x9d\xb1\x1d\x90\xae\x76\xb8\x14\x57\x3d\x05\xff\x35\xc9\xed\x5a\xf2\x89\x3d\xb7\x62\x33\x0c\x3a\x8d\xe5\xa1\x0b\xe0\x3c\xcd\xa3\xf9\x11\x00\x00\xff\xff\x9c\xac\x9a\xce\x70\x02\x00\x00")

func kmseditkeypolicyTxtBytes() ([]byte, error) {
	return bindataRead(
		_kmseditkeypolicyTxt,
		"kmseditkeypolicy.txt",
	)
}

func kmseditkeypolicyTxt() (*asset, error) {
	bytes, err := kmseditkeypolicyTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "kmseditkeypolicy.txt", size: 624, mode: os.FileMode(436), modTime: time.Unix(1462400352, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _kmsgrantcreateTxt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x56\x5d\xcb\xdb\x46\x13\xbd\xce\xfe\x8a\xc1\xef\x0b\x49\xc0\x96\xd3\x94\xde\xe4\x2e\x49\x3f\x28\x6d\x5a\x48\x52\x4a\xa1\x50\xd6\xd2\xc8\x5e\xb2\xda\x75\x77\x57\x76\xdd\x5f\xdf\x33\xb3\xfa\x70\x9e\x34\x85\x84\xc7\x92\x46\xb3\x67\xce\x9c\x39\xa3\xd7\x89\x6d\x61\xb2\xf4\xc3\x9b\x77\xf4\x5d\xb2\xa1\x90\xf5\x3e\x5e\x5d\x38\x52\x3e\x73\xeb\x7a\xd7\x52\x3c\x73\xb2\xc5\xc5\x90\x29\x06\xc4\x2e\x0f\x32\xb7\x89\x4b\x63\x8c\xbe\x99\x89\x83\x3d\x78\xa6\x5b\x1c\xa9\x44\xea\xd8\xf3\x51\xb2\x97\x13\xd3\x98\x99\x62\x3f\x1d\xf4\x81\x6f\xd4\xc7\x74\x9f\xea\x3c\xa6\x73\xcc\x6c\x76\x3b\x84\x8e\xc0\x70\xa3\xaf\x91\xfd\x76\x2e\x84\x5b\x47\x77\xe1\x40\x2d\xa7\x62\x1d\xfe\x02\x49\x49\xf8\x55\x72\x43\xd3\xd1\x36\x31\x0d\x31\x17\x3c\x1c\x86\x18\xfc\xcd\xe0\xc4\x4e\x60\x68\x3d\xcb\xbb\xe7\xe4\x42\xeb\xce\xd6\xe7\x0a\xb1\x1e\x51\x0b\x91\x02\xf4\x1a\x2f\x8e\xa1\x63\x00\x34\x0b\x40\x60\x6e\xe8\x3d\x2a\x59\x32\x50\x6b\x03\x1d\x40\x5e\xa0\xef\x5f\xbe\xa1\x5f\x32\xa7\x2d\xd9\x9c\xc7\x01\xef\xa7\xe8\x79\x4b\x52\x63\x7d\xfa\x16\xd7\x20\xea\x65\x96\x1b\xfc\x97\x1d\xce\xf2\x5c\x2a\x71\x7a\x50\xa0\x6f\x5e\x3f\x27\x87\xc2\x6c\x68\x99\xd2\x18\x82\xf4\xc0\xd2\x95\x0f\x48\x7c\x41\x0c\xd0\xd7\x64\x46\x92\x81\x55\x8b\x66\x8d\x20\x37\x14\xd7\x82\xe7\x4c\xae\x64\xf6\xbd\x16\x4d\x9d\x2d\xf6\x60\xc1\xfa\xd5\x95\x13\xae\xcf\x40\x76\x8d\xa9\xd3\x22\xcc\x9a\x35\x30\x77\x59\x5b\xb4\xbc\x31\x87\xd2\x13\xb9\x5d\xb9\x79\xba\xa5\xeb\xc9\xb5\x27\x72\x99\x72\x89\x09\x25\x0a\x1e\x6a\x6e\x76\xf0\xa6\x77\x00\xd4\xaa\x96\x3a\x3a\xdc\x68\x7d\x4f\x68\x8e\xbe\xa1\xcd\x87\x21\xd3\x51\x05\x56\xe3\x36\x92\xe9\xe3\x1e\x49\x45\x2b\xe1\x49\x8b\x5c\x7b\xf4\x89\xee\xb6\x8a\xc0\x7b\xd1\x95\x1c\x98\xf8\xa8\x1a\x55\x62\x70\xc3\xd4\x30\x39\xe7\x41\x5f\x85\x03\xb9\x5b\xdb\xa0\x92\xb1\xa1\x03\x53\x48\x56\xd1\x55\xa8\x59\x55\x2a\xa9\x36\x33\x39\x7f\xcc\xe4\x6c\x68\x4e\xff\x79\x10\x4e\x0f\x6f\xc7\x94\xd0\x23\x28\x7a\x81\x61\x5c\x78\x61\xcc\xa3\xff\xd3\xc1\xe5\x76\x44\xd8\xc2\x4e\x9e\x01\xec\x76\x7a\xcd\xbc\x5b\x05\x27\x94\xec\x97\xd6\x65\xfa\xdd\x3c\x7a\xb4\xeb\xd1\x0f\x9b\x4f\xcd\x6d\xf0\xf4\x09\x4a\x63\x7e\x16\x39\x4d\xad\x81\x3c\xc3\xed\xdf\x85\x86\x2a\x04\xfe\x5d\x72\xe5\x5f\x24\x1e\xd0\x1a\x84\x19\x79\xbe\x99\x01\x1f\x51\xf9\x7f\x1f\xbd\x59\x7d\x61\x56\x86\x1e\x02\x46\xc0\x76\x80\xf4\x4e\xf6\x22\xd6\x33\x06\xf7\xe7\xc8\x14\xec\x00\xda\x39\x88\xd7\x20\x76\x62\x7e\x98\x7a\xe5\xf2\x3a\xd0\x97\xe8\x30\x60\xdc\xa1\x97\xa2\xa7\xca\x9b\xd4\xd6\xa1\x13\x39\x12\xfe\x29\xfd\xe2\x43\x52\x40\xe2\x4b\xfc\x50\x6d\x48\x63\x91\x49\x4a\x7e\x7c\x47\x3a\x1a\xe9\x12\x3f\x06\xe4\xf7\xa7\xa5\xf9\xe2\x29\x89\x61\x35\xae\x2d\xf5\x68\x49\xa1\x38\xa7\x6e\x4f\x12\x98\x84\x89\x98\x58\x69\x9c\x14\x65\xbc\x0b\xdc\xd0\x5b\x46\x7d\xb9\x7c\xe4\x38\xd2\x88\x88\xd0\x34\xe7\x78\x68\x3d\x9a\x5e\x8e\x12\xb3\x14\x65\x9a\xde\x3a\x2f\x6c\x8c\x60\xa3\xd7\xe2\x94\xc0\x2f\x9e\x2d\xa3\x36\x35\x51\xa6\x71\xab\xa2\xfc\x5c\x52\x33\xb9\xf0\xc4\x9a\xe4\xba\xaa\xf7\x2f\x53\x27\xbe\x76\x78\x30\x80\x8b\xca\x87\xad\xbe\x32\x8c\x79\x9e\x66\x03\x14\x13\x6b\x58\x16\xd5\x2c\x7f\x8a\x65\x72\xa9\x97\xbf\xbe\xa3\xb3\xb7\x2d\xeb\xd3\x9d\x54\xe4\xdd\x00\xb7\x9a\xf9\xfa\x9f\x26\x8e\x05\x1a\xaf\x59\xcc\x93\x75\x6a\x9e\x7f\xf5\xec\xa9\xc2\xbc\x3b\x60\x9d\x89\xbb\xc0\x2f\x9f\x3d\x9d\x2d\xda\xb6\xe2\x89\xde\x38\x0c\xb8\x98\x23\x66\xb2\x62\x57\x29\xad\xf2\xc8\xa7\x38\xfa\x4e\xaa\x1d\x5c\xe8\xfa\xd1\x4b\xc1\x89\x87\x78\x59\x54\x50\x7d\x36\xb1\x09\x91\x7c\x0c\xc7\xea\xc4\x10\xe3\x16\x4b\x4b\xb3\x24\x2d\x10\x4a\xc1\x4f\xa9\x11\x8b\x29\x2e\xa6\x3f\x51\xec\xb0\x17\xd6\x94\x11\xc3\x84\x6d\xa7\x41\xc2\x38\xf6\xd8\xb7\xd0\xfb\x00\x5f\x85\x35\x40\xfa\x83\x42\xde\xa2\xb3\x4c\xa7\x52\xce\x2f\xf6\xfb\x2e\xb6\xb9\xb1\x57\xfc\x1f\xec\xdf\x31\x34\xd0\xd8\x1e\x0a\xde\x7b\xb1\xfe\xb2\xef\xf8\xc2\x5e\x36\xf5\x71\xc4\x46\xd9\x57\x7e\x9b\x53\x81\x3b\x9b\x57\x37\x34\xb1\xb7\xa3\x87\x67\x96\x55\xdf\x77\x86\xfd\x6a\x1a\x6b\x9d\x9f\xd9\x8d\xa7\x40\xe6\x4a\x8a\x0c\x91\x59\xee\x2e\x73\x29\x54\x08\x71\x20\x25\xb7\xf8\x02\x48\x2e\x66\xec\x0a\x4e\x7c\xaf\xab\xc1\xb6\x27\x0c\x83\x64\x52\xef\x81\x96\xec\x62\xd0\x41\x37\x16\xae\x46\xbc\xa4\x42\xa1\x27\x79\xc4\xb2\xb1\xba\x2e\x5d\xd8\x0d\x68\x49\xba\xd1\xfb\x1f\xdf\xe9\x26\x97\x2d\x00\xe8\x55\x16\xb2\x01\x4d\xe7\xf2\x0a\x5b\x45\xdc\xa7\x38\x50\x1e\xe1\x67\x98\x3e\xd5\xc7\x24\x64\x99\x7d\x6d\xea\x74\xbc\x3d\xe2\xc3\xa0\xa1\xdf\xaa\x5b\x98\xc4\xbd\x02\x45\x1e\xef\x20\x70\xc8\xe6\xee\x03\x68\xfe\xa8\xc1\x44\xa9\x89\xcc\xbb\x43\x26\x9d\x7a\x6f\x8f\xb9\x31\xff\x04\x00\x00\xff\xff\x33\x3a\xea\x7e\x55\x09\x00\x00")

func kmsgrantcreateTxtBytes() ([]byte, error) {
	return bindataRead(
		_kmsgrantcreateTxt,
		"kmsgrantcreate.txt",
	)
}

func kmsgrantcreateTxt() (*asset, error) {
	bytes, err := kmsgrantcreateTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "kmsgrantcreate.txt", size: 2389, mode: os.FileMode(436), modTime: time.Unix(1462493337, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _kmsgrantslistTxt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x90\xbd\x8e\x13\x31\x14\x85\xeb\xcc\x53\xdc\x02\x69\x1b\xb2\xeb\x7f\x8f\x5d\x83\x68\xa8\x68\x11\x42\xd7\xd7\xd7\x60\x31\x33\x19\x8d\x1d\x85\xbc\x3d\x13\x12\xad\x56\x5b\x59\xfa\x7c\xce\x77\x64\x7f\xad\xad\xc3\xaf\x0d\x97\xde\x80\x36\xc6\xce\x19\xd2\x15\x9e\xfe\xcc\xed\x8e\x1f\xf4\x09\xca\x69\x03\x84\xb6\x32\xd5\x52\x09\x1a\xef\x17\xfd\x79\x18\x3e\xff\xc5\x79\x9d\x38\x0e\xc3\xe1\x03\xa4\xda\xe8\x5c\x3b\xbc\xd6\x1b\x4c\xb7\x85\x63\x81\xd6\xb1\xfd\x7e\xbe\xe2\x3c\x41\xc6\x8e\x09\x1b\xff\x5c\xb1\xb5\xcb\x69\xcb\xc3\x01\xa7\x8a\xed\xe5\x51\x3f\x66\x2e\x78\x9e\x7a\x1c\x0e\x00\xaf\xcc\x2a\x3d\x1a\x59\x98\xfe\x63\xb8\xfb\x99\xd7\xad\x2e\x54\x57\x9c\x22\xe0\xb6\x44\xbc\xb4\x58\x71\x8e\x51\x2a\x6d\xac\xf3\x63\x10\x52\xc5\x73\xe3\xed\x65\xae\x4b\x9d\xf1\x5e\x3e\xad\xbc\x61\xaf\xa7\xa5\x45\xf8\xfe\x69\x7f\xcb\x75\xed\x1f\xe1\x1b\xf7\xba\xf1\x97\x9b\xf8\xc7\x9b\x91\x9a\xdb\x63\x13\xe0\xdc\x8e\x8c\xad\x1f\x65\x04\xf2\xde\x67\xb4\xfb\x21\x72\xb2\x42\x05\x24\x23\x8c\x36\x61\x54\x5e\x48\x2f\x5d\x60\x72\xa3\xf7\x24\x0c\x7a\x92\xa3\xf6\x85\x9c\x47\xad\xc9\x2a\xe7\x02\xbd\x51\x5e\xf8\xae\x34\x8c\x42\x92\x36\x18\xd8\xe8\x44\x5a\x04\x12\x85\xbc\x14\x68\x33\x1b\xab\xd1\x39\x92\xce\xa7\x51\xec\x5f\x74\xcb\x14\x27\xbd\xc6\x34\x5a\x16\x92\x1d\xbe\x57\xaa\x08\xb2\x48\x69\xa4\x41\xd6\x86\xed\x98\xdd\xee\x77\x46\x59\xab\x93\xf3\xd2\xd8\x40\x3a\x90\x4b\x29\x59\x15\x94\x51\x7b\xb6\x58\xc1\x39\x38\x95\xcd\xa8\x51\xfb\xe1\x5f\x00\x00\x00\xff\xff\x68\xe3\x96\x4d\x23\x02\x00\x00")

func kmsgrantslistTxtBytes() ([]byte, error) {
	return bindataRead(
		_kmsgrantslistTxt,
		"kmsgrantslist.txt",
	)
}

func kmsgrantslistTxt() (*asset, error) {
	bytes, err := kmsgrantslistTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "kmsgrantslist.txt", size: 547, mode: os.FileMode(436), modTime: time.Unix(1462468202, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _kmsgrantsretireTxt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x92\x4f\x6b\xdc\x30\x10\xc5\xcf\xeb\x4f\x31\xa4\x85\x5c\xea\x44\xff\x65\xe9\xdc\xd2\x43\xe9\xa5\x3d\x96\x52\x46\xd2\x28\x11\xb5\xbd\x46\xd2\x92\xee\xb7\xef\x66\x37\x84\x25\x50\x7a\x32\xf3\x3c\xef\x37\xf0\x9e\xbe\x51\x2f\x95\x1a\x20\x7c\xf9\xfa\x1d\x3e\x57\x5c\x3b\xc4\x4a\xd8\x29\x41\x38\xc2\xed\xef\xa5\xc1\xc3\x95\x7a\x7b\x37\x0c\x9f\xfe\xe0\xb2\xcd\xe4\x87\x61\xf7\x0e\x1e\xa8\x43\x7f\x24\xb8\x59\x71\xa1\x1b\xd8\xe7\xf3\x74\xb1\xe4\xba\x5f\xae\x10\x0d\xe6\xd2\xfa\xed\xb0\x7b\x0f\xa1\xb4\x78\x28\x1d\xde\xfc\x83\x31\x43\xeb\xd8\x1e\xef\x8e\xb8\xcc\x90\xb0\x63\xc0\x46\xbf\x36\x6c\xed\x69\x5f\xd3\xb0\xc3\xb9\x60\xbb\x7f\xb1\x8f\x89\x32\x1e\xe6\xee\x87\x1d\xc0\xab\xa6\x85\x9c\x14\xcf\x14\xcf\x32\x5c\xf8\x44\x5b\x2d\x6b\x2c\x1b\xce\x1e\xb0\xae\x1e\x9f\x9a\x2f\xb8\x78\xcf\x85\x54\xda\xd8\xc9\x31\x2e\xfc\xa1\x51\xbd\x5f\xca\x5a\x16\xbc\x98\xf7\x1b\x55\xec\x65\xbf\x36\x0f\x3f\x3e\x52\xac\xc7\xad\x7f\x80\x4b\x6a\xe7\xb4\x7e\x5e\x1d\x29\xa9\xbd\xdc\x04\x38\xb4\x91\xb0\xf5\x91\x7b\x88\xd6\xda\x84\xfa\xf4\x61\x29\x68\x26\x1c\x46\xc5\x94\x54\x6e\x12\x96\x71\xcb\x8d\xa3\x68\x26\x6b\x23\x53\x68\x23\x9f\xa4\xcd\xd1\x58\x94\x32\x6a\x61\x8c\x8b\x57\xc8\x27\xba\x20\x15\x21\xe3\x51\x2a\x74\xa4\x64\x88\x92\xb9\xc8\x72\xb4\x9c\xa1\x4e\xa4\xb4\x44\x63\x22\x37\x36\x4c\xec\x14\xd1\xf3\x4e\x36\xdc\x4a\x0c\x93\x26\xc6\xc9\xe0\x5b\xa4\xf0\xc0\x33\xe7\x8a\x2b\x24\xa9\x48\x4f\xc9\x9c\xf8\x46\x09\xad\x65\x30\x96\x2b\xed\xa2\x74\xd1\x84\x10\xb4\x70\x42\x89\xd3\x6e\xd6\x8c\x92\x33\x22\xa9\x49\xa2\xb4\xff\xa8\xb5\x9e\xa3\x82\x71\x3c\xcf\xe3\xf3\x2b\x79\xad\x2a\xca\x64\xa7\x2c\x03\xe3\xff\x2d\x7e\xf8\x1b\x00\x00\xff\xff\x36\xea\x84\xfe\xa9\x02\x00\x00")

func kmsgrantsretireTxtBytes() ([]byte, error) {
	return bindataRead(
		_kmsgrantsretireTxt,
		"kmsgrantsretire.txt",
	)
}

func kmsgrantsretireTxt() (*asset, error) {
	bytes, err := kmsgrantsretireTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "kmsgrantsretire.txt", size: 681, mode: os.FileMode(436), modTime: time.Unix(1462468128, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _kmsinitTxt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x54\xc1\x8e\xe3\x36\x0c\xbd\xfb\x2b\x08\x9f\x5a\x60\x9c\x7e\xc3\x6e\x81\x02\x83\xc5\x9e\xa6\x40\xd1\xa3\x22\xd3\x31\x31\xb2\x64\x88\x52\xa6\x6e\xd1\x7f\x2f\x49\xd9\x49\x26\xdd\x5b\x62\x49\x8f\x8f\xef\x3d\xf2\x35\x52\x21\x17\xe8\x6f\x64\x48\x19\xea\x3a\xba\x22\x3f\x1d\x4c\x14\x10\x3e\xa8\xcc\xf0\x8e\x1b\xf8\x14\x27\xba\xd4\xec\x0a\xa5\x08\x93\xde\xe4\xfd\xf8\xcb\x1f\x6f\xf0\xed\xfb\xdb\xa9\xeb\xbe\x12\xfb\x4a\x05\x22\xe2\xa8\x10\xf2\xd5\x1e\x53\x04\x74\x7e\x86\x8c\x17\x7d\x5d\x66\x57\x60\x4b\x55\x9e\xf3\x0c\x25\x41\x5a\x51\x80\x51\xee\x9d\xe0\x1b\x6e\xdc\xb9\x2c\x7f\x46\x8c\x85\x26\xc2\x11\x9c\xcf\x89\x79\x7f\xce\x70\x25\x07\xc2\xd8\xb1\xf0\x24\x85\x43\x25\xb4\x74\xbd\x7d\xfc\xe5\xdc\x58\x0c\xff\x04\x77\xc6\xf0\x6f\x2f\xc4\x5e\x27\xe5\xc1\x8d\x6f\x99\x93\x50\x3f\x00\xf0\x2f\xe2\xf2\xa2\x20\x9b\x1c\x87\x00\x67\xd4\xd6\xa4\x68\x1c\x1b\xb4\xea\xc0\x2b\x7a\xe3\xd2\xed\x10\x08\xc3\xa0\x27\xd1\x2d\x72\x25\xb8\xcb\xfd\xb1\x29\x38\x6a\x5f\x19\x27\xcc\xfa\x43\xee\x2f\x8d\x86\xbe\x34\x2a\x63\x82\x98\xca\xa7\xf2\x3f\x69\x49\x4e\x02\xe8\x98\x93\x27\x83\x59\x53\x20\x4f\xc8\x3f\x1f\x05\x3a\x9f\xd1\x4e\xce\x9b\x10\xa5\x78\x11\xa1\xcf\x95\x42\x19\x44\x8b\x5f\x43\xaa\xe3\x6f\x22\x46\xf3\x89\x8b\xf3\xef\x4f\x85\xad\xa2\xea\x66\x95\x0e\x4d\xcf\xb5\x18\x9f\x24\xf7\x32\xbf\x98\x3d\xde\x45\x98\xdd\x15\xed\xed\x42\xac\xc5\x3a\xc3\x38\x28\x98\x18\xc3\xd0\xfe\x0e\xfb\x95\x41\xaf\x48\xd1\xdf\x8f\x8a\x86\x61\x66\x70\xf3\x5e\xed\x35\x8d\x45\x9a\xdd\xe6\xcd\x34\x7a\x72\xfa\x64\x18\x23\x4e\xae\x86\xd2\x10\x80\x18\xfa\xfd\x4b\x7f\x82\x3f\x77\x9a\xcd\x9f\x4d\x69\x67\x48\x1f\xf1\x28\x47\xd3\x2d\x68\x9d\x14\xd3\xcc\x2e\xf2\x92\xd6\x70\x63\xa4\x3d\x88\x1a\x0e\x94\xbb\x7c\x76\xde\xa7\x1a\xcb\xde\x40\x46\x16\x48\x8f\xf7\x9e\x45\xf6\xfe\x7d\xd1\xe4\x51\xe9\xad\x95\x35\xa7\x2b\xb1\x10\x3e\x14\x91\x32\x53\x0d\x07\xf1\x8e\xb1\x14\xc1\xe6\x13\x28\xe2\x8f\x2c\xb2\x00\xdc\xac\x06\x57\xe4\xe3\xdc\xf4\x39\x7c\x7b\xd1\x40\x76\xd6\xdf\xeb\x97\xef\xf7\xbb\xda\xfd\x43\xea\x1c\x37\x0d\xda\x08\xfa\xd9\xc5\x4b\x1b\xd3\x24\x06\x67\x1c\x72\x8d\x51\x6d\xbc\xb7\x70\x82\xaf\xdb\x41\xb5\xf1\x38\xc6\xd6\x02\x67\xe6\xb9\x56\x6f\xfb\xc4\x4c\x06\x5d\xdd\xec\x5c\x08\xe9\x63\x2f\xbb\x66\x8a\x9e\x56\x17\x64\x61\x2c\x22\xb2\x0c\xb5\x6c\x8e\x92\x53\x80\x74\xd5\x59\x68\xdd\xdc\x8d\x73\xe3\xd8\x32\xa7\xa2\x65\xee\xf6\x8e\x83\x66\x34\x4d\xd0\xbb\x71\x11\x96\x5c\x64\x3f\xa4\xcc\xbd\xae\xa8\xde\x6e\xf6\xf7\xfc\x4b\xa9\x45\x47\x27\x50\x6c\xb3\xa8\xe9\x7b\xf0\x68\xc1\x7c\x41\x36\xdc\xcf\x6b\x4c\x3a\xb8\x60\xb4\xdd\x73\xdb\x0d\xd8\x06\x44\x35\x7a\xda\x79\x39\x2d\x76\x7e\xdb\x05\xb6\x1a\x5e\x34\x63\xb4\x0f\xb2\x59\x5c\x3f\xc5\x4e\x94\xea\x30\xfa\xbc\xad\xa5\xbd\xd6\x85\x71\x75\xa1\x4a\xc9\x1a\x47\x69\xfc\x29\x90\xf7\xd9\x93\x15\xc1\x49\xdb\xfc\xbf\x69\x90\xa2\xb7\xad\xd7\x16\x6b\x1b\x8d\xc7\x95\x2a\x1a\x3d\x88\x1c\x04\xa7\x65\xa1\x2b\x0f\x13\xe5\xc2\x25\x65\x69\x7b\x31\xa4\xa9\x96\x2a\x69\x66\x94\xa4\x17\xb6\xba\xb7\xbc\x3c\x46\xfe\x10\xaa\xb3\x56\x5a\x69\x5d\x1f\x0e\x46\x9a\x64\xe1\xc9\x3c\x3f\x20\xfb\x39\x91\xc7\x53\xf7\x5f\x00\x00\x00\xff\xff\x7e\xae\x79\x02\x68\x06\x00\x00")

func kmsinitTxtBytes() ([]byte, error) {
	return bindataRead(
		_kmsinitTxt,
		"kmsinit.txt",
	)
}

func kmsinitTxt() (*asset, error) {
	bytes, err := kmsinitTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "kmsinit.txt", size: 1640, mode: os.FileMode(436), modTime: time.Unix(1462488452, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _usageTxt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x91\x41\xeb\xd3\x40\x10\xc5\xcf\xe6\x53\x0c\x3d\xb5\xd0\x14\x11\xf4\xe0\xad\x16\x11\xd1\x0a\x52\x41\x3c\xc9\x74\x33\x49\x96\x6e\x76\x97\x99\xd9\xfa\xcf\xb7\x77\x36\x46\xfe\xf5\x60\x4e\x9b\xcc\xe6\xf7\xe6\xbd\xf7\xce\x8b\x2b\x5e\x61\xa4\x90\x05\xe6\x54\xa0\x08\xc1\xf1\xfb\x05\x3e\x9d\x2f\xa0\x09\x26\x8c\x38\x10\xf8\xd8\x33\x8a\x72\x71\x5a\x98\x40\xc8\x31\xa9\x1c\x9a\xe6\x2f\xc0\x7e\x13\x10\xc5\xd8\x21\x77\x20\xf3\x34\x91\xb2\x77\xe0\x78\xce\x86\xc1\x30\x24\xf6\x3a\x4e\x02\xdb\x2f\x78\x0a\x2b\xe1\x9a\x9e\x20\x71\x73\x7c\x7f\x69\x3f\x9c\xce\xed\xab\xd7\x6f\x76\x60\x08\xa0\x3c\xd2\x44\x8c\x01\x3a\x54\x84\x1b\xcd\x02\x99\xd3\xdd\x77\xd4\xc1\x75\x5e\x96\xdb\x16\xf1\x71\x80\x8d\x0d\x9b\x5f\x8c\x39\xdb\xdb\xc6\x68\xb0\xa1\x78\xa7\x90\x32\x01\xc5\x45\xde\xa7\xb8\xd9\x1d\xe0\xdb\x48\xe0\xbc\x91\x59\xe9\x49\x05\xb0\x1a\xd1\xc4\x86\xf4\x11\xb0\xf9\x71\x3c\x7f\x86\xde\x07\x02\x1d\x51\xc1\x9b\x1f\xec\xa9\x86\xd0\x79\xb3\xee\xaf\x45\x6b\x10\x20\xa9\xb0\x33\x54\x8a\xca\x29\xec\xa1\xa3\x1c\xd2\x3c\x51\x54\xd9\x37\xa7\x8f\x66\x5e\x94\x26\xd9\x03\xa9\x7b\x48\x68\xdd\xdf\x74\xab\x9f\xf6\x8e\xa1\xac\xfa\x06\x55\xe2\x1e\x0d\xba\xcd\x45\x97\x04\x06\xd2\xdd\x7e\x39\x69\x4a\xa1\x1a\xed\x2d\xa8\x81\x31\xfe\x99\x67\xfb\xe8\xe6\xb5\x9e\xaa\x0d\xe8\x38\x89\x2c\xdd\x31\x0d\xe6\xb9\xd6\xf3\xb5\x78\x77\xb3\x5a\x58\xdf\x36\xcd\x8b\xeb\xba\xca\xcd\x6a\x30\x81\xd6\x61\x08\xc4\xad\x6d\x15\xd5\xeb\xfc\xef\x05\x1f\xed\xd0\xf6\xb5\x54\x19\x0f\xf3\x14\x9e\xc7\x75\xc9\xc7\x09\x04\x2c\xd1\x8d\x3f\x5d\xaa\xfe\x5e\xda\xf3\x7c\xd7\x74\xfe\x7f\xb7\xf9\x1d\x00\x00\xff\xff\x69\x1d\x49\x1d\x80\x02\x00\x00")

func usageTxtBytes() ([]byte, error) {
	return bindataRead(
		_usageTxt,
		"usage.txt",
	)
}

func usageTxt() (*asset, error) {
	bytes, err := usageTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "usage.txt", size: 640, mode: os.FileMode(436), modTime: time.Unix(1462485873, 0)}
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
	"awskms-key.template":  awskmsKeyTemplate,
	"kmseditkeypolicy.txt": kmseditkeypolicyTxt,
	"kmsgrantcreate.txt":   kmsgrantcreateTxt,
	"kmsgrantslist.txt":    kmsgrantslistTxt,
	"kmsgrantsretire.txt":  kmsgrantsretireTxt,
	"kmsinit.txt":          kmsinitTxt,
	"usage.txt":            usageTxt,
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
	"awskms-key.template":  {awskmsKeyTemplate, map[string]*bintree{}},
	"kmseditkeypolicy.txt": {kmseditkeypolicyTxt, map[string]*bintree{}},
	"kmsgrantcreate.txt":   {kmsgrantcreateTxt, map[string]*bintree{}},
	"kmsgrantslist.txt":    {kmsgrantslistTxt, map[string]*bintree{}},
	"kmsgrantsretire.txt":  {kmsgrantsretireTxt, map[string]*bintree{}},
	"kmsinit.txt":          {kmsinitTxt, map[string]*bintree{}},
	"usage.txt":            {usageTxt, map[string]*bintree{}},
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
