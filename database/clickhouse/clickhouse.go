package clickhouse

import (
	"database/sql"
	"fmt"
	"io"
	"sync/atomic"

	"github.com/golang-migrate/migrate/v4/database"
)

var (
	multiStmtDelimiter = []byte(";")

	DefaultMigrationsTable       = "schema_migrations"
	DefaultMigrationsTableEngine = "TinyLog"
	DefaultMultiStatementMaxSize = 10 * 1 << 20

	ErrNilConfig = fmt.Errorf("no config")
)

type Config struct {
	DatabaseName          string
	ClusterName           string
	MigrationsTable       string
	MigrationsTableEngine string
	MultiStatementEnabled bool
	MultiStatementMaxSize int
}

func init() {
	database.Register("clickhouse", &ClickHouse{})
}

func WithInstance(conn *sql.DB, config *Config) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

type ClickHouse struct {
	conn     *sql.DB
	config   *Config
	isLocked atomic.Bool
}

func (ch *ClickHouse) Open(dsn string) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func (ch *ClickHouse) init() error {
	if len(ch.config.DatabaseName) == 0 {
		if err := ch.conn.QueryRow("SELECT currentDatabase()").Scan(&ch.config.DatabaseName); err != nil {
			return err
		}
	}

	if len(ch.config.MigrationsTable) == 0 {
		ch.config.MigrationsTable = DefaultMigrationsTable
	}

	if ch.config.MultiStatementMaxSize <= 0 {
		ch.config.MultiStatementMaxSize = DefaultMultiStatementMaxSize
	}

	if len(ch.config.MigrationsTableEngine) == 0 {
		ch.config.MigrationsTableEngine = DefaultMigrationsTableEngine
	}

	return ch.ensureVersionTable()
}

func (ch *ClickHouse) Run(r io.Reader) error { _ = "STUB: not implemented"; return nil }

func (ch *ClickHouse) Version() (int, bool, error) { _ = "STUB: not implemented"; return 0, false, nil }

func (ch *ClickHouse) SetVersion(version int, dirty bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (ch *ClickHouse) ensureVersionTable() (err error) { _ = "STUB: not implemented"; return nil }

func (ch *ClickHouse) Drop() (err error) { _ = "STUB: not implemented"; return nil }

func (ch *ClickHouse) Lock() error { _ = "STUB: not implemented"; return nil }

func (ch *ClickHouse) Unlock() error { _ = "STUB: not implemented"; return nil }

func (ch *ClickHouse) Close() error { _ = "STUB: not implemented"; return nil }

func quoteIdentifier(name string) string { _ = "STUB: not implemented"; return "" }
