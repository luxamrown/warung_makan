package repository

import "enigmacamp.com/warung_makan/model"

type CustomersRepo interface {
	Insert(newCustomer model.Customers)
	GetHargaMakanan(kodeMakanan string) string
}
