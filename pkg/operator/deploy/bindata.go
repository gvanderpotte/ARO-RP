// Code generated for package deploy by go-bindata DO NOT EDIT. (@generated)
// sources:
// deploy/staticresources/aro.openshift.io_clusters.yaml
// deploy/staticresources/master/deployment.yaml
// deploy/staticresources/master/rolebinding.yaml
// deploy/staticresources/master/service.yaml
// deploy/staticresources/master/serviceaccount.yaml
// deploy/staticresources/namespace.yaml
// deploy/staticresources/worker/deployment.yaml
// deploy/staticresources/worker/role.yaml
// deploy/staticresources/worker/rolebinding.yaml
// deploy/staticresources/worker/serviceaccount.yaml
package deploy

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

var _aroOpenshiftIo_clustersYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x58\x4f\x6f\xeb\xc8\x0d\xbf\xfb\x53\x10\xe9\x21\x87\xc6\xca\x3e\xec\xa5\xf5\x2d\xc8\xdb\x2d\x82\x6e\xbb\xc1\x4b\xf0\x2e\x9b\x3d\x50\x23\x5a\x62\x33\x9a\x51\x67\xa8\xe4\xf9\x15\xfd\xee\x05\x67\x24\x5b\x76\xa4\x38\x2d\x50\x5d\x0c\x71\x86\xff\xc9\x1f\x69\xad\xd6\xeb\xf5\x0a\x3b\xfe\x4a\x21\xb2\x77\x1b\xc0\x8e\xe9\x9b\x90\xd3\xb7\x58\x3c\xff\x29\x16\xec\xaf\x5f\x3e\x95\x24\xf8\x69\xf5\xcc\xae\xda\xc0\x6d\x1f\xc5\xb7\x5f\x28\xfa\x3e\x18\xfa\x4c\x5b\x76\x2c\xec\xdd\xaa\x25\xc1\x0a\x05\x37\x2b\x00\x74\xce\x0b\x2a\x39\xea\x2b\x80\xf1\x4e\x82\xb7\x96\xc2\xba\x26\x57\x3c\xf7\x25\x95\x3d\xdb\x8a\x42\xd2\x30\xea\x7f\xf9\xa1\xf8\xb1\xf8\x61\x05\x60\x02\x25\xf6\x47\x6e\x29\x0a\xb6\xdd\x06\x5c\x6f\xed\x0a\xc0\x61\x4b\x1b\x30\xb6\x8f\x42\x21\x16\x18\x7c\xe1\x3b\x72\xb1\xe1\xad\x14\xec\x57\xb1\x23\xa3\x3a\xeb\xe0\xfb\x6e\x03\x6f\xce\xb3\x84\xc1\xac\xc1\xa5\x2c\x2c\x51\x2c\x47\xf9\xeb\x94\xfa\x0b\x47\x49\x27\x9d\xed\x03\xda\x83\xea\x44\x8c\xec\xea\xde\x62\xd8\x93\x57\x00\xd1\xf8\x8e\xa6\x52\x63\x5f\x86\x21\x5e\x83\xde\x28\x28\x7d\xdc\xc0\xbf\xfe\xbd\x02\x78\x41\xcb\x55\xf2\x36\x1f\xaa\xb9\x37\xf7\x77\x5f\x7f\x7c\x30\x0d\xb5\x98\x89\x00\x15\x45\x13\xb8\x4b\xf7\x46\xe1\xc0\x11\xa4\x21\xc8\x37\x61\xeb\x43\x7a\x1d\x4d\x84\x9b\xfb\xbb\x81\xbb\x0b\xbe\xa3\x20\x3c\x5a\xa0\xcf\x24\xf3\x7b\xda\x89\x9e\x4b\x35\x24\xdf\x81\x4a\x73\x4d\x59\xe1\x90\x31\xaa\x20\x66\xd5\x7e\x0b\xd2\x70\x84\x40\x5d\xa0\x48\x2e\x67\x7f\x22\x16\xf4\x0a\x3a\xf0\xe5\x3f\xc8\x48\x01\x0f\x14\x54\x08\xc4\xc6\xf7\xb6\xd2\x02\x79\xa1\x20\x10\xc8\xf8\xda\xf1\xf7\xbd\xe4\x08\xe2\x93\x4a\x8b\x42\x43\x2a\xc6\x87\x9d\x50\x70\x68\x35\x84\x3d\x5d\x01\xba\x0a\x5a\xdc\x41\x20\xd5\x01\xbd\x9b\x48\x4b\x57\x62\x01\x7f\xf3\x81\x80\xdd\xd6\x6f\xa0\x11\xe9\xe2\xe6\xfa\xba\x66\x19\x6b\xdd\xf8\xb6\xed\x1d\xcb\xee\x3a\x55\x2c\x97\xbd\xf8\x10\xaf\x2b\x7a\x21\x7b\x1d\xb9\x5e\x63\x30\x0d\x0b\x19\xe9\x03\x5d\x63\xc7\xeb\x64\xb8\x4b\xa5\x5e\xb4\xd5\x1f\xf6\x89\xbe\x9c\x58\x2a\x3b\x2d\x88\x28\x81\x5d\xbd\x27\xa7\xda\x5b\x8c\xbb\xd6\xa0\x66\x17\x07\xb6\x6c\xff\x21\xbc\x4a\xd2\xa8\x7c\xf9\xe9\xe1\x11\x46\xa5\x29\x05\xc7\x31\x4f\xd1\x3e\xb0\xc5\x43\xe0\x35\x50\xec\xb6\x14\x72\xe2\xb6\xc1\xb7\x49\x22\xb9\xaa\xf3\xec\x64\xa8\x24\x26\x77\x1c\xf4\xd8\x97\x2d\x8b\x66\xfa\x9f\x3d\x45\xd1\xfc\x14\x70\x9b\x3a\x1e\x4a\x82\xbe\xab\x50\xa8\x2a\xe0\xce\xc1\x2d\xb6\x64\x6f\x31\xd2\xff\x3d\xec\x1a\xe1\xb8\xd6\x90\x9e\x0f\xfc\x14\xa8\x8e\x2f\xe6\x68\xed\xc9\x23\x94\xcc\x66\x68\xe8\xc0\x87\x8e\xcc\x51\x67\x54\x14\x39\x68\xf5\x0a\x0a\x69\xcd\x4f\xd1\x65\xb9\x17\x53\x3f\x9a\xf0\xd9\xb7\xc8\xee\x98\xbc\xe0\x46\xe2\xf8\xfe\x93\x7b\xe1\xe0\x5d\x4b\x4e\x3e\xcc\xb5\x25\xd4\x38\xc6\x53\x86\x23\xff\x7e\x1e\x2e\x1d\x39\x78\xf3\xe5\x57\x85\xa7\x80\xe2\xc3\x28\x06\x6a\x6d\xcc\x13\x51\x4b\x2e\xa6\x33\xc5\x8e\x28\xe4\xe4\x3e\xf8\x96\xa4\xa1\x7e\xe6\xd6\x68\x7f\xe9\xbd\x25\x74\xb3\xbe\x9d\xe4\x4b\x9f\x9a\x1c\xbd\xe0\x2f\xbe\xae\xd9\xd5\xa7\x52\xdf\xb3\xca\x78\xb7\xe5\x7a\x06\x0f\xf7\xcc\x28\x8a\x36\x1b\xb8\xfc\xed\x87\xf5\x9f\x7f\xff\x63\x91\x7f\x2e\x17\x2d\x9f\x89\xbc\x3e\xad\x77\x2c\x5e\x8f\xfe\x72\xfb\xf0\x4e\xfa\xf4\x21\xd7\xb7\x73\xf4\x35\x7c\x66\xac\x9d\x8f\xc2\x26\xde\x07\x5f\xcd\xde\x79\x3c\xc5\xcb\xb3\xd6\x2d\x86\x35\x23\x2d\xc9\x6d\x43\xe6\x99\xc2\x7f\x13\xd8\x3e\xd8\xd9\xf4\xb2\x50\x3b\x7b\x70\x26\x7e\xe3\x31\x86\x80\xbb\x8f\xda\x6f\xbd\x99\x4c\xd8\x0f\x68\x1a\x21\xf5\xae\x7a\xb7\x4b\xc6\x3d\xe8\xee\xf3\x38\x8a\x6f\xbe\x6b\x4f\x1c\xd8\xf3\x64\xa4\xc9\x7e\xf0\x21\xfd\x2f\x8e\xe4\xad\xee\x85\xeb\xf3\xe8\x95\x77\x8c\x73\xf8\x95\x6e\x1d\x21\x98\x2f\xa3\x8e\x88\xff\x09\xc2\x8c\x77\x15\x4f\xf6\xbe\x25\xe5\xfb\x6b\xc3\x8c\x23\x49\x7a\x46\x32\xb0\x8b\x82\xce\x50\x2c\x4e\xc4\x2c\x54\xcd\x91\xf4\x8b\x83\x9c\xc3\xe0\xcb\xbb\x87\x7a\x96\x0a\xe1\x68\x1b\xb9\x8c\xd9\xd7\x53\x65\xfa\x4c\x4c\xc5\x40\xca\xb3\xdf\x90\xa1\x25\xd3\xa0\xe3\xd8\xa6\xfe\x70\x15\x55\xba\xac\xe8\x10\x8c\x54\xc1\x6b\x43\x4e\x03\x3a\x23\xb4\x22\x41\xb6\x71\x6f\xc4\xc1\x2c\xd5\xa1\x93\x14\xa1\x0b\xec\x03\xc3\xb3\xf3\xaf\x0e\x7c\x80\xd7\xb4\x27\xa5\xb3\xae\xb3\xa7\x95\x9f\xca\xc0\x03\x5a\x7b\x88\x5d\x12\x0f\x35\xbf\x90\x03\xdd\x27\x0a\x78\x72\x53\x7f\x86\xd5\xab\x24\xc0\xaa\xa2\x39\x0c\x11\x0f\xf4\xad\xb3\x6c\x58\xec\x2e\xef\x68\xbb\x49\xee\x41\x1a\x14\x75\x36\xc4\xb4\x7b\x19\xdf\x76\xde\xa5\x68\x9b\x14\xac\xd2\xf7\x73\x08\x14\x50\x9a\xb4\x77\xa0\x4b\x6b\x04\x87\xbc\xce\xf8\x48\x47\xd2\x53\x2c\xd3\x8e\xa2\x13\x35\x6d\x28\x5e\x39\x67\x44\x4e\x62\x18\x0b\xf8\xd5\x19\x1a\x6a\xba\xba\x4a\x45\xdd\x12\x3a\x55\x92\x42\x72\xa8\x0f\x83\x0e\xf2\xe2\x32\x23\x53\x93\x5b\x53\x05\x18\x4a\x96\x80\x81\xed\x0e\xd6\xc0\x7a\xdb\xf8\x96\x22\x74\x18\x64\xec\xef\x9b\xfb\xbb\xbc\x80\x36\x98\xdb\x28\x62\x3b\x27\xb4\x44\xf3\xfc\x8a\xa1\x8a\xeb\x74\x7b\xeb\x43\x7e\xd3\xd8\xa1\x70\xc9\x96\x25\x85\xda\x50\x70\x43\x85\xec\xb2\xdb\x49\xdf\x9c\xef\x7b\x0b\x8a\x8b\xb7\xd3\xf6\x1d\x68\x06\xb0\x18\xe5\x31\xa0\x8b\x3c\xfe\xdb\x9a\x47\xe4\xad\x0f\x2d\xca\x06\x74\xb7\x5b\x0b\xcf\x7a\x76\x16\xb7\x5b\x8a\x11\xeb\x05\x0d\x67\x78\x03\x61\x9c\x9f\xcb\x4b\xd0\xf2\x25\x71\x28\xbe\x9c\x34\x27\x82\x77\xb4\x7e\xf5\xa1\xba\x3a\x6c\xa8\xb3\x82\xe1\xe4\xef\xcc\x1e\xcb\x51\xa8\xf6\x61\xa7\xef\x06\xfb\x48\xfb\x83\x3e\x04\x72\x32\x60\xef\x1c\x9c\xe8\x73\x27\x33\x56\x25\xc8\x60\x97\x32\xcf\x2a\xb1\x97\xae\x97\x2b\x88\xbd\x69\x00\x63\xb2\xd9\xb2\x5b\x32\x54\xff\x51\x1b\xb1\x50\x2b\x92\x0e\xac\x5a\x5f\xec\x20\xf6\x6d\x8b\x81\xbf\xa7\xf2\x37\xd9\xc4\x01\x1d\x92\xf1\x0b\x76\x9e\x49\xc8\xdb\xf1\xf2\x61\xd6\x74\x7c\x3e\x93\x07\x18\x7f\xdc\x75\x34\xce\x57\x65\xde\x87\x7b\xdf\xc7\xc9\xd5\xd3\x25\x74\xa2\x8f\x0d\x5a\xbb\xd3\xd6\x1f\x13\x5e\x81\x56\x80\x02\x6b\x6c\x7c\x10\xe8\x9a\x90\xfe\xa8\x4c\x21\x32\x29\x5b\x92\x3a\xa0\x27\xbb\x8a\xb5\x1e\x86\x69\xc9\x09\xf2\xe1\xe9\x02\x4b\xa7\x3d\x63\xd7\x12\x7a\x7a\xba\x80\xce\x5b\x0c\x2c\xbb\x02\x7e\xf6\x73\x00\xa6\x0f\x7d\xc3\xb6\xb3\x74\x05\x7c\xea\xdf\xa8\x25\xe6\xa9\x82\x2a\x8e\xcd\x2e\xd7\x51\xfa\x80\x70\xb5\xe4\x7c\xb2\x86\x63\xfe\xcc\xf0\x74\x01\x06\x63\x0a\x66\x17\x7c\x89\xa5\xdd\xa5\x1b\x6a\xeb\x15\x44\x7f\xac\xf6\x7d\xcf\x4b\x6d\x04\x6b\xa9\x82\xa7\x8b\x3b\x37\x88\x9f\x41\x20\x38\x57\x11\x79\x04\xd0\x9b\x6d\x47\x77\xd8\x5c\x66\x33\x07\x2a\xf1\x0d\x79\x71\xff\x5b\x5e\x1a\xc7\xbf\x33\x0b\x9b\xff\xc7\x37\xae\x13\xd2\xe1\x9b\xd6\x27\xb4\x5d\x83\x9f\x0e\xb4\xd4\x35\xeb\xe1\x4b\xd6\xe4\x18\x20\x2f\x5e\x1b\xd0\x9a\x19\x3e\x14\xf9\xa0\x90\x99\x29\x87\x9e\x43\x63\xa8\x13\xaa\xfe\x7e\xfa\x2d\xeb\xe2\xe2\xe8\x63\x55\x7a\x9d\x2c\x65\xf0\xdb\xef\xab\x2c\x95\xaa\xaf\xa3\x35\x4a\xfc\x4f\x00\x00\x00\xff\xff\x99\x65\x98\xfa\x0b\x14\x00\x00")

