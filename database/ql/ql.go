package ql

import (
	"database/sql"
	"fmt"
	"io"
	"sync/atomic"

	"github.com/golang-migrate/migrate/v4/database"
	_ "modernc.org/ql/driver"
)

func init() {
	database.Register("ql", &Ql{})
}

var DefaultMigrationsTable = "schema_migrations"
var (
	ErrDatabaseDirty  = fmt.Errorf("database is dirty")
	ErrNilConfig      = fmt.Errorf("no config")
	ErrNoDatabaseName = fmt.Errorf("no database name")
	ErrAppendPEM      = fmt.Errorf("failed to append PEM")
)

type Config struct {
	MigrationsTable string
	DatabaseName    string
}

type Ql struct {
	db       *sql.DB
	isLocked atomic.Bool

	config *Config
}

func WithInstance(instance *sql.DB, config *Config) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func (m *Ql) ensureVersionTable() (err error) { _ = "STUB: not implemented"; return nil }

func (m *Ql) Open(url string) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func (m *Ql) Close() error { _ = "STUB: not implemented"; return nil }

func (m *Ql) Drop() (err error) { _ = "STUB: not implemented"; return nil }

func (m *Ql) Lock() error { _ = "STUB: not implemented"; return nil }

func (m *Ql) Unlock() error { _ = "STUB: not implemented"; return nil }

func (m *Ql) Run(migration io.Reader) error { _ = "STUB: not implemented"; return nil }

func (m *Ql) executeQuery(query string) error { _ = "STUB: not implemented"; return nil }

func (m *Ql) SetVersion(version int, dirty bool) error { _ = "STUB: not implemented"; return nil }

func (m *Ql) Version() (version int, dirty bool, err error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}
