package cockroachdb

import (
	"database/sql"
	"fmt"
	"io"
	"sync/atomic"

	"github.com/golang-migrate/migrate/v4/database"
)

func init() {
	db := CockroachDb{}
	database.Register("cockroach", &db)
	database.Register("cockroachdb", &db)
	database.Register("crdb-postgres", &db)
}

var DefaultMigrationsTable = "schema_migrations"
var DefaultLockTable = "schema_lock"

var (
	ErrNilConfig      = fmt.Errorf("no config")
	ErrNoDatabaseName = fmt.Errorf("no database name")
)

type Config struct {
	MigrationsTable string
	LockTable       string
	ForceLock       bool
	DatabaseName    string
}

type CockroachDb struct {
	db       *sql.DB
	isLocked atomic.Bool

	config *Config
}

func WithInstance(instance *sql.DB, config *Config) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func (c *CockroachDb) Open(url string) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func (c *CockroachDb) Close() error { _ = "STUB: not implemented"; return nil }

func (c *CockroachDb) Lock() error { _ = "STUB: not implemented"; return nil }

func (c *CockroachDb) Unlock() error { _ = "STUB: not implemented"; return nil }

func (c *CockroachDb) Run(migration io.Reader) error { _ = "STUB: not implemented"; return nil }

func (c *CockroachDb) SetVersion(version int, dirty bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CockroachDb) Version() (version int, dirty bool, err error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

func (c *CockroachDb) Drop() (err error) { _ = "STUB: not implemented"; return nil }

func (c *CockroachDb) ensureVersionTable() (err error) { _ = "STUB: not implemented"; return nil }

func (c *CockroachDb) ensureLockTable() error { _ = "STUB: not implemented"; return nil }
