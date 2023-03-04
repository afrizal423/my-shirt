package response

import (
	"kemejaku/business/kemejakeranjangs"
	"kemejaku/business/keranjangs"
	"time"

	"gorm.io/gorm"
)

type KeranjangResponse struct {
	Id              int                                `json:"id"`
	IdUser          int                                `json:"idUser"`
	Status          bool                               `json:"status"`
	KemejaKeranjang []kemejakeranjangs.KemejaKeranjang `json:"kemejaKeranjang"`
	CreatedAt       time.Time                          `json:"createdAt"`
	UpdatedAt       time.Time                          `json:"updatedAt"`
	DeletedAt       gorm.DeletedAt                     `json:"deletedAt"`
}

func FromUsecase(keranjang keranjangs.Keranjang) KeranjangResponse {
	return KeranjangResponse{
		Id:              keranjang.Id,
		IdUser:          keranjang.IdUser,
		Status:          keranjang.Status,
		KemejaKeranjang: keranjang.KemejaKeranjang,
		CreatedAt:       keranjang.CreatedAt,
		UpdatedAt:       keranjang.UpdatedAt,
		DeletedAt:       keranjang.DeletedAt,
	}
}

func FromUsecaseList(keranjang []keranjangs.Keranjang) []KeranjangResponse {
	var keranjangResponse []KeranjangResponse

	for _, v := range keranjang {
		keranjangResponse = append(keranjangResponse, FromUsecase(v))
	}

	return keranjangResponse
}
