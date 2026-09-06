package rqlite

import (
	"fmt"
	"io"
	nurl "net/url"
	"sync/atomic"

	"github.com/golang-migrate/migrate/v4/database"
	"github.com/rqlite/gorqlite"
)

func init() {
	database.Register("rqlite", &Rqlite{})
}

const (
	DefaultMigrationsTable = "schema_migrations"

	DefaultConnectInsecure = false
)

var ErrNilConfig = fmt.Errorf("no config")

var ErrBadConfig = fmt.Errorf("bad parameter")

type Config struct {
	ConnectInsecure bool

	MigrationsTable string
}

type Rqlite struct {
	db       *gorqlite.Connection
	isLocked atomic.Bool

	config *Config
}

func WithInstance(instance *gorqlite.Connection, config *Config) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func OpenURL(url string) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func (r *Rqlite) ensureVersionTable() (err error) { _ = "STUB: not implemented"; return nil }

func (r *Rqlite) Open(url string) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func (r *Rqlite) Close() error { _ = "STUB: not implemented"; return nil }

func (r *Rqlite) Lock() error { _ = "STUB: not implemented"; return nil }

func (r *Rqlite) Unlock() error { _ = "STUB: not implemented"; return nil }

func (r *Rqlite) Run(migration io.Reader) error { _ = "STUB: not implemented"; return nil }

func (r *Rqlite) SetVersion(version int, dirty bool) error { _ = "STUB: not implemented"; return nil }

func (r *Rqlite) Version() (version int, dirty bool, err error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

func (r *Rqlite) Drop() error { _ = "STUB: not implemented"; return nil }

func parseUrl(url string) (*nurl.URL, *Config, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func parseConfigFromQuery(queryVals nurl.Values) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
