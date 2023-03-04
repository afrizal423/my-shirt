package response

import (
	"kemejaku/business/kemejas"
	"kemejaku/business/sales"
	"time"

	"gorm.io/gorm"
)

type SaleResponse struct {
	Id               int              `json:"id"`
	Percent          float64          `json:"percent"`
	MinimumPembelian int              `json:"minimumPembelian"`
	StartDate        time.Time        `json:"startDate"`
	EndDate          time.Time        `json:"endDate"`
	Kemejas          []kemejas.Kemeja `json:"kemejas"`
	CreatedAt        time.Time        `json:"createdAt"`
	UpdatedAt        time.Time        `json:"updatedAt"`
	DeletedAt        gorm.DeletedAt   `gorm:"index" json:"deletedAt"`
}

func FromUsecase(sale sales.Sale) SaleResponse {
	return SaleResponse{
		Id:               sale.Id,
		Percent:          sale.Percent,
		MinimumPembelian: sale.MinimumPembelian,
		StartDate:        sale.StartDate,
		EndDate:          sale.EndDate,
		Kemejas:          sale.Kemejas,
		CreatedAt:        sale.CreatedAt,
		UpdatedAt:        sale.UpdatedAt,
		DeletedAt:        sale.DeletedAt,
	}
}

func FromUsecaseList(sale []sales.Sale) []SaleResponse {
	var saleResponse []SaleResponse

	for _, v := range sale {
		saleResponse = append(saleResponse, FromUsecase(v))
	}

	return saleResponse
}
