package kemejakeranjangs

import (
	"time"

	"github.com/afrizal423/my-shirt/business/kemejakeranjangs"

	"gorm.io/gorm"
)

type KemejaKeranjang struct {
	Id          int `gorm:"primaryKey;unique;autoIncrement:true"`
	IdKemeja    int
	IdKeranjang int
	Jumlah      int
	Size        string
	// Keranjang   keranjangs.Keranjang `gorm:"foreignKey:IdKeranjang"`
	// Kemeja      kemejas.Kemeja       `gorm:"foreignKey:IdKemeja"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (kk KemejaKeranjang) ToUsecase() kemejakeranjangs.KemejaKeranjang {
	return kemejakeranjangs.KemejaKeranjang{
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

func ToUsecaseList(kk []KemejaKeranjang) []kemejakeranjangs.KemejaKeranjang {
	var newKKs []kemejakeranjangs.KemejaKeranjang

	for _, v := range kk {
		newKKs = append(newKKs, v.ToUsecase())
	}

	return newKKs
}

func FromUsecase(kk kemejakeranjangs.KemejaKeranjang) KemejaKeranjang {
	return KemejaKeranjang{
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

func FromUseCaseList(kemejakeranjangs []kemejakeranjangs.KemejaKeranjang) []KemejaKeranjang {
	var newKK []KemejaKeranjang

	for _, v := range kemejakeranjangs {
		newKK = append(newKK, FromUsecase(v))
	}

	return newKK
}
