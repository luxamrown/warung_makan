package model

type Food struct {
	KodeMakanan  string `db:"kode_makanan"`
	PaketMakanan string `db:"paket_makanan"`
	HargaMakanan int    `db:"harga_makanan"`
}

func (f *Food) GetKodeMakanan() string {
	return f.KodeMakanan
}

func (f *Food) GetPaketMakanan() string {
	return f.PaketMakanan
}

func (f *Food) GetHargaMakanan() int {
	return f.HargaMakanan
}

func NewFood(kodeMakanan string, paketMakanan string, hargaMakanan int) Food {
	return Food{
		KodeMakanan:  kodeMakanan,
		PaketMakanan: paketMakanan,
		HargaMakanan: hargaMakanan,
	}
}
