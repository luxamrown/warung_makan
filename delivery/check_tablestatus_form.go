package delivery

import (
	"fmt"

	"enigmacamp.com/warung_makan/usecase"
)

func CheckTableStatus(usecase usecase.GetStatusTableUseCase) {
	var tableNum string
	fmt.Print("Masukan No Meja: ")
	fmt.Scan(&tableNum)
	tableStatus := usecase.GetStatusTable(tableNum)
	if tableStatus == true {
		fmt.Print("MEJA KOSONG, DAPAT DIISI")
	} else {
		fmt.Println("MEJA TERISI, SILAHKAN CARI MEJA LAIN")
	}
	BackToMain()
}
