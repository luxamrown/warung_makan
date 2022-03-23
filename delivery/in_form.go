package delivery

import (
	"fmt"
	"strings"

	"enigmacamp.com/warung_makan/config"
)

func InForm() {
	appConfig := config.NewConfig()
	var choiceInForm string
	fmt.Println(strings.Repeat("*", 37))
	fmt.Println("SELAMAT DATANG DI WARUNG MAKAN BAHARI")
	fmt.Println(strings.Repeat("*", 37))
	fmt.Println("1. LIHAT DAFTAR MENU")
	fmt.Println("2. CARI MENU BY KODE")
	fmt.Println("3. CARI MENU BY NAMA")
	fmt.Println("4. CEK STATUS MEJA")
	fmt.Println("5. PESAN MAKANAN")

	fmt.Scan(&choiceInForm)
	switch choiceInForm {
	case "1":
		ListFood(appConfig.UseCaseManager.ListFoodUseCase())
	case "2":
		SearchFoodByCode(appConfig.UseCaseManager.SearchFoodByCodeUseCase())
	case "3":
		SearchFoodByName(appConfig.UseCaseManager.SearchFoodByNameUseCase())
	case "4":
		CheckTableStatus(appConfig.UseCaseManager.GetStatusTableUseCase())
	case "5":
		BuyForm(appConfig.UseCaseManager.BuyFoodUseCase())
	}

}