func aroOpenshiftIo_clustersYamlBytes() ([]byte, error) {
	return bindataRead(
		_aroOpenshiftIo_clustersYaml,
		"aro.openshift.io_clusters.yaml",
	)
}

func aroOpenshiftIo_clustersYaml() (*asset, error) {
	bytes, err := aroOpenshiftIo_clustersYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "aro.openshift.io_clusters.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _masterDeploymentYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x52\xc1\x6e\xdb\x30\x0c\xbd\xfb\x2b\x88\xde\xdd\xa4\xb7\x42\xb7\x62\x0d\x7a\x19\x82\x62\x59\x77\x67\x64\x26\x16\x22\x8b\x02\x49\x07\x75\xbf\x7e\xd0\x92\x28\xce\x06\x64\x3c\x19\x7c\x8f\xef\x3d\xd2\xc2\x1c\x7e\x91\x68\xe0\xe4\x00\x73\xd6\xc5\xf1\xa9\x39\x84\xd4\x39\x78\xa5\x1c\x79\x1a\x28\x59\x33\x90\x61\x87\x86\xae\x01\x88\xb8\xa5\xa8\xe5\x0b\xca\x80\x03\x14\x6e\x39\x93\xa0\xb1\xb4\x03\xaa\x91\x34\x00\x09\x07\xba\x87\x69\x46\x4f\x0e\x38\x53\xd2\x3e\xec\xac\xc5\xaf\x51\xa8\x92\x1b\xcd\xe4\x8b\x89\x50\x8e\xc1\xa3\x3a\x78\x6a\x00\x94\x22\x79\x63\x39\xd9\x0f\x68\xbe\xff\x3e\xcb\x73\x37\x91\x9a\xa0\xd1\x7e\x3a\x51\x85\x63\x0c\x69\xff\x91\x3b\x34\xba\x4c\x0f\xf8\xb9\x19\x65\x4f\x27\xb3\x73\xe7\x23\xe1\x11\x43\xc4\x6d\x24\x07\xcb\x06\xc0\x68\xc8\xb1\x4e\xcd\x6f\x53\x2a\xde\xe4\xb9\x9b\x08\xe0\xb2\x65\x29\xcf\xc9\x30\x24\x92\x3a\xdc\x82\xe7\x61\xc0\xd4\x5d\xd5\xda\x22\x75\xd5\x96\xbd\xce\xb1\x7a\xbd\x6b\x6b\x66\x56\x2a\x0c\x58\xd6\x7b\x5b\xad\x57\x3f\x5e\x7e\xae\x5e\x2b\xf0\xef\xff\xaa\x50\x66\xb1\x1b\x9b\x9a\xf4\x9d\xc5\x1c\x3c\x2f\x9f\x97\x15\xbd\x28\xf5\x66\xb9\x36\x63\x38\x52\x22\xd5\x77\xe1\x2d\xb9\x19\xb7\xb0\xde\xc8\xe6\x2d\x80\x8c\xd6\x3b\x58\xf4\x84\xd1\xfa\xaf\x85\x10\x76\xd3\x2d\xe1\x6f\xdb\xc4\x1d\x6d\x6e\x9e\xc6\xa5\xdb\x0a\x47\x7a\x3c\x8c\x5b\x92\x44\x46\xfa\x18\x78\x71\x3a\x89\x83\x87\x87\x33\x55\x49\x8e\xc1\xd3\x8b\xf7\x3c\x26\x5b\xdf\x79\xb9\x7f\xdc\x25\xb0\x04\x9b\xbe\x45\x54\x3d\x91\x75\x52\xa3\xa1\xf5\x71\x2c\xbc\xd6\x4b\xb0\xe0\x31\x9e\x07\x8c\x63\xd1\x09\x9c\x66\x7f\xf6\x40\x93\xfb\x4f\xc2\xba\xc8\x25\x87\x83\xd5\x67\x50\xd3\x0a\xd0\x6e\x47\xde\x1c\xac\x79\xe3\x7b\xea\xc6\x48\xcd\xef\x00\x00\x00\xff\xff\x33\xe2\xa4\x2c\xd0\x03\x00\x00")

