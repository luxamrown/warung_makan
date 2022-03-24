package delivery

import (
	"fmt"
	"strings"

	"enigmacamp.com/warung_makan/usecase"
)

func CheckOut(usecase usecase.GetCustUseCase, usecaseTable usecase.UpdateTableUseCase) {
	var custName string
	var payConfirmation string
	fmt.Println(strings.Repeat("*", 10))
	fmt.Println("PEMBAYARAN")
	fmt.Println(strings.Repeat("*", 10))
	fmt.Print("Nama: ")
	fmt.Scan(&custName)

	fmt.Printf("ATAS NAMA %s DENGAN DATA SEBAGAI BERIKUT \n", custName)

	custChoiceFood := usecase.GetFoodCust(custName)
	custChoiceTable := usecase.GetTableNum(custName)
	custBill := usecase.GetTotalBill(custName)

	fmt.Println("MAKANAN YANG DIBELI DENGAN CODE : ", custChoiceFood)
	fmt.Println("NOMOR MEJA : ", custChoiceTable)
	fmt.Println("TAGIHAN : ", custBill)
	fmt.Print("Bayar (y/n)? ")
	fmt.Scan(&payConfirmation)

	if payConfirmation == "y" || payConfirmation == "Y" {
		usecaseTable.SetStatusTable(custChoiceTable, true)
		fmt.Println("TERIMA KASIH, SELAMAT DATANG KEMBALI")
		BackToMain()
	}

}
