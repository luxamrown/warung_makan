package usecase

import (
	"enigmacamp.com/warung_makan/repository"
)

type GetCustUseCase interface {
	GetHargaMakanan(kodeMakanan string) string
	GetTotalBill(namaCust string) string
	GetTableNum(namaCust string) string
	GetFoodCust(namaCust string) string
}

type getCustUseCase struct {
	repo repository.CustomersRepo
}

func (u *getCustUseCase) GetHargaMakanan(namaCust string) string {
	return u.repo.GetHargaMakanan(namaCust)
}

func (u *getCustUseCase) GetTotalBill(namaCust string) string {
	return u.repo.GetTotalBill(namaCust)
}
func (u *getCustUseCase) GetTableNum(namaCust string) string {
	return u.repo.GetTableNum(namaCust)
}

func (u *getCustUseCase) GetFoodCust(namaCust string) string {
	return u.repo.GetFoodCust(namaCust)
}

func NewGetCustUseCase(repo repository.CustomersRepo) GetCustUseCase {
	return &getCustUseCase{
		repo: repo,
	}
}
