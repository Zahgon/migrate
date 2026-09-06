package sqlserver

import (
	"database/sql"
	"fmt"
	"io"
	nurl "net/url"
	"sync/atomic"

	"github.com/golang-migrate/migrate/v4/database"
)

func init() {
	database.Register("sqlserver", &SQLServer{})
}

var DefaultMigrationsTable = "schema_migrations"

var (
	ErrNilConfig                 = fmt.Errorf("no config")
	ErrNoDatabaseName            = fmt.Errorf("no database name")
	ErrNoSchema                  = fmt.Errorf("no schema")
	ErrDatabaseDirty             = fmt.Errorf("database is dirty")
	ErrMultipleAuthOptionsPassed = fmt.Errorf("both password and useMsi=true were passed")
)

var lockErrorMap = map[int]string{
	-1:   "The lock request timed out.",
	-2:   "The lock request was canceled.",
	-3:   "The lock request was chosen as a deadlock victim.",
	-999: "Parameter validation or other call error.",
}

type Config struct {
	MigrationsTable string
	DatabaseName    string
	SchemaName      string
}

type SQLServer struct {
	conn     *sql.Conn
	db       *sql.DB
	isLocked atomic.Bool

	config *Config
}

func WithInstance(instance *sql.DB, config *Config) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func (ss *SQLServer) Open(url string) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func (ss *SQLServer) Close() error { _ = "STUB: not implemented"; return nil }

func (ss *SQLServer) Lock() error { _ = "STUB: not implemented"; return nil }

func (ss *SQLServer) Unlock() error { _ = "STUB: not implemented"; return nil }

func (ss *SQLServer) Run(migration io.Reader) error { _ = "STUB: not implemented"; return nil }

func (ss *SQLServer) SetVersion(version int, dirty bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (ss *SQLServer) Version() (version int, dirty bool, err error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

func (ss *SQLServer) Drop() error { _ = "STUB: not implemented"; return nil }

func (ss *SQLServer) ensureVersionTable() (err error) { _ = "STUB: not implemented"; return nil }

func (ss *SQLServer) getMigrationTable() string { _ = "STUB: not implemented"; return "" }

func getMSITokenProvider(resource string) (func() (string, error), error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getAADResourceFromServerUri(purl *nurl.URL) string { _ = "STUB: not implemented"; return "" }
