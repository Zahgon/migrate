package migrate

import (
	"errors"
	"sync"
	"time"

	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/source"
)

var DefaultPrefetchMigrations = uint(10)

var DefaultLockTimeout = 15 * time.Second

var (
	ErrNoChange       = errors.New("no change")
	ErrNilVersion     = errors.New("no migration")
	ErrInvalidVersion = errors.New("version must be >= -1")
	ErrLocked         = errors.New("database locked")
	ErrLockTimeout    = errors.New("timeout: can't acquire database lock")
)

type ErrShortLimit struct {
	Short uint
}

func (e ErrShortLimit) Error() string { _ = "STUB: not implemented"; return "" }

type ErrDirty struct {
	Version int
}

func (e ErrDirty) Error() string { _ = "STUB: not implemented"; return "" }

type Migrate struct {
	sourceName         string
	sourceDrv          source.Driver
	databaseDriverName string
	databaseDrv        database.Driver

	Log Logger

	GracefulStop chan bool
	isLockedMu   *sync.Mutex

	isGracefulStop bool
	isLocked       bool

	PrefetchMigrations uint

	LockTimeout time.Duration
}

func New(sourceURL, databaseURL string) (*Migrate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewWithDatabaseInstance(sourceURL string, databaseDriverName string, databaseInstance database.Driver) (*Migrate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewWithSourceInstance(sourceName string, sourceInstance source.Driver, databaseURL string) (*Migrate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewWithInstance(sourceName string, sourceInstance source.Driver, databaseDriverName string, databaseInstance database.Driver) (*Migrate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newCommon() *Migrate { _ = "STUB: not implemented"; return nil }

func (m *Migrate) Close() (source error, database error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Migrate) Migrate(version uint) error { _ = "STUB: not implemented"; return nil }

func (m *Migrate) Steps(n int) error { _ = "STUB: not implemented"; return nil }

func (m *Migrate) Up() error { _ = "STUB: not implemented"; return nil }

func (m *Migrate) Down() error { _ = "STUB: not implemented"; return nil }

func (m *Migrate) Drop() error { _ = "STUB: not implemented"; return nil }

func (m *Migrate) Run(migration ...*Migration) error { _ = "STUB: not implemented"; return nil }

func (m *Migrate) Force(version int) error { _ = "STUB: not implemented"; return nil }

func (m *Migrate) Version() (version uint, dirty bool, err error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

func (m *Migrate) read(from int, to int, ret chan<- interface{}) { _ = "STUB: not implemented"; return }

func (m *Migrate) readUp(from int, limit int, ret chan<- interface{}) {
	_ = "STUB: not implemented"
	return
}

func (m *Migrate) readDown(from int, limit int, ret chan<- interface{}) {
	_ = "STUB: not implemented"
	return
}

func (m *Migrate) runMigrations(ret <-chan interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Migrate) versionExists(version uint) (result error) { _ = "STUB: not implemented"; return nil }

func (m *Migrate) stop() bool { _ = "STUB: not implemented"; return false }

func (m *Migrate) newMigration(version uint, targetVersion int) (*Migration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Migrate) lock() error { _ = "STUB: not implemented"; return nil }

func (m *Migrate) unlock() error { _ = "STUB: not implemented"; return nil }

func (m *Migrate) unlockErr(prevErr error) error { _ = "STUB: not implemented"; return nil }

func (m *Migrate) logPrintf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (m *Migrate) logVerbosePrintf(format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (m *Migrate) logErr(err error) { _ = "STUB: not implemented"; return }
