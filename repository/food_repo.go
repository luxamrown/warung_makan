package repository

import "enigmacamp.com/warung_makan/model"

type FoodRepo interface {
	GetFoodByName(foodName string) model.Food
	GetFoodByCode(foodCode string) model.Food
	GetAll() []model.Food
}
