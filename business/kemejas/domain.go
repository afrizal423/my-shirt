package kemejas

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Kemeja struct {
	Id        int
	Nama      string
	Deskripsi string
	Harga     int
	Stock_L   int
	Stock_M   int
	Stock_S   int
	IdSale    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

//ini interface antara controller dan usecase
type KemejaUseCaseInterface interface {
	//perlu konteks biar ada timeout
	InsertKemeja(kemeja Kemeja, ctx context.Context) (Kemeja, error)
	GetAllKemeja(ctx context.Context) ([]Kemeja, error)
	GetKemejaDetail(id int, ctx context.Context) (Kemeja, error)
	EditKemeja(kemeja Kemeja, id int, ctx context.Context) (Kemeja, error)
	DeleteKemeja(id int, ctx context.Context) (Kemeja, error)
}

//ini interface antara usecase dan repo
type KemejaRepoInterface interface {
	InsertKemeja(kemeja Kemeja, ctx context.Context) (Kemeja, error)
	GetAllKemeja(ctx context.Context) ([]Kemeja, error)
	GetKemejaDetail(id int, ctx context.Context) (Kemeja, error)
	EditKemeja(kemeja Kemeja, id int, ctx context.Context) (Kemeja, error)
	DeleteKemeja(id int, ctx context.Context) (Kemeja, error)
}
