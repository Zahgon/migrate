//go:build go1.16

package iofs

import (
	"io"
	"io/fs"

	"github.com/golang-migrate/migrate/v4/source"
)

type driver struct {
	PartialDriver
}

func New(fsys fs.FS, path string) (source.Driver, error) {
	_ = "STUB: not implemented"
	return *new(source.Driver), nil
}

func (d *driver) Open(url string) (source.Driver, error) {
	_ = "STUB: not implemented"
	return *new(source.Driver), nil
}

type PartialDriver struct {
	migrations *source.Migrations
	fsys       fs.FS
	path       string
}

func (d *PartialDriver) Init(fsys fs.FS, path string) error { _ = "STUB: not implemented"; return nil }

func (d *PartialDriver) Close() error { _ = "STUB: not implemented"; return nil }

func (d *PartialDriver) First() (version uint, err error) { _ = "STUB: not implemented"; return 0, nil }

func (d *PartialDriver) Prev(version uint) (prevVersion uint, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *PartialDriver) Next(version uint) (nextVersion uint, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *PartialDriver) ReadUp(version uint) (r io.ReadCloser, identifier string, err error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), "", nil
}

func (d *PartialDriver) ReadDown(version uint) (r io.ReadCloser, identifier string, err error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), "", nil
}

func (d *PartialDriver) open(path string) (fs.File, error) {
	_ = "STUB: not implemented"
	return *new(fs.File), nil
}
