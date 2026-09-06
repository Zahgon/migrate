package database

type Error struct {
	Line uint

	Query []byte

	Err string

	OrigErr error
}

func (e Error) Error() string { _ = "STUB: not implemented"; return "" }
