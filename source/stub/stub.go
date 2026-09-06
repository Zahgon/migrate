package stub

import (
	"io"

	"github.com/golang-migrate/migrate/v4/source"
)

func init() {
	source.Register("stub", &Stub{})
}

type Config struct{}

type Stub struct {
	Url        string
	Instance   interface{}
	Migrations *source.Migrations
	Config     *Config
}

func (s *Stub) Open(url string) (source.Driver, error) {
	_ = "STUB: not implemented"
	return *new(source.Driver), nil
}

func WithInstance(instance interface{}, config *Config) (source.Driver, error) {
	_ = "STUB: not implemented"
	return *new(source.Driver), nil
}

func (s *Stub) Close() error { _ = "STUB: not implemented"; return nil }

func (s *Stub) First() (version uint, err error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Stub) Prev(version uint) (prevVersion uint, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Stub) Next(version uint) (nextVersion uint, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Stub) ReadUp(version uint) (r io.ReadCloser, identifier string, err error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), "", nil
}

func (s *Stub) ReadDown(version uint) (r io.ReadCloser, identifier string, err error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), "", nil
}
