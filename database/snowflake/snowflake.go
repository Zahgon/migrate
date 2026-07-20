package snowflake

import (
	"database/sql"
	"fmt"
	"io"
	"sync/atomic"

	"github.com/golang-migrate/migrate/v4/database"
)

func init() {
	db := Snowflake{}
	database.Register("snowflake", &db)
}

var DefaultMigrationsTable = "schema_migrations"

var (
	ErrNilConfig          = fmt.Errorf("no config")
	ErrNoDatabaseName     = fmt.Errorf("no database name")
	ErrNoPassword         = fmt.Errorf("no password")
	ErrNoSchema           = fmt.Errorf("no schema")
	ErrNoSchemaOrDatabase = fmt.Errorf("no schema/database name")
)

type Config struct {
	MigrationsTable string
	DatabaseName    string
}

type Snowflake struct {
	isLocked atomic.Bool
	conn     *sql.Conn
	db       *sql.DB

	config *Config
}

func WithInstance(instance *sql.DB, config *Config) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func (p *Snowflake) Open(url string) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func (p *Snowflake) Close() error { _ = "STUB: not implemented"; return nil }

func (p *Snowflake) Lock() error { _ = "STUB: not implemented"; return nil }

func (p *Snowflake) Unlock() error { _ = "STUB: not implemented"; return nil }

func (p *Snowflake) Run(migration io.Reader) error { _ = "STUB: not implemented"; return nil }

func computeLineFromPos(s string, pos int) (line uint, col uint, ok bool) {
	_ = "STUB: not implemented"
	return 0, 0, false
}

const newLine = '\n'

func runesCount(input []rune, target rune) int { _ = "STUB: not implemented"; return 0 }

func runesLastIndex(input []rune, target rune) int { _ = "STUB: not implemented"; return 0 }

func (p *Snowflake) SetVersion(version int, dirty bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Snowflake) Version() (version int, dirty bool, err error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

func (p *Snowflake) Drop() (err error) { _ = "STUB: not implemented"; return nil }

func (p *Snowflake) ensureVersionTable() (err error) { _ = "STUB: not implemented"; return nil }
