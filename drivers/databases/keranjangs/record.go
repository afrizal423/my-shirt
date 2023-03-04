package keranjangs

import (
	"kemejaku/business/keranjangs"
	"kemejaku/drivers/databases/kemejakeranjangs"
	"time"

	"gorm.io/gorm"
)

type Keranjang struct {
	Id              int `gorm:"primaryKey;unique;autoIncrement:true"`
	IdUser          int
	Status          bool
	KemejaKeranjang []kemejakeranjangs.KemejaKeranjang `gorm:"foreignKey:IdKeranjang"`
	CreatedAt       time.Time                          `gorm:"autoCreateTime"`
	UpdatedAt       time.Time                          `gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt                     `gorm:"index"`
}

func (keranjang Keranjang) ToUsecase() keranjangs.Keranjang {
	newKK := kemejakeranjangs.ToUsecaseList(keranjang.KemejaKeranjang)

	return keranjangs.Keranjang{
		Id:              keranjang.Id,
		IdUser:          keranjang.IdUser,
		Status:          keranjang.Status,
		KemejaKeranjang: newKK,
		CreatedAt:       keranjang.CreatedAt,
		UpdatedAt:       keranjang.UpdatedAt,
		DeletedAt:       keranjang.DeletedAt,
	}
}

func ToUsecaseList(keranjang []Keranjang) []keranjangs.Keranjang {
	var newKeranjangs []keranjangs.Keranjang

	for _, v := range keranjang {
		newKeranjangs = append(newKeranjangs, v.ToUsecase())
	}

	return newKeranjangs
}

func FromUsecase(keranjang keranjangs.Keranjang) Keranjang {
	newKK := kemejakeranjangs.FromUseCaseList(keranjang.KemejaKeranjang)

	return Keranjang{
		Id:              keranjang.Id,
		IdUser:          keranjang.IdUser,
		Status:          keranjang.Status,
		KemejaKeranjang: newKK,
		CreatedAt:       keranjang.CreatedAt,
		UpdatedAt:       keranjang.UpdatedAt,
		DeletedAt:       keranjang.DeletedAt,
	}
}
