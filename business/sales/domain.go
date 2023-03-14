package sales

import (
	"context"
	"time"

	"github.com/afrizal423/my-shirt/business/kemejas"

	"gorm.io/gorm"
)

// stuct untuk keluar masuk usecase
type Sale struct {
	Id               int
	Percent          float64
	MinimumPembelian int
	StartDate        time.Time
	EndDate          time.Time
	Kemejas          []kemejas.Kemeja
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
}

// ini interface antara controller dan usecase
type SaleUseCaseInterface interface {
	//perlu konteks biar ada timeout
	InsertSale(sale Sale, ctx context.Context) (Sale, error)
	GetAllSale(ctx context.Context) ([]Sale, error)
	GetSaleDetail(id int, ctx context.Context) (Sale, error)
	EditSale(sale Sale, id int, ctx context.Context) (Sale, error)
	DeleteSale(id int, ctx context.Context) (Sale, error)
}

// ini interface antara usecase dan repo
type SaleRepoInterface interface {
	InsertSale(sale Sale, ctx context.Context) (Sale, error)
	GetAllSale(ctx context.Context) ([]Sale, error)
	GetSaleDetail(id int, ctx context.Context) (Sale, error)
	EditSale(sale Sale, id int, ctx context.Context) (Sale, error)
	DeleteSale(id int, ctx context.Context) (Sale, error)
}
