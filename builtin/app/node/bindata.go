// Code generated by go-bindata.
// sources:
// data/aws-simple/build/build-node.sh
// data/aws-simple/build/template.json.tpl
// data/aws-simple/deploy/main.tf.tpl
// data/common/dev/Vagrantfile.tpl
// DO NOT EDIT!

package nodeapp

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
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

var _dataAwsSimpleBuildBuildNodeSh = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x56\x7f\x4f\x1b\x3d\x12\xfe\x7f\x3f\xc5\x34\xa0\x72\x27\xb1\xbb\x80\xae\x77\x6d\x68\xab\x03\x0a\x2d\x52\x15\x2a\xda\xd3\xbd\x52\xdb\x17\x6d\xd6\x93\xc4\x65\x63\xbb\xb6\x37\x10\x28\xdf\xfd\x7d\xc6\x1b\x12\xf2\xf6\x87\xfa\x0f\x38\xb6\x67\xe6\x99\x99\x67\x1e\xef\xc6\xa3\x72\xa8\x4d\x39\xac\xc2\x24\xcb\x02\x47\xca\x2d\x19\xdb\x9a\xc5\x92\xbd\xe7\x6b\x9d\x96\x4e\x3b\x1e\x55\xba\x59\x6c\x47\x5f\xd5\x9c\x65\x58\x59\xff\x8f\x7f\xd2\x6d\x46\x44\x8d\xad\xab\x86\x82\x6d\x7d\xcd\x23\xdd\xf0\x8b\xcd\xdd\xd5\x76\xa3\x0d\x1b\xfb\x62\x73\x4f\xb6\xb8\x9e\x58\xea\x1d\x9f\x9f\x9f\x9d\x53\x15\x69\xf3\x76\x65\x74\xd7\xdf\xbc\xed\xee\xde\xed\xd3\xdb\x2a\x44\xd8\x8f\x43\xbf\x27\x66\x63\xcf\x8e\x6c\x8c\x96\xca\x59\xe5\x4b\x1c\x94\x61\x1e\xf0\x8f\xbe\x51\x4c\xd8\x0c\xed\xed\x64\x77\x19\xd0\x39\xda\x4a\xe0\xa8\xb7\x79\x7b\x78\xf0\xfe\xcd\xc5\xfb\xb3\xff\x9d\x1f\x1d\xdf\xf5\x64\xe3\xed\xe9\xe0\x78\x70\x76\xd7\xdb\x22\x60\xc8\x32\xcb\x92\x02\x0e\xfe\xdb\xa3\xbd\x97\x8f\x77\xe1\x0e\x4e\xc7\xec\x29\x8f\x5d\xbc\x97\x54\x2a\x9e\x95\xa6\x6d\x9a\x7d\xba\xcb\x6c\x93\x0c\xba\x34\x3e\xca\x8d\xcf\x04\x63\x39\xca\x36\xa8\x6e\x6c\xab\xf2\xda\x9a\x91\x1e\x53\x5d\x19\xd2\x26\xb2\x1f\xb1\x67\xba\xd2\x71\x42\x95\x8b\x54\xdb\xe9\xb4\x32\x2a\x90\x1e\x91\x8e\x5b\x81\x42\xd4\x4d\x83\x9b\xe4\xbc\x45\x9e\x21\x20\x08\xf5\xfe\x5f\xe9\xa8\xcd\x98\x46\x48\x64\xcd\x2d\x30\xc1\x85\x6b\x38\x72\x51\x14\xbd\xac\x35\xb0\xa7\x8f\x1f\x29\x1f\x2d\x8a\xa3\x87\x65\xb2\x28\xb5\x09\xb1\x32\x35\x97\x43\x6b\x63\x3e\xd2\x46\x87\x09\x2b\xfa\xfc\x79\x9f\x94\x45\x59\x43\xc3\x28\xeb\x4e\xf1\x24\x53\xd6\xa0\xa7\x12\xf7\x40\x29\x09\x2b\x48\x81\xc5\xd9\xa0\xa3\xf5\x9a\x03\x01\x33\xb5\x4e\x55\x82\x2a\x05\xb6\x4c\xa1\x55\x56\xae\xe6\x63\xb0\x26\x1d\x32\xe5\xf3\xef\x4e\x12\x0e\xe4\x98\xcf\xc1\x91\x51\xbc\xaa\x3c\xe7\x48\xd6\xb1\x8f\xf0\x9c\x4b\x45\xac\x59\x59\x29\x95\x8b\x25\x5a\xde\x45\x9f\x8b\xa1\x73\x55\xbf\x9e\x78\x1d\xf2\x86\xab\xd2\x58\xc5\xc5\x97\xb0\x16\xe9\x92\xe7\xb0\x9d\x51\x2e\xab\xc0\x7e\x86\x2e\x4e\x2e\x5d\xbf\x2c\x97\xbf\x8b\x76\x88\x72\xb5\x05\x22\xf6\x9f\xee\xe0\xa6\xe7\x7a\x96\xae\xd3\x93\x7f\xef\x9e\x3c\x3b\x7c\x76\x74\x70\xf4\xaf\x9d\xc3\xbd\x93\xff\x64\xa9\xc5\x5b\x8a\x87\x34\x89\xd1\x05\xb8\xb1\x21\xe4\x98\x99\x4a\xca\x51\xb8\x49\x1b\xb4\x35\xae\x0a\x81\x0d\x08\x23\x3e\x4b\xc0\x28\x97\x3b\x14\x7d\x1b\xe2\x9c\xa6\x95\x36\x5b\x20\x56\x02\x1a\x99\xa9\xe4\x58\xa7\xab\x1d\xf7\x43\xd1\xe8\x10\x0b\xb5\xb2\x4c\x1b\x0f\x99\xf7\x93\x5a\x77\x1d\x3b\xed\xaa\x2b\x5d\x1b\xa0\x2c\xdb\xf4\xee\xde\xcf\x36\x0d\xc6\xda\x5c\x6f\xa7\xde\xd9\x38\x01\x26\x57\xd5\x97\xd5\x18\x31\xa5\x83\x7c\xed\xac\x8f\xf4\xea\xf8\xf0\xf4\x60\x70\x71\x72\x7e\x36\xf8\x70\x3c\x78\xf5\xc2\x58\x93\x78\x5b\xd5\x51\xcf\xf8\x57\xdd\x1c\xde\x78\x1a\x43\x25\xa6\xec\xeb\xd6\x6b\x0c\xfb\xb0\xd5\x8d\xca\x59\x00\x44\xf9\xfd\x09\x34\x03\x21\xdd\xd7\x1c\xb9\xd0\x0d\x96\xbb\xe3\xb4\xfc\x39\x11\x92\x8d\x04\xc3\x2c\x9b\x20\x08\xf3\xd4\x81\xb4\x2f\x8d\xff\xb2\x58\x4a\x6e\x39\x5f\xe3\x5a\xa0\x65\xed\xba\x9a\x1c\xcb\x6e\x1d\x3b\x26\xbb\x94\x6c\xca\x61\x7a\xa9\x34\x66\xdb\x51\x19\xfc\xac\x94\xf1\x05\xd5\x5c\x77\x16\x2b\x4f\x37\xd7\x18\xa2\x38\x75\xcb\xa3\x22\x8e\x6f\x28\x3f\xfa\xdb\xfd\xf5\x49\x71\x8d\xae\x31\x16\x40\xde\x82\x66\xeb\xb3\xa1\x94\xec\x81\x69\x4a\x87\x6a\xd8\xb0\xca\x05\xe9\x95\xf5\x0a\x7b\x63\xae\x6d\xa0\x5e\x8f\xd6\x1d\xbf\xe7\x98\x90\xa3\x2c\x53\x1d\x84\x65\x61\xcd\x29\x88\x79\x65\x28\x3f\x5f\x9a\xf5\x7f\x04\xef\x28\xc9\x05\xba\x02\x4f\xa9\x54\xc9\x07\x44\xea\xc3\x44\x43\x7c\x02\xc6\xfb\x6b\xab\x3d\xf4\x40\x24\xe6\x01\x6b\x2d\x49\x97\x2b\x9c\x57\xc1\x1a\x01\x4d\x6c\x66\xda\x5b\x33\x45\x53\xe9\x6a\x22\x72\x86\xa6\x43\xdf\xe0\x0d\xaa\xa2\x88\xaf\xb9\x6e\xa3\x5c\x0d\xe8\xf6\x25\x28\xde\x06\x9f\xde\x17\x58\x6e\xaf\x7e\x81\x24\xcd\x36\x81\xfe\x05\x9d\x22\x44\x13\x84\x55\x0e\x2c\x30\xb1\x99\xc3\x99\x61\x86\x30\x02\x81\xad\x71\x95\x26\x7a\x3c\x11\x61\x04\x73\xa9\x53\xbf\x82\x0e\x9c\x63\x93\x0a\x0f\x08\x92\x88\x09\xed\x68\xa4\x6b\x0d\x1f\x05\xf5\xf3\x6f\x5d\x33\x03\xf2\xca\x35\x6d\xed\x86\xf2\x4f\x01\x41\xef\x0e\x3e\xbc\xd9\xff\x64\xca\xad\x6e\xfc\x52\x45\xba\xbf\x85\xb8\xfe\x81\xd5\x06\x9d\xa1\xa0\x7d\x92\x17\x51\xac\x41\xd9\x07\x65\x12\x75\x0f\x18\xe6\x7b\x41\xf8\x85\x6b\x24\x36\x40\x62\x92\x97\xe7\xa9\x9d\x31\x12\x02\x72\xb8\x4b\x97\xa4\xd0\xc8\x1a\x5a\x40\x90\x3b\xee\x90\xf8\xe9\x43\x67\xb2\x1f\x72\x4e\xcd\x50\x50\x85\x51\xd5\x36\xb1\xeb\x25\x07\x4e\x2f\x2c\x04\x1a\x6d\x71\xd0\x6c\x69\x12\x26\x42\x86\x09\xcb\x70\x5f\xc0\x25\xf4\x7c\xa1\x01\x8a\x56\x18\xb7\x31\xb8\x11\xfe\xd2\x83\x84\xbe\xeb\x8e\x08\x0a\xb3\x09\x26\x04\x46\x8f\x44\xdb\xef\x9f\xa0\x09\x92\x8f\x5d\xb9\x6c\x8b\x68\x60\xbe\x59\xc4\x2b\x32\xcc\x02\x3d\x7f\x3e\x78\x7d\x3a\xf8\xe3\xe8\x6c\x70\xf2\x9d\xf2\x75\x29\x89\xab\x35\xcd\x93\x8d\x35\xcd\xdb\xa0\xd7\x6c\x58\xe2\x2a\x1a\xce\x53\x33\xb2\xe5\xf5\x0b\x8f\x27\xad\x23\x96\xbc\x77\xbe\x1d\xce\xcb\x19\x88\x61\x71\x22\xeb\x85\x3c\x5f\x2c\x0d\x4a\xf9\x18\x89\x69\x96\xf0\x0e\xee\x67\x4b\x7c\xd9\x6f\x02\x5e\xef\xc1\x52\x1f\x7e\x0b\xf6\xe2\x35\x4a\xdf\x4a\x24\xe2\xce\x86\x9e\xee\xec\xa7\x9f\x5d\x22\x0f\xe7\xb7\x74\xed\x10\x92\xd2\x1d\xaf\x32\x5e\x84\x26\x6b\xf6\xf1\xa1\xf3\x00\xbf\x4c\xfb\x79\x6b\x4c\x9a\x74\x37\x5d\xa9\x5d\xde\x2e\x25\x42\x68\xbd\xfc\xe0\xa3\xbc\xa9\xa9\x57\xab\xf5\xb0\xf4\xf8\xb1\x98\xaf\xf4\x5d\xa4\x59\xb5\xb5\x14\xad\xd7\x45\x81\x6b\xf9\x52\x78\xd4\xcb\xfe\x0a\x00\x00\xff\xff\x88\x43\xa9\x98\x3f\x0a\x00\x00"

