package neo4j

import (
	"fmt"
	"io"

	"github.com/golang-migrate/migrate/v4/database"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func init() {
	db := Neo4j{}
	database.Register("neo4j", &db)
}

const DefaultMigrationsLabel = "SchemaMigration"

var (
	StatementSeparator           = []byte(";")
	DefaultMultiStatementMaxSize = 10 * 1 << 20
)

var (
	ErrNilConfig = fmt.Errorf("no config")
)

type Config struct {
	MigrationsLabel       string
	MultiStatement        bool
	MultiStatementMaxSize int
}

type Neo4j struct {
	driver neo4j.Driver
	lock   uint32

	config *Config
}

func WithInstance(driver neo4j.Driver, config *Config) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func (n *Neo4j) Open(url string) (database.Driver, error) {
	_ = "STUB: not implemented"
	return *new(database.Driver), nil
}

func (n *Neo4j) Close() error { _ = "STUB: not implemented"; return nil }

func (n *Neo4j) Lock() error { _ = "STUB: not implemented"; return nil }

func (n *Neo4j) Unlock() error { _ = "STUB: not implemented"; return nil }

func (n *Neo4j) Run(migration io.Reader) (err error) { _ = "STUB: not implemented"; return nil }

func (n *Neo4j) SetVersion(version int, dirty bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

type MigrationRecord struct {
	Version int
	Dirty   bool
}

func (n *Neo4j) Version() (version int, dirty bool, err error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

func (n *Neo4j) Drop() (err error) { _ = "STUB: not implemented"; return nil }

func (n *Neo4j) ensureVersionConstraint() (err error) { _ = "STUB: not implemented"; return nil }
