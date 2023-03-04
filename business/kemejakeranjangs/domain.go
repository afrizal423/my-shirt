package kemejakeranjangs

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type KemejaKeranjang struct {
	Id          int
	IdKemeja    int
	IdKeranjang int
	Jumlah      int
	Size        string
	// Keranjang   keranjangs.Keranjang `gorm:"foreignKey:IdKeranjang"`
	// Kemeja      kemejas.Kemeja       `gorm:"foreignKey:IdKemeja"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

//ini interface antara controller dan usecase
type KemejaKeranjangUseCaseInterface interface {
	//perlu konteks biar ada timeout
	InsertKemejaKeranjang(kk KemejaKeranjang, ctx context.Context) (KemejaKeranjang, error)
	GetAllKemejaKeranjang(ctx context.Context) ([]KemejaKeranjang, error)
	GetKemejaKeranjangDetail(id int, ctx context.Context) (KemejaKeranjang, error)
	EditKemejaKeranjang(kk KemejaKeranjang, id int, ctx context.Context) (KemejaKeranjang, error)
	DeleteKemejaKeranjang(id int, ctx context.Context) (KemejaKeranjang, error)
}

//ini interface antara usecase dan repo
type KemejaKeranjangRepoInterface interface {
	InsertKemejaKeranjang(kk KemejaKeranjang, ctx context.Context) (KemejaKeranjang, error)
	GetAllKemejaKeranjang(ctx context.Context) ([]KemejaKeranjang, error)
	GetKemejaKeranjangDetail(id int, ctx context.Context) (KemejaKeranjang, error)
	EditKemejaKeranjang(kk KemejaKeranjang, id int, ctx context.Context) (KemejaKeranjang, error)
	DeleteKemejaKeranjang(id int, ctx context.Context) (KemejaKeranjang, error)
}