func masterDeploymentYamlBytes() ([]byte, error) {
	return bindataRead(
		_masterDeploymentYaml,
		"master/deployment.yaml",
	)
}

func masterDeploymentYaml() (*asset, error) {
	bytes, err := masterDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _masterRolebindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8e\xb1\x4e\x03\x31\x0c\x40\x77\x7f\x85\x7f\x20\x87\xd8\x50\x36\x60\x60\x2f\x12\xbb\x9b\xb8\xd4\xf4\x62\x47\x8e\xd3\xa1\x5f\x8f\xaa\xa2\x5b\x90\x6e\xb5\xdf\xf3\x33\x75\xf9\x62\x1f\x62\x9a\xd1\x8f\x54\x16\x9a\x71\x36\x97\x1b\x85\x98\x2e\x97\x97\xb1\x88\x3d\x5d\x9f\xe1\x22\x5a\x33\xbe\xaf\x73\x04\xfb\xc1\x56\x7e\x13\xad\xa2\xdf\xd0\x38\xa8\x52\x50\x06\x44\xa5\xc6\x19\xc9\x2d\x59\x67\xa7\x30\x4f\x8d\xee\x02\xb8\xad\x7c\xe0\xd3\x1d\xa2\x2e\x1f\x6e\xb3\xef\x04\x01\xf1\x5f\x6f\x3b\x5f\x1e\xb3\x44\xb5\x89\xc2\x98\xc7\x1f\x2e\x31\x32\xa4\x3f\xe7\x93\xfd\x2a\x85\x5f\x4b\xb1\xa9\xb1\xfb\xd5\x63\x37\x3a\x15\xce\x68\x9d\x75\x9c\xe5\x14\x89\x6e\xd3\x79\x83\xe1\x37\x00\x00\xff\xff\x4f\x98\xa4\x7c\x24\x01\x00\x00")

func masterRolebindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_masterRolebindingYaml,
		"master/rolebinding.yaml",
	)
}

