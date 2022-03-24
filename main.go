package main

import (
	"fmt"

	"enigmacamp.com/warung_makan/config"
	"enigmacamp.com/warung_makan/delivery"
)

func main() {
	appConfig := config.NewConfig()
	delivery.MainForm()
	for {
		var choiceMainForm string
		fmt.Scanln(&choiceMainForm)

		switch choiceMainForm {
		case "1":
			delivery.InForm()
		case "2":
			delivery.CheckOut(appConfig.UseCaseManager.GetCustUseCase(), appConfig.UseCaseManager.UpdateTableUseCase())
		case "3":
			delivery.ExitApp()
		}
	}
}
