package delivery

import (
	"fmt"
	"strconv"
	"strings"

	"enigmacamp.com/warung_makan/usecase"
)

var custName string
var custChoiceTable int
var custChoiceFood string

func BuyForm(usecase usecase.BuyFoodUseCase) {

	fmt.Println(strings.Repeat("*", 10))
	fmt.Println("FORM PEMBELIAN")
	fmt.Println(strings.Repeat("*", 10))

	fmt.Print("Nama: ")
	fmt.Scan(&custName)

	fmt.Print("Pilih Nomor Meja: ")
	fmt.Scan(&custChoiceTable)
	if custChoiceTable > 30 {
		fmt.Println("MEJA HANYA ADA 30")
		BackToMain()
		return
	}
	custChoiceTable := strconv.Itoa(custChoiceTable)
	fmt.Print("Pilih Kode Menu: ")
	fmt.Scan(&custChoiceFood)
	custbill := usecase.GetHargaMakanan(custChoiceFood)
	usecase.Insert(custName, custChoiceTable, custChoiceFood, custbill)
}
