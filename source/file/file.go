package file

import (
	"github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

func init() {
	source.Register("file", &File{})
}

type File struct {
	iofs.PartialDriver
	url  string
	path string
}

func (f *File) Open(url string) (source.Driver, error) {
	_ = "STUB: not implemented"
	return *new(source.Driver), nil
}

func parseURL(url string) (string, error) { _ = "STUB: not implemented"; return "", nil }
