package sales

import (
	"context"
	"errors"
	"time"
)

//buat struct supaya bisa pake interface biar bisa dipasangkan dengan yang lainnya
type SaleUsecase struct {
	// interface repo
	repo SaleRepoInterface
	ctx  time.Duration
}

//generate usecase baru
func NewSaleUsecase(saleRepo SaleRepoInterface, contextTimeout time.Duration) SaleUseCaseInterface {
	return &SaleUsecase{
		repo: saleRepo,
		ctx:  contextTimeout,
	}
}

func (saleUseCase *SaleUsecase) InsertSale(sale Sale, ctx context.Context) (Sale, error) {

	if sale.Percent == 0 {
		return Sale{}, errors.New("Percent empty")
	}

	if sale.MinimumPembelian == 0 {
		return Sale{}, errors.New("Minimum pembelian empty")
	}

	saleRepo, err := saleUseCase.repo.InsertSale(sale, ctx)

	if err != nil {
		return Sale{}, err
	}

	return saleRepo, nil
}

func (saleUseCase *SaleUsecase) GetAllSale(ctx context.Context) ([]Sale, error) {
	saleRepo, err := saleUseCase.repo.GetAllSale(ctx)

	if err != nil {
		return []Sale{}, err
	}

	return saleRepo, nil
}

func (saleUseCase *SaleUsecase) GetSaleDetail(id int, ctx context.Context) (Sale, error) {
	if id == 0 {
		return Sale{}, errors.New("Sale ID empty")
	}

	saleRepo, err := saleUseCase.repo.GetSaleDetail(id, ctx)

	if err != nil {
		return Sale{}, err
	}

	return saleRepo, nil
}

func (saleUseCase *SaleUsecase) EditSale(sale Sale, id int, ctx context.Context) (Sale, error) {
	if id == 0 {
		return Sale{}, errors.New("Sale ID empty")
	}

	if sale.Percent == 0 {
		return Sale{}, errors.New("Percent empty")
	}

	if sale.MinimumPembelian == 0 {
		return Sale{}, errors.New("Minimum pembelian empty")
	}

	saleRepo, err := saleUseCase.repo.EditSale(sale, id, ctx)

	if err != nil {
		return Sale{}, err
	}

	return saleRepo, nil
}

func (saleUseCase *SaleUsecase) DeleteSale(id int, ctx context.Context) (Sale, error) {
	if id == 0 {
		return Sale{}, errors.New("Sale ID empty")
	}

	saleRepo, err := saleUseCase.repo.DeleteSale(id, ctx)

	if err != nil {
		return Sale{}, err
	}

	return saleRepo, nil
}