func masterRolebindingYaml() (*asset, error) {
	bytes, err := masterRolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master/rolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _masterServiceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8d\x41\xca\xc2\x40\x0c\x46\xf7\x73\x8a\x5c\x60\xa0\xff\xae\xcc\x29\x7e\x10\xdc\x87\xe9\xa7\x1d\xb4\x93\x90\xc4\x2e\x3c\xbd\xd4\x16\x5d\xb9\x0b\xef\x7b\xbc\xb0\xb6\x33\xcc\x9b\xf4\x42\xeb\x5f\xba\xb5\x3e\x15\x3a\xc1\xd6\x56\x91\x16\x04\x4f\x1c\x5c\x12\x51\xe7\x05\x85\xd8\x24\x8b\xc2\x38\xc4\xf2\xc2\x1e\xb0\x63\x73\xe5\x8a\x42\xa2\xe8\x3e\xb7\x4b\x64\x7e\x3e\x0c\x1f\x39\xb9\xa2\x6e\x1d\xc7\x1d\x35\xc4\xb6\x9b\x88\x55\x7f\x45\x55\x2c\x7c\xb7\xf2\xf1\x7d\x8e\xd0\x37\xd8\xd7\x42\xe3\x30\x0e\x07\x08\xb6\x2b\xe2\xff\x8b\x5f\x01\x00\x00\xff\xff\x10\x70\xf6\x36\xda\x00\x00\x00")

func masterServiceYamlBytes() ([]byte, error) {
	return bindataRead(
		_masterServiceYaml,
		"master/service.yaml",
	)
}

func masterServiceYaml() (*asset, error) {
	bytes, err := masterServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master/service.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _masterServiceaccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xca\x31\x8e\x02\x31\x0c\x05\xd0\x3e\xa7\xf0\x05\x52\x6c\xeb\x6e\xcf\x80\x44\xff\x95\xf9\x08\x0b\xc5\x8e\x1c\xcf\x14\x9c\x9e\x06\x51\xbf\x87\x65\x77\xe6\xb6\x70\x95\xeb\xaf\xbd\xcc\x0f\x95\x1b\xf3\xb2\xc1\xff\x31\xe2\xf4\x6a\x93\x85\x03\x05\x6d\x22\x8e\x49\x15\x64\xf4\x58\x4c\x54\x64\x9f\xd8\xc5\xfc\xda\x5e\x18\x54\x89\x45\xdf\x4f\x7b\x54\xc7\xfb\x4c\xfe\x72\xfb\x04\x00\x00\xff\xff\xe4\xf5\x04\x25\x70\x00\x00\x00")

func masterServiceaccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_masterServiceaccountYaml,
		"master/serviceaccount.yaml",
	)
}

