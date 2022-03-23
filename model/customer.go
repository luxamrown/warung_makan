package model

type Customers struct {
	NamaCustomer        string
	MejaCustomer        string
	DataPesananCustomer string
	TagihanCustomer     string
}

func (c *Customers) SetNamaCustomer(namaCustomer string) {
	c.NamaCustomer = namaCustomer
}

func (c *Customers) SetMejaCustomer(tableNum string) {
	c.MejaCustomer = tableNum
}

func (c *Customers) SetDataPesananCustomer(makanan string) {
	c.DataPesananCustomer = makanan
}

func (c *Customers) SetTagihanCustomer(tagihanCustromer string) {
	c.TagihanCustomer = tagihanCustromer
}

func (c *Customers) GetNamaCustomer() string {
	return c.NamaCustomer
}

func (c *Customers) GetMejaCustomer() string {
	return c.MejaCustomer
}

func (c *Customers) GetDataPesananCustomer() string {
	return c.DataPesananCustomer
}

func (c *Customers) GetTagihanCustomer() string {
	return c.TagihanCustomer
}

func NewCustomer(namaCust, mejaCust, dataCust, tagihanCust string) Customers {
	return Customers{
		NamaCustomer:        namaCust,
		MejaCustomer:        mejaCust,
		DataPesananCustomer: dataCust,
		TagihanCustomer:     tagihanCust,
	}
}
