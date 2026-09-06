package bitbucket

import (
	"fmt"
	"io"

	"github.com/golang-migrate/migrate/v4/source"
	"github.com/ktrysmt/go-bitbucket"
)

func init() {
	source.Register("bitbucket", &Bitbucket{})
}

var (
	ErrNoUserInfo             = fmt.Errorf("no username:password provided")
	ErrNoAccessToken          = fmt.Errorf("no password/app password")
	ErrInvalidRepo            = fmt.Errorf("invalid repo")
	ErrInvalidBitbucketClient = fmt.Errorf("expected *bitbucket.Client")
	ErrNoDir                  = fmt.Errorf("no directory")
)

type Bitbucket struct {
	config     *Config
	client     *bitbucket.Client
	migrations *source.Migrations
}

type Config struct {
	Owner string
	Repo  string
	Path  string
	Ref   string
}

func (b *Bitbucket) Open(url string) (source.Driver, error) {
	_ = "STUB: not implemented"
	return *new(source.Driver), nil
}

func WithInstance(client *bitbucket.Client, config *Config) (source.Driver, error) {
	_ = "STUB: not implemented"
	return *new(source.Driver), nil
}

func (b *Bitbucket) readDirectory() error { _ = "STUB: not implemented"; return nil }

func (b *Bitbucket) ensureFields() { _ = "STUB: not implemented"; return }

func (b *Bitbucket) Close() error { _ = "STUB: not implemented"; return nil }

func (b *Bitbucket) First() (version uint, er error) { _ = "STUB: not implemented"; return 0, nil }

func (b *Bitbucket) Prev(version uint) (prevVersion uint, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *Bitbucket) Next(version uint) (nextVersion uint, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *Bitbucket) ReadUp(version uint) (r io.ReadCloser, identifier string, err error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), "", nil
}

func (b *Bitbucket) ReadDown(version uint) (r io.ReadCloser, identifier string, err error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), "", nil
}
