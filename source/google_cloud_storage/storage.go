package googlecloudstorage

import (
	"io"

	"cloud.google.com/go/storage"
	"github.com/golang-migrate/migrate/v4/source"
)

func init() {
	source.Register("gcs", &gcs{})
}

type gcs struct {
	bucket     *storage.BucketHandle
	prefix     string
	migrations *source.Migrations
}

func (g *gcs) Open(folder string) (source.Driver, error) {
	_ = "STUB: not implemented"
	return *new(source.Driver), nil
}

func (g *gcs) loadMigrations() error { _ = "STUB: not implemented"; return nil }

func (g *gcs) Close() error { _ = "STUB: not implemented"; return nil }

func (g *gcs) First() (uint, error) { _ = "STUB: not implemented"; return 0, nil }

func (g *gcs) Prev(version uint) (uint, error) { _ = "STUB: not implemented"; return 0, nil }

func (g *gcs) Next(version uint) (uint, error) { _ = "STUB: not implemented"; return 0, nil }

func (g *gcs) ReadUp(version uint) (io.ReadCloser, string, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), "", nil
}

func (g *gcs) ReadDown(version uint) (io.ReadCloser, string, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), "", nil
}

func (g *gcs) open(m *source.Migration) (io.ReadCloser, string, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), "", nil
}
