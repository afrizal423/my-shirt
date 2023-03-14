package request

import "github.com/afrizal423/my-shirt/business/kemejas"

type Kemeja struct {
	Nama      string `json:"nama"`
	Deskripsi string `json:"deskripsi"`
	Harga     int    `json:"harga"`
	Stock_L   int    `json:"stock_L"`
	Stock_M   int    `json:"stock_M"`
	Stock_S   int    `json:"stock_S"`
	IdSale    int    `json:"idSale"`
}

func (kemeja *Kemeja) ToUsecase() *kemejas.Kemeja {
	return &kemejas.Kemeja{
		Nama:      kemeja.Nama,
		Deskripsi: kemeja.Deskripsi,
		Harga:     kemeja.Harga,
		Stock_L:   kemeja.Stock_L,
		Stock_M:   kemeja.Stock_M,
		Stock_S:   kemeja.Stock_S,
		IdSale:    kemeja.IdSale,
	}
}
