package multistmt

import (
	"io"
)

var StartBufSize = 4096

type Handler func(migration []byte) bool

func splitWithDelimiter(delimiter []byte) func(d []byte, atEOF bool) (int, []byte, error) {
	_ = "STUB: not implemented"
	return nil
}

func Parse(reader io.Reader, delimiter []byte, maxMigrationSize int, h Handler) error {
	_ = "STUB: not implemented"
	return nil
}
