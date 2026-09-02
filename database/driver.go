package database

import (
	"fmt"
	"io"
	"sync"
)

var (
	ErrLocked    = fmt.Errorf("can't acquire lock")
	ErrNotLocked = fmt.Errorf("can't unlock, as not currently locked")
)

const NilVersion int = -1

var driversMu sync.RWMutex
var drivers = make(map[string]Driver)

type Driver interface {
	Open(url string) (Driver, error)

	Close() error

	Lock() error

	Unlock() error

	Run(migration io.Reader) error

	SetVersion(version int, dirty bool) error

	Version() (version int, dirty bool, err error)

	Drop() error
}

func Open(url string) (Driver, error) { _ = "STUB: not implemented"; return *new(Driver), nil }

func Register(name string, driver Driver) { _ = "STUB: not implemented"; return }

func List() []string { _ = "STUB: not implemented"; return nil }
