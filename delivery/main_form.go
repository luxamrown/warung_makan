package delivery

import (
	"fmt"
	"strings"
)

func MainForm() {
	fmt.Println(strings.Repeat("*", 20))
	fmt.Println("WARUNG MAKAN BAHARI")
	fmt.Println(strings.Repeat("*", 20))
	fmt.Println("1. Masuk")
	fmt.Println("2. Keluar, bayar")
	fmt.Println("3. Keluar Aplikasi")
	fmt.Print("Pilih: ")
}
