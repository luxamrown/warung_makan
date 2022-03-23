package repository

type TableRepo interface {
	GetStatusTable(tableNum string) bool
}
