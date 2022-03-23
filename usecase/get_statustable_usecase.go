package usecase

import "enigmacamp.com/warung_makan/repository"

type GetStatusTableUseCase interface {
	GetStatusTable(tableNum string) bool
}

type getStatusTableUseCase struct {
	repo repository.TableRepo
}

func (a *getStatusTableUseCase) GetStatusTable(tableNum string) bool {
	return a.repo.GetStatusTable(tableNum)
}

func NewGetStatusTableUseCase(repo repository.TableRepo) GetStatusTableUseCase {
	return &getStatusTableUseCase{
		repo: repo,
	}
}
