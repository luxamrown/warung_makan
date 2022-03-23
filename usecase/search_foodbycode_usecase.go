package usecase

import (
	"enigmacamp.com/warung_makan/model"
	"enigmacamp.com/warung_makan/repository"
)

type SearchFoodByCodeUseCase interface {
	GetFoodByCode(foodCode string) []model.Food
}

type searchFoodByCodeUsecase struct {
	repo repository.FoodRepo
}

func (a *searchFoodByCodeUsecase) GetFoodByCode(codeFood string) []model.Food {
	if len(codeFood) == 0 {
		return a.repo.GetAll()
	}
	result := a.repo.GetFoodByCode(codeFood)
	if len(result.GetKodeMakanan()) == 0 {
		return nil
	} else {
		return []model.Food{result}
	}
}

func NewSearchFoodByCodeUseCase(repo repository.FoodRepo) SearchFoodByCodeUseCase {
	return &searchFoodByCodeUsecase{
		repo: repo,
	}
}
