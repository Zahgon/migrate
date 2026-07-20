//go:build go1.9

package redshift

import (
	"database/sql"
	"fmt"
	"io"
	"sync/atomic"

	"github.com/golang-migrate/migrate/v4/database"
)

func init() {
	db := Redshift{}
	database.Register("redshift", &db)
}

var DefaultMigrationsTable = "schema_migrations"

var (
	ErrNilConfig      = fmt.Errorf("no config")
	ErrNoDatabaseName = fmt.Errorf("no database name")
)

type Config struct {
	MigrationsTable string
	DatabaseName    string
}

type Redshift struct {
	isLocked atomic.Bool
	conn     *sql.Conn
	db       *sql.DB

	config *Config
}

func WithInstance(instance *sql.DB, config *Config) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func (p *Redshift) Open(url string) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func (p *Redshift) Close() error { _ = "STUB: not implemented"; return nil }

func (p *Redshift) Lock() error { _ = "STUB: not implemented"; return nil }

func (p *Redshift) Unlock() error { _ = "STUB: not implemented"; return nil }

func (p *Redshift) Run(migration io.Reader) error { _ = "STUB: not implemented"; return nil }

func computeLineFromPos(s string, pos int) (line uint, col uint, ok bool) {
	_ = "STUB: not implemented"
	return 0, 0, false
}

const newLine = '\n'

func runesCount(input []rune, target rune) int { _ = "STUB: not implemented"; return 0 }

func runesLastIndex(input []rune, target rune) int { _ = "STUB: not implemented"; return 0 }

func (p *Redshift) SetVersion(version int, dirty bool) error { _ = "STUB: not implemented"; return nil }

func (p *Redshift) Version() (version int, dirty bool, err error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

func (p *Redshift) Drop() (err error) { _ = "STUB: not implemented"; return nil }

func (p *Redshift) ensureVersionTable() (err error) { _ = "STUB: not implemented"; return nil }
