package repository

import "enigmacamp.com/warung_makan/model"

type CustomersRepo interface {
	Insert(newCustomer model.Customers)
	GetHargaMakanan(kodeMakanan string) string
	GetTotalBill(namaCust string) string
	GetTableNum(namaCust string) string
	GetFoodCust(namaCust string) string
}
