package awss3

import (
	"io"

	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/golang-migrate/migrate/v4/source"
)

func init() {
	source.Register("s3", &s3Driver{})
}

type s3Driver struct {
	s3client   s3iface.S3API
	config     *Config
	migrations *source.Migrations
}

type Config struct {
	Bucket string
	Prefix string
}

func (s *s3Driver) Open(folder string) (source.Driver, error) {
	_ = "STUB: not implemented"
	return *new(source.Driver), nil
}

func WithInstance(s3client s3iface.S3API, config *Config) (source.Driver, error) {
	_ = "STUB: not implemented"
	return *new(source.Driver), nil
}

func parseURI(uri string) (*Config, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *s3Driver) loadMigrations() error { _ = "STUB: not implemented"; return nil }

func (s *s3Driver) Close() error { _ = "STUB: not implemented"; return nil }

func (s *s3Driver) First() (uint, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *s3Driver) Prev(version uint) (uint, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *s3Driver) Next(version uint) (uint, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *s3Driver) ReadUp(version uint) (io.ReadCloser, string, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), "", nil
}

func (s *s3Driver) ReadDown(version uint) (io.ReadCloser, string, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), "", nil
}

func (s *s3Driver) open(m *source.Migration) (io.ReadCloser, string, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), "", nil
}
