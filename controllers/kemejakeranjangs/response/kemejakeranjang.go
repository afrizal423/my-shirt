package response

import (
	"time"

	"github.com/afrizal423/my-shirt/business/kemejakeranjangs"

	"gorm.io/gorm"
)

type KemejaKeranjangResponse struct {
	Id          int            `json:"id"`
	IdKemeja    int            `json:"idKemeja"`
	IdKeranjang int            `json:"idKeranjang"`
	Jumlah      int            `json:"jumlah"`
	Size        string         `json:"size"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt"`
}

func FromUsecase(kk kemejakeranjangs.KemejaKeranjang) KemejaKeranjangResponse {
	return KemejaKeranjangResponse{
		Id:          kk.Id,
		IdKemeja:    kk.IdKemeja,
		IdKeranjang: kk.IdKeranjang,
		Jumlah:      kk.Jumlah,
		Size:        kk.Size,
		CreatedAt:   kk.CreatedAt,
		UpdatedAt:   kk.UpdatedAt,
		DeletedAt:   kk.DeletedAt,
	}
}

func FromUsecaseList(kk []kemejakeranjangs.KemejaKeranjang) []KemejaKeranjangResponse {
	var kkResponse []KemejaKeranjangResponse

	for _, v := range kk {
		kkResponse = append(kkResponse, FromUsecase(v))
	}

	return kkResponse
}
