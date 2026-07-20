package mongodb

import (
	"context"
	"fmt"
	"io"
	"sync/atomic"
	"time"

	"github.com/golang-migrate/migrate/v4/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	db := Mongo{}
	database.Register("mongodb", &db)
	database.Register("mongodb+srv", &db)
}

var DefaultMigrationsCollection = "schema_migrations"

const DefaultLockingCollection = "migrate_advisory_lock"
const lockKeyUniqueValue = 0
const DefaultLockTimeout = 15
const DefaultLockTimeoutInterval = 10
const DefaultAdvisoryLockingFlag = true
const LockIndexName = "lock_unique_key"
const contextWaitTimeout = 5 * time.Second

var (
	ErrNoDatabaseName            = fmt.Errorf("no database name")
	ErrNilConfig                 = fmt.Errorf("no config")
	ErrLockTimeoutConfigConflict = fmt.Errorf("both x-advisory-lock-timeout-interval and x-advisory-lock-timout-interval were specified")
)

type Mongo struct {
	client   *mongo.Client
	db       *mongo.Database
	config   *Config
	isLocked atomic.Bool
}

type Locking struct {
	CollectionName string
	Timeout        int
	Enabled        bool
	Interval       int
}
type Config struct {
	DatabaseName         string
	MigrationsCollection string
	TransactionMode      bool
	Locking              Locking
}
type versionInfo struct {
	Version int  `bson:"version"`
	Dirty   bool `bson:"dirty"`
}

type lockObj struct {
	Key       int       `bson:"locking_key"`
	Pid       int       `bson:"pid"`
	Hostname  string    `bson:"hostname"`
	CreatedAt time.Time `bson:"created_at"`
}
type findFilter struct {
	Key int `bson:"locking_key"`
}

func WithInstance(instance *mongo.Client, config *Config) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func (m *Mongo) Open(dsn string) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func parseBoolean(urlParam string, defaultValue bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func parseInt(urlParam string, defaultValue int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *Mongo) SetVersion(version int, dirty bool) error { _ = "STUB: not implemented"; return nil }

func (m *Mongo) Version() (version int, dirty bool, err error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

func (m *Mongo) Run(migration io.Reader) error { _ = "STUB: not implemented"; return nil }

func (m *Mongo) executeCommandsWithTransaction(ctx context.Context, cmds []bson.D) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Mongo) executeCommands(ctx context.Context, cmds []bson.D) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Mongo) Close() error { _ = "STUB: not implemented"; return nil }

func (m *Mongo) Drop() error { _ = "STUB: not implemented"; return nil }

func (m *Mongo) ensureLockTable() error { _ = "STUB: not implemented"; return nil }

func (m *Mongo) ensureVersionTable() (err error) { _ = "STUB: not implemented"; return nil }

func (m *Mongo) Lock() error { _ = "STUB: not implemented"; return nil }

func (m *Mongo) Unlock() error { _ = "STUB: not implemented"; return nil }
