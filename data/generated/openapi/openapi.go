// Code generated for package openapi by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../../openapi/openapi.dinosaurs.yaml
// ../../openapi/openapi.subscriptions.yaml
// ../../openapi/openapi.yaml
// ../../openapi/openapi2.yaml
// ../../openapi/subscriptions_openapi.yaml
package openapi

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

// Mode return file modify time
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

var _openapiDinosaursYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x5f\x6f\x1b\xb9\x11\x7f\xf7\xa7\x18\xe0\x5a\xe8\xee\x20\x59\x4e\x9b\x3e\x54\x88\x0f\x70\x2e\x4e\x9b\xc2\xb1\x53\xdb\x69\x1f\x8a\xc2\x1e\x91\xb3\x5a\xc6\xbb\xe4\x86\xc3\xb5\xb5\x69\xfb\xdd\x0b\x92\xfa\xb3\x2b\xad\xf5\x27\xb1\x50\xa5\xf0\x53\x2c\x72\x66\x38\x3f\xce\xcc\x8f\x43\x6e\x0a\x74\x29\x0f\x0e\x00\x7e\x80\xf3\xd3\xbf\xc3\xe9\xf9\x9b\x0f\x17\xef\xce\xaf\xe1\xea\xfa\xe4\xf2\xfa\x00\xa0\x8f\x85\xea\xdb\xb4\xe7\x2c\x8d\xfb\xf7\x2f\xfa\x52\x69\xc3\x58\xda\x36\x9d\xd3\xf3\x37\x07\x00\x00\x23\x72\x83\xf0\x07\x00\x97\x79\x8e\xb6\x1a\xc0\x25\xb9\xd2\x6a\x06\x84\x4c\xb1\x03\x93\xc0\xcc\xd2\x54\x94\x44\x69\x95\xab\xa6\xaa\x00\x3d\x78\x4d\x68\xc9\x0e\xe0\x1f\xff\x9c\x0c\x5a\xe2\xc2\x68\x26\x9e\x4b\x75\x7e\x77\x74\xd4\x99\xff\x04\x90\xc4\xc2\xaa\xc2\x29\xa3\x07\x70\x02\x7f\xb9\xba\x38\x07\xb4\x16\xab\xfa\xaa\x60\x86\x9f\x48\x38\xae\xe9\x09\xa3\x1d\x69\x57\x37\x05\x80\x45\x91\x29\x81\xde\x58\xff\x13\x1b\xdd\x9c\x05\x60\x91\x52\x8e\x8b\xa3\x00\xbf\xb1\x94\x0c\xa0\xf3\x43\x5f\x98\xbc\x30\x9a\xb4\xe3\x7e\x94\xe5\xfe\x9b\x89\x0f\x67\x8a\x5d\x67\x8e\xe3\xe5\xd1\x8b\x15\x38\x4a\x97\x82\x33\x77\xa4\x41\x31\x28\x7d\x8f\x99\x92\x3b\x76\xde\x14\xa4\xb1\x50\x87\x15\xe6\x59\x2b\x90\x53\x6b\x8d\x6d\x20\xf8\xfd\xe3\x08\x3e\x6a\x2c\x5d\x6a\xac\xfa\x42\x12\x9c\x81\x82\x6c\x62\x6c\x0e\xa6\x20\x1b\x5c\xdc\x37\x34\x7f\x58\x95\x57\x1f\x35\x8d\x0b\x12\x8e\x24\x90\xd7\x03\x23\x44\x69\x2d\xed\x57\x4c\x0a\xb4\x98\x93\x23\xcb\xf5\xaa\x6a\xcb\xcd\xb9\x64\xbf\xc0\x11\x75\x36\x17\x67\xf5\x65\x2b\x71\x42\x2b\xd2\x2d\x14\x8c\x95\x64\x5f\x57\x5b\x68\x24\x8a\x32\xc9\x51\xa1\x30\xbc\x4c\x46\xbf\x5a\x42\x47\x80\xa0\xe9\x61\xc6\x08\xdb\xd1\xd0\xe7\x92\xd8\xbd\x36\xb2\x26\xd7\xc8\x8f\x69\x8d\x83\x44\x87\x33\x11\xaf\xa7\x2c\xc9\x01\x38\x5b\xd2\xc1\x8a\x44\x59\x9d\x26\xed\x49\xb2\x09\xe7\x74\x56\x12\xe9\x0a\x02\x8a\x7b\xb6\xeb\xf4\xde\xc8\xf7\xc0\x34\x2b\x6a\xf3\x6f\x9e\x1b\x83\x0b\xb1\x36\x79\x3f\x8b\xf3\x99\xf2\xf7\x8e\xf2\x5f\x1e\xfd\xf1\x71\x34\xb3\x92\xc6\xcc\x12\xca\x0a\x68\xac\x78\xe7\x1d\xc4\xd3\x1e\x5b\x27\x1a\xca\xc7\x4e\x2e\x10\xbe\xc4\x95\x1e\x81\x4b\x69\x91\x16\xff\xf7\xf0\xb6\x6d\x50\xfb\xff\x52\xf2\x3f\x5b\x76\xa9\x7f\x22\x07\xa8\xe7\x4d\xe2\xb0\x82\x59\xc9\xed\xa6\x3f\x9d\x25\x55\x62\x4a\x2d\x1b\x0b\xee\x6a\xcb\x37\xe7\xd8\x67\x72\xda\x2f\x34\x2f\x1f\x47\x73\x6e\xe6\x59\xfb\xa0\x5c\x0a\x5c\x90\x50\x89\x22\x09\x4a\x7e\x8f\x4c\xf5\x3d\x34\xd8\x05\x3a\x91\x2e\x91\xc8\xc7\x42\x86\xee\x52\xef\xa8\xb5\x8c\xf6\xe5\x3c\xde\x7b\xd6\x62\x7e\xf0\xbb\x72\x19\x61\xac\x6e\x37\x37\xe1\xc5\x72\x82\x96\x4b\x21\x88\x39\x29\xb3\xac\xda\x1b\x82\x7c\x6e\x42\xf7\x04\xc1\x33\xcf\xef\x31\xcf\xff\xdf\x77\xd5\x4b\x67\x55\x20\x2d\xdf\x49\xef\x5d\x17\xdd\xf6\x26\xb4\xf6\x41\x45\xc9\xce\xc1\x7c\xc6\xab\x4d\xec\x46\x0b\xb1\xc1\xbe\xfa\xf5\xcf\xa7\xef\x4f\x66\x7d\x39\xcc\x22\xdb\x22\x34\xed\xc1\x01\x30\xcb\x2e\x92\x96\xc7\xa9\xb5\x88\x2e\xc2\xfb\xed\x25\x25\x64\x49\x8b\xc6\x0b\x94\xab\x0a\x1a\x4c\x1e\x78\x6b\x3b\x56\x58\x5f\xf2\x4e\xd5\x0f\xa1\xb0\xb7\xbe\x82\x16\x07\x61\x62\x85\x9d\x55\x7a\xd4\x98\x12\xf1\x25\xe4\x06\xdd\xc6\x2a\x00\x9e\x72\xd0\x0d\xfc\x59\x4d\x3d\xa7\x72\x6a\xcc\x4f\x0e\xb9\xa7\x30\xb9\x26\x18\x67\x6a\xfa\x1e\xf6\xd4\x01\x69\xbe\x66\x6f\x1d\x05\xe5\x28\x7f\x24\x06\xe1\xf1\x7e\x61\xa6\x55\x7c\x9b\x13\x7c\xcd\x3e\xd5\x9b\x98\x95\xfb\xd5\x02\xb3\x0d\x62\x4b\x92\x2d\x84\x76\xb9\x2e\x95\x9c\x8b\x6b\xcc\x69\x50\xbf\x1b\x2a\x3d\xf0\xfd\x67\xda\xde\x23\x5e\xa7\xe4\xcf\x03\x93\x80\x25\x61\xac\x5c\xd7\x1e\x2e\xf3\x4b\x4b\xe2\x15\x38\xa2\x45\x8f\xfc\x58\xc3\xa7\xcf\x25\xd9\xaa\xdd\xa9\x0f\x38\x22\xd0\x65\x3e\x24\x3b\xf7\x2c\x7e\x0f\x7a\x48\x49\x37\x06\x68\x2c\x88\x24\xd7\x8e\x37\xbf\x12\xb0\xfa\xb2\xde\x69\xa5\x1d\x8d\xc8\x36\xd8\x3a\xc1\x32\x73\x03\x78\x51\x1b\xcc\x95\x56\x79\x99\xd7\x07\xe7\xbb\x93\x60\xc6\xd3\x95\xfc\xa2\x8b\xb8\x1b\x8e\xac\xc1\xfd\x1e\xc7\x7e\xa1\x25\xe8\xec\xdb\x11\x1b\x3e\x8c\x7d\x13\xa6\xa3\xa3\x36\x54\x47\xeb\x50\x85\x87\xf8\x25\x5c\x61\x74\x05\xb2\x76\x63\x0b\x88\xff\xdd\xab\x79\x74\x35\x09\x21\x87\x77\xa5\xb8\x00\x08\xab\x1c\x59\x85\x87\x21\x55\xb9\xd2\x0e\xc7\x7e\x67\x5c\xaa\x78\x5e\x0a\xa0\xea\x8d\x00\xab\x5c\x65\x68\xfd\xae\xb9\x05\x25\x82\x9b\x87\x94\x2c\xdd\x80\xc8\xb0\x64\xf2\xa3\xa8\xe1\xea\xaf\x67\xc0\x0e\x1d\xe5\xa4\x5d\xb7\x66\xaa\xe4\xe9\x3b\x97\x87\xcd\x53\x23\xfe\xe8\x05\x74\xce\xaa\x61\xe9\x88\xa1\x0f\xc2\x64\x65\xae\x9b\x52\x28\x84\x29\xb5\x3b\xac\xd9\x7b\x6b\x2c\xd0\x18\xf3\x22\xa3\x2e\x28\x0d\xe1\xa3\xc5\x24\xbc\x56\xd1\x3d\x79\x5e\xad\x2b\x73\xec\xde\x10\x4a\x26\xeb\xad\xd7\x71\x3a\xb4\xa1\x7b\x08\x22\xb7\x79\x75\x3b\x38\xa8\x4d\xdf\xde\xde\xf2\xe7\xac\x01\x26\x9a\x80\x4c\xdd\x11\x74\xf2\xea\xb7\x9d\xa6\x78\x5d\xfb\x7a\x39\x06\x20\x50\x03\x66\x6c\x60\x48\xb1\x13\x21\x09\xc6\xd7\x63\x16\xae\x5e\x96\xd8\x94\x56\xd0\xd7\x01\xe6\x72\x38\xcb\x0c\x86\x0c\x87\x94\x51\x78\xe8\xba\x4d\x8c\x39\x1e\xa2\xbd\xed\xae\x44\x57\xd7\xbf\x09\xea\x7c\x78\x47\x15\x1c\x43\x27\x31\xa6\x03\xa8\x65\xab\xcc\x3d\x66\x25\x79\xa9\x21\xda\x15\xfb\xf1\x2e\xc6\xb4\x9e\x72\xba\xe3\x3c\x87\xdf\x2b\x49\xb2\x0b\xc6\x82\x8a\x32\xd1\xa2\x62\xa0\xbc\x70\x55\xd7\x8f\xd5\x2f\x0f\x4b\x11\x76\x29\xba\x30\xe2\x03\x04\x29\xb2\xbf\x7b\xe4\x8a\xd9\x5f\x0a\x9d\x01\x26\x82\x07\x95\x65\x30\xac\x47\x3f\x12\x02\xc9\xc3\xed\xc8\x79\xf2\x99\x6c\xb1\x9e\x27\xc3\x3b\x2c\xe8\x18\xf9\x61\xb5\x83\x92\x9e\x9a\xde\xb4\xaa\x87\xa5\xdb\xba\xb2\xbf\xb5\xa6\x67\xd1\x0e\xd3\x31\xb3\xa7\x05\xb9\x61\xd9\x22\x8b\xc7\xf3\xf3\xc2\x7e\xdd\xda\x70\x83\x5a\xde\x40\xa2\x2c\x3b\xd8\xce\x99\x6e\xd4\x3a\x5f\xeb\xdb\xd3\xd5\x8e\x36\x40\x63\x7f\x05\x52\x2e\x82\x89\xf4\x17\x6a\x63\x4a\x49\x5b\x16\x44\xfc\x0a\xbc\x58\x0f\x71\xf4\x69\xcb\xa1\x0c\xfe\x31\x20\x08\x93\xe7\xd8\x63\xf2\x3b\xe2\x99\x73\xfa\x7f\x5d\xe2\xaa\x3e\x7e\x43\x6a\x29\x70\x80\xb7\x51\xc0\x24\x9e\xca\x7a\xec\x6c\x29\x5c\x69\xbd\x4d\x1d\xba\xb8\xd0\x02\xb3\x8f\x12\xbc\x9a\xcd\xfe\x72\xf8\x2a\x18\xfe\x05\xb4\x71\xe1\xee\x58\x37\xf9\x8a\xdd\x54\xec\x67\xc8\x09\x35\x87\x9c\x09\x1a\xc1\x24\xcc\x0c\xd5\xb4\x4e\x63\xc2\x0f\x62\xf6\xa3\x48\xe1\xaa\xc6\xad\x1e\xc1\x88\x1c\x28\xd9\x85\xd4\x52\xd2\x85\x22\x43\xfd\xa3\x92\xc1\xcf\x3b\xa5\xe5\x4f\xe1\xaf\x48\xc1\xf0\xe3\x6c\x41\xfe\x69\x21\xfb\x6a\xbf\x8c\xc8\x83\xd1\xe6\x41\xd1\xeb\xcd\xd3\x2a\x9a\x38\x56\xb2\x1b\x16\xf5\x6b\x1e\x2a\x19\xff\xf5\x8b\x76\x27\x94\xff\x73\x53\x8b\x9c\x48\xcf\xc2\xcc\x71\xa3\xd9\x6d\x3a\xb0\x3a\x99\xfe\x1b\x00\x00\xff\xff\xe5\x5b\x31\xa9\x20\x25\x00\x00")

