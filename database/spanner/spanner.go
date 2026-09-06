package spanner

import (
	"errors"
	"io"
	"regexp"
	"sync/atomic"

	"cloud.google.com/go/spanner"
	sdb "cloud.google.com/go/spanner/admin/database/apiv1"

	"github.com/golang-migrate/migrate/v4/database"
)

func init() {
	db := Spanner{}
	database.Register("spanner", &db)
}

const DefaultMigrationsTable = "SchemaMigrations"

var (
	ErrNilConfig      = errors.New("no config")
	ErrNoDatabaseName = errors.New("no database name")
	ErrNoSchema       = errors.New("no schema")
	ErrDatabaseDirty  = errors.New("database is dirty")
	ErrLockHeld       = errors.New("unable to obtain lock")
	ErrLockNotHeld    = errors.New("unable to release already released lock")
)

type Config struct {
	MigrationsTable string
	DatabaseName    string

	CleanStatements bool
}

type Spanner struct {
	db *DB

	config *Config

	lock atomic.Bool
}

type DB struct {
	admin *sdb.DatabaseAdminClient
	data  *spanner.Client
}

func NewDB(admin sdb.DatabaseAdminClient, data spanner.Client) *DB {
	_ = "STUB: not implemented"
	return nil
}

func WithInstance(instance *DB, config *Config) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func (s *Spanner) Open(url string) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func (s *Spanner) Close() error { _ = "STUB: not implemented"; return nil }

func (s *Spanner) Lock() error { _ = "STUB: not implemented"; return nil }

func (s *Spanner) Unlock() error { _ = "STUB: not implemented"; return nil }

func (s *Spanner) Run(migration io.Reader) error { _ = "STUB: not implemented"; return nil }

func (s *Spanner) SetVersion(version int, dirty bool) error { _ = "STUB: not implemented"; return nil }

func (s *Spanner) Version() (version int, dirty bool, err error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

var nameMatcher = regexp.MustCompile(`(CREATE TABLE\s(\S+)\s)|(CREATE.+INDEX\s(\S+)\s)`)

func (s *Spanner) Drop() error { _ = "STUB: not implemented"; return nil }

func (s *Spanner) ensureVersionTable() (err error) { _ = "STUB: not implemented"; return nil }

func cleanStatements(migration []byte) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
