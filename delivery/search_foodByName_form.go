package delivery

import (
	"fmt"

	"enigmacamp.com/warung_makan/usecase"
)

func SearchFoodByName(usecase usecase.SearchFoodByNameUseCase) {
	var foodName string
	fmt.Print("Nama Makanan: ")
	fmt.Scan(&foodName)
	food := usecase.GetFoodByName(foodName)
	if food != nil {
		f := food[0]
		fmt.Println(f.KodeMakanan, f.PaketMakanan, f.HargaMakanan)
	} else {
		fmt.Println("Makanan Tidak Ditemukan..")
	}
	BackToMain()
}
