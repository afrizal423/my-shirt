package kemejas

import (
	"time"

	"github.com/afrizal423/my-shirt/business/kemejas"

	"gorm.io/gorm"
)

type Kemeja struct {
	Id        int `gorm:"primaryKey;unique;autoIncrement:true"`
	Nama      string
	Deskripsi string
	Harga     int
	Stock_L   int
	Stock_M   int
	Stock_S   int
	IdSale    int
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (kemeja Kemeja) ToUsecase() kemejas.Kemeja {
	return kemejas.Kemeja{
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

func ToUsecaseList(kemeja []Kemeja) []kemejas.Kemeja {
	var newkemejas []kemejas.Kemeja

	for _, v := range kemeja {
		newkemejas = append(newkemejas, v.ToUsecase())
	}

	return newkemejas
}

func FromUsecase(kemeja kemejas.Kemeja) Kemeja {
	return Kemeja{
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

func FromUseCaseList(kemejas []kemejas.Kemeja) []Kemeja {
	var newkemejas []Kemeja

	for _, v := range kemejas {
		newkemejas = append(newkemejas, FromUsecase(v))
	}

	return newkemejas
}
