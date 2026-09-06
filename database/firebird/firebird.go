//go:build go1.9

package firebird

import (
	"database/sql"
	"fmt"
	"io"
	"sync/atomic"

	"github.com/golang-migrate/migrate/v4/database"
	_ "github.com/nakagami/firebirdsql"
)

func init() {
	db := Firebird{}
	database.Register("firebird", &db)
	database.Register("firebirdsql", &db)
}

var DefaultMigrationsTable = "schema_migrations"

var (
	ErrNilConfig = fmt.Errorf("no config")
)

type Config struct {
	DatabaseName    string
	MigrationsTable string
}

type Firebird struct {
	conn     *sql.Conn
	db       *sql.DB
	isLocked atomic.Bool

	config *Config
}

func WithInstance(instance *sql.DB, config *Config) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func (f *Firebird) Open(dsn string) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func (f *Firebird) Close() error { _ = "STUB: not implemented"; return nil }

func (f *Firebird) Lock() error { _ = "STUB: not implemented"; return nil }

func (f *Firebird) Unlock() error { _ = "STUB: not implemented"; return nil }

func (f *Firebird) Run(migration io.Reader) error { _ = "STUB: not implemented"; return nil }

func (f *Firebird) SetVersion(version int, dirty bool) error { _ = "STUB: not implemented"; return nil }

func (f *Firebird) Version() (version int, dirty bool, err error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

func (f *Firebird) Drop() (err error) { _ = "STUB: not implemented"; return nil }

func (f *Firebird) ensureVersionTable() (err error) { _ = "STUB: not implemented"; return nil }

func btoi(v bool) int { _ = "STUB: not implemented"; return 0 }

func itob(v int) bool { _ = "STUB: not implemented"; return false }
