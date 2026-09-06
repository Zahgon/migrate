package source

import "os"

type ErrDuplicateMigration struct {
	Migration
	os.FileInfo
}

func (e ErrDuplicateMigration) Error() string { _ = "STUB: not implemented"; return "" }
