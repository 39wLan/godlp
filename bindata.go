// Code generated by go-bindata. DO NOT EDIT.
// sources:
// conf.yml (20.327kB)

package dlp

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	// clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _confYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3c\xfd\x77\xdb\xc6\x91\xbf\xeb\xaf\xd8\x07\xbd\xeb\x23\x53\x50\x47\x90\xd4\x87\xf9\xd2\xaa\x34\x49\xcb\x3c\xeb\xeb\x48\xc9\xbd\x96\xe2\xf1\x2d\x81\x25\x85\x13\x09\xf0\x00\x50\x96\x22\xf0\x5e\x94\x26\xb1\xd3\xca\x76\xda\xc8\xcd\x35\xb1\xeb\xb8\x97\xd8\x49\x5a\x5b\x76\xd3\xfa\x43\x72\xec\x7f\x46\x20\xa5\x9f\xfc\x2f\xdc\xdb\x5d\x00\x04\x40\x50\xa2\x64\xba\xbd\xbb\x77\x79\x2f\xe2\x7e\xcc\xcc\xee\xce\xcc\xce\xce\xcc\x2e\x3c\x0c\xa6\xe6\x52\xd3\xf3\x80\x97\xa5\xb2\x58\x01\x65\xb1\x8a\x86\x86\xc1\x0a\x5a\x57\x01\x54\x10\x58\xac\xd7\x91\x92\x84\x35\x54\x4d\x42\x15\xb1\x00\x4a\x02\x69\x57\x61\x0d\x01\xa8\x82\x54\xb5\x9e\x94\xa5\x32\x50\x35\xa5\xc1\x6b\x40\x94\x08\xa1\x7f\xc4\x7f\x46\x2a\xf2\xd0\x54\x55\x2e\xc1\x6a\x7c\x08\x80\x14\xd4\x50\x1c\x44\xc2\x11\x2e\xc4\x85\x43\x91\xf1\x21\x00\x12\x75\xf1\x22\x52\x54\x51\x96\xe2\x60\x35\x32\x04\xc0\x8c\x2c\xa0\x38\x50\x50\x15\x41\x15\x81\x61\x20\xa0\x52\xa3\xa2\x9b\x75\x8c\x50\xad\xca\x97\xb2\xf3\xc9\x38\x00\x65\x58\x25\x20\x9a\xd2\x40\xa0\x2c\x2b\x40\x41\x35\x59\x43\x40\x45\xca\xaa\xc8\x23\x70\x49\xd4\x96\x81\x52\xe7\x59\x13\x12\x83\xd4\x1b\x0a\x02\x7c\x55\x44\x92\x06\x72\xa9\x0b\x2c\x10\x50\x19\x36\xaa\x1a\x10\x55\x0a\x35\x04\xc0\x30\x10\xcb\x20\x2d\xc1\x52\x15\x65\x1b\x55\xa4\xe2\x3e\x54\xab\x6b\xeb\x2c\x10\x35\x50\x43\x50\x52\x01\xac\x56\x81\x42\x3a\x31\x27\x10\x01\x16\x58\x50\x6a\x68\x5e\x64\x5e\x96\x34\x28\x4a\x2a\x50\xe5\x1a\x22\x38\x99\x94\xca\x02\x59\xaa\xae\x03\x6d\x19\xa9\xa8\x9b\xce\x08\x99\xc4\xc2\x32\x92\x00\x16\xcb\x25\x11\x0f\x86\x6a\xf2\x2a\x72\x11\xc1\x23\x69\xcb\x68\x9d\xa0\x8a\x12\x48\x89\xaa\x3d\x2c\x26\xe1\x98\x45\x1c\xe4\x0b\x84\xa8\x40\x61\x00\x04\x3c\x52\xf0\xbc\x08\x31\x50\x5a\x07\xf5\x86\xba\x6c\x52\xc6\xc4\x04\x07\x31\x2c\x3a\x47\xd5\x24\x36\x03\xd7\xa6\xe5\x4a\x46\xaa\x37\xb4\x38\x88\x85\xcf\x8c\xd1\xb6\x2c\xaa\xa0\xb5\x2c\xa1\x13\x07\xe1\xa1\x19\xa8\xae\x50\x2c\x32\x7e\x7a\x0d\xd6\xea\x55\x04\xac\x66\xa0\x6a\x50\xd1\x86\x00\x08\x01\x5c\x9d\x85\x35\x14\xb7\x80\x92\xe7\x13\x59\x30\x0c\x70\x1b\x90\xcb\x36\xca\x10\x00\x80\x54\x16\xd6\xeb\x28\x0e\x4c\x28\x59\x22\x40\x79\x5c\x65\xc1\x42\x62\x8a\x05\xd9\xf4\xfc\x74\x22\x99\x66\x41\x62\x7a\x6a\x0e\x14\x08\xde\x45\x58\x6d\xa0\x38\x60\xde\x62\x48\x75\xae\x5c\x56\x91\x16\x07\x1c\xa9\xcd\x43\x41\x10\xa5\x4a\x1c\x84\x31\x41\xd2\x05\xca\x8a\x5c\xc3\x5c\x06\x1a\x14\xab\x04\x6a\x1a\x49\x15\x6d\x39\x0e\x46\x49\x2d\x8b\x56\x91\xa2\xa2\x38\xd1\x42\xd2\x92\xa9\x48\xb2\x82\x92\xcb\x50\xc9\x61\xd2\xcc\x4f\x18\x47\xf3\x05\x51\x12\xe2\x20\x0f\x66\x17\x67\xd2\xd9\x4c\x12\x14\x1c\x53\x37\xdb\xf0\x7c\xe7\xcf\x27\x8a\x8b\xf3\xf3\xe9\x6c\x31\x99\xc8\xa5\xad\x96\xe9\xb9\x9f\xda\x2d\x3f\x3d\x9f\x59\x48\xe7\xe6\xc9\xfa\xe6\x17\x67\x93\x0b\x8b\x89\x85\xcc\xdc\x6c\xc1\xc3\xca\xc4\xf4\xb4\x0f\xbf\xba\x59\xe1\xc3\xfe\x85\xc4\x94\x07\x95\xb6\xf8\x80\x9a\x9c\xf6\x80\x3b\x5b\xad\xc1\xde\x36\x1b\x53\x3f\xee\x31\x6a\x7a\x66\x7e\xe1\x67\x7d\x10\xea\x81\x7e\x36\x91\x4b\x8f\xc5\x3c\xf8\x58\xfe\x2e\x64\x0a\xc5\x38\x78\x4f\x5b\x58\x30\x93\x1a\x65\x41\x32\x9b\x8c\x46\x58\x90\x48\xa5\xb2\xe9\x5c\x8e\xc5\xc2\x3a\x9b\xce\xb2\x20\x95\xce\xa4\xd2\xb3\x0b\x99\x73\x3f\xf3\xb2\xd9\x1c\x7d\x26\x35\x7a\xdc\xd0\x33\xa9\xd1\x01\x8e\xdb\xe9\x39\x6e\xdc\x0e\x24\xe3\xbf\x11\x91\x24\x78\x88\x33\xb3\x8b\xd3\xd3\x4c\x1f\xb2\xb0\xe0\x9c\xc8\xc9\xf3\x99\xd9\xc4\xfc\xf9\xb9\x59\xaf\x56\xf8\xeb\x5f\x67\x2b\x46\x5d\x9b\x6c\xcc\x77\x4b\x85\xbc\xa3\x9d\x66\xa0\x88\x7b\xcf\x47\xfa\x1b\x89\xac\x2b\x93\x3a\xe1\x58\x1e\xfb\xc2\x79\x88\x66\x52\xc9\x44\x76\xc0\x34\xd3\x33\x89\x4c\x7f\x7b\xdf\x4b\xd2\xd7\x80\x39\x49\x2f\x9e\x6a\xfd\x4e\x0a\x67\x13\xb3\x17\x4e\x48\x22\xe6\x6b\x6e\x5d\x4a\x90\xc8\xe5\xe6\xe7\xb2\x0b\xaf\xaf\x07\x2e\x0b\x4a\x77\xe3\x31\xfb\xab\x03\xe5\xc4\x9d\x4d\xcc\x9c\x46\xfd\x5d\x24\x88\x0d\x38\x66\x74\x1b\xc8\x89\x39\x93\x48\xe2\x59\xf5\x39\x7e\xf7\x39\x66\x6d\xc1\x09\x5f\xa5\x88\x77\xed\x8c\xc4\xd9\xc4\x6b\x8f\xd5\xe7\x76\x3f\x9b\x59\x48\xce\x65\x66\xfb\x1c\xce\x22\x1e\x89\xb9\x38\x3d\xea\xdd\xd9\x89\x7e\x79\xe5\x56\xc9\x9e\x9a\x93\x3a\xf1\x36\xf1\x10\x8c\xf5\xcb\x8d\xec\xc2\xf9\xd7\x66\x7d\x9f\xf6\x2f\x31\xd5\xaf\x46\xf7\x1e\xaa\xcb\x54\xa5\x16\xfb\xa4\x69\xb3\x66\xcc\x43\x82\x04\x30\xfd\xf8\x1f\x34\xd2\x49\xca\xf5\x75\x45\xac\x2c\x6b\x24\x1a\xf9\x31\x33\x34\x0c\xb2\xb6\x17\xce\xcb\xb5\x92\x28\x21\x81\x06\x10\x66\x80\x60\x76\x4b\x02\xa8\x2b\xe2\x2a\xd4\x4c\x37\x1b\x7b\xe9\x76\x0c\xd1\x71\xe4\x6b\x50\x82\x15\x24\x60\xbf\x1a\x0f\xa7\x21\x58\x63\x81\x20\x03\x49\xd6\x40\x4d\x16\xc4\xf2\x3a\x46\x6a\x40\x8b\xae\x20\x2a\x88\xd7\xaa\xeb\x6c\x2f\x07\x5d\xc3\xc1\x00\x2f\xd7\xd7\xc9\x14\x4c\x12\x22\x09\xb6\x9c\xf3\x19\x1a\x06\xe1\xb7\xdd\x53\xa6\x2e\xfd\xdb\x5c\x38\x1c\x0e\x03\x16\x63\x93\xe2\xdb\x3f\xb2\x10\x4d\x88\xa1\x61\x50\x51\x10\xd4\x90\x02\xa8\xf7\x4e\xe3\x0e\x79\x15\x29\x97\x14\x11\x87\x55\x35\x58\xad\x76\x7a\x15\xa4\x92\xb8\xa9\x0c\x04\xa4\x21\x5e\x03\x75\x59\x15\x35\x51\x96\x70\xb8\x84\x1d\x66\x1c\x1e\x0e\x39\x5c\x7f\xc7\xac\x3c\x5e\x3f\x8e\x14\xcc\xa3\x47\x2a\xcb\x54\x78\x9d\x13\x3d\x85\x54\x5e\x11\xeb\x1a\x89\x11\x5b\x1f\xfd\xaa\x75\x73\xd7\xb8\xfe\x84\x74\xa5\x25\x2a\x7d\x0d\x55\x51\x7d\x59\x96\x50\x51\x6a\xd4\x4a\x48\x21\x9d\x49\xb3\xb3\xbd\xfd\xd7\x83\x9d\x5b\xc6\xf5\x27\xed\x2f\x36\x4d\x35\x5c\x45\xd5\x38\x98\xa6\xdb\x6b\x18\xa4\xe8\xfc\xcb\x48\xac\x0a\x78\xf2\x50\x02\x50\x51\xe0\x3a\x09\x18\xcd\xc5\xd5\x90\xb6\x2c\x0b\x2a\x4b\x56\xa6\xa0\x2a\x24\x2b\x95\xcb\x00\x41\x7e\x19\x88\x1a\xaa\x91\xa0\x89\x02\x8b\x2a\x98\xcb\xda\x50\x23\xe6\x30\x24\x30\xc2\xa1\x6d\x05\xad\x01\xb4\x56\x57\x90\x8a\xe3\x5e\x16\x48\x32\x90\x10\x12\x80\x26\x03\xa4\xf2\xb0\x8e\x4c\x84\x0b\x59\x54\x61\x2f\xe2\x3f\x17\x52\x22\xaf\xb1\x17\xf1\x5f\x6b\xce\x22\xaf\xe1\x38\xe2\x92\xac\x08\x1c\x4b\x7e\x22\x2c\x18\x19\x19\x29\xb0\x26\x44\x00\xa3\x03\x5d\x07\x04\x39\x08\x7e\xf0\x03\x10\xb8\x68\x36\x11\x4a\x41\x93\xbb\x78\xca\x71\x52\x06\x64\x48\x33\xbc\x23\x55\x73\x18\xab\x8e\xf1\x2d\x50\x2c\x3c\x2e\x30\x19\x0f\x04\x02\xd1\x7c\x38\x74\xa6\x10\xd4\x03\xb1\xfc\x28\x2d\x8c\xe6\xc3\xa1\xa8\x59\x1e\xcb\x47\xd8\xd1\xd0\x38\x2e\x8e\xe7\xc3\x5c\x74\x34\x34\x81\xcb\x13\x16\xd2\x19\x1b\x36\x98\x07\xa1\xc2\xe4\x92\xb0\x11\x6b\x76\x4a\x41\x3d\x10\x18\x8f\xd1\xae\x7c\x38\x34\x5a\x58\x12\x36\xa2\xce\xfe\xe0\x52\xc9\x9a\x9e\x6b\xba\xc3\xe0\x9c\x58\xc5\xea\x6c\x47\xe0\xa5\x2a\xe4\x57\xaa\xa2\xaa\x51\x91\xd0\x6e\x6b\x3d\x67\x13\xd5\x8a\x1c\x07\xf9\x99\x44\xee\x42\x3a\x85\xe3\x32\xb5\x51\xaf\xcb\x8a\xa6\x02\xda\xc4\x76\x74\x1d\x09\x60\x15\x1b\x94\x0e\xe9\xb7\x2c\xcd\x20\x9b\x82\xec\x1c\xbc\xd7\x4b\xb8\x49\x6b\x28\x12\x12\xac\x61\xac\x39\xd2\x34\x83\x19\x07\x50\x3c\x15\xeb\x90\x3d\x49\x20\x60\xa1\x1f\x47\x76\xc4\xa2\x6b\x4a\x0e\x0c\x3b\x28\x50\x5d\xc3\x45\x93\x23\x49\x59\xd2\xd0\x1a\x1e\xdf\xa9\x38\x05\x6b\x1a\x3c\xed\x26\xad\x2a\x58\x86\x2a\xd6\xc9\x12\xc2\xdb\x41\x6e\x48\x82\x7b\x2e\xda\xb2\x28\x59\x04\x2f\x22\x45\x2c\xaf\x67\xa1\x54\xa1\xca\x4b\xeb\x16\x6b\x93\xe6\x9a\x19\xc2\x2e\x5e\x2b\x92\xbd\xca\xb0\x80\x51\x50\x0d\x2a\x2b\xc5\x9a\x5c\x12\xab\x48\x65\x58\x26\x49\x21\xe6\x2d\x00\x13\x92\xfe\x62\x00\xba\xc1\x19\x96\xb1\xf7\xbc\xb3\x8c\x21\x78\x54\xad\x32\x2c\x43\x69\x32\x2c\x23\x97\xcb\x22\x8f\x0b\x3c\x24\x1d\xb8\xdf\x42\xb4\xcb\x18\x51\xad\x41\x45\xb3\x7a\x3a\x15\x73\x54\xfc\x57\xa6\x63\x31\x2c\x53\x15\xa5\x15\x5a\x32\xd7\xd4\x29\x89\x52\x59\xb6\x66\x2c\x75\x4a\x84\x82\x55\xf4\x2c\xa1\x68\x51\x46\x1d\xac\x4e\x8d\x60\x3a\xab\x14\x9b\x2e\xb0\x83\x50\x93\x4b\x55\xd1\x89\xe2\x04\x70\x23\x39\xe0\x5d\xb0\x1e\x30\x5e\x16\x30\x2b\x6c\x9b\xcb\xb0\xcc\xfe\xf3\x2f\xda\x37\x6f\xdb\x8d\x0c\xcb\x50\xc3\xca\xb0\xcc\xc1\xe6\x76\xfb\xbb\x3d\x86\x65\xa8\xc1\x65\x80\x65\x35\x92\x96\x6a\x5a\x46\x90\x68\x24\xb1\xae\xa6\xc2\x59\xfb\xd7\xdc\x82\x58\x89\xe9\xfe\x22\x1a\x5f\x42\x60\x15\xeb\x93\x48\x0f\x54\x52\x5e\x07\xe5\x86\xc4\x6b\xc4\x7e\xaa\x0d\x7e\x19\x40\xd5\x8c\xa2\x58\x70\xb0\xfb\xc7\xfd\xbd\xef\x0f\x76\x36\x5b\x5f\xdc\x39\xfc\x76\xcb\xb8\xfc\x7d\xeb\xc6\x43\xdb\x37\x70\x86\xa6\x60\xd8\x8e\x7f\xd5\x11\xcb\x99\xa0\x27\xcb\x9a\x86\xcf\xa2\x38\x18\x06\x68\x4d\x53\x20\xc0\x52\x55\x6a\x90\x8e\xb8\xb2\x0a\x68\x0d\x99\x33\x4f\x4b\x53\x8a\xdc\xa8\xc7\x41\x43\x45\x4a\x51\x80\x1a\xb4\xd6\x6e\x75\xb4\xb7\xbf\x6e\x5d\x79\xd2\xba\xf1\xb0\x75\xf5\x81\xeb\xe0\x8b\x78\x0e\xbe\x4e\xdc\xe6\x3a\xf8\xda\xdb\x7f\x35\xee\x7f\x7c\xf8\xde\x83\xfd\xbd\xc7\xc6\xcd\x87\xc6\xad\x77\x5d\x27\x20\xc1\x2a\x42\x41\xc0\xa7\x8a\xf7\xf8\xa3\x88\xed\x07\x8f\x48\x07\x99\x10\xed\x73\xcf\xd6\x7d\x2e\xba\x0f\x06\xaf\xe5\x5f\x2a\x05\x02\x81\x40\xfe\xad\x1f\x2e\x85\x7e\x34\xf9\xaf\xc5\x0d\xbd\xf9\x1f\x4b\x97\xb0\x35\xf7\x36\xb9\xeb\x23\x4b\x97\x0a\x1b\x61\xb6\xd9\x85\x18\xcc\xff\xa4\xb0\x74\xe9\x87\x81\x7c\x68\x04\xff\x06\xdf\x5a\x1a\xc9\x27\x42\x3f\x87\xa1\x77\x0a\x1b\x11\x76\xa2\x63\xe5\x8f\x34\xda\x3e\xb6\xc7\xd2\xa9\xd4\xdc\x4c\x22\x33\x5b\x70\xa8\x41\x87\xd1\x96\xb0\x5f\x5f\x98\x31\x8f\x30\x89\xaa\x15\x1d\xe1\xbd\x4b\xa6\xfb\x4f\xef\x1b\x9f\x7f\x6f\xab\xeb\xab\xe7\x5b\xc6\xf5\x6f\x5b\xdb\x3b\xad\xad\x4d\x6e\x62\xff\xfb\xab\xaf\x9e\x6f\x1d\xde\x7c\xf7\xe0\xee\xe6\xe1\xbb\x9f\x1d\xbc\xbc\x4c\xd5\xb9\xfd\xe0\xd3\xd6\x77\x37\x5c\xc2\xe7\x97\x45\x09\x16\x45\xa1\xc8\x43\x45\x70\x09\xdf\x33\xc2\xa9\xc4\xcc\xe5\xb9\xd0\x68\x41\x8f\xe4\xb9\x50\xb4\xa0\x47\xf3\x5c\x68\xbc\xa0\xc7\xf2\x5c\x68\xac\xa0\xe3\x73\x3e\x56\xd0\xc7\x28\x48\x7e\x3c\x74\xa6\xc0\x05\xc9\xb1\x1c\xe0\x26\x74\xee\x8c\x1e\x09\xe3\x6a\xa4\x19\x08\x84\xf3\x1c\x3d\xef\xb9\x7c\x38\x14\x29\x04\x83\x81\x00\x29\x98\xcd\x5c\x58\x8f\x84\xf5\x68\x58\x8f\x12\x02\xd1\x26\x76\x0f\xfe\x65\xad\x60\xc9\xdd\x5f\xaa\x94\xb3\x05\xef\xe6\x36\x03\xaf\xc1\xc9\x75\xd4\x23\xd7\x54\xfa\x6c\x66\xa1\xe8\x2f\x55\xe3\xdd\xdb\x07\x0f\x1e\x1a\x57\xef\x18\xd7\x9f\xb0\x87\x9f\x3c\x3e\xd8\xdc\x76\x49\x4b\x40\x25\x51\x23\xa2\x2a\x42\x9e\x97\x1b\x92\xe6\xe7\xb5\x3a\xa9\x0c\x72\xdb\x8e\x45\x96\x84\x0d\x8e\x63\xb9\xf1\xa6\x2f\x67\xed\xb3\x9a\x4c\x93\x9c\x9a\x8a\xc0\xb0\xcc\xaa\xa8\x42\x60\xb5\x35\x24\x51\x96\xea\x70\x1d\x9b\x7f\x6b\x9e\x8c\x53\x0a\x76\x7e\x67\x70\x22\x18\xf3\x6e\xad\x6c\x3a\xd5\x53\x06\xfb\x2f\xef\xb4\xb7\xbf\x76\x70\xcf\xde\x2a\x0a\x12\xfa\xe0\x7e\x17\xfe\x89\x6d\x23\x51\x6b\xa2\xc8\xc1\xfc\x92\x1a\x2a\x04\xa8\xb3\xda\xab\xec\xf0\x62\xdd\x72\x02\x5c\xac\xe9\xe9\xb2\x48\x73\x11\x96\x9b\xf0\x97\xa2\xb5\x3f\x28\x93\x3a\x7b\xa4\x23\x5f\xc0\x50\x56\x30\xac\x4d\xdb\xfc\xcf\x94\x78\x57\x33\xd6\x00\x9f\xe6\x8e\x2e\x74\x75\xd5\xa0\xaa\x21\xa5\x07\x39\x58\x43\x6b\x3e\xcd\x82\xa8\xf2\x38\xe6\xf4\xe9\xfa\x37\xbe\xe4\x8b\x20\x21\x45\xf5\x1d\x1d\xa9\x9a\x22\xfb\xf4\x88\x92\xaa\xc1\x3a\x5c\xaf\x21\xc9\x6f\xfd\xb6\xf0\x19\xb3\xeb\x0d\x6b\xf6\xb8\xef\xa1\xe1\xca\x66\xba\x63\xe0\x5f\x7e\xd9\xfe\xe0\x1e\x4b\x6d\x3b\xad\xb8\x74\xbc\x0e\x55\x15\x07\x2e\x2e\x8d\x76\xc0\x9d\x58\x97\xb9\x7c\x8c\x84\x5d\xe3\x38\x18\xcb\xcf\xeb\x75\x3d\xa7\xab\x9d\x86\x9c\xae\xea\x53\x7a\x45\x4f\xeb\x08\x37\x4e\x90\xc6\xa9\x8a\xbe\xa0\xe9\x39\x55\x9f\xae\xea\xff\xfc\xef\x7a\x4a\xd0\x13\x50\x3f\x57\xee\x00\x9c\xd7\x97\xf5\x19\xbd\x46\x1a\x58\x2e\xec\xd8\x02\x3d\x0c\x92\xb5\x2c\xec\x46\x9b\xc5\x61\xec\x18\x2b\x70\x95\x38\xe2\x82\xcc\x37\xa8\x44\x99\x92\x2c\xaf\x98\x3f\xa2\x40\x6c\x98\x06\xab\x72\x05\x97\x44\x4d\x7c\x07\x49\xea\xb2\x58\xc7\x4e\x2c\x61\x0a\xf6\x5d\x77\x36\xf7\xf7\x1e\x63\xdf\xf5\xfe\x8b\x83\x9d\x4d\x97\x2d\x73\x09\x62\x70\x52\x9f\xf0\x48\xdd\x99\x67\xf6\x7a\x09\xad\xdf\x5e\xf6\xf1\xf9\x4c\x6f\xaf\xc8\x4b\xde\x33\xdf\x0d\x7f\xec\xc9\xc1\xf5\xa1\x08\x81\xc0\xc8\x06\xc7\x8e\x35\x03\xc6\xd6\xae\x7e\x78\xe3\x72\x70\xd2\xac\x1f\x3c\xd9\xd1\x0f\xee\x7c\x1a\xa4\x55\xe3\xfa\x13\xbb\xd0\xfa\xea\xb9\x59\xbe\x7a\xc3\xf8\xe0\x17\x16\xfe\xa3\xf7\x74\xe3\xc1\x97\x7a\xeb\xca\x93\xe0\x64\x50\x37\xc9\x1a\xd7\x5e\x9a\xfd\x87\x37\x2e\xeb\xfb\xcf\xee\x0c\x98\xbe\x77\xda\xc6\xc3\xeb\xc6\xd6\xee\x60\xa7\xfa\x1a\x34\xad\x55\xea\x87\x97\xb7\xbc\x2b\xed\x85\x62\x8d\x4e\xab\xad\x5b\xbf\x36\xdb\xdb\x7b\xef\xeb\xc6\xa3\x9d\x0e\x99\x6e\x1c\x07\x43\x03\xad\x5b\xbf\x76\x80\x53\xec\xeb\x74\x20\x13\xab\x7d\x73\xd3\x5e\xe2\xd3\xf7\x1c\xfc\xc4\x13\x3e\x78\x62\xa3\x1a\x0f\x1e\x1b\x9f\xdf\x77\x2f\x21\x18\xf4\xe8\x50\x84\x1d\x6b\x4e\x06\xda\x37\x37\xf5\x83\xcb\xdf\xb6\xfe\xbc\x67\x6c\xed\x06\xed\xf9\x4d\x06\x8c\xa7\xef\x75\x3a\xac\xd2\x93\xdf\xbb\x40\xae\xbd\xd4\x4d\x69\x92\xf5\x07\x83\x1b\x1c\x1b\x6d\x06\x7c\x58\x89\x7f\x0f\x37\x3f\xd1\xcd\x35\xea\xed\xbd\xf7\x83\x0e\x95\x20\xf2\xd2\x8d\x2f\xef\x19\xd7\xee\xe2\x45\xeb\xc6\xb3\x97\xc6\xcd\xdd\x2e\x8a\x54\x98\x4e\x4c\x22\x51\x67\x43\x47\x44\x84\x50\xeb\xca\x4b\xdc\xbb\x24\xfc\x30\x64\xfe\x1f\x0c\x6e\x84\xd9\x68\xf3\x88\xb3\xda\x69\x73\x9c\xc6\x60\x70\x26\xe7\x8c\xc7\xe4\xd8\xd7\x53\x6e\x7b\xb3\xbb\x6b\x7c\x7c\xd5\x65\x69\x24\x2b\x2c\xb6\x6d\x4c\x07\xe6\x54\x7e\xa9\x95\x58\x64\x5a\xdb\x8f\xf7\xf7\x1e\xef\xef\xee\x32\x47\x05\x6f\x4e\xd6\xd8\x93\x1e\x1c\x5f\xb8\xb0\x87\x31\xce\xab\x33\x17\x6f\x66\x12\x49\x1f\x43\x3c\x93\x48\xfa\x86\xde\x6e\xe8\x63\x19\x15\xed\xeb\x3c\xc6\x71\x11\x0c\x95\x13\xa1\x73\x05\x1c\x57\xc5\xdd\xf5\xe0\xc6\xa8\xd7\x5d\xc4\x61\xd6\x99\x44\xe8\x1c\x0c\x95\x31\x44\x3e\x1e\x2a\x60\x28\x4f\xb3\x75\xf8\x0e\x9b\x12\x00\x24\x37\xc2\x27\x04\x41\xb1\x2e\x2e\x1a\x2a\xb2\xf2\x7d\x56\x12\x66\x1d\x59\x99\xc5\xce\x39\x8d\xb4\x65\x9a\x43\x82\x3c\xc3\x02\xc6\x64\x0c\x2e\x52\x66\xe0\x52\x0d\xf2\xb8\xdd\x51\x34\x41\xec\x56\xde\x75\x06\x3b\x05\x32\x40\xb9\x7b\x2f\x1d\xfa\x3a\x83\xd9\xd6\x17\xcf\x5a\x57\x1f\xac\xa0\xf5\xd6\xad\xaf\x0e\x76\x3e\x34\xae\xfc\xf1\x54\xc7\xf2\xf1\xfb\xc3\x21\x46\x9a\x5c\x73\x60\xbf\x49\x3b\xc1\x79\x73\x52\x89\xb3\x89\x62\x76\x6e\x71\x21\x33\x3b\xd5\xcd\x9b\xc4\xd9\x04\x30\x3b\x01\x4d\x06\xbe\x7a\xbe\xd5\xda\xde\x69\x7f\xf5\xb5\xb1\x7b\xe3\xf0\x17\x5f\xb7\x3f\x7b\xbf\xfd\xfc\xb7\xd6\xed\x8b\xc5\xa6\x12\x94\x56\x8a\x25\x05\x4a\xfc\x72\x91\x97\x05\xb7\x7d\xc1\x61\xf3\x9d\x2d\xe3\xca\x87\xad\xed\x9d\xd6\xcd\xdd\xd6\xef\xdf\x3f\x65\x20\x96\x0f\x73\x91\xe8\xd8\xf8\x84\xe9\x77\x76\x6f\x8d\x4e\x77\xb4\x19\x22\xb1\x58\x68\x49\x38\x3a\xe9\xc0\x24\xce\x26\xcc\x05\x33\xdd\x61\x95\x77\x5d\xd8\xcd\xdc\xde\x39\xb8\xb3\xb5\xbf\xf7\x5f\x34\x51\x0a\x4b\x90\x61\x19\x45\x6e\x68\xa2\x54\x71\xa9\xb9\x75\x85\x3e\x40\x51\x46\x7d\x83\x8b\x54\x36\x73\x31\x9d\x2d\x4e\x67\x92\xe9\xd9\x9c\xdf\x19\x40\x62\x8b\xc3\x6f\x5e\xb4\x3f\xb8\xf7\xea\xf9\xd6\xc1\xbd\xf7\x8d\x2b\xbf\x6b\x6d\x7d\x64\x67\x92\x0e\xbf\xdd\xb2\xf2\x49\x76\x5e\x43\x11\x57\x45\xa9\x52\xac\x8a\x3c\x92\x54\xe4\x55\x7e\x9b\xde\xa9\x62\x90\xff\x89\x49\xa8\x4e\xaa\x44\x11\x49\xc4\x0a\x18\x73\xed\xb8\x68\xb2\x03\x17\x0f\xbf\x79\x71\xf8\xcd\x63\x1c\x57\xd0\x0a\x8e\x39\x0a\x1e\xa5\xfa\x5b\x65\xb2\x38\x6f\x8a\xd2\xf9\x96\xc2\x1d\x66\xee\x6c\xb7\x3f\x7a\x66\x3c\xdd\x3c\xfc\xcd\x23\x63\xeb\x03\x9f\x73\xaf\x24\x6a\xbc\x2c\xba\xcd\x5c\x6f\xac\x93\x89\x3c\xcf\x45\x0b\x79\x18\x5a\xa9\x85\xde\x49\x84\xce\xff\x53\x68\x76\x3e\xf4\x73\x2c\xa4\x8d\xc8\x18\x1b\x8d\x1e\x9d\xf5\x30\x97\xe4\x4a\xf6\xba\x1f\xe8\x0d\x90\x9f\x5d\xa9\x41\x92\x68\xee\x66\xa7\x71\xfb\xb6\xd7\xad\x12\xe4\x1a\x14\xa5\x62\x97\x77\xe5\x00\x3d\x49\x88\x46\x73\xf2\x81\x3c\xc4\x2c\xfb\xb9\x79\xbb\xea\xa8\x75\x8a\x4b\xa1\xc2\x46\x98\x9d\x18\x6b\x3a\x61\x83\x4b\x23\x27\x40\x1e\x8f\x1e\x89\xbc\x11\x61\xb9\x48\x73\x69\xc4\x6c\x32\xeb\x1d\x92\xb8\x21\x32\x8a\x83\x7e\xfd\x24\x53\xe6\xc6\x22\x03\x1a\xb6\x47\xb6\xc1\xef\xba\x60\x98\xaa\x10\x18\x36\x05\xe6\xba\x87\xad\x41\x75\x05\x09\x6c\xa9\xa1\x01\xa8\x02\x68\x3f\x8b\xd0\x50\xed\x48\x4d\xab\xe0\x1f\x5f\x55\x3b\xfc\xfc\x43\xe3\xf3\x3f\xf8\xa9\x9a\x37\x05\x9a\x99\xef\x56\xb3\xcc\x3c\xdd\x72\xaf\x9e\x6f\xe1\xed\xf7\xf1\x1f\x57\x63\xc6\x6f\xb6\x56\xc7\x5c\x8a\x97\x99\xf7\xf5\x57\x2d\xdc\x53\x78\xa5\x81\xc0\x64\x3c\x32\x4a\xee\xe7\xf5\x08\xb1\xc7\xe4\x7a\x5f\xcf\x87\x39\x72\x6d\x7f\x86\xd6\x27\xb1\xc0\xfe\x8e\x90\x5d\x0e\xc0\x64\x3c\xb0\xa4\xea\x4b\x09\x7c\x12\x04\x9c\xae\x34\xc7\xc6\x9a\xf1\xe0\xc6\x38\x3b\xde\xf4\x36\xeb\x65\x34\x11\x8e\x07\xe2\xba\xdb\xf9\xc6\x5d\x38\xc8\x8b\x35\x83\xff\x40\xdb\xa9\xc6\x71\x6c\x53\x8f\xc7\x03\xe5\x72\xb9\x1c\x88\x87\x6d\x30\x0e\x93\xc7\x3f\x81\x80\x35\xf5\x00\x9d\xbb\xce\x91\x76\xfa\x3e\xc2\x51\x5c\x1a\x09\x6e\x44\x71\x5c\xda\x27\xbc\xee\xb7\x20\xf2\xf3\xe6\x86\xec\x1a\x31\xe0\xc7\x24\x1c\x33\xf7\x9a\x9e\x37\xa6\xb1\x30\x46\x9b\x3d\x10\xa2\x3d\x10\x62\xbd\x10\x62\x3d\x10\xa2\xbd\x10\x46\x7b\x20\x44\x7a\x21\x8c\x35\xbb\xe0\x7b\x40\x8e\x37\xe3\x41\xbd\x17\x93\xc6\x9b\x7a\x3c\x18\xb4\x74\xf4\x1d\x5b\x7b\x3b\x06\x29\x33\x3f\x10\x63\x74\x8a\x63\xcf\x9b\xb4\x5e\xcc\x1d\x91\xb1\x6e\xbf\xb8\xe6\x9f\xaa\x6e\xa8\x45\xdf\x6c\x75\x17\xc2\x09\x9d\x07\x72\x16\x9c\x39\xe6\x86\x6b\x31\x97\x70\xe4\x92\xff\x8f\xa4\x95\x39\x6f\x5e\x79\x31\x57\x3c\x9b\x98\xbd\x50\x74\x3c\x24\xf6\x91\x0d\x0d\xb8\x0e\xfe\x72\xd7\x7b\x61\xd6\x50\x8b\x24\xa2\x39\xe2\xb2\xac\x07\x85\xd3\x88\x6c\xe2\xf8\x7b\xc9\x86\x8a\xa3\x27\x3c\x27\xcc\xf7\x65\xc4\xe3\x5f\x73\x76\x9d\xd2\x30\x2d\xe2\x16\x15\xae\xe2\x20\x8c\xde\x5d\xbe\xe1\xab\x1c\xce\x9b\x62\x5b\xcc\x15\x33\x0b\x7e\xee\x20\x65\x5a\x7b\xf7\xbb\xf6\xd7\xd7\xf6\x77\x77\x69\x1e\xa1\xeb\x5d\x23\x5c\xab\xc3\x75\xa4\x14\x45\x01\x49\x9a\x58\x16\x79\xf2\x00\xa5\xb7\x10\x7a\xd0\x3b\x69\xb4\x75\x86\x84\x4b\xc1\x40\x60\x9c\x8a\x85\x6b\xea\x13\xf9\x70\x68\x02\x97\xcc\x17\x7b\x11\xbb\x1c\xa3\x10\xc1\x60\xc0\xfb\x1e\xcf\x43\x2e\x1f\x02\x18\xf0\x84\x54\x2d\x2c\x37\xf1\x1e\xca\x21\x4a\x82\xb8\x2a\x0a\x0d\x88\x77\xae\xc5\x3d\x86\x65\x44\x4d\x94\x68\x0b\xd9\xe6\xb4\x51\x83\x6b\x64\x1b\xe3\xbe\x37\xac\x17\x11\x6f\x86\xb1\xc7\xf3\x56\x53\x8a\xde\x17\xab\x8e\xcd\x78\xf4\x53\x57\x7f\xf4\x93\x8a\x9f\x5e\x5b\x2f\xa9\x6f\xd1\x18\x38\xb4\x34\xb2\xa4\x16\x26\xe9\x73\x4b\xdd\xd9\xe6\x29\xc7\xba\x13\x2c\xbd\xa0\x4d\x6a\xc7\x5d\xf9\xfd\xff\xeb\xbf\xff\xad\xaf\xff\x5c\x47\x9d\xad\xec\x03\xdc\x51\xde\xdc\x6d\x6e\xaa\x38\x9b\xcd\x24\x8b\xe7\x7c\x73\x19\xbf\x7d\x68\xfc\xf2\x0b\xe3\xd6\x1d\xf7\x53\x28\x6b\x5f\xa9\x15\xdf\xd7\x53\x3d\xb0\x4e\xba\x9d\x26\xc5\x60\x20\x9f\x5b\x38\x37\x45\xa3\x90\x8d\xf1\x66\x1e\x07\x06\xc7\xd9\xb2\x32\x31\x59\x65\x51\xc2\x67\x99\xa4\x88\xbc\xf9\x33\xec\xe2\xad\xe3\x6d\xd9\x00\x99\xeb\xcd\x1a\x1e\xe9\x41\x38\x4f\x7e\x76\x65\xb5\xfd\x68\xcf\xf8\xfd\xaf\xdc\x89\xa2\x63\x3c\x88\x13\xf9\x0e\xf6\xdd\x11\xa6\x6a\xbe\x4e\x22\x03\x38\xcb\xee\xc1\x18\x96\x31\x33\xca\x57\xef\xd8\x65\x3a\x1c\xd6\x6a\xfb\xbd\x87\xe9\x25\x00\x5f\x4a\xf2\x1b\x3f\x24\xba\x5e\x0f\x26\xb2\x3d\x99\x7e\xf0\xfd\xdd\xf6\x47\x5b\x3d\x39\x6e\x26\x22\x8b\xf5\x2a\xd4\x7c\x8f\x0a\x1b\xbf\x7f\x7e\xbb\x68\xd2\x97\x61\x1d\xfe\x52\x7a\xc4\x36\x12\x28\x0a\xe4\xca\x67\x26\x06\x7d\x7f\x13\xf1\xe6\xde\x8e\x7b\x33\xd3\x93\x5d\x47\x3c\x99\x39\x11\x8b\x1c\xc1\x84\x1d\x04\xb4\x9e\xee\xb4\x5e\x7c\x77\xf8\xee\x67\x58\xe7\x76\x36\x89\xd1\x7c\xd8\x7a\xf6\xa2\xd3\xd2\xeb\xf3\x8a\xbf\xc7\xe3\x9b\x37\x1a\x9c\x44\xbc\x19\x2c\xeb\x13\x39\xb7\x6a\x3f\x78\x61\x7c\x79\x39\x93\xf2\x17\x95\x20\xba\x8d\xb4\x05\x7d\x6c\xd2\xca\x16\x92\x40\xbc\x3e\x01\xad\x8a\x3c\x72\x14\x8b\xb4\xdc\xbb\xd3\xa9\xcc\xa9\x81\x27\xe6\x23\xde\x88\xda\xff\x76\xde\xb8\xf7\x89\xf1\xf1\x55\x7f\xd6\x74\x27\x92\x09\xb0\x8b\x33\x7e\x89\x64\x93\x33\x20\xcf\x60\x0a\x58\x41\x09\x1e\x2e\x7c\x7c\xd5\xb8\xff\x29\x09\xa2\xaa\xa8\x68\xf6\xd2\x25\x60\x80\x37\x7c\x2f\x1f\xf1\xc6\xb2\xe4\x4b\xc7\x54\xe2\x67\x3e\x8e\xf3\xf6\xed\xd6\xa7\x5f\xf5\x38\x7d\x44\x45\x5b\x16\xe0\xba\xdb\x55\x26\x08\x27\x60\x8d\x45\x85\x30\x00\xe3\xe2\x03\x02\xb7\xe1\x2d\xfe\x9f\xb7\x8d\xdd\x7b\xee\xf3\xc1\xfe\x28\x73\x80\xfc\xf0\x46\x97\xd6\xc7\x98\x6e\x0d\x79\xf6\x97\xc3\x17\xef\xfb\xb3\x02\x56\x3c\x0a\x42\x60\xfb\xe5\x42\x9e\x81\x15\xa2\x1e\x04\x0b\x17\xfe\xbc\xd9\xba\xf1\xf0\x38\xfb\x15\x0e\x8d\x16\x82\x3a\xbd\x62\x9b\xec\x5c\xae\x9a\x57\x9e\x53\x83\xd6\x9a\xa8\x37\xd6\x4a\xa7\x16\x93\xe4\x1f\xc7\xf0\xe1\xd5\xfd\xbb\xc6\xb5\x0f\xfd\x79\x85\x84\x86\x19\x6e\xa3\xb5\x3a\x52\x44\x24\xf1\x1e\xe6\x11\xe4\xfe\xed\x8e\x4d\x10\x73\x8e\xe0\x32\x6c\xa7\x11\x56\x41\x09\xf2\x2b\x15\xf2\xa5\x15\x85\x38\xfc\xdd\x1f\xb0\x7b\xf2\xf4\x93\xfd\xa7\x9f\xb9\x94\xcb\xfa\x60\x76\x80\x4c\xf3\xba\xd3\xb3\x84\x63\x89\xe9\xcc\x82\xcf\x6e\xc3\x41\xe6\xa3\x9e\x46\x88\xae\x46\xd4\xdc\x1b\x8e\xe2\xf4\xaf\x6a\x0e\x3a\x98\x1b\x04\x9b\xf1\xb9\xd2\xb3\xfe\x99\x92\x01\xf2\xc2\xfb\xfc\x21\x97\xf3\x4b\xe0\x7c\xf9\x62\xff\xf9\x67\xfb\x2f\x6f\x1d\xfe\xee\x1b\xe3\xea\x1d\x7f\x66\xa8\x32\x2f\xc2\x6a\x51\x94\xd4\x86\x02\x25\x1e\x75\xc7\x18\x1e\x32\xfd\xbb\x1a\x18\xf1\xe5\x2d\xea\xc1\xe6\x72\xb3\xd8\x46\xcb\x0d\x3c\x1a\x50\x11\xdf\x50\x44\x6d\x1d\xd8\xce\x99\x6b\x90\xeb\x4f\xde\x7c\x00\x11\xf5\x06\x10\xd3\x73\xbd\xb6\x60\x7b\xef\x7a\x7b\xf7\x4f\xc6\xee\xdd\xfd\x97\x77\x5a\x9b\x3b\x3d\xdc\x59\xa8\x89\x5a\x43\x40\x45\x28\x09\xc5\xaa\x2c\x55\x68\xcd\xf1\x51\x96\x9b\xa9\x6e\x9a\xfd\x2b\x9d\x35\x0e\xf6\x63\xad\x51\x70\x19\x62\x8f\xae\x2a\x55\x30\x2f\xf7\xae\x1b\xbb\x77\xc9\xbe\xbc\xd9\xde\xbb\x8e\x5d\xdf\xaf\x5e\xd2\x02\x1d\x13\x6b\xeb\xd5\x4f\xdb\xbb\x7f\xc2\x85\x2d\x52\xf8\xdb\xa8\xad\x37\x7c\xe8\x7c\x13\xef\x62\x78\xa7\xd9\xe2\x6e\xa7\x25\xd9\xd5\x72\x64\xb0\x6b\xb1\x2d\x16\x8d\x8d\xf3\x42\x2c\x3c\xc1\x73\x25\x21\x1a\x1d\x83\x13\x61\x6e\x62\x6c\x5c\x88\x86\x21\xe4\xd1\x58\x98\x61\x99\x58\x79\x7c\x3c\x3a\x21\x94\x63\x28\x3a\xca\x9d\x11\x26\xc6\xc2\x28\x36\x3a\x1a\x83\x31\x1e\x46\xc6\x84\xd1\x30\x53\x00\xc3\xa0\x26\x8c\x06\x18\x32\x36\x13\x74\xb0\xcc\x31\xe3\xc1\x71\xab\x2b\x76\xe8\x91\x91\x23\xe9\x8c\x57\xcf\xb7\x7c\xf5\xf2\x75\x3e\x3c\xef\xb5\xaf\x69\xfe\x04\x07\x10\x37\x0e\x76\x3e\xb2\xab\x76\x46\x8b\xe6\xa8\xa8\x4a\x1d\xff\x35\x5c\xaf\x0c\xcc\x9b\xff\xb0\x31\xea\x75\xf6\x17\xfd\x9c\x7d\x8a\x8e\xa9\x03\xd3\xb1\xef\x24\x3b\x91\xe2\xf1\xf5\x3b\xc0\x45\xb3\xe7\x54\x8f\x2a\x2d\x5e\x83\x86\x28\xb0\x26\x35\x27\xb7\x16\x07\xe2\xdc\xbb\xff\xc5\x03\x24\x09\xff\x1d\x00\x00\xff\xff\x1d\x17\x43\xa9\x67\x4f\x00\x00")

func confYmlBytes() ([]byte, error) {
	return bindataRead(
		_confYml,
		"conf.yml",
	)
}

func confYml() (*asset, error) {
	bytes, err := confYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "conf.yml", size: 20327, mode: os.FileMode(0644), modTime: time.Unix(1664267981, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x3d, 0x2a, 0xa4, 0x56, 0xac, 0x42, 0xdc, 0x98, 0x7f, 0xb5, 0xc6, 0xf9, 0xe, 0xc1, 0x76, 0xd7, 0x8, 0x55, 0xd7, 0xd7, 0xfe, 0xd1, 0xc0, 0x9e, 0x81, 0x46, 0x25, 0x19, 0x63, 0x5f, 0x73, 0x68}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"conf.yml": confYml,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"conf.yml": {confYml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
