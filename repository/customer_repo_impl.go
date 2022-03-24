package repository

import (
	"fmt"

	"enigmacamp.com/warung_makan/model"
	"github.com/jmoiron/sqlx"
)

type CustomersRepoImpl struct {
	custDb *sqlx.DB
}

func (c *CustomersRepoImpl) Insert(newCustomer model.Customers) {
	tx := c.custDb.MustBegin()
	_, err := tx.Exec("INSERT INTO list_pembeli(nama_pembeli, nomor_meja,data_pesanan_pembeli, tagihan_pembeli) values($1, $2, $3, $4)", newCustomer.NamaCustomer, newCustomer.MejaCustomer, newCustomer.DataPesananCustomer, newCustomer.TagihanCustomer)
	if err != nil {
		fmt.Println(err)
	}
	tx.Commit()
}

func (c *CustomersRepoImpl) GetHargaMakanan(kodeMakanan string) string {
	var selectedPrice string
	c.custDb.Get(&selectedPrice, "SELECT harga_makanan FROM list_makanan WHERE kode_makanan = $1", kodeMakanan) // Get seperti first pada silverstripe
	return selectedPrice
}

func (c *CustomersRepoImpl) GetTotalBill(namaCust string) string {
	var selectedPrice string
	c.custDb.Get(&selectedPrice, "SELECT tagihan_pembeli FROM list_pembeli WHERE nama_pembeli = $1", namaCust) // Get seperti first pada silverstripe
	return selectedPrice
}

func (c *CustomersRepoImpl) GetTableNum(namaCust string) string {
	var selectedTable string
	c.custDb.Get(&selectedTable, "SELECT nomor_meja FROM list_pembeli WHERE nama_pembeli = $1", namaCust) // Get seperti first pada silverstripe
	return selectedTable
}

func (c *CustomersRepoImpl) GetFoodCust(namaCust string) string {
	var selectedFood string
	c.custDb.Get(&selectedFood, "SELECT data_pesanan_pembeli FROM list_pembeli WHERE nama_pembeli = $1", namaCust) // Get seperti first pada silverstripe
	return selectedFood
}

func NewCustomersRepo(custDb *sqlx.DB) CustomersRepo {
	customersRepo := CustomersRepoImpl{
		custDb: custDb,
	}
	return &customersRepo
}
