package source

type Direction string

const (
	Down Direction = "down"
	Up   Direction = "up"
)

type Migration struct {
	Version uint

	Identifier string

	Direction Direction

	Raw string
}

type Migrations struct {
	index      uintSlice
	migrations map[uint]map[Direction]*Migration
}

func NewMigrations() *Migrations { _ = "STUB: not implemented"; return nil }

func (i *Migrations) Append(m *Migration) (ok bool) { _ = "STUB: not implemented"; return false }

func (i *Migrations) buildIndex() { _ = "STUB: not implemented"; return }

func (i *Migrations) First() (version uint, ok bool) { _ = "STUB: not implemented"; return 0, false }

func (i *Migrations) Prev(version uint) (prevVersion uint, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (i *Migrations) Next(version uint) (nextVersion uint, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (i *Migrations) Up(version uint) (m *Migration, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (i *Migrations) Down(version uint) (m *Migration, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (i *Migrations) findPos(version uint) int { _ = "STUB: not implemented"; return 0 }

type uintSlice []uint

func (s uintSlice) Search(x uint) int { _ = "STUB: not implemented"; return 0 }
