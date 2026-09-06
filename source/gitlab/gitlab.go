package gitlab

import (
	"fmt"
	"io"

	"github.com/golang-migrate/migrate/v4/source"
	"github.com/xanzy/go-gitlab"
)

func init() {
	source.Register("gitlab", &Gitlab{})
}

const DefaultMaxItemsPerPage = 100

var (
	ErrNoUserInfo       = fmt.Errorf("no username:token provided")
	ErrNoAccessToken    = fmt.Errorf("no access token")
	ErrInvalidHost      = fmt.Errorf("invalid host")
	ErrInvalidProjectID = fmt.Errorf("invalid project id")
	ErrInvalidResponse  = fmt.Errorf("invalid response")
)

type Gitlab struct {
	client *gitlab.Client
	url    string

	projectID   string
	path        string
	listOptions *gitlab.ListTreeOptions
	getOptions  *gitlab.GetFileOptions
	migrations  *source.Migrations
}

type Config struct {
}

func (g *Gitlab) Open(url string) (source.Driver, error) {
	_ = "STUB: not implemented"
	return *new(source.Driver), nil
}

func WithInstance(client *gitlab.Client, config *Config) (source.Driver, error) {
	_ = "STUB: not implemented"
	return *new(source.Driver), nil
}

func (g *Gitlab) readDirectory() error { _ = "STUB: not implemented"; return nil }

func (g *Gitlab) nodeToMigration(node *gitlab.TreeNode) (*source.Migration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Gitlab) Close() error { _ = "STUB: not implemented"; return nil }

func (g *Gitlab) First() (version uint, er error) { _ = "STUB: not implemented"; return 0, nil }

func (g *Gitlab) Prev(version uint) (prevVersion uint, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (g *Gitlab) Next(version uint) (nextVersion uint, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (g *Gitlab) ReadUp(version uint) (r io.ReadCloser, identifier string, err error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), "", nil
}

func (g *Gitlab) ReadDown(version uint) (r io.ReadCloser, identifier string, err error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), "", nil
}
