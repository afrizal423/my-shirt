package keranjangs

import (
	"context"
	"time"

	"github.com/afrizal423/my-shirt/business/kemejakeranjangs"

	"gorm.io/gorm"
)

type Keranjang struct {
	Id              int
	IdUser          int
	Status          bool
	KemejaKeranjang []kemejakeranjangs.KemejaKeranjang
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

type KeranjangUseCaseInterface interface {
	InsertKeranjang(keranjang Keranjang, ctx context.Context) (Keranjang, error)
	GetAllKeranjang(ctx context.Context) ([]Keranjang, error)
	GetKeranjangDetail(id int, ctx context.Context) (Keranjang, error)
	EditKeranjang(keranjang Keranjang, id int, ctx context.Context) (Keranjang, error)
	DeleteKeranjang(id int, ctx context.Context) (Keranjang, error)
}

type KeranjangRepoInterface interface {
	InsertKeranjang(keranjang Keranjang, ctx context.Context) (Keranjang, error)
	GetAllKeranjang(ctx context.Context) ([]Keranjang, error)
	GetKeranjangDetail(id int, ctx context.Context) (Keranjang, error)
	EditKeranjang(keranjang Keranjang, id int, ctx context.Context) (Keranjang, error)
	DeleteKeranjang(id int, ctx context.Context) (Keranjang, error)
}
