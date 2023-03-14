package sales

import (
	"context"

	"github.com/afrizal423/my-shirt/business/sales"

	"gorm.io/gorm"
)

type SaleRepository struct {
	db *gorm.DB
}

func NewSaleRepo(gormDb *gorm.DB) sales.SaleRepoInterface {
	//yang direturn adalah interfacenya repo
	return &SaleRepository{
		db: gormDb,
	}
}

func (repo *SaleRepository) InsertSale(sale sales.Sale, ctx context.Context) (sales.Sale, error) {
	saleDB := FromUsecase(sale)

	result := repo.db.Preload("Kemejas").Create(&saleDB)

	if result.Error != nil {
		return sales.Sale{}, result.Error
	}
	return saleDB.ToUsecase(), nil
}

func (repo *SaleRepository) GetAllSale(ctx context.Context) ([]sales.Sale, error) {
	var saleDB []Sale

	result := repo.db.Preload("Kemejas").Find(&saleDB)

	if result.Error != nil {
		return []sales.Sale{}, result.Error
	}
	return ToUsecaseList(saleDB), nil
}

func (repo *SaleRepository) GetSaleDetail(id int, ctx context.Context) (sales.Sale, error) {
	var saleDB Sale

	result := repo.db.Preload("Kemejas").First(&saleDB, id)

	if result.Error != nil {
		return sales.Sale{}, result.Error
	}
	return saleDB.ToUsecase(), nil
}

func (repo *SaleRepository) EditSale(kk sales.Sale, id int, ctx context.Context) (sales.Sale, error) {
	saleDB := FromUsecase(kk)
	var newSale Sale

	result := repo.db.Preload("Kemejas").First(&newSale, id)

	if result.Error != nil {
		return sales.Sale{}, result.Error
	}
	newSale.Percent = saleDB.Percent
	newSale.MinimumPembelian = saleDB.MinimumPembelian
	newSale.StartDate = saleDB.StartDate
	newSale.EndDate = saleDB.EndDate

	repo.db.Save(&newSale)
	return newSale.ToUsecase(), nil
}

func (repo *SaleRepository) DeleteSale(id int, ctx context.Context) (sales.Sale, error) {
	var saleDB Sale

	resultFind := repo.db.First(&saleDB, id)

	if resultFind.Error != nil {
		return sales.Sale{}, resultFind.Error
	}

	result := repo.db.Delete(&saleDB, id)

	//kalo ngecek ga ada id kayak gitu pake result kah?
	if result.Error != nil {
		return sales.Sale{}, result.Error
	}

	return saleDB.ToUsecase(), nil
}
