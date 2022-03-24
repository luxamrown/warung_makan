package delivery

import (
	"fmt"
	"strconv"
	"strings"

	"enigmacamp.com/warung_makan/manager"
)

var custName string
var custChoiceTable int
var custChoiceFood string

func BuyForm(usecase manager.UseCaseManager) {

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

	tableStatus := usecase.GetStatusTableUseCase().GetStatusTable(custChoiceTable)
	if tableStatus == false {
		fmt.Println("MEJA SUDAH PENUH, SILAHKAN PILIH MEJA LAIN")
		BackToMain()
	}
	fmt.Print("Pilih Kode Menu: ")
	fmt.Scan(&custChoiceFood)
	custbill := usecase.BuyFoodUseCase().GetHargaMakanan(custChoiceFood)
	usecase.BuyFoodUseCase().Insert(custName, custChoiceTable, custChoiceFood, custbill)
	usecase.UpdateTableUseCase().SetStatusTable(custChoiceTable, false)
	fmt.Println("TERIMA KASIH!")
	BackToMain()
}