func masterServiceaccountYaml() (*asset, error) {
	bytes, err := masterServiceaccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master/serviceaccount.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _namespaceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xca\xb1\x0d\x02\x31\x0c\x05\xd0\x3e\x53\x58\xd7\x07\x44\x9b\x21\x28\xe9\xbf\x2e\x1f\x61\x41\xec\x28\x36\x14\x4c\x8f\xa8\xae\x7f\x98\x7a\xe3\x0a\x75\x6b\xf2\xb9\x94\xa7\x5a\x6f\x72\xc5\x60\x4c\xec\x2c\x83\x89\x8e\x44\x2b\x22\x86\xc1\x26\x3e\x69\xf1\xd0\x7b\x56\x7c\xdf\x8b\xd5\x27\x17\xd2\x57\x11\x81\x99\x27\x52\xdd\xe2\xef\xe5\xb0\x27\xf5\xb3\x79\x67\x0d\xbe\xb8\xa7\xaf\x26\xdb\x56\x7e\x01\x00\x00\xff\xff\xc1\xaf\xa6\x4c\x7c\x00\x00\x00")

func namespaceYamlBytes() ([]byte, error) {
	return bindataRead(
		_namespaceYaml,
		"namespace.yaml",
	)
}

func namespaceYaml() (*asset, error) {
	bytes, err := namespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "namespace.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _workerDeploymentYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x52\xcb\x6e\xdb\x40\x0c\xbc\xeb\x2b\x88\xdc\x15\x27\xb7\x60\x6f\x41\x63\xe4\x52\x04\x45\xd3\xf4\x4e\xaf\xa6\xd6\xc2\xfb\x02\x49\xb9\x55\xbe\xbe\x50\x65\xcb\x32\x0a\x98\x27\x61\x38\xc3\x19\x2e\xc5\x35\xfc\x84\x68\x28\xd9\x11\xd7\xaa\x9b\xe3\x63\x73\x08\xb9\x73\xf4\x82\x1a\xcb\x98\x90\xad\x49\x30\xee\xd8\xd8\x35\x44\x91\x77\x88\x3a\x7d\xd1\x24\x70\xc4\x52\xda\x52\x21\x6c\x45\xda\xdf\x45\x0e\x90\x86\x28\x73\xc2\xad\x9e\x56\xf6\x70\x54\x2a\xb2\xf6\xe1\x97\xb5\xfc\x39\x08\x16\x72\xa3\x15\x7e\x32\x11\xd4\x18\x3c\xab\xa3\xc7\x86\x48\x11\xe1\xad\xc8\x6c\x9f\xd8\x7c\xff\x75\x95\xe7\x66\x22\x35\x61\xc3\x7e\x9c\xa9\x52\x62\x0c\x79\xff\x51\x3b\x36\x9c\xd5\x89\xff\xbc\x0f\xb2\xc7\x6c\x76\x42\x3e\x32\x1f\x39\x44\xde\x45\x38\x7a\x68\x88\x0c\xa9\xc6\x45\xb5\x7e\x9b\xa9\xe2\x55\x9e\x9b\x89\x88\xce\x5b\x4e\xe5\x4b\x36\x0e\x19\xb2\x88\x5b\xf2\x25\x25\xce\xdd\x65\x5a\x3b\x8d\xba\xcc\x96\xbd\xae\x7b\xcb\xeb\x5d\xa0\x95\xd9\x54\x21\xf1\xb4\xde\xeb\xf6\x6d\xfb\xfd\xf9\xc7\xf6\x65\x69\xfc\x7f\xaf\xa5\x15\xc3\x11\x19\xaa\xdf\xa4\xec\x70\xb1\x23\xea\xcd\xea\x2b\x6c\x0d\x11\x55\xb6\xde\xd1\xa6\x07\x47\xeb\x3f\x37\x02\xee\xc6\x6b\x42\x11\x73\xf4\xf4\xf0\xf4\x70\x82\x73\xe9\xf0\x7e\x75\xd8\x33\xda\x4a\x89\xb8\x3f\x0c\x3b\x48\x86\x41\xef\x43\xd9\xcc\x0b\x39\xba\xbb\x3b\x51\x15\x72\x0c\x1e\xcf\xde\x97\x21\xdb\xdb\x8d\xff\xee\x9f\xbb\x84\x22\xc1\xc6\x2f\x91\x55\x67\xb2\x8e\x6a\x48\xad\x8f\x83\x1a\xa4\xf5\x12\x2c\x78\x8e\xcd\xdf\x00\x00\x00\xff\xff\xce\x78\x87\x64\x1b\x03\x00\x00")

func workerDeploymentYamlBytes() ([]byte, error) {
	return bindataRead(
		_workerDeploymentYaml,
		"worker/deployment.yaml",
	)
}

func workerDeploymentYaml() (*asset, error) {
	bytes, err := workerDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "worker/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _workerRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x8e\xb1\x6e\x2c\x31\x08\x45\x7b\xbe\x82\x1f\xb0\x57\xaf\x7b\x72\x9b\x22\x7d\x14\xa5\x67\x3d\x24\x83\xc6\x63\x2c\xc0\xbb\x52\xbe\x3e\x9a\xd9\x6d\x53\xa5\xe2\x0a\x1d\x0e\x17\x52\x4a\x40\x43\x3e\xd8\x5c\xb4\x17\xb4\x2b\xd5\x4c\x33\x56\x35\xf9\xa6\x10\xed\x79\xfb\xef\x59\xf4\x72\xfb\x07\x9b\xf4\xa5\xe0\x4b\x9b\x1e\x6c\x6f\xda\x18\x76\x0e\x5a\x28\xa8\x00\x62\x35\x3e\x0f\xde\x65\x67\x0f\xda\x47\xc1\x3e\x5b\x03\xc4\x4e\x3b\x17\x24\xd3\xa4\x83\x8d\x42\x2d\xdd\xd5\x36\x36\xb0\xd9\xd8\x0b\x24\xa4\x21\xaf\xa6\x73\xf8\x61\x4a\x07\x9b\x75\x70\xf7\x55\x3e\x23\x8b\x02\xa2\xb1\xeb\xb4\xca\x4f\xa2\x3e\x5a\x38\x20\xde\xd8\xae\xcf\xed\x17\xc7\x39\x9b\xf8\x23\xdc\x29\xea\xfa\x17\xff\xc5\x83\x62\xfe\xf2\x66\x9c\xf6\x23\xcd\xb1\x50\x30\xfc\x04\x00\x00\xff\xff\x30\x78\x19\x41\x50\x01\x00\x00")

func workerRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_workerRoleYaml,
		"worker/role.yaml",
	)
}

