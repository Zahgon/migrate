package httpfs

import (
	"io"
	"net/http"

	"github.com/golang-migrate/migrate/v4/source"
)

type PartialDriver struct {
	migrations *source.Migrations
	fs         http.FileSystem
	path       string
}

func (p *PartialDriver) Init(fs http.FileSystem, path string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *PartialDriver) Close() error { _ = "STUB: not implemented"; return nil }

func (p *PartialDriver) First() (version uint, err error) { _ = "STUB: not implemented"; return 0, nil }

func (p *PartialDriver) Prev(version uint) (prevVersion uint, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (p *PartialDriver) Next(version uint) (nextVersion uint, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (p *PartialDriver) ReadUp(version uint) (r io.ReadCloser, identifier string, err error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), "", nil
}

func (p *PartialDriver) ReadDown(version uint) (r io.ReadCloser, identifier string, err error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), "", nil
}

func (p *PartialDriver) open(path string) (http.File, error) {
	_ = "STUB: not implemented"
	return *new(http.File), nil
}
