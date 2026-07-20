package github

import (
	"fmt"
	"io"

	"github.com/golang-migrate/migrate/v4/source"
	"github.com/google/go-github/v39/github"
)

func init() {
	source.Register("github", &Github{})
}

var (
	ErrNoUserInfo          = fmt.Errorf("no username:token provided")
	ErrNoAccessToken       = fmt.Errorf("no access token")
	ErrInvalidRepo         = fmt.Errorf("invalid repo")
	ErrInvalidGithubClient = fmt.Errorf("expected *github.Client")
	ErrNoDir               = fmt.Errorf("no directory")
)

type Github struct {
	config     *Config
	client     *github.Client
	options    *github.RepositoryContentGetOptions
	migrations *source.Migrations
}

type Config struct {
	Owner string
	Repo  string
	Path  string
	Ref   string
}

func (g *Github) Open(url string) (source.Driver, error) {
	_ = "STUB: not implemented"
	return *new(source.Driver), nil
}

func WithInstance(client *github.Client, config *Config) (source.Driver, error) {
	_ = "STUB: not implemented"
	return *new(source.Driver), nil
}

func (g *Github) readDirectory() error { _ = "STUB: not implemented"; return nil }

func (g *Github) ensureFields() { _ = "STUB: not implemented"; return }

func (g *Github) Close() error { _ = "STUB: not implemented"; return nil }

func (g *Github) First() (version uint, err error) { _ = "STUB: not implemented"; return 0, nil }

func (g *Github) Prev(version uint) (prevVersion uint, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (g *Github) Next(version uint) (nextVersion uint, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (g *Github) ReadUp(version uint) (r io.ReadCloser, identifier string, err error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), "", nil
}

func (g *Github) ReadDown(version uint) (r io.ReadCloser, identifier string, err error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), "", nil
}
