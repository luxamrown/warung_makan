package model

type Table struct {
	NoMeja    int  `db:"no_meja"`
	Available bool `db:"tersedia"`
}

func (t *Table) GetNumber() int {
	return t.NoMeja
}

func (t *Table) SetStatus(tersedia bool) {
	t.Available = tersedia
}

func NewTable(noMeja int, available bool) Table {
	return Table{
		NoMeja:    noMeja,
		Available: available,
	}
}
