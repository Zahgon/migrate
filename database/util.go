package database

import (
	"sync/atomic"
)

const advisoryLockIDSalt uint = 1486364155

func GenerateAdvisoryLockId(databaseName string, additionalNames ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func CasRestoreOnErr(lock *atomic.Bool, o, n bool, casErr error, f func() error) error {
	_ = "STUB: not implemented"
	return nil
}
