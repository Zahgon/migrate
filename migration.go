package migrate

import (
	"io"
	"time"
)

var DefaultBufferSize = uint(100000)

type Migration struct {
	Identifier string

	Version uint

	TargetVersion int

	Body io.ReadCloser

	BufferedBody io.Reader

	BufferSize uint

	bufferWriter io.WriteCloser

	Scheduled time.Time

	StartedBuffering time.Time

	FinishedBuffering time.Time

	FinishedReading time.Time

	BytesRead int64
}

func NewMigration(body io.ReadCloser, identifier string,
	version uint, targetVersion int) (*Migration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Migration) String() string { _ = "STUB: not implemented"; return "" }

func (m *Migration) LogString() string { _ = "STUB: not implemented"; return "" }

func (m *Migration) Buffer() (berr error) { _ = "STUB: not implemented"; return nil }
