package bindata

import (
	"fmt"
	"io"

	"github.com/golang-migrate/migrate/v4/source"
)

type AssetFunc func(name string) ([]byte, error)

func Resource(names []string, afn AssetFunc) *AssetSource { _ = "STUB: not implemented"; return nil }

type AssetSource struct {
	Names     []string
	AssetFunc AssetFunc
}

func init() {
	source.Register("go-bindata", &Bindata{})
}

type Bindata struct {
	path        string
	assetSource *AssetSource
	migrations  *source.Migrations
}

func (b *Bindata) Open(url string) (source.Driver, error) {
	_ = "STUB: not implemented"
	return *new(source.Driver), nil
}

var (
	ErrNoAssetSource = fmt.Errorf("expects *AssetSource")
)

func WithInstance(instance interface{}) (source.Driver, error) {
	_ = "STUB: not implemented"
	return *new(source.Driver), nil
}

func (b *Bindata) Close() error { _ = "STUB: not implemented"; return nil }

func (b *Bindata) First() (version uint, err error) { _ = "STUB: not implemented"; return 0, nil }

func (b *Bindata) Prev(version uint) (prevVersion uint, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *Bindata) Next(version uint) (nextVersion uint, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *Bindata) ReadUp(version uint) (r io.ReadCloser, identifier string, err error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), "", nil
}

func (b *Bindata) ReadDown(version uint) (r io.ReadCloser, identifier string, err error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), "", nil
}