func workerRoleYaml() (*asset, error) {
	bytes, err := workerRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "worker/role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _workerRolebindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8d\x31\x6e\xc3\x30\x0c\x45\x77\x9d\x82\x17\x90\x8b\x6e\x85\xb6\xb6\x43\x77\x17\xe8\x4e\xcb\x74\xcd\xda\x26\x05\x8a\x72\x01\x9f\x3e\x08\x12\x64\x09\xe0\xf9\xbf\xf7\x1f\x16\xfe\x21\xab\xac\x92\xc0\x06\xcc\x1d\x36\x9f\xd5\xf8\x40\x67\x95\x6e\x79\xab\x1d\xeb\xcb\xfe\x1a\x16\x96\x31\xc1\xe7\xda\xaa\x93\xf5\xba\xd2\x07\xcb\xc8\xf2\x1b\x36\x72\x1c\xd1\x31\x05\x00\xc1\x8d\x12\xa0\x69\xd4\x42\x86\xae\x16\xff\xd5\x16\xb2\x60\xba\x52\x4f\xd3\x15\xc2\xc2\x5f\xa6\xad\x9c\x04\x03\xc0\x53\xef\xf4\xbe\xb6\xe1\x8f\xb2\xd7\x14\xe2\xdd\xfc\x26\xdb\x39\xd3\x7b\xce\xda\xc4\x4f\xe5\xdb\x56\x0b\x66\x4a\xa0\x85\xa4\xce\x3c\x79\xc4\xa3\x19\x3d\xe0\x70\x09\x00\x00\xff\xff\x73\xce\x57\x9b\x2a\x01\x00\x00")

func workerRolebindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_workerRolebindingYaml,
		"worker/rolebinding.yaml",
	)
}

func workerRolebindingYaml() (*asset, error) {
	bytes, err := workerRolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "worker/rolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _workerServiceaccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xca\x31\x8a\xc3\x40\x0c\x05\xd0\x7e\x4e\xa1\x0b\x4c\xb1\xad\xba\x3d\x43\x20\xfd\x67\xfc\x43\x84\xb1\x34\x68\x64\x07\x72\xfa\x34\x21\xf5\x7b\x98\x76\x67\x2e\x0b\x57\xb9\xfe\xda\x6e\xbe\xa9\xdc\x98\x97\x0d\xfe\x8f\x11\xa7\x57\x3b\x58\xd8\x50\xd0\x26\xe2\x38\xa8\x82\x8c\x1e\x93\x89\x8a\xec\xaf\xc8\x9d\xf9\xb5\x35\x31\xa8\x12\x93\xbe\x9e\xf6\xa8\x8e\xf7\x99\xfc\xe5\xf6\x09\x00\x00\xff\xff\xe3\x3c\x43\x66\x70\x00\x00\x00")

func workerServiceaccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_workerServiceaccountYaml,
		"worker/serviceaccount.yaml",
	)
}

func workerServiceaccountYaml() (*asset, error) {
	bytes, err := workerServiceaccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "worker/serviceaccount.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"aro.openshift.io_clusters.yaml": aroOpenshiftIo_clustersYaml,
	"master/deployment.yaml":         masterDeploymentYaml,
	"master/rolebinding.yaml":        masterRolebindingYaml,
	"master/service.yaml":            masterServiceYaml,
	"master/serviceaccount.yaml":     masterServiceaccountYaml,
	"namespace.yaml":                 namespaceYaml,
	"worker/deployment.yaml":         workerDeploymentYaml,
	"worker/role.yaml":               workerRoleYaml,
	"worker/rolebinding.yaml":        workerRolebindingYaml,
	"worker/serviceaccount.yaml":     workerServiceaccountYaml,
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
	"aro.openshift.io_clusters.yaml": {aroOpenshiftIo_clustersYaml, map[string]*bintree{}},
	"master": {nil, map[string]*bintree{
		"deployment.yaml":     {masterDeploymentYaml, map[string]*bintree{}},
		"rolebinding.yaml":    {masterRolebindingYaml, map[string]*bintree{}},
		"service.yaml":        {masterServiceYaml, map[string]*bintree{}},
		"serviceaccount.yaml": {masterServiceaccountYaml, map[string]*bintree{}},
	}},
	"namespace.yaml": {namespaceYaml, map[string]*bintree{}},
	"worker": {nil, map[string]*bintree{
		"deployment.yaml":     {workerDeploymentYaml, map[string]*bintree{}},
		"role.yaml":           {workerRoleYaml, map[string]*bintree{}},
		"rolebinding.yaml":    {workerRolebindingYaml, map[string]*bintree{}},
		"serviceaccount.yaml": {workerServiceaccountYaml, map[string]*bintree{}},
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
