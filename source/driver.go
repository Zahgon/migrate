package source

import (
	"io"
	"sync"
)

var driversMu sync.RWMutex
var drivers = make(map[string]Driver)

type Driver interface {
	Open(url string) (Driver, error)

	Close() error

	First() (version uint, err error)

	Prev(version uint) (prevVersion uint, err error)

	Next(version uint) (nextVersion uint, err error)

	ReadUp(version uint) (r io.ReadCloser, identifier string, err error)

	ReadDown(version uint) (r io.ReadCloser, identifier string, err error)
}

func Open(url string) (Driver, error) { _ = "STUB: not implemented"; return *new(Driver), nil }

func Register(name string, driver Driver) { _ = "STUB: not implemented"; return }

func List() []string { _ = "STUB: not implemented"; return nil }
