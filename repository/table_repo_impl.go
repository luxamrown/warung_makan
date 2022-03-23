package repository

import (
	"github.com/jmoiron/sqlx"
)

type TableRepoImpl struct {
	tableDb *sqlx.DB
}

func (t *TableRepoImpl) GetStatusTable(tableNum string) bool {
	var tableStatus bool
	t.tableDb.Get(&tableStatus, "SELECT tersedia FROM list_meja WHERE no_meja = $1", tableNum)
	return tableStatus
}

func NewTableRepo(tableDb *sqlx.DB) TableRepo {
	tableRepo := TableRepoImpl{
		tableDb: tableDb,
	}
	return &tableRepo
}
