package sales

import (
	"time"

	"github.com/afrizal423/my-shirt/business/sales"
	"github.com/afrizal423/my-shirt/drivers/databases/kemejas"

	"gorm.io/gorm"
)

type Sale struct {
	Id               int `gorm:"primaryKey;unique;autoIncrement:true"`
	Percent          float64
	MinimumPembelian int
	StartDate        time.Time
	EndDate          time.Time
	Kemejas          []kemejas.Kemeja `gorm:"foreignKey:IdSale"`
	CreatedAt        time.Time        `gorm:"autoCreateTime"`
	UpdatedAt        time.Time        `gorm:"autoUpdateTime"`
	DeletedAt        gorm.DeletedAt   `gorm:"index"`
}

func (sale Sale) ToUsecase() sales.Sale {
	newKemejas := kemejas.ToUsecaseList(sale.Kemejas)

	return sales.Sale{
		Id:               sale.Id,
		Percent:          sale.Percent,
		MinimumPembelian: sale.MinimumPembelian,
		StartDate:        sale.StartDate,
		EndDate:          sale.EndDate,
		Kemejas:          newKemejas,
		CreatedAt:        sale.CreatedAt,
		UpdatedAt:        sale.UpdatedAt,
		DeletedAt:        sale.DeletedAt,
	}
}

func ToUsecaseList(sale []Sale) []sales.Sale {
	var newSales []sales.Sale

	for _, v := range sale {
		newSales = append(newSales, v.ToUsecase())
	}

	return newSales
}

func FromUsecase(sale sales.Sale) Sale {
	newKemejas := kemejas.FromUseCaseList(sale.Kemejas)

	return Sale{
		Id:               sale.Id,
		Percent:          sale.Percent,
		MinimumPembelian: sale.MinimumPembelian,
		StartDate:        sale.StartDate,
		EndDate:          sale.EndDate,
		Kemejas:          newKemejas,
		CreatedAt:        sale.CreatedAt,
		UpdatedAt:        sale.UpdatedAt,
		DeletedAt:        sale.DeletedAt,
	}
}
