package repository

type TableRepo interface {
	GetStatusTable(tableNum string) bool
	SetStatusTable(tableNum string, tableStatus bool)
}
