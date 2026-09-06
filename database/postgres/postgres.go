//go:build go1.9

package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"sync/atomic"
	"time"

	"github.com/golang-migrate/migrate/v4/database"
)

func init() {
	db := Postgres{}
	database.Register("postgres", &db)
	database.Register("postgresql", &db)
}

var (
	multiStmtDelimiter = []byte(";")

	DefaultMigrationsTable       = "schema_migrations"
	DefaultMultiStatementMaxSize = 10 * 1 << 20
)

var (
	ErrNilConfig      = fmt.Errorf("no config")
	ErrNoDatabaseName = fmt.Errorf("no database name")
	ErrNoSchema       = fmt.Errorf("no schema")
	ErrDatabaseDirty  = fmt.Errorf("database is dirty")
)

type Config struct {
	MigrationsTable       string
	MigrationsTableQuoted bool
	MultiStatementEnabled bool
	DatabaseName          string
	SchemaName            string
	migrationsSchemaName  string
	migrationsTableName   string
	StatementTimeout      time.Duration
	MultiStatementMaxSize int
}

type Postgres struct {
	conn     *sql.Conn
	db       *sql.DB
	isLocked atomic.Bool

	config *Config
}

func WithConnection(ctx context.Context, conn *sql.Conn, config *Config) (*Postgres, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func WithInstance(instance *sql.DB, config *Config) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func (p *Postgres) Open(url string) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func (p *Postgres) Close() error { _ = "STUB: not implemented"; return nil }

func (p *Postgres) Lock() error { _ = "STUB: not implemented"; return nil }

func (p *Postgres) Unlock() error { _ = "STUB: not implemented"; return nil }

func (p *Postgres) Run(migration io.Reader) error { _ = "STUB: not implemented"; return nil }

func (p *Postgres) runStatement(statement []byte) error { _ = "STUB: not implemented"; return nil }

func computeLineFromPos(s string, pos int) (line uint, col uint, ok bool) {
	_ = "STUB: not implemented"
	return 0, 0, false
}

const newLine = '\n'

func runesCount(input []rune, target rune) int { _ = "STUB: not implemented"; return 0 }

func runesLastIndex(input []rune, target rune) int { _ = "STUB: not implemented"; return 0 }

func (p *Postgres) SetVersion(version int, dirty bool) error { _ = "STUB: not implemented"; return nil }

func (p *Postgres) Version() (version int, dirty bool, err error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

func (p *Postgres) Drop() (err error) { _ = "STUB: not implemented"; return nil }

func (p *Postgres) ensureVersionTable() (err error) { _ = "STUB: not implemented"; return nil }
