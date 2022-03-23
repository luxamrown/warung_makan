package main

import (
	"fmt"

	"enigmacamp.com/warung_makan/delivery"
)

func main() {
	delivery.MainForm()
	for {
		var choiceMainForm string
		fmt.Scanln(&choiceMainForm)

		switch choiceMainForm {
		case "1":
			delivery.InForm()
		}
	}
}
