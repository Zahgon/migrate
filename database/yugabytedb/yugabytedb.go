package yugabytedb

import (
	"context"
	"database/sql"
	"errors"
	"io"
	"sync/atomic"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/golang-migrate/migrate/v4/database"
)

const (
	DefaultMaxRetryInterval    = time.Second * 15
	DefaultMaxRetryElapsedTime = time.Second * 30
	DefaultMaxRetries          = 10
	DefaultMigrationsTable     = "migrations"
	DefaultLockTable           = "migrations_locks"
)

var (
	ErrNilConfig          = errors.New("no config")
	ErrNoDatabaseName     = errors.New("no database name")
	ErrMaxRetriesExceeded = errors.New("max retries exceeded")
)

func init() {
	db := YugabyteDB{}
	database.Register("yugabyte", &db)
	database.Register("yugabytedb", &db)
	database.Register("ysql", &db)
}

type Config struct {
	MigrationsTable     string
	LockTable           string
	ForceLock           bool
	DatabaseName        string
	MaxRetryInterval    time.Duration
	MaxRetryElapsedTime time.Duration
	MaxRetries          int
}

type YugabyteDB struct {
	db       *sql.DB
	isLocked atomic.Bool

	config *Config
}

func WithInstance(instance *sql.DB, config *Config) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func (c *YugabyteDB) Open(dbURL string) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func (c *YugabyteDB) Close() error { _ = "STUB: not implemented"; return nil }

func (c *YugabyteDB) Lock() error { _ = "STUB: not implemented"; return nil }

func (c *YugabyteDB) Unlock() error { _ = "STUB: not implemented"; return nil }

func (c *YugabyteDB) Run(migration io.Reader) error { _ = "STUB: not implemented"; return nil }

func (c *YugabyteDB) SetVersion(version int, dirty bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *YugabyteDB) Version() (version int, dirty bool, err error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

func (c *YugabyteDB) Drop() (err error) { _ = "STUB: not implemented"; return nil }

func (c *YugabyteDB) ensureVersionTable() (err error) { _ = "STUB: not implemented"; return nil }

func (c *YugabyteDB) ensureLockTable() error { _ = "STUB: not implemented"; return nil }

func (c *YugabyteDB) doTxWithRetry(
	ctx context.Context,
	txOpts *sql.TxOptions,
	fn func(tx *sql.Tx) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *YugabyteDB) newBackoff(ctx context.Context) backoff.BackOff {
	_ = "STUB: not implemented"
	return *new(backoff.BackOff)
}

func errIsRetryable(err error) bool { _ = "STUB: not implemented"; return false }
