package godoc_vfs

import (
	"github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/httpfs"

	"golang.org/x/tools/godoc/vfs"
)

func init() {
	source.Register("godoc-vfs", &VFS{})
}

type VFS struct {
	httpfs.PartialDriver
	fs   vfs.FileSystem
	path string
}

func (b *VFS) Open(url string) (source.Driver, error) {
	_ = "STUB: not implemented"
	return *new(source.Driver), nil
}

func WithInstance(fs vfs.FileSystem, searchPath string) (source.Driver, error) {
	_ = "STUB: not implemented"
	return *new(source.Driver), nil
}
