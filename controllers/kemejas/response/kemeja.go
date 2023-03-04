package response

import (
	"kemejaku/business/kemejas"
	"time"

	"gorm.io/gorm"
)

type KemejaResponse struct {
	Id        int            `json:"id"`
	Nama      string         `json:"nama"`
	Deskripsi string         `json:"deskripsi"`
	Harga     int            `json:"harga"`
	Stock_L   int            `json:"stock_L"`
	Stock_M   int            `json:"stock_M"`
	Stock_S   int            `json:"stock_S"`
	IdSale    int            `json:"idSale"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

func FromUsecase(kemeja kemejas.Kemeja) KemejaResponse {
	return KemejaResponse{
		Id:        kemeja.Id,
		Nama:      kemeja.Nama,
		Deskripsi: kemeja.Deskripsi,
		Harga:     kemeja.Harga,
		Stock_L:   kemeja.Stock_L,
		Stock_M:   kemeja.Stock_M,
		Stock_S:   kemeja.Stock_S,
		IdSale:    kemeja.IdSale,
		CreatedAt: kemeja.CreatedAt,
		UpdatedAt: kemeja.UpdatedAt,
		DeletedAt: kemeja.DeletedAt,
	}
}

func FromUsecaseList(kemeja []kemejas.Kemeja) []KemejaResponse {
	var kemejaResponse []KemejaResponse

	for _, v := range kemeja {
		kemejaResponse = append(kemejaResponse, FromUsecase(v))
	}

	return kemejaResponse
}
