package sqlite

import (
	"database/sql"
	"fmt"
	"io"
	"sync/atomic"

	"github.com/golang-migrate/migrate/v4/database"
	_ "modernc.org/sqlite"
)

func init() {
	database.Register("sqlite", &Sqlite{})
}

var DefaultMigrationsTable = "schema_migrations"
var (
	ErrDatabaseDirty  = fmt.Errorf("database is dirty")
	ErrNilConfig      = fmt.Errorf("no config")
	ErrNoDatabaseName = fmt.Errorf("no database name")
)

type Config struct {
	MigrationsTable string
	DatabaseName    string
	NoTxWrap        bool
}

type Sqlite struct {
	db       *sql.DB
	isLocked atomic.Bool

	config *Config
}

func WithInstance(instance *sql.DB, config *Config) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func (m *Sqlite) ensureVersionTable() (err error) { _ = "STUB: not implemented"; return nil }

func (m *Sqlite) Open(url string) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func (m *Sqlite) Close() error { _ = "STUB: not implemented"; return nil }

func (m *Sqlite) Drop() (err error) { _ = "STUB: not implemented"; return nil }

func (m *Sqlite) Lock() error { _ = "STUB: not implemented"; return nil }

func (m *Sqlite) Unlock() error { _ = "STUB: not implemented"; return nil }

func (m *Sqlite) Run(migration io.Reader) error { _ = "STUB: not implemented"; return nil }

func (m *Sqlite) executeQuery(query string) error { _ = "STUB: not implemented"; return nil }

func (m *Sqlite) executeQueryNoTx(query string) error { _ = "STUB: not implemented"; return nil }

func (m *Sqlite) SetVersion(version int, dirty bool) error { _ = "STUB: not implemented"; return nil }

func (m *Sqlite) Version() (version int, dirty bool, err error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}
