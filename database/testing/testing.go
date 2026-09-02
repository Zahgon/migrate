package testing

import (
	"io"
	"testing"

	"github.com/golang-migrate/migrate/v4/database"
)

func Test(t *testing.T, d database.Driver, migration []byte) { _ = "STUB: not implemented"; return }

func TestNilVersion(t *testing.T, d database.Driver) { _ = "STUB: not implemented"; return }

func TestLockAndUnlock(t *testing.T, d database.Driver) { _ = "STUB: not implemented"; return }

func TestRun(t *testing.T, d database.Driver, migration io.Reader) {
	_ = "STUB: not implemented"
	return
}

func TestDrop(t *testing.T, d database.Driver) { _ = "STUB: not implemented"; return }

func TestSetVersion(t *testing.T, d database.Driver) { _ = "STUB: not implemented"; return }