func dataAwsSimpleBuildBuildNodeShBytes() ([]byte, error) {
	return bindataRead(
		_dataAwsSimpleBuildBuildNodeSh,
		"data/aws-simple/build/build-node.sh",
	)
}

func dataAwsSimpleBuildBuildNodeSh() (*asset, error) {
	bytes, err := dataAwsSimpleBuildBuildNodeShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws-simple/build/build-node.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataAwsSimpleBuildTemplateJsonTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x54\x5d\x6f\x9b\x30\x14\x7d\xe7\x57\x58\x48\xe9\x53\x80\x6c\xad\xa6\x69\xaf\xfb\x19\x55\x44\x0d\xb8\xe1\x2a\xb6\xb1\x7c\x4d\xa6\x16\xf9\xbf\xef\x9a\xef\x4c\x0d\xcd\x9a\x17\x23\x9f\xc3\x39\xc7\x87\xeb\x74\x11\xa3\x5f\xac\x40\xe7\x86\x97\x67\x61\xf3\x8b\xb0\x08\x8d\x8e\x7f\xb1\xf8\x90\xfe\x4c\x0f\xf1\x3e\x1a\x38\x17\x6e\x81\x17\x52\x20\x41\xc3\x6b\xb4\xc9\xff\x60\xce\xcb\x52\x20\xe6\x67\xf1\x46\x88\x6e\xa5\xdc\xaf\x51\x14\xa5\x15\xee\x16\x6a\xc5\x69\x30\xbb\x42\x50\xb6\x27\xca\xe3\xea\x11\xe8\xf7\xfd\x14\xc4\xd8\xe6\x02\x21\x23\x25\x25\xc2\xf3\xf8\x56\xb7\x63\xaf\x8d\x65\x15\x58\x06\x9a\x1e\x5b\x5d\x71\x47\xac\x9c\x76\x30\x2d\x5a\x90\x15\xdb\xf9\x89\x3c\xae\x24\xe7\xde\x8c\x08\xa7\xc5\x5a\x48\x19\xef\x17\x00\xb4\x04\x1d\xa0\xe7\x58\x9d\x83\x6c\x62\x58\xe6\x94\xc9\x1a\xe7\x9a\x6c\x31\x48\xba\x2e\x38\xcb\xa6\x31\xe9\x6f\xda\x75\xc2\x32\xef\xe3\xe3\xa8\xe4\xf7\xb7\x3d\x5f\x41\x8a\xb5\x25\x36\xad\x2d\x7b\x84\x34\x83\xa5\xf7\xd9\x1a\xaf\x04\x3a\xd0\xbd\x6b\x20\xfd\x47\x9a\x3b\xc2\x6c\x15\x50\x56\xf7\x1e\xdd\x7b\xf6\xf0\xc0\x0a\x8e\x35\x4b\x33\xc5\x41\xa7\x58\x7f\xd0\xc5\x8e\x09\x5d\x85\xef\xb5\xf5\x49\x36\xea\xd9\x31\x1a\xd4\x82\x42\x28\x52\xa0\x14\x2d\xd2\x39\x5f\xe6\xc1\x79\xa1\x33\x0f\x1e\x2b\xda\x3d\x4d\x26\xdc\x98\xd4\x9d\xde\xbf\x54\x18\x96\x16\x8c\x0b\x50\x3f\x6e\x89\x6e\x2a\x11\x8e\x3f\x69\xf5\xeb\x71\x9a\xe3\x9e\x33\xce\xf0\x7c\xa1\x34\x57\xbd\x76\xc8\x32\x4b\xcf\x8e\x5c\xf1\x77\x6a\x5d\x14\xb8\x60\x57\xd7\xef\x56\x31\xd7\xf7\x74\xbb\x9d\xf8\xea\xca\x6e\x29\x2e\xc4\x4f\x14\xe7\x6b\xbe\xa5\x36\x90\x3e\xcb\xd6\x8f\x40\xce\x15\x0c\x7d\x40\xf2\xfd\xdb\x8f\xc7\x43\xf5\xf4\xb4\x70\x40\xa3\xe3\x9a\x58\x53\x6d\xe5\x63\x2a\xb9\x3d\x89\x95\x0c\xd6\x79\x70\x9e\xea\x6e\x0b\x1a\xde\x76\x55\xaa\x82\x7c\xc2\xba\x2e\x3c\xd1\x5c\xff\x9b\x9d\x56\x9a\x22\xae\xcc\x47\x89\x87\xff\xac\x63\x14\xf9\xe8\x6f\x00\x00\x00\xff\xff\xa2\x67\x64\x41\x65\x05\x00\x00"

func dataAwsSimpleBuildTemplateJsonTplBytes() ([]byte, error) {
	return bindataRead(
		_dataAwsSimpleBuildTemplateJsonTpl,
		"data/aws-simple/build/template.json.tpl",
	)
}

func dataAwsSimpleBuildTemplateJsonTpl() (*asset, error) {
	bytes, err := dataAwsSimpleBuildTemplateJsonTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws-simple/build/template.json.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataAwsSimpleDeployMainTfTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x53\xc1\x8e\xdc\x2a\x10\xbc\xcf\x57\xb4\xd8\x3d\xbe\xf5\xcc\xcb\x31\x52\x6e\x91\x72\x4b\x3e\x20\x5a\x21\x8c\x99\x09\x5a\x0c\x08\x9a\x89\x2c\xcb\xff\x9e\x06\x62\x31\x38\x9b\x6b\x3c\x27\xaa\x8b\x2e\xba\xba\xe6\x09\xbe\x28\xab\x82\x40\x35\xc1\xb8\xc0\x37\x44\xf7\x1f\x4c\x0e\xac\x43\x50\x93\x46\x98\x85\x4d\xc2\x98\xe5\x74\xba\x8b\xa0\xc5\x68\x14\x30\x6d\xaf\x41\x70\x3d\x31\x58\xb7\x07\x58\xfc\x8c\x5c\x48\xa9\x62\xe4\x6f\x6a\x79\xa7\x18\x95\x0c\x0a\xff\x52\x0c\xea\xa6\x9d\x3d\x14\x88\xca\xad\x98\x55\x81\x1f\x2f\xcc\xfa\xc0\xd4\x36\xa2\xb0\x52\x71\x5c\x7c\xa6\xc3\xa4\xae\x22\x19\x84\x4f\xc0\xf0\xc3\x30\x6b\x19\x1c\x83\xc7\x1b\x31\x8d\x96\x5e\xe3\xd3\x68\xb4\x3c\x74\xbb\x7b\xc9\xa5\x9e\xc2\x3b\xf0\xef\xb1\x4f\x3e\xb8\xbb\x9e\x54\x28\xaf\x27\xe8\x04\xd0\x86\xcf\xaa\xcf\x2b\x5d\x1c\x7a\x53\x36\x46\xb4\x66\x43\x4f\x6b\x78\xa1\x55\x43\x20\x7f\x1d\xad\xe2\x44\xa1\x47\x04\x15\x5d\x0a\xb2\xf9\x9b\x82\xc6\x85\xdf\x82\x4b\x9e\x11\xe8\x7d\x7d\x59\xf6\xb0\xf6\x59\xd7\x7a\xd8\xb6\x97\xda\x72\x5f\x66\xd1\xac\x03\x36\xbd\x7a\xa6\x12\xd5\xb4\xbd\x91\x5c\x2c\xfd\x00\x68\x7c\x74\xd2\x99\xfa\xbc\x97\xff\x0b\x78\x0d\x6e\xe6\xde\x05\x2c\xe0\xa5\x60\xe8\x76\xa4\x61\xd9\x5a\x3e\x1a\x27\xdf\x22\x61\xdf\xd9\x65\x28\xbf\xf3\x85\xbd\x52\x7d\xcb\x6a\xea\x9f\x89\x91\xdc\x13\x7c\x56\xde\xb8\x05\x04\x2d\x07\xc1\x5d\x61\x8f\x53\x3c\x58\xbc\xe3\x8f\xe6\x52\x18\xa1\x7d\x6d\x57\xb3\x2e\x9e\x76\xc9\x6c\xe5\x0e\xae\xb9\xa8\x81\x24\xff\xbb\x3e\x5d\x4e\x0b\x71\xff\x57\x1c\x04\x77\xb8\xae\x2b\xaf\xae\x0f\x04\x75\xae\x0e\x3c\xaf\x7f\xa6\x65\xa0\x71\x86\xbc\xea\xd7\x7c\x19\xc5\x8d\xbc\x87\xaf\x59\xa4\x0b\x0d\xab\x86\xb9\x84\x3e\x21\xb0\x14\x4c\xf5\xe0\x2e\x4c\x2a\xd4\x1f\x88\xfe\xe3\xf9\x5c\x25\xf6\x19\x4b\xf3\x3a\x00\x9f\x6c\xdc\xce\x39\xbc\xbf\x02\x00\x00\xff\xff\x05\xad\xee\xc5\x7b\x04\x00\x00"

func dataAwsSimpleDeployMainTfTplBytes() ([]byte, error) {
	return bindataRead(
		_dataAwsSimpleDeployMainTfTpl,
		"data/aws-simple/deploy/main.tf.tpl",
	)
}

func dataAwsSimpleDeployMainTfTpl() (*asset, error) {
	bytes, err := dataAwsSimpleDeployMainTfTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws-simple/deploy/main.tf.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataCommonDevVagrantfileTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x55\x7f\x6f\xd4\x46\x10\xfd\xdf\x9f\x62\x70\x00\x25\x4a\x6c\x53\x84\x10\x4a\x21\x2a\x4d\xf8\x11\xa9\x0d\x88\x84\xfe\x53\x55\xc7\x9e\x3d\xb6\xb7\xec\xed\x2e\xbb\xeb\xcb\x1d\xc9\x7d\xf7\xce\xac\xed\xe4\x0e\x42\x00\x55\xca\x45\xf6\x78\xf6\xcd\xcc\x9b\x79\xb3\x5b\xf0\x0a\x35\x3a\x11\xb0\x82\xe9\x12\xde\x84\x60\xf6\xa0\x32\xa0\x4d\x00\xac\x64\xb8\x93\x6c\x25\x5b\x70\xd6\x4a\x0f\xf4\x17\x5a\x84\xbf\x44\xe3\x84\x0e\xb5\x54\x08\xcd\x97\x67\xa1\x36\x2e\x7a\x55\x38\x47\x65\xec\x0c\x75\x00\x53\x13\x44\x60\x08\x61\xad\x92\xa5\x08\xd2\xe8\xc2\xa3\x9b\xcb\x12\x73\x38\x0e\xe0\x5b\xd3\xa9\x2a\x06\x9d\x22\xb4\x42\x57\x19\x07\xc7\x2a\x87\x33\x03\x33\x53\xc9\x7a\xc9\xb0\x84\xb3\x16\x7e\x0f\x3a\x8f\x31\xda\x73\x6b\xd9\x90\x27\xc9\xf0\x39\x2f\x8d\xae\x65\xd3\x39\xdc\x4e\x1f\xa6\x3b\x5c\xd1\x65\x6f\xba\x4c\x00\xfa\xa7\x7c\x3e\xcb\xa7\x66\x01\xcf\x20\x6d\x85\x6f\x65\x69\x9c\x2d\xac\xc3\x52\x7a\x7c\xfc\x28\x4d\xc8\x71\x0b\x5e\x1b\x4f\x05\x68\xb5\x04\x8d\xe1\xdc\xb8\x8f\x1b\xc7\x07\x1b\xa4\xd6\xc9\x39\xf1\x30\x19\x0c\xe9\x1e\x48\xbb\x0f\xe9\xc5\x05\x13\x31\x91\x76\x22\xaa\xca\xa1\xf7\xb0\x5a\x0d\xc0\xa7\x18\x3a\x0b\x02\xfc\x52\x97\xc4\x5f\x6d\x54\x85\x0e\x6a\x67\x66\x60\x3a\x07\x8c\x22\x75\x03\x95\xa4\x84\x82\x71\x54\xbe\x81\x62\xde\x57\xb7\x91\x43\x0f\x30\x19\x00\x38\xa4\x15\xa1\xcd\x47\x00\x0a\xb8\x07\xe9\x78\x32\xdd\xa3\xb3\x00\xe6\x9c\xfa\x46\xf9\x5d\x59\xa1\x71\xa6\xb3\x6b\x96\x3e\xc9\x17\x5a\x4c\xa9\xcd\xa7\xa7\xaf\x41\x34\xdc\x4a\x6a\xef\xb9\x70\x15\x03\x7b\x43\xed\x0f\x81\x1f\x87\xea\xa9\x56\x8b\xba\x42\x5d\x4a\xf4\xb1\x02\x7f\x9d\xa9\xf7\x6d\x3e\x9c\x9e\xf4\x58\xcf\x20\xb8\x0e\xfb\x40\x2f\x4d\xa7\xab\x38\x17\x30\x76\xae\x7f\xdb\x96\x35\x08\xbd\xdc\x21\xaf\x8b\x7b\x71\xba\x88\x11\x90\x9a\x1e\xc7\x13\x13\xb2\xf8\x9c\x78\x86\x7b\x2b\x72\xe3\xef\xd4\xd2\xc2\xd0\x38\x16\xd7\x5e\x19\x11\x43\xc7\x95\x31\x36\x3f\x24\x6b\x20\xb2\xb8\x19\xb7\x53\xc9\x60\x91\x41\x7a\xd8\x70\xb5\xce\xcc\xa5\xe7\x0c\x53\xdf\xa2\x52\xdc\x71\xad\xa4\x46\xe2\xb0\xac\x60\xeb\x82\x0e\xac\xe0\xfe\x7d\x98\xd2\x68\x0d\xaf\xc5\x4c\x48\x9d\xfb\x36\xed\x8b\x21\xaa\xb8\x1e\x4a\x3a\x52\xf0\x87\x11\x15\x08\xa5\x62\xfb\x6b\x27\x1a\xd6\x8e\x87\x16\x1d\xc6\xba\x89\x85\x0d\x82\xf3\x6b\x4a\x46\x6f\xe6\x85\xe7\xed\xfa\x74\x64\x84\x2b\x1f\x2c\x97\x0e\x29\xca\x6a\x75\x63\x06\xc7\xda\x07\x4e\x60\xda\x49\x12\x23\xea\xb9\x74\x46\xf3\xa9\x1f\xad\xfc\xae\x2f\x9d\xb4\x61\x42\x32\x4f\x08\x3b\x49\xd6\x0c\xd4\x93\xa7\x4f\x4f\x0f\xdf\x1d\xbf\x3d\x4b\x3c\x06\xc8\xa8\xf3\x5b\xc0\x4d\xca\x70\x81\xe5\x3e\xf0\xff\x8e\x86\xa8\x34\xb3\x19\x2d\x00\x38\x97\xa1\x25\x2e\x82\xed\x02\x28\xd3\x34\xbc\x64\xe8\x91\x77\x44\x25\xbd\x55\x62\x89\x55\x62\x70\x7b\x07\x2e\xe0\xee\x6f\xf0\xf0\xe0\xfe\x2f\x70\xd9\x7b\x3a\xc8\x42\x84\x86\x03\x28\x88\x90\x42\x77\x4a\xfd\x0a\xab\xab\x88\xe4\xb5\x3f\x62\x0b\x1a\x5f\xac\xe5\x82\xf0\x67\xa4\x50\x9a\xcd\xc4\xa8\x88\x8a\x65\x6b\x20\xfd\x9b\x4f\xfc\x43\x21\xd2\x01\xe1\x4f\xf1\x11\x41\x06\x16\x40\x68\x45\x80\x0f\x83\x66\x80\x46\xfc\x03\x34\x86\x66\xbf\x57\xad\x8a\xa2\xe5\xfd\x44\xab\x85\x0d\x71\x8a\x7a\x54\x9a\x91\x2b\x4d\xc2\x01\xa5\xd9\x9a\x19\x8e\x96\x22\xe7\xa9\x71\x25\x47\x3b\x1c\xe4\xc0\x3a\x63\x1d\xc6\x7e\x0b\xcf\xe3\x4b\x55\x48\x9d\x90\x40\xee\x90\x7c\xd1\x42\xfa\xde\xe3\xd1\xc9\x29\x51\x94\x42\x81\xa1\x2c\x28\x21\xfe\x55\x93\xbe\x7b\x70\xb0\x46\x06\xa5\xa5\xa9\xaf\x7d\x36\x6b\x07\x2f\xc1\x77\xb4\x2d\x03\x22\x64\xe2\x7b\x30\x04\x60\xb0\x3f\x30\xac\x73\x26\x01\x68\xd3\x05\xe1\x42\x52\xcb\x24\xc1\x85\x35\x2e\xc0\xd1\x8b\xdf\x8f\x9f\x9f\x4c\x5e\xbe\x7b\x73\x72\xf6\xe2\xe4\xe8\x99\x36\x5a\xb2\x06\x45\x19\xe4\x9c\x18\x1f\x50\x84\x0d\x19\x6d\x15\xe8\x6c\xc5\x0b\x25\x5b\x26\xd4\x0c\x48\x8f\x68\x61\x29\xd2\x07\x93\x70\x62\x2a\x84\x61\xb1\x6a\x7a\x9e\xcc\xd1\xc5\x71\x5c\xad\xf2\x3c\x4f\x19\xea\x9c\x21\xb2\x4f\x90\xbd\xf9\x82\x57\xf6\xcf\x29\xb5\xbc\xf9\x0c\x6d\x08\xd6\xef\x17\xd1\xf6\xaf\xcf\x8d\x6b\x0a\x9a\xab\x50\xcc\x6f\xc6\x8e\x7e\xd9\x37\x3e\x66\x34\xff\xdd\x22\x5b\x3c\x7e\x34\xa0\xf7\x69\xbf\xd7\xf4\xe6\xc6\xa4\xc7\xec\x7a\x7e\x05\x0d\xe9\x21\x14\xc6\x52\xa6\x8b\xcf\xf5\xb7\x13\xed\xa1\x4e\x87\x55\x4b\x77\xc6\xdb\xe7\x67\xaf\x37\xb0\x94\x86\xcc\x47\xa8\x1f\x4c\xb2\x98\x4a\x1d\x5d\xa1\xe8\xbc\x2b\x94\x29\x85\xba\xb2\xfd\x5f\x58\x3b\xfb\x0a\xd5\xce\xfa\x22\x86\x05\xc3\x75\xc4\x1d\x93\x91\xde\x68\xbd\x48\xa1\xe2\x58\x6b\xc1\xb3\x40\xd7\x57\xf9\x91\x54\xe8\x37\x6a\x1c\x07\x43\x0e\x3b\x2a\x5b\x7e\x09\xf1\x55\x88\x57\x87\x87\xc5\xab\xdd\x5d\x78\x90\x3f\x81\x6d\x87\x9f\x3a\x12\x65\xd5\x07\xc2\x73\x92\x4f\x1c\xa4\xa1\x0e\xbf\xf3\xbd\x68\x76\x19\x5a\xba\x45\xbc\xa9\x03\xdd\x61\x98\xd1\x1a\xb4\xe8\x02\xdf\x74\x37\xd8\x32\xde\x62\x46\x5f\x03\x56\x55\xc6\xa0\x24\x53\xe3\x65\xbc\xcf\x19\xd3\x8a\xfd\x6e\x4a\x77\x51\x97\x05\x63\x54\xd9\xd2\xfd\x90\xb9\x22\x90\x7c\x6e\x51\xc4\xf8\xa5\xb7\x64\x42\x91\x8c\x7a\xea\x3c\x64\x14\x61\x66\xe6\x6c\x55\xd0\x94\xe5\x4f\x38\xef\xee\xde\x56\x3e\x61\x65\x8f\xf2\x27\xb7\xba\xec\xee\x6e\xb8\xdc\x1c\x72\x3c\x10\x67\x84\xa7\x83\x90\x19\x7d\xc3\xc0\x38\xf0\xf0\xc1\x4f\x43\x51\xb7\xf9\xb7\x6e\xf8\x31\xa8\x61\xb3\x7d\x9f\xb1\xd1\x91\xd8\x4a\x86\xdb\xec\xbf\x00\x00\x00\xff\xff\xc0\x9b\xe4\x6f\x43\x0b\x00\x00"

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

	info := bindataFileInfo{name: "data/common/dev/Vagrantfile.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"data/aws-simple/build/build-node.sh": dataAwsSimpleBuildBuildNodeSh,
	"data/aws-simple/build/template.json.tpl": dataAwsSimpleBuildTemplateJsonTpl,
	"data/aws-simple/deploy/main.tf.tpl": dataAwsSimpleDeployMainTfTpl,
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
		"aws-simple": &bintree{nil, map[string]*bintree{
			"build": &bintree{nil, map[string]*bintree{
				"build-node.sh": &bintree{dataAwsSimpleBuildBuildNodeSh, map[string]*bintree{
				}},
				"template.json.tpl": &bintree{dataAwsSimpleBuildTemplateJsonTpl, map[string]*bintree{
				}},
			}},
			"deploy": &bintree{nil, map[string]*bintree{
				"main.tf.tpl": &bintree{dataAwsSimpleDeployMainTfTpl, map[string]*bintree{
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

