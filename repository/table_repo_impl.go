package repository

import (
	"fmt"

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

func (t *TableRepoImpl) SetStatusTable(tableNum string, tableStatus bool) {
	tx := t.tableDb.MustBegin()
	_, err := tx.Exec("UPDATE list_meja SET tersedia = $1 WHERE no_meja = $2", tableStatus, tableNum)
	if err != nil {
		fmt.Println(err)
	}
	tx.Commit()
}

func NewTableRepo(tableDb *sqlx.DB) TableRepo {
	tableRepo := TableRepoImpl{
		tableDb: tableDb,
	}
	return &tableRepo
}
