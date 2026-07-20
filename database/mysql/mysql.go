//go:build go1.9

package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"sync/atomic"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4/database"
)

var _ database.Driver = (*Mysql)(nil)

func init() {
	database.Register("mysql", &Mysql{})
}

var DefaultMigrationsTable = "schema_migrations"

var (
	ErrDatabaseDirty    = fmt.Errorf("database is dirty")
	ErrNilConfig        = fmt.Errorf("no config")
	ErrNoDatabaseName   = fmt.Errorf("no database name")
	ErrAppendPEM        = fmt.Errorf("failed to append PEM")
	ErrTLSCertKeyConfig = fmt.Errorf("to use TLS client authentication, both x-tls-cert and x-tls-key must not be empty")
)

type Config struct {
	MigrationsTable  string
	DatabaseName     string
	NoLock           bool
	StatementTimeout time.Duration
}

type Mysql struct {
	conn     *sql.Conn
	db       *sql.DB
	isLocked atomic.Bool

	config *Config
}

func WithConnection(ctx context.Context, conn *sql.Conn, config *Config) (*Mysql, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func WithInstance(instance *sql.DB, config *Config) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func extractCustomQueryParams(c *mysql.Config) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func urlToMySQLConfig(url string) (*mysql.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Mysql) Open(url string) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func (m *Mysql) Close() error { _ = "STUB: not implemented"; return nil }

func (m *Mysql) Lock() error { _ = "STUB: not implemented"; return nil }

func (m *Mysql) Unlock() error { _ = "STUB: not implemented"; return nil }

func (m *Mysql) Run(migration io.Reader) error { _ = "STUB: not implemented"; return nil }

func (m *Mysql) SetVersion(version int, dirty bool) error { _ = "STUB: not implemented"; return nil }

func (m *Mysql) Version() (version int, dirty bool, err error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

func (m *Mysql) Drop() (err error) { _ = "STUB: not implemented"; return nil }

func (m *Mysql) ensureVersionTable() (err error) { _ = "STUB: not implemented"; return nil }

func readBool(input string) (value bool, valid bool) {
	_ = "STUB: not implemented"
	return false, false
}
