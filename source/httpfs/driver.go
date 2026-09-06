package httpfs

import (
	"net/http"

	"github.com/golang-migrate/migrate/v4/source"
)

type driver struct {
	PartialDriver
}

func New(fs http.FileSystem, path string) (source.Driver, error) {
	_ = "STUB: not implemented"
	return *new(source.Driver), nil
}

func (d *driver) Open(url string) (source.Driver, error) {
	_ = "STUB: not implemented"
	return *new(source.Driver), nil
}