func openapiDinosaursYamlBytes() ([]byte, error) {
	return bindataRead(
		_openapiDinosaursYaml,
		"openapi.dinosaurs.yaml",
	)
}

func openapiDinosaursYaml() (*asset, error) {
	bytes, err := openapiDinosaursYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "openapi.dinosaurs.yaml", size: 9504, mode: os.FileMode(420), modTime: time.Unix(1700695476, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _openapiSubscriptionsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\xdd\x6e\x1b\xbb\x11\xbe\xf7\x53\x0c\x90\x16\x4a\x02\xd9\x72\xda\xf4\xa2\x42\x1c\xc0\x49\x9c\x36\x85\x63\xa7\xb6\xd3\x5e\x14\x85\x4d\x91\xb3\x5a\xc6\xbb\xe4\x86\x33\x6b\x6b\xd3\x9e\x77\x3f\x20\xa9\x9f\x5d\x69\x2d\x4b\x89\x85\xa3\xe0\xd8\x37\x96\xc8\x99\xe1\x7c\x9c\x99\x8f\xb3\x5c\x15\x82\x53\xea\xef\x00\x3c\x81\x93\xa3\x7f\xc3\xd1\xc9\xbb\x4f\xa7\x1f\x4e\x2e\xe0\xfc\xe2\xf0\xec\x62\x07\xa0\x27\x0a\xdd\x73\xe9\x2e\x3b\x1c\xf5\x6e\x5e\xf4\xa8\x1c\x90\x74\xba\x60\x6d\x4d\x9b\xde\xd1\xc9\xbb\x1d\x00\x80\x21\x72\x3f\x7c\x00\xa0\x32\xcf\x85\xab\xfa\x70\x86\x5c\x3a\x43\x20\x20\xd3\xc4\x60\x13\x68\x58\x9b\x88\xa3\x2c\x9d\xe6\x6a\xa2\x0e\xb0\x0b\x6f\x50\x38\x74\x7d\xf8\xcf\x7f\xc7\x83\x0e\xa9\xb0\x86\x90\x66\x52\x9d\x3f\xed\xef\x77\x66\x5f\x01\x14\x4e\x6d\xf7\xe1\x10\xfe\x71\x7e\x7a\x02\xc2\x39\x51\xcd\xaf\x0c\x76\xf0\x05\x25\x53\x4d\x57\x5a\xc3\x68\xb8\x6e\x0e\x40\x14\x45\xa6\xa5\xf0\x2a\xbd\x2f\x64\x4d\x73\x16\x80\x64\x8a\xb9\x98\x1f\x05\xf8\x83\xc3\xa4\x0f\x9d\x27\x3d\x69\xf3\xc2\x1a\x34\x4c\xbd\x28\x4b\xbd\xf3\x9a\x1f\xc7\x9a\xb8\x33\xc3\xf3\x72\xff\xc5\x12\x3c\x25\xa7\xc0\xf6\x1a\x0d\x68\x02\x6d\x6e\x44\xa6\xd5\x86\x01\xd8\x02\x8d\x28\xf4\x5e\x25\xf2\xac\x15\xcc\x91\x73\xd6\x35\x10\xfc\xf9\x6e\x04\x9f\x8d\x28\x39\xb5\x4e\x7f\x43\x05\x6c\xa1\x40\x97\x58\x97\x83\x2d\xd0\x05\x17\xb7\x0d\xcd\x5f\x96\xe5\xd7\x67\x83\xa3\x02\x25\xa3\x02\xf4\x7a\x60\xa5\x2c\x9d\xc3\xed\x8a\x49\x21\x9c\xc8\x91\xd1\x51\xbd\xba\xda\xf2\x73\x26\xd9\x2b\xc4\x10\x3b\xab\x8b\x93\xfe\xb6\x96\x38\x0a\x27\xd3\x35\x14\xac\x53\xe8\xde\x54\x6b\x68\x24\x1a\x33\x45\x51\xa1\xb0\xb4\x48\x4c\x6f\x1d\x0a\x46\x10\x60\xf0\xb6\xc1\x0c\xeb\x51\xd2\xd7\x12\x89\xdf\x58\x55\x93\x6b\xe4\x48\xbd\xd6\x41\x09\x16\x53\x31\xaf\xab\x1d\xaa\x3e\xb0\x2b\x71\x67\x49\xc2\x2c\x4f\x97\xf6\x64\x59\x95\x7f\x3a\x4b\xc9\x75\x09\x19\xc5\xfd\xdb\x74\xaa\xaf\xec\x7f\x60\x9e\x25\xb5\xfa\x2f\xcf\x95\xc1\x8d\x58\xab\xb4\x9d\xc5\xfa\x78\x04\x6c\xdd\x11\xf0\x72\xff\xaf\x77\xa3\x69\x94\xb7\xc8\x1c\x0a\x55\x01\x8e\x34\x6d\xbc\xb3\x78\xd8\xa3\xec\xd0\x40\x79\xd7\x69\x06\xd2\x97\xba\x36\x43\xe0\x14\xdb\xa8\xf2\xb7\x87\xf8\x3d\x8d\x6c\xef\x7f\x5a\xfd\xb2\x66\x37\xfb\x37\x64\x10\xa6\xd9\x48\x0e\x2a\x98\x96\xe0\x66\xfa\xd8\x46\x92\x25\xb6\x34\xaa\xb1\xe8\xa6\xb6\x7f\x3d\xee\x7d\x24\xad\xed\x42\xf3\xf2\x6e\x34\x27\xb6\x99\xc1\xb7\x9a\x53\xa0\x02\xa5\x4e\x34\x2a\xd0\xea\x67\x64\xb0\x9f\xa1\x19\x2f\x04\xcb\x74\x81\x54\x3e\x17\x2a\x74\xa2\x66\x83\x6d\x68\x5c\x43\x35\xe3\xbe\x85\xed\xe8\x27\xbf\x43\x67\x11\xce\xf2\xd6\x74\x55\xbe\x2c\xa7\xc8\xa5\x44\xa2\xa4\xcc\xb2\x6a\xab\x88\xf3\xb1\x69\xdd\x12\x04\x8f\xfc\xbf\xe5\xfc\xff\xbb\xe8\xc4\x17\xce\xb1\x40\x60\xbe\xfb\xde\xca\xce\xbb\xed\x7e\xe9\xde\xcb\x19\xad\x3a\x3b\xb3\x19\xaf\x36\xb6\x1b\x2d\xc4\x86\xfc\xfc\xed\xdf\x8f\x3e\x1e\x4e\x7b\x79\x68\x44\xb8\x45\x70\xd2\xb7\x03\x88\x2c\x3b\x4d\x5a\x2e\xbb\xee\x45\x75\x1a\xee\x84\xcf\x30\x41\x87\x46\x36\x6e\xb4\xb8\x2a\xb0\x3f\xbe\x34\xae\xed\x5a\xe1\x3c\x05\xb0\xae\x1f\x4e\xfe\x2f\xb1\x76\x7e\x73\xa3\x05\x62\xa7\xcd\x70\x55\x98\xc7\x7a\x72\x73\xf5\xd0\x50\x9b\xf7\xce\x6b\xe3\xd3\x8c\x39\xb5\x23\x0c\xd7\xed\x73\x33\xad\xe2\xeb\x9e\x97\x2b\xec\x57\xbd\x7d\x58\xba\x6f\x2d\x70\xdb\xa0\xce\x85\x71\x2e\x84\x8b\x99\xaf\xd5\x4c\xdc\x88\x1c\xfb\xf5\xa7\x34\x6d\xfa\xbe\x03\x4c\xdb\x3b\xb4\x8b\x14\x3d\xf3\xda\x04\x1c\x4a\xeb\xd4\x7d\x8d\xd9\x62\x05\x2f\x24\x98\xf7\x70\x88\xf3\x1e\xf9\xb1\x86\x4f\x5f\x4b\x74\x55\xbb\x53\x9f\xc4\x10\xc1\x94\xf9\x00\xdd\xcc\xb3\xf8\x26\xe7\x36\x45\xd3\x18\xc0\x91\x44\x54\x54\x3b\x48\xfc\x4a\x40\xfa\xdb\xfd\x4e\x6b\xc3\x38\x44\xd7\xe0\xc4\x44\x94\x19\xf7\xe1\x45\x6d\x30\xd7\x46\xe7\x65\x5e\x1f\x9c\xed\x4e\x22\x32\x9a\xac\xe4\x17\x9d\xc7\xdd\x70\xe4\x1e\xdc\x1f\xc5\xc8\x2f\xb4\x00\x9d\xfc\xe1\xef\xc2\x2b\xad\x1f\xc2\xb4\xbf\xdf\x86\x6a\xff\x3e\x54\xe1\xda\x7c\x01\x57\x18\x5d\x82\xac\xdd\xd8\x1c\xe2\xff\xef\xd6\x3c\x3a\x1f\x87\x90\xe2\x8d\x4f\x58\x00\xa4\xd3\x8c\x4e\x8b\xbd\x90\xaa\x54\x19\x16\x23\xbf\x33\x9c\x6a\x9a\x95\x02\xe8\xfa\x71\x4b\x3a\xd7\x99\x70\x7e\xd7\x78\x4e\x09\xe1\xf2\x36\x45\x87\x97\x20\x33\x51\x12\xfa\x51\x61\xe0\xfc\x9f\xc7\x40\x2c\x18\x73\x34\xdc\xad\x99\x2a\x69\x72\x03\xe5\x61\xd3\xc4\x88\x3f\xdc\x40\x30\x3b\x3d\x28\x19\x09\x7a\x20\x6d\x56\xe6\xa6\x29\x25\xa4\xb4\xa5\xe1\xbd\x9a\xbd\xf7\xd6\x01\x8e\x44\x5e\x64\xd8\x05\x6d\x20\xbc\x62\x18\x87\xd7\x69\xbc\x41\xcf\xad\x75\x65\x8a\x7d\x92\x80\x92\xd0\x79\xeb\x75\x9c\x2c\x5c\x38\xa3\x83\xc8\x55\x5e\x5d\xf5\x77\x6a\xd3\x57\x57\x57\xf4\x35\x6b\x80\x89\x26\x20\xd3\xd7\x08\x9d\xbc\xfa\x63\xa7\x29\x5e\xd7\xbe\x58\x8c\x01\x48\xe1\x9b\x1c\xb2\x30\xc0\x78\xd6\xa3\x02\xeb\xeb\x31\x0b\x0f\x3b\x0e\xc9\x96\x4e\xe2\xf7\x01\x6e\x5c\x93\x41\x26\x06\x98\x61\xb8\x72\xba\x4a\xac\x3d\x18\x08\x77\xd5\x5d\x8a\xae\xae\x7f\x19\xd4\x69\xef\x1a\x2b\x38\x80\x4e\x62\x6d\x07\x84\x51\xad\x32\x37\x22\x2b\xd1\x4b\x0d\x84\x5b\xb2\x1f\x1f\x62\x4c\xeb\x29\x67\x3a\xec\xf9\xfb\x46\x2b\x54\x5d\xb0\x0e\x74\x94\x89\x16\x35\x01\xe6\x05\x57\x5d\x3f\x56\x6f\x9d\x16\x22\xcc\xa9\xe0\x30\xe2\x03\x04\xa9\x20\xdf\xe9\xe7\x9a\xc8\x37\x95\x6c\x81\x10\xe1\x56\x67\x19\x0c\xea\xd1\x8f\x84\x80\x6a\x6f\x3d\x72\x1e\xbf\xd4\x9a\xaf\xe7\xf1\xf0\x06\x0b\x3a\x46\x7e\x50\x6d\xa0\xa4\x27\xa6\x57\xad\xea\x41\xc9\x6b\x57\xf6\x8f\xd6\xf4\x34\xda\x61\x3a\x66\xf6\xa4\x20\x57\x2c\x5b\x41\xf2\xee\xfc\x3c\x75\xdf\xb7\x36\x5c\x0a\xa3\x2e\x21\xd1\x8e\x18\xd6\x73\xa6\x1b\xb5\x4e\xee\xf5\xed\xe1\x6a\xc7\x58\xc0\x91\x7f\xc8\xd0\x1c\xc1\x44\xfa\x0b\xb5\x31\xa1\xa4\x35\x0b\x22\xbe\xb3\x9d\xaf\x87\x38\xfa\xb0\xe5\x50\x06\xff\x08\x04\x48\x9b\xe7\x62\x97\xd0\xef\x88\x67\xce\xc9\xaf\x54\xe2\xaa\x3e\x7e\x03\x6c\x29\x70\x80\xf7\x51\x20\xfe\xac\x64\x97\xd8\x95\x92\x4b\xe7\x6d\x9a\xd0\xc5\x85\x36\x98\x7c\x94\xe0\xd5\x74\xf6\xf5\xde\xab\x60\xf8\x35\x18\xcb\xe1\xe9\xac\x6e\xf2\x15\xf1\x44\xec\x39\xe4\x28\x0c\x85\x9c\x09\x1a\xc1\x24\x4c\x0d\xd5\xb4\x8e\x62\xc2\xf7\x63\xf6\x0b\x99\x36\x1f\x85\xd9\xc2\x10\x19\xb4\xea\x42\xea\x30\xe9\x42\x91\x09\xf3\x54\xab\xe0\xe7\xb5\x36\xea\x59\xf8\x14\x29\x18\x9e\x4e\x17\xa4\x67\x73\xd9\x57\xfb\x66\x65\x1e\x8c\x36\x0f\x8a\xdd\xdd\x59\x5a\x45\x13\x07\x5a\x75\xc3\xa2\x7e\xcd\x3d\xad\xe2\x7f\xbf\x68\x77\x4c\xf9\xcf\x9b\x5a\xc8\x32\x3d\x0e\x33\x07\x8d\x66\xb7\xe9\xc0\xf2\x64\xfa\x35\x00\x00\xff\xff\x98\xf1\x39\x1a\xde\x24\x00\x00")

func openapiSubscriptionsYamlBytes() ([]byte, error) {
	return bindataRead(
		_openapiSubscriptionsYaml,
		"openapi.subscriptions.yaml",
	)
}

func openapiSubscriptionsYaml() (*asset, error) {
	bytes, err := openapiSubscriptionsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "openapi.subscriptions.yaml", size: 9438, mode: os.FileMode(420), modTime: time.Unix(1701101775, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _openapiYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\xdd\x6f\x1b\x37\x0c\x7f\xf7\x5f\x41\xa0\x1b\xdc\x16\xf6\xd9\xc6\x5e\x06\xa3\x29\x90\xae\x29\xda\xa1\x6d\xb2\x3a\x43\x1f\x13\x5a\xc7\xf3\xa9\xbd\x93\x2e\x12\x2f\x89\xf7\x91\xbf\x7d\x90\x74\xbe\x0f\x7f\xc5\x49\xb3\xbc\xc4\xa2\xc8\x1f\xc9\x9f\x48\x4a\xa7\x0b\x52\x58\xc8\x29\xfc\x12\x8d\xa3\x71\x4f\xaa\x44\x4f\x7b\x00\x2c\x39\xa3\x29\x98\x74\xc8\x86\x6e\x61\x46\xe6\x5a\x0a\x82\xe3\xb3\x0f\x3d\x80\x98\xac\x30\xb2\x60\xa9\xd5\x2e\x95\x6b\x32\xd6\x6f\x8f\xa3\x71\x34\xe9\x59\x32\x4e\xe2\x90\x87\x50\x9a\x6c\x0a\x29\x73\x31\x1d\x8d\x32\x2d\x30\x4b\xb5\xe5\xe9\xaf\xe3\xf1\xb8\x07\xb0\x86\x2e\x4a\x63\x48\x31\xc4\x3a\x47\xa9\xba\xe6\x76\x3a\x1a\x61\x21\x23\x97\x82\x4d\x65\xc2\x91\xd0\xf9\x26\xc4\x27\x94\x0a\x9e\x17\x46\xc7\xa5\x70\x92\x17\x10\xa2\xd9\x0e\x66\x19\x17\x74\x1f\xe4\x8c\x71\x21\xd5\x62\x05\x54\x20\xa7\x3e\x37\x87\x30\xaa\x08\x19\x5d\x4f\x46\xb1\x54\xda\x62\x19\x12\x07\xf8\xc9\x50\x32\x85\x7e\x45\x79\x54\xef\x46\x4b\xcc\xb3\x67\x23\x0f\x33\xba\x9b\x60\x21\xef\x26\x15\xca\xdd\xe4\x7a\x72\x37\xa9\x35\xfb\xfb\x9c\x8c\xfe\x96\xf1\xbf\x4f\xe5\xe9\x6e\xe2\xd0\xb6\xf9\xb3\xe5\xbc\xa6\x62\x7b\x62\x1d\x8d\x7b\x5d\x76\xb4\xef\x75\xb8\x3b\xc9\x1f\xf0\xda\x24\xfb\x0c\x8e\xff\x3c\x3f\x1d\x1e\xbf\x7d\x0b\x9f\x4f\xbe\xc2\xd9\xf1\xf9\xfb\x59\x4f\xe8\xbc\xd0\x8a\x14\xfb\x74\x2d\x89\xd2\x48\x5e\xce\x44\x4a\x39\x55\x0c\xbc\x21\x34\x64\xc2\x6f\x00\x5e\x16\x14\xca\xaa\x12\x58\xaf\x3b\x85\xb9\x57\xab\x84\x61\xf1\x4e\x9b\x1c\x79\x0a\xbf\x7f\x3d\xef\x55\x8a\x58\x81\x9e\xce\xbf\x91\xe0\x2f\x94\x90\x21\x25\xa8\x8b\xae\xfd\x66\x25\x2a\x8c\x2e\xc8\xb0\x5c\x85\xe3\xfe\x64\xdc\xfc\x5e\x19\x59\x36\x52\x2d\x6a\xf1\x77\xa9\xee\x57\x4a\x1d\xcd\xfb\x94\x3e\x4a\xcb\x0f\x8c\xed\x20\xc7\x05\x2e\x68\x53\x49\x2a\xa6\x45\xcd\x21\x80\x95\x7f\x1d\xa0\xc5\x9a\x31\xbb\x4f\xcd\xd0\x55\x29\x0d\xb5\x22\x1b\xfa\x48\x5b\x4b\x17\x53\x6b\xe9\x9c\xb7\x96\xde\x4b\x6b\x2d\x99\x72\xeb\xd7\x27\xc6\xe8\xba\x3c\x30\xcb\x4e\x93\xb6\x93\xaa\x92\x9f\x8d\x9a\x4a\x1b\x55\x95\x30\x5a\x2b\x82\x7e\xdb\xdd\x26\xd9\xbb\x08\x77\x7f\x42\xc7\xd4\x95\xec\x20\x3e\x70\x81\x56\xab\x83\xd5\x9d\x47\x74\xbd\x74\xd1\x2d\xbb\xad\x46\x9e\x8c\x76\xd5\x3c\x88\x10\x67\xf8\x03\x2c\xf8\x33\xd9\x1e\x22\x1a\x83\xcb\xb5\x9d\xad\xea\xb0\x37\x40\x9f\x5e\x88\xf0\x6d\x35\x4e\x57\x00\xf7\xcc\xe5\x2d\x60\x2b\x84\x2e\x5e\x9b\xbc\xc7\x63\x36\x4c\xae\x24\x67\xc8\x22\xfd\x42\x57\x25\x3d\x05\x7e\x1b\x2d\xf8\x99\xb5\xa6\xee\x0e\xfc\x6d\x63\x7c\x8b\x8f\x36\xd2\x26\xf6\x1e\x7e\x1e\x81\xdf\xf0\xd4\x96\x1e\xc0\xd5\x23\x7c\x6d\x72\xb6\x76\x25\xcd\x7e\x7b\x7f\xf2\xe9\x78\xd6\x73\xf3\xd1\x60\x4e\x4c\xab\xe7\x45\xd3\x79\x0a\xdd\x65\x23\x57\x83\x4b\xaa\x29\xb8\x9b\xb0\x5a\x76\x9e\x32\xe7\x29\x81\x8c\x41\x27\x60\x48\x68\x13\xaf\x8f\x42\x60\x53\x52\xfb\x16\xc3\xa6\x19\x36\x1a\xbb\x3d\xb1\x43\x0c\xad\x79\xe9\xa2\xb8\x2a\xc9\x2c\xb7\x85\x71\x86\x0b\x02\x55\xe6\x73\x32\x4d\x2c\x90\x49\xcb\x70\x93\x92\xea\x08\xe8\x56\x10\xc5\x16\x6c\x41\x42\x26\x92\x62\xef\xa5\x3d\x8b\xb7\x07\xba\x7e\x27\xc4\x94\x60\x99\xf1\x14\x26\xb5\x28\x97\x4a\xe6\x65\xde\x88\x1a\x1e\x12\xcc\x6c\xc0\x6f\xdf\x38\x21\xcb\x96\xeb\xbd\x59\x7e\xc2\x5b\x07\xbf\x91\xa8\x05\xd6\x60\x88\x4b\xa3\x1e\x99\x41\xf5\x76\xee\xe4\x30\xde\x97\x03\xa1\x11\xe9\x5a\x16\x5e\xb6\x23\x8f\x6d\x20\x6b\xd9\xfd\x33\xac\x63\x98\x55\x47\x63\x81\x53\xaa\x80\x41\x18\xc9\x64\x24\x46\xbe\xe8\xec\x52\x31\xde\x3a\x0e\x38\x95\xb6\x29\x66\x90\xb6\x75\xb7\xe7\x32\x43\xe3\xd8\xe1\x35\x13\x82\x8b\x9b\x94\x0c\x5d\x80\xc8\xb0\xb4\xe4\xa4\xa8\x60\xf6\xc7\x47\xb0\x8c\x4c\x39\x29\x1e\xd4\x40\xa5\x75\xcf\x75\x67\xe5\x52\xb5\x2b\x88\x6f\x56\x2b\x40\x66\x23\xe7\x25\x93\x85\x11\x08\x9d\x95\xb9\xea\x6a\xa1\x10\xba\x54\x1c\x41\x0d\xf7\x4e\x1b\xa0\x5b\xcc\x8b\x8c\x06\x20\x15\x68\x13\x93\xa9\xce\xd0\x48\xba\x26\x77\xa3\xb5\x6d\x2d\xdc\x48\x4e\x01\xa1\xb4\x64\x1c\x78\x93\x22\xa3\x61\x17\x9c\x57\xb8\xcc\x97\x97\xd3\x5e\xbd\x79\x79\x79\x69\xaf\xb2\x56\x16\xc1\x18\x32\xf9\x9d\xa0\x9f\x2f\x7f\xee\xb7\x55\x1b\xbb\xf3\x4d\xd2\x41\xa0\x02\xcc\xac\x86\x39\x01\x16\x45\xe6\xfa\x46\xbb\xc6\xca\x90\x29\x06\x43\x56\x97\x46\x50\xf4\x88\x24\x3b\x63\x0e\x32\x9c\x53\x46\x31\xcc\x97\x70\x99\x68\x7d\x34\x47\x73\x39\xd8\x99\x53\xdb\xf6\xc2\x9b\xda\xe8\x3b\x2d\xe1\x08\xfa\x89\xd6\x7d\x40\x15\x6f\xd5\xb9\xc6\xac\x24\xa7\x35\x47\xb3\x83\x85\x0f\xe1\xf8\xda\x95\xa5\xfa\xec\x1e\x06\xd7\x32\xa6\x78\x00\xda\x80\x0c\x3a\x01\x4d\x5a\xa0\xbc\xe0\xe5\xc0\xc9\x54\x8d\xb3\x71\x96\x9c\x22\x7b\x89\x3b\x10\x48\xd1\x42\x41\x26\x97\xd6\x7d\xf1\x3a\x82\x2c\x11\xdc\xc8\x2c\x83\x79\x73\xce\xa1\xbb\x29\x8e\x0e\x9d\xa5\x9e\xee\x37\xcb\x6e\x8b\x56\xc2\xff\xa1\x47\xc3\xe9\xce\x97\x4f\xde\xa5\x2b\xe0\xc3\x1a\x75\x5e\xf2\x83\x9b\x75\xad\x4d\x1f\x58\xc0\xf5\xa9\xfa\xed\x50\xb7\xab\x46\x3b\xa0\x15\xd1\x8a\xed\xd5\x77\x6a\x1e\xe7\x13\x2e\x50\xc5\x17\x90\x48\x63\x19\x0e\x0f\x62\x10\x2c\x3e\xef\x8d\xe9\xa9\x3a\x42\x69\xa0\xdb\x22\x93\x42\x72\x48\x21\x0c\x30\x5f\xf1\xab\xe1\x72\x70\xa1\x27\x92\xb2\xd8\x76\xeb\x3c\xc8\x9e\xa6\xcc\x4b\x1f\x8f\x05\x04\xa1\xf3\x1c\x87\x96\x5c\xfe\x6e\xe6\xf9\xd7\x84\x4e\x2a\x6f\xee\x94\xe6\xb4\xd1\xa8\x00\xef\xc2\xb6\x4e\xdc\x20\x1a\x5a\x36\xa5\xe0\xd2\x38\x44\xe5\x1f\x4e\xfe\xb3\xc1\xba\xd3\x80\x57\xf5\xee\xeb\xe8\x95\x87\x7d\x0d\x4a\xb3\xff\x32\x6a\x00\x5f\x59\x5e\x29\xbd\x84\x9c\x50\x59\x5f\x15\x5e\xdf\x03\x42\x0d\x53\xdb\x9c\x84\x42\x9e\x86\xaa\x46\x91\x76\xde\xa2\x2e\xf6\x05\x31\xc8\x78\xe0\xbf\xd7\x07\x50\x64\xa8\x9e\xcb\xd8\xc7\xe8\xbe\x61\x5f\xf8\x5f\x61\x78\xc2\xf3\xda\x9d\x7d\xd1\xa9\xae\xfa\xb7\x16\xb9\x07\xec\x8e\xf6\xe1\xb0\x29\x9d\x60\x7e\x24\xe3\x81\x77\xe8\xfc\x45\x32\x0e\xff\x9d\xc3\x41\x35\xa8\x5f\x76\xad\x88\x45\xfa\xd1\xef\x1c\xb5\x5e\x96\x6d\xe7\x7b\x0b\xe6\xbf\x00\x00\x00\xff\xff\xff\x3d\x34\xf6\xaf\x14\x00\x00")

func openapiYamlBytes() ([]byte, error) {
	return bindataRead(
		_openapiYaml,
		"openapi.yaml",
	)
}

func openapiYaml() (*asset, error) {
	bytes, err := openapiYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "openapi.yaml", size: 5295, mode: os.FileMode(420), modTime: time.Unix(1700695607, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _openapi2Yaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x6d\x8f\x1b\xb7\x11\xfe\xae\x5f\x31\x40\x5a\xc8\x0e\xf4\xe6\x26\x05\x5a\xc1\x0e\x70\x8e\xed\x26\x81\xed\x73\xad\x73\xf3\xa1\x28\xee\x46\xe4\xac\x96\xf1\x2e\xb9\x26\x67\xef\x4e\x6e\xfb\xdf\x0b\x92\xab\x7d\x91\x56\xb2\xce\xb5\x73\x97\xc0\x9f\x4e\xcb\x9d\x19\xce\x33\x9c\x79\x38\xe4\x9e\x29\x48\x63\xa1\xe6\xf0\xcd\x64\x36\x99\x0d\x94\x4e\xcc\x7c\x00\xc0\x8a\x33\x9a\x83\x4d\xc7\x6c\xe9\x1a\x16\x64\x2f\x95\x20\x38\x79\xf5\xe3\x00\x40\x92\x13\x56\x15\xac\x8c\xde\x27\x72\x49\xd6\x85\xd7\xb3\xc9\x6c\xf2\x60\xe0\xc8\xfa\x11\x6f\x79\x0c\xa5\xcd\xe6\x90\x32\x17\xf3\xe9\x34\x33\x02\xb3\xd4\x38\x9e\xff\x65\x36\x9b\x0d\x00\xb6\xac\x8b\xd2\x5a\xd2\x0c\xd2\xe4\xa8\x74\x57\xdd\xcd\xa7\x53\x2c\xd4\xc4\x43\x70\xa9\x4a\x78\x22\x4c\xbe\x6b\xe2\x05\x2a\x0d\xf7\x0a\x6b\x64\x29\xfc\xc8\x7d\x88\xde\xf4\x1b\x73\x8c\x2b\xfa\x90\xc9\x05\xe3\x4a\xe9\xd5\xc6\x50\x81\x9c\x06\x6c\xde\xc2\xb4\x0a\xc8\xf4\xf2\xc1\x54\x2a\x6d\x1c\x96\x11\x38\xc0\x8a\x38\xfe\x00\x70\x65\x9e\xa3\x5d\xcf\xe1\x35\x71\x69\xb5\x03\x84\x4c\x39\x06\x93\x40\xad\xb4\x11\x25\x51\x5a\xc5\xeb\x8d\xaa\x77\xfb\x31\xa1\x25\x3b\x87\x7f\xfe\xab\x1a\xb4\xe4\x0a\xa3\x1d\xb9\x46\x6a\xf8\xa7\xd9\x6c\xd8\x3c\x6e\x41\x38\x81\x9f\x16\xa7\x2f\x01\xad\xc5\x75\x7b\x56\x30\xcb\x5f\x48\xb0\x6b\xe9\x09\xa3\x99\x34\xb7\x4d\x01\x60\x51\x64\x4a\xa0\x37\x36\xfd\xc5\x19\xdd\x7d\x0b\xe0\x44\x4a\x39\x6e\x8f\x02\xfc\xc1\x52\x32\x87\xe1\x57\x53\x61\xf2\xc2\x68\xd2\xec\xa6\x51\xd6\x4d\x9f\x54\x3e\x3c\x57\x8e\x87\x0d\x8e\x6f\x67\x0f\x0e\xe0\x28\x39\x05\x36\x6f\x49\x83\x72\xa0\xf4\x25\x66\x4a\xde\x86\xf3\x4f\xad\x35\xb6\xe3\xf5\x37\xfb\xbd\x7e\xa3\xb1\xe4\xd4\x58\xf5\x9e\x24\xb0\x81\x82\x6c\x62\x6c\x0e\xa6\x20\x1b\xdc\xba\x0b\x08\xfe\x7c\x28\x7f\xde\x68\xba\x2e\x48\x30\x49\x20\xaf\x07\x46\x84\x5a\xbd\xfd\xd8\x17\x68\x31\x27\x26\x5b\xd7\xc2\xb8\x57\xb9\x91\x9b\x16\xb8\xa2\xe1\xb1\xc2\x4e\xbd\xbf\x81\x30\xa1\x15\xe9\xd1\xe2\xc6\x4a\xb2\x8f\xd7\x47\xcb\x27\x8a\x32\xe9\xa2\x78\xe1\x59\x74\x9b\x5e\xbe\xb7\x84\x4c\x80\xa0\xe9\xaa\xae\xf1\x9b\x11\xcb\xbb\x92\x1c\x3f\x36\xb2\x25\xd7\xc9\x84\x4d\xd5\x82\x44\xc6\x5a\xc4\xeb\x29\x4b\x72\x0e\x6c\x4b\x1a\x1c\x48\x89\xc3\x09\xd1\x9f\x0e\xc7\xb0\xc8\xf0\x20\x35\x1e\xa0\x94\x18\xb3\x5b\x49\xe4\x6d\xdf\x03\x8f\x1c\xa8\xc2\x7f\x78\xb6\x0b\x2e\xc4\x2a\x74\x77\xa7\x0c\xbf\x10\xf7\x2d\x22\xf8\xeb\x7e\x04\x75\xb9\x62\x66\x09\xe5\x1a\xe8\x5a\xb9\xdb\xd9\xef\x6f\xb4\xe1\x9c\x68\x28\xf7\xed\x39\x20\x7c\xc9\xfa\x8e\x8c\x53\xda\xa6\xb9\xdb\x81\xb4\xb7\x15\x9c\xfe\x5b\xc9\xff\xee\xef\x07\xff\x46\x0c\xa8\x9b\x76\x6c\xb9\x86\xba\x2c\x3e\x4f\x27\x58\x27\x44\x62\x4a\x2d\x3b\x13\xfe\xaa\xa1\xeb\xe5\xbe\x2f\x04\x72\x3b\x08\xbe\xdd\x8f\xe0\xa5\x69\xb2\xf3\x4a\x71\x0a\xae\x20\xa1\x12\x45\x12\x94\xfc\xad\xb0\xc9\x5d\x6d\x5f\x0b\x64\x91\xee\x90\xc2\x9b\x42\x86\x2e\x4e\x7f\xa6\x16\x2e\xda\x97\xcd\xba\xde\xb1\x56\xee\x95\x8f\xca\xeb\x08\xe3\x70\x5b\x77\x0c\xcf\x95\x15\x5a\x57\x0a\x41\xce\x25\x65\x96\xad\xef\x0c\xe1\x7d\x69\xf6\x7e\x65\xaf\xbf\x70\xf5\x9d\x00\xf1\x3b\xec\x58\x77\xf6\x98\x40\x3c\xbe\x4b\xbd\x13\x1d\xea\xee\x4d\xc9\x07\xaf\x1b\x94\xf4\x8a\x5f\xc1\xc9\x9b\xb3\xd3\xf1\xc9\x93\x27\xf0\xf2\xe9\xcf\xf0\xea\xe4\xec\x87\xc5\xa0\x51\xf0\xb6\x36\xfb\xd2\xc2\x4f\xba\xa1\xe8\x6a\x63\xaa\xbc\xe5\x75\x41\xf1\xea\x75\xd0\x02\x43\x73\x58\x06\xb1\x6a\x30\x3e\x3c\x33\x36\x47\x9e\xc3\x4f\x3f\x9f\x0d\x36\xa8\x2b\xa3\xa7\xe1\xba\xf2\x35\x25\x64\x49\x0b\xea\x5a\x8f\x77\x99\xd5\x50\x61\x7d\xdd\xb2\x6a\xef\x18\x4a\xb6\x83\x17\x95\x1c\x5b\xa5\x57\xf5\xf0\x5b\xa5\x3f\x2c\x94\xfa\xb8\x1d\x12\x7a\xae\x9a\xcb\x99\x23\x7d\x3b\x6a\xe2\x02\x57\xb4\x2b\xa4\x34\xd3\x8a\x9a\xf4\x72\xea\xfd\x11\x52\x6c\x18\xb3\x0f\x89\xd5\xcd\x40\xab\xe3\xf0\x9e\xb6\x1e\xbd\x4f\xad\x47\x3f\x79\xeb\x31\xcc\xd2\x7a\x56\x4c\x79\xac\xe5\x90\x99\x1b\xbb\x98\x65\xa7\xc9\xe1\x0b\xbc\x4d\x46\x6f\xa5\x40\x73\x79\xd6\x13\xe8\xfe\x50\xfb\xf2\x93\xd4\xad\xa3\xde\x70\x7b\xfc\xb8\x53\x88\x7b\x44\xeb\x6d\xe2\xbc\x9b\x66\x3d\x0a\x01\x7a\x3b\x47\x6e\x00\xbf\x7d\x5b\x7e\x23\xcc\x21\xf2\x7d\x8e\x85\x8f\x02\x9d\xf1\x1e\xd1\xa3\x59\x66\x43\xdf\xb7\xb4\xb2\x61\xc7\xa3\x5e\x9c\x3b\x2b\x26\xe2\xcd\xdf\x39\xf2\x51\xe2\x00\x49\x45\x4c\xbe\xa1\x1c\xb3\xca\xa9\xf5\xb6\x6a\x33\xff\x5f\x63\xed\x8f\x22\xbf\x97\xec\xe8\x76\xbe\x7d\x5d\xfe\x0d\xb9\xb2\x67\x91\x77\xc2\xbc\xb5\x6d\x2d\xbe\xff\xe1\xe9\x8b\x93\xc5\x60\x77\x1b\x6c\xaa\x55\xa3\xdf\x90\xea\x4e\x55\xe9\xb9\x3f\xa0\xa5\xd5\x63\x67\xb3\x3f\x4b\xc9\xb7\x54\x26\x01\x4b\xc2\x58\xb9\x4d\x97\xed\xb3\xd3\xf6\xb6\xbd\xe3\x6a\x9b\xd5\xa3\x0f\x2d\x4e\xf5\x5e\xbc\x2b\xc9\xae\xfb\xdc\x78\x85\x2b\x02\x5d\xe6\x4b\xb2\x8d\x2f\xf1\xd3\xe2\x55\x4a\xba\x33\x40\xd7\x82\x48\xba\x56\x4f\xe8\x67\x69\xf3\x75\xbf\xa3\xdb\xfb\x86\xa4\x04\xcb\x8c\xe7\xf0\xa0\x1e\xca\x95\x56\x79\x99\x37\x43\x4d\x1c\x12\xcc\x5c\xb4\xdf\xde\x95\x22\xca\xd6\xd4\x07\x51\xbe\xc0\x6b\x6f\x7e\x07\xa8\xf3\x5d\xba\x0d\x5f\x54\x3f\x12\x41\xf5\x0d\xba\x83\x61\x76\x08\x43\xf8\xbe\xb3\x85\x22\x8c\xed\xc1\xd1\x67\x64\x0b\xdd\x7f\xc6\xb5\x0f\x8b\x6a\x69\x5c\xb8\xd8\x8c\x86\x41\x58\xc5\x64\x15\x4e\x42\xd2\xb9\xb5\x66\xbc\xf6\x31\xe0\x54\xb9\x26\x99\x41\xb9\xd6\xfe\x9f\xab\x0c\xad\x8f\x0e\x6f\xa9\x10\x9c\x5f\xa5\x64\xe9\x1c\x44\x86\xa5\x23\x3f\x8a\x1a\x16\x7f\x7f\x0e\x8e\x91\x29\x27\xcd\xa3\xda\x50\xe9\x36\x97\xac\x1e\xaa\xdb\x98\xf0\xbd\x29\x20\xb3\x55\xcb\x92\xc9\xc1\x14\x84\xc9\xca\x5c\x77\xa5\x50\x08\x53\x6a\x9e\x40\x6d\xee\x99\xb1\x40\xd7\x98\x17\x19\x8d\x40\x69\x08\x9f\xbf\xaa\x35\xb4\x8a\x2e\xc9\x33\x5d\x5b\xd7\xc5\x53\x0c\x42\xe9\xc8\x7a\xe3\x0d\x44\x46\x1b\x7a\xeb\x20\x70\x91\xaf\x2f\xe6\x83\xfa\xe5\xc5\xc5\x85\x7b\x97\xb5\x50\x44\x65\xc8\xd4\x5b\x82\x61\xbe\xfe\xe3\xb0\x2d\xda\xe8\x9d\xed\x06\x1d\x04\x6a\xc0\xcc\x19\x58\x52\xec\xcf\x49\x82\xf1\x85\x95\x85\x4b\x05\x4b\xce\x94\x56\xd0\xe4\x23\x40\xba\x72\x59\xa7\x81\x83\x0c\x97\x94\x51\xb8\x90\xbd\x48\x8c\x79\xb4\x44\x7b\x31\xda\x8b\xa9\xad\x7b\x1e\x54\xdd\xe4\x2d\xad\xe1\x11\x0c\x13\x63\x86\x80\x5a\xf6\xca\x5c\x62\x56\x92\x97\x5a\xa2\xdd\x13\x85\x1f\xe3\xf2\xb5\x33\x4b\x0f\xd9\xf3\xf0\xa5\x92\x24\x47\x60\x2c\xa8\x28\x13\xad\x29\x07\x94\x17\xbc\x1e\xf9\xb1\xe6\x90\xbc\xb3\x96\x9c\x22\x87\x11\xbf\x20\x90\xa2\xf3\x27\xec\x5c\x39\xa7\x8c\xf6\x01\x72\x44\x70\xa5\xb2\x0c\x96\xcd\x3a\xc7\xea\x26\x39\x39\x96\x4b\xab\x4f\xaa\xdd\x12\xad\x06\x3f\x43\x8d\xc6\xd5\x5d\xae\x3f\x79\x95\x6e\x0c\x1f\x57\xa8\xcb\x92\x6f\x5c\xac\x5b\x65\x7a\xc3\x04\xae\x57\x35\xbc\x8e\x79\xbb\x29\xb4\x23\x4a\x11\x9d\xe8\xcf\xbe\x53\xfb\x71\x73\xc2\x39\x6a\x79\x0e\x89\xb2\x8e\xe1\x78\x27\x46\x51\xe3\xe5\x41\x9f\x3e\x55\x45\x68\x03\x74\xed\x8f\xf8\x8a\x23\x84\x48\x60\x21\xe3\x37\xe4\x72\x74\xa2\xc7\xff\x05\xe8\xe6\x79\x1c\xfb\x34\x69\x5e\x06\x7f\x1c\x20\x08\x93\xe7\x38\x76\xe4\xf1\x7b\xce\xdb\xfc\xe7\x52\x9c\xcd\xaf\xd2\x92\x76\x0a\x15\xe0\x59\x7c\x6d\x12\x4f\x44\x63\xc7\xb6\x14\x5c\x5a\x6f\x51\x87\xc6\x29\x34\x96\xce\xaf\x06\x3c\xac\xdf\x7e\x37\x79\x18\xcc\x7e\x07\xda\x70\x38\x4d\x35\x06\x1f\x3a\xde\x08\x7d\x0d\x39\xa1\x76\x21\x2b\x82\x7c\x30\x08\xb5\x99\x5a\xe7\x69\x4c\xe4\x79\xcc\x6a\x14\x29\x2c\x5a\xac\xe8\x7d\x5f\x11\x83\x92\xa3\x70\xa6\x1f\x41\x91\xa1\xbe\xa7\x64\xf0\xd1\x9f\x73\xef\x87\x5f\x91\x3c\xe1\x5e\x3d\x9d\xbb\xdf\xc9\xae\xfa\xb7\x11\x79\x30\xd8\xa5\xf6\xf1\xb8\x49\x9d\xa8\xfe\x48\xc9\x51\x98\xd0\xcf\x37\x51\x32\xfe\xf5\x13\x8e\x2a\xa2\xfe\xba\xab\x45\x2c\xd2\xe7\xe1\xcd\xa3\xce\xad\x7c\x33\xf9\xc1\x84\xf9\x5f\x00\x00\x00\xff\xff\xde\xca\x72\xd9\xf7\x27\x00\x00")

func openapi2YamlBytes() ([]byte, error) {
	return bindataRead(
		_openapi2Yaml,
		"openapi2.yaml",
	)
}

func openapi2Yaml() (*asset, error) {
	bytes, err := openapi2YamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "openapi2.yaml", size: 10231, mode: os.FileMode(420), modTime: time.Unix(1700695467, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _subscriptions_openapiYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x4d\x4f\xe3\x48\x10\xbd\xe7\x57\x94\xc4\x4a\x39\xb1\x0e\xbb\xec\x61\x7d\x0b\x10\xed\x87\x98\x04\x05\x98\x39\x8c\xe6\xd0\x74\x97\xe3\x66\xec\xee\x9e\xea\x36\x60\x46\xf3\xdf\x47\xb6\xc9\x47\x07\xc7\xc4\x23\xa2\x31\x9a\x70\x22\x76\x55\x75\xbf\x72\xd5\xab\xa7\x32\xcc\xc5\x36\xec\x01\x1c\xc0\x78\xf4\x01\x46\xe3\xb3\x8b\xc9\x7f\xe3\x2b\xb8\xbc\x1a\x4e\xaf\x7a\x00\x01\x33\x32\xa0\xf8\xd0\x11\x3e\x04\x77\x47\x81\xcd\x6e\x2c\x27\x69\x9c\xd4\xaa\xce\x6f\x34\x3e\xeb\x01\x00\xcc\xd0\x85\xe5\x3f\x00\x36\x4b\x53\x46\x79\x08\x53\x74\x19\x29\x0b\x0c\x12\x69\x1d\xe8\x08\xbc\x68\x73\x73\xe4\x19\x49\x97\xcf\xdd\x01\x0e\xe1\x04\x19\x21\x85\xf0\xf1\xd3\xd3\x43\x42\x6b\xb4\xb2\x68\x97\x56\xfd\x3f\x06\x83\xfe\xf2\x27\x80\xc0\x45\xec\x10\x86\xf0\xff\xe5\x64\x0c\x8c\x88\xe5\xeb\x27\x83\xbe\xb9\x45\xee\xec\x8a\x2f\xd7\xca\xa1\x72\xab\xe1\x00\x98\x31\x89\xe4\xac\x70\x09\x6e\xad\x56\xfe\x5b\x00\xcb\x63\x4c\xd9\xfa\x53\x80\xdf\x08\xa3\x10\xfa\x07\x01\xd7\xa9\xd1\x0a\x95\xb3\x41\x65\x6b\xbd\x7c\x9e\x4b\xeb\xfa\x4b\x3c\xc7\x83\xa3\x06\x3c\x99\x8b\xc1\xe9\xcf\xa8\x40\x5a\x90\xea\x8e\x25\x52\xec\x18\x80\x36\xa8\x98\x91\xbf\xe7\x2c\x4d\x6a\xc1\x8c\x88\x34\x79\x08\xfe\xdc\x8c\xe0\x5a\xb1\xcc\xc5\x9a\xe4\x23\x0a\x70\x1a\x0c\x52\xa4\x29\x05\x6d\x90\xca\x2b\x76\x0d\xcd\x5f\x4d\xf5\x75\xad\xf0\xc1\x20\x77\x28\x00\x0b\x3f\xd0\x9c\x67\x44\xd8\xad\x6f\x62\x18\xb1\x14\x1d\x92\x5d\xed\xae\xba\xfa\x5c\x5a\x06\x86\xcd\xb0\xbf\xbd\xb9\x95\x8f\xad\xcc\x91\x11\x8f\x5b\x38\x68\x12\x48\x27\x79\x0b\x8f\x48\x62\x22\x6c\xe5\x60\xb4\x7d\x4e\x4c\xa7\x84\xcc\x21\x30\x50\x78\xef\x31\x43\x3b\x4a\xfa\x92\xa1\x75\x27\x5a\xac\xd8\x79\x35\xe2\x71\x8e\x60\x8e\x2d\xcc\x0a\x5f\x49\x28\x42\x70\x94\x61\xaf\xa1\x60\x9a\xcb\xa5\xbe\x58\xb6\xe5\x9f\x7e\x23\xb9\x36\x90\x51\x95\xbf\x5d\x97\xfa\xd6\xf7\x2f\x99\xa7\xa1\x57\xdf\x17\x5c\x59\x5e\xa3\xea\x55\xdb\xcd\x66\xdd\x8f\x80\xce\x8d\x80\xe3\xc1\xdf\x9b\xd1\x78\xed\xcd\x12\x42\x26\x72\xc0\x07\x69\x77\xae\x2c\x5e\x77\x94\x0d\x15\x64\x9b\xa6\x19\xf0\xa2\xd5\xa5\x9a\x81\x8b\xb1\x8e\x2a\x7f\x3e\xc4\x1f\x11\xb2\xc1\x57\x29\xbe\xb5\x54\xb3\xff\xa0\x03\xa6\xfc\xaf\x7e\x93\xc3\xa2\x05\x77\xa3\x63\xbd\xe3\x22\x9d\x29\xe1\x1d\xba\xab\xf4\xb7\xe3\xde\x3d\x69\x75\x0b\xcd\xf1\x66\x34\x63\xed\x97\xd4\xbd\x74\x31\x58\x83\x5c\x46\x12\x05\x48\xf1\x16\x19\xec\x2d\x88\x71\xc3\x1c\x8f\x9f\x91\xca\xb5\x11\xa5\x12\x55\x3b\x94\xa1\xd5\x19\xa2\xf3\x72\xf4\xa2\xc8\xd0\xb4\x82\xd3\x2c\x4d\xb7\xe5\xcb\x6c\x81\x9c\x73\xb4\x36\xca\x92\x24\xef\x14\x71\xee\x45\x6b\x47\x10\xec\xf9\xbf\xe3\xfc\xff\x4b\x28\xf1\x67\x73\xac\x24\xb0\x42\x7d\x77\x52\x79\xd7\xed\x97\x5e\x5c\xce\x48\xd1\xef\x2d\xdf\x14\x6e\x4f\x71\xab\x08\x95\x20\xbf\x3c\xfd\x77\xf4\x6e\xb8\xd0\xf2\xe0\xe1\xaf\x31\x9c\xeb\x76\x00\x96\x24\x93\xa8\x66\xd9\xf5\x22\xaa\x49\xb9\x13\x9e\x62\x84\x84\x8a\x7b\x1b\x2d\x97\x1b\x0c\x9f\x96\xc6\x2b\x59\x33\x54\x50\x80\x93\xab\xc3\xa9\xf8\x8b\xb4\x5e\x4f\x6e\x15\xc1\x3a\x92\x6a\xb6\x2d\xcc\x73\x39\xdf\x5c\xbd\x36\x54\x7f\xef\xdc\x1a\x9f\x74\x98\xda\x7a\x84\xe5\xba\x7d\xed\x4d\xad\x79\xdb\x79\xb9\x45\xbe\x56\xe5\x43\x63\xde\x3c\xb8\xdf\x03\x00\x00\xff\xff\x41\x05\xcc\xa6\x10\x19\x00\x00")

func subscriptions_openapiYamlBytes() ([]byte, error) {
	return bindataRead(
		_subscriptions_openapiYaml,
		"subscriptions_openapi.yaml",
	)
}

func subscriptions_openapiYaml() (*asset, error) {
	bytes, err := subscriptions_openapiYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "subscriptions_openapi.yaml", size: 6416, mode: os.FileMode(420), modTime: time.Unix(1700684336, 0)}
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
	"openapi.dinosaurs.yaml":     openapiDinosaursYaml,
	"openapi.subscriptions.yaml": openapiSubscriptionsYaml,
	"openapi.yaml":               openapiYaml,
	"openapi2.yaml":              openapi2Yaml,
	"subscriptions_openapi.yaml": subscriptions_openapiYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
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
	"openapi.dinosaurs.yaml":     &bintree{openapiDinosaursYaml, map[string]*bintree{}},
	"openapi.subscriptions.yaml": &bintree{openapiSubscriptionsYaml, map[string]*bintree{}},
	"openapi.yaml":               &bintree{openapiYaml, map[string]*bintree{}},
	"openapi2.yaml":              &bintree{openapi2Yaml, map[string]*bintree{}},
	"subscriptions_openapi.yaml": &bintree{subscriptions_openapiYaml, map[string]*bintree{}},
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
