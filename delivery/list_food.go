package delivery

import (
	"fmt"
	"strings"

	"enigmacamp.com/warung_makan/usecase"
)

func ListFood(usecase usecase.ListFoodUseCase) {
	fmt.Println(strings.Repeat("*", 10))
	fmt.Println("DAFTAR MENU")
	fmt.Println(strings.Repeat("*", 10))

	for _, food := range usecase.GetAll() {
		fmt.Println(food.KodeMakanan, food.PaketMakanan, food.HargaMakanan)
	}
	BackToMain()
}
