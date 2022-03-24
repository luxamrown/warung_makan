package manager

import "enigmacamp.com/warung_makan/usecase"

type UseCaseManager interface {
	ListFoodUseCase() usecase.ListFoodUseCase
	SearchFoodByCodeUseCase() usecase.SearchFoodByCodeUseCase
	SearchFoodByNameUseCase() usecase.SearchFoodByNameUseCase
	GetStatusTableUseCase() usecase.GetStatusTableUseCase
	BuyFoodUseCase() usecase.BuyFoodUseCase
	UpdateTableUseCase() usecase.UpdateTableUseCase
	GetCustUseCase() usecase.GetCustUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) ListFoodUseCase() usecase.ListFoodUseCase {
	return usecase.NewListFoodUseCase(u.repo.FoodRepo())
}

func (u *useCaseManager) SearchFoodByCodeUseCase() usecase.SearchFoodByCodeUseCase {
	return usecase.NewSearchFoodByCodeUseCase(u.repo.FoodRepo())
}

func (u *useCaseManager) SearchFoodByNameUseCase() usecase.SearchFoodByNameUseCase {
	return usecase.NewSearchFoodByNameUseCase(u.repo.FoodRepo())
}

func (u *useCaseManager) GetStatusTableUseCase() usecase.GetStatusTableUseCase {
	return usecase.NewGetStatusTableUseCase(u.repo.TableRepo())
}

func (u *useCaseManager) BuyFoodUseCase() usecase.BuyFoodUseCase {
	return usecase.NewBuyFoodUseCase(u.repo.CustomersRepo())
}

func (u *useCaseManager) UpdateTableUseCase() usecase.UpdateTableUseCase {
	return usecase.NewUpdateTableUseCase(u.repo.TableRepo())
}

func (u *useCaseManager) GetCustUseCase() usecase.GetCustUseCase {
	return usecase.NewGetCustUseCase(u.repo.CustomersRepo())
}

func NewUseCaseManager(manager RepoManager) UseCaseManager {
	return &useCaseManager{
		repo: manager,
	}
}
