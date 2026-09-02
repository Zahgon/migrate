package pkger

import (
	"net/http"

	"github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	"github.com/markbates/pkger/pkging"
)

func init() {
	source.Register("pkger", &Pkger{})
}

type Pkger struct {
	httpfs.PartialDriver
}

func (p *Pkger) Open(url string) (source.Driver, error) {
	_ = "STUB: not implemented"
	return *new(source.Driver), nil
}

func WithInstance(instance pkging.Pkger, path string) (source.Driver, error) {
	_ = "STUB: not implemented"
	return *new(source.Driver), nil
}

type fsFunc func(name string) (http.File, error)

func (f fsFunc) Open(name string) (http.File, error) {
	_ = "STUB: not implemented"
	return *new(http.File), nil
}
