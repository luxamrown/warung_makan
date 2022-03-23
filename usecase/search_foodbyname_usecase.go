package usecase

import (
	"enigmacamp.com/warung_makan/model"
	"enigmacamp.com/warung_makan/repository"
)

type SearchFoodByNameUseCase interface {
	GetFoodByName(foodName string) []model.Food
}

type searchFoodByNameUsecase struct {
	repo repository.FoodRepo
}

func (a *searchFoodByNameUsecase) GetFoodByName(foodName string) []model.Food {
	if len(foodName) == 0 {
		return a.repo.GetAll()
	}
	result := a.repo.GetFoodByName(foodName)
	if len(result.GetPaketMakanan()) == 0 {
		return nil
	} else {
		return []model.Food{result}
	}
}

func NewSearchFoodByNameUseCase(repo repository.FoodRepo) SearchFoodByNameUseCase {
	return &searchFoodByNameUsecase{
		repo: repo,
	}
}
