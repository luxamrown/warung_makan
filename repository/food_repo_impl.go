package repository

import (
	"enigmacamp.com/warung_makan/model"
	"github.com/jmoiron/sqlx"
)

type FoodRepoImpl struct {
	foodDb *sqlx.DB
}

func (f *FoodRepoImpl) GetFoodByName(foodName string) model.Food {
	var selectedFood model.Food
	f.foodDb.Get(&selectedFood, "SELECT * FROM list_makanan WHERE paket_makanan = $1", foodName) // Get seperti first pada silverstripe
	return selectedFood
}

func (f *FoodRepoImpl) GetFoodByCode(foodCode string) model.Food {
	var selectedFood model.Food
	f.foodDb.Get(&selectedFood, "SELECT * FROM list_makanan WHERE kode_makanan = $1", foodCode) // Get seperti first pada silverstripe
	return selectedFood
}

// func (f *FoodRepoImpl) GetPrice(foodCode string) model.Food {
// 	var selectedFood model.Food
// 	f.foodDb.Get(&selectedFood, "SELECT * FROM list_makanan WHERE kode_makanan = $1", foodCode))
// }

func (f *FoodRepoImpl) GetAll() []model.Food {
	var dataFood []model.Food
	sql := `SELECT * FROM list_makanan`
	f.foodDb.Select(&dataFood, sql)
	return dataFood
}

func NewFoodRepo(foodDb *sqlx.DB) FoodRepo {
	foodRepo := FoodRepoImpl{
		foodDb: foodDb,
	}
	return &foodRepo
}
