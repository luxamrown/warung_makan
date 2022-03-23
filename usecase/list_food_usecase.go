package usecase

import (
	"enigmacamp.com/warung_makan/model"
	"enigmacamp.com/warung_makan/repository"
)

type ListFoodUseCase interface {
	GetAll() []model.Food
}

type listFoodUseCase struct {
	repo repository.FoodRepo
}

func (a *listFoodUseCase) GetAll() []model.Food {
	return a.repo.GetAll()
}

func NewListFoodUseCase(repo repository.FoodRepo) ListFoodUseCase {
	return &listFoodUseCase{
		repo: repo,
	}
}
