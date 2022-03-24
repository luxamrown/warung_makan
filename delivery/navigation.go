package delivery

import (
	"fmt"
	"os"
)

func BackToMain() {
	var mainMenuConfirmation string
	fmt.Print("KEMBALI KE MENU (y/n)?")
	fmt.Scan(&mainMenuConfirmation)
	if mainMenuConfirmation == "y" || mainMenuConfirmation == "Y" {
		MainForm()
	}
}

func ExitApp() {
	fmt.Println("TERIMA KASIH SELAMAT DATANG KEMBALI")
	os.Exit(0)
}
