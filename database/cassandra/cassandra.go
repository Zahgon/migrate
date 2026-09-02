package cassandra

import (
	"errors"
	"io"
	"sync/atomic"

	"github.com/gocql/gocql"
	"github.com/golang-migrate/migrate/v4/database"
)

func init() {
	db := new(Cassandra)
	database.Register("cassandra", db)
}

var (
	multiStmtDelimiter = []byte(";")

	DefaultMultiStatementMaxSize = 10 * 1 << 20
)

var DefaultMigrationsTable = "schema_migrations"

var (
	ErrNilConfig     = errors.New("no config")
	ErrNoKeyspace    = errors.New("no keyspace provided")
	ErrDatabaseDirty = errors.New("database is dirty")
	ErrClosedSession = errors.New("session is closed")
)

type Config struct {
	MigrationsTable       string
	KeyspaceName          string
	MultiStatementEnabled bool
	MultiStatementMaxSize int
}

type Cassandra struct {
	session  *gocql.Session
	isLocked atomic.Bool

	config *Config
}

func WithInstance(session *gocql.Session, config *Config) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func (c *Cassandra) Open(url string) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func (c *Cassandra) Close() error { _ = "STUB: not implemented"; return nil }

func (c *Cassandra) Lock() error { _ = "STUB: not implemented"; return nil }

func (c *Cassandra) Unlock() error { _ = "STUB: not implemented"; return nil }

func (c *Cassandra) Run(migration io.Reader) error { _ = "STUB: not implemented"; return nil }

func (c *Cassandra) SetVersion(version int, dirty bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cassandra) Version() (version int, dirty bool, err error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

func (c *Cassandra) Drop() error { _ = "STUB: not implemented"; return nil }

func (c *Cassandra) ensureVersionTable() (err error) { _ = "STUB: not implemented"; return nil }

func parseConsistency(consistencyStr string) (consistency gocql.Consistency, err error) {
	_ = "STUB: not implemented"
	return *new(gocql.Consistency), nil
}
