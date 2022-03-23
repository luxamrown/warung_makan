package delivery

import (
	"fmt"

	"enigmacamp.com/warung_makan/usecase"
)

func SearchFoodByCode(usecase usecase.SearchFoodByCodeUseCase) {
	var codeFood string
	fmt.Print("Kode Makanan: ")
	fmt.Scan(&codeFood)
	food := usecase.GetFoodByCode(codeFood)
	if food != nil {
		f := food[0]
		fmt.Println(f.KodeMakanan, f.PaketMakanan, f.HargaMakanan)
	} else {
		fmt.Println("Makanan Tidak Ditemukan..")
	}
	BackToMain()
}
