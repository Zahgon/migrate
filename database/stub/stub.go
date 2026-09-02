package stub

import (
	"io"
	"sync/atomic"

	"github.com/golang-migrate/migrate/v4/database"
)

func init() {
	database.Register("stub", &Stub{})
}

type Stub struct {
	Url               string
	Instance          interface{}
	CurrentVersion    int
	MigrationSequence []string
	LastRunMigration  []byte
	IsDirty           bool
	isLocked          atomic.Bool

	Config *Config
}

func (s *Stub) Open(url string) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

type Config struct{}

func WithInstance(instance interface{}, config *Config) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func (s *Stub) Close() error { _ = "STUB: not implemented"; return nil }

func (s *Stub) Lock() error { _ = "STUB: not implemented"; return nil }

func (s *Stub) Unlock() error { _ = "STUB: not implemented"; return nil }

func (s *Stub) Run(migration io.Reader) error { _ = "STUB: not implemented"; return nil }

func (s *Stub) SetVersion(version int, state bool) error { _ = "STUB: not implemented"; return nil }

func (s *Stub) Version() (version int, dirty bool, err error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

const DROP = "DROP"

func (s *Stub) Drop() error { _ = "STUB: not implemented"; return nil }

func (s *Stub) EqualSequence(seq []string) bool { _ = "STUB: not implemented"; return false }
