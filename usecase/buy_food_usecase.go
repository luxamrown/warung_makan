package usecase

import (
	"enigmacamp.com/warung_makan/model"
	"enigmacamp.com/warung_makan/repository"
)

type BuyFoodUseCase interface {
	Insert(namaCust, mejaCust, dataPesananCust, tagihanCust string) model.Customers
	GetHargaMakanan(kodeMakanan string) string
}

type buyFoodUseCase struct {
	repo repository.CustomersRepo
}

func (b *buyFoodUseCase) Insert(namaCust, mejaCust, dataPesananCust, tagihanCust string) model.Customers {
	newCustomer := model.NewCustomer(namaCust, mejaCust, dataPesananCust, tagihanCust)
	b.repo.Insert(newCustomer)
	return newCustomer
}

func (b *buyFoodUseCase) GetHargaMakanan(kodeMakanan string) string {
	return b.repo.GetHargaMakanan(kodeMakanan)
}
func NewBuyFoodUseCase(repo repository.CustomersRepo) BuyFoodUseCase {
	return &buyFoodUseCase{
		repo: repo,
	}
}
