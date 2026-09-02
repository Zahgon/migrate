package migrate

import (
	nurl "net/url"
)

type MultiError struct {
	Errs []error
}

func NewMultiError(errs ...error) MultiError { _ = "STUB: not implemented"; return *new(MultiError) }

func (m MultiError) Error() string { _ = "STUB: not implemented"; return "" }

func suint(n int) uint { _ = "STUB: not implemented"; return 0 }

func FilterCustomQuery(u *nurl.URL) *nurl.URL { _ = "STUB: not implemented"; return nil }
