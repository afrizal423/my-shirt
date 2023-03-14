package request

import (
	"time"

	"github.com/afrizal423/my-shirt/business/sales"
)

type Sale struct {
	Percent          float64 `json:"percent"`
	MinimumPembelian int     `json:"minimumPembelian"`
	StartDate        string  `json:"startDate"`
	EndDate          string  `json:"endDate"`
}

func (sale *Sale) ToUsecase() *sales.Sale {
	layout := "Jan 2, 2006"
	startDate, _ := time.Parse(layout, sale.StartDate)
	endDate, _ := time.Parse(layout, sale.EndDate)

	return &sales.Sale{
		Percent:          sale.Percent,
		MinimumPembelian: sale.MinimumPembelian,
		StartDate:        startDate,
		EndDate:          endDate,
	}
}
