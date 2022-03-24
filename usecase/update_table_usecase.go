package usecase

import (
	"enigmacamp.com/warung_makan/repository"
)

type UpdateTableUseCase interface {
	SetStatusTable(tableNum string, tableStatus bool)
}

type updateTableUseCase struct {
	repo repository.TableRepo
}

func (u *updateTableUseCase) SetStatusTable(tableNum string, tableStatus bool) {
	u.repo.SetStatusTable(tableNum, tableStatus)
}

func NewUpdateTableUseCase(repo repository.TableRepo) UpdateTableUseCase {
	return &updateTableUseCase{
		repo: repo,
	}
}
