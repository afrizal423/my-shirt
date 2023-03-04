package sales_test

import (
	"context"
	"errors"
	"kemejaku/business/sales"
	"kemejaku/business/sales/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

//buat mock yang seolah-olah interface dari database
var saleRepoInterfaceMock mocks.SaleRepoInterface
var saleUsecaseInterface sales.SaleUseCaseInterface
var saleDataDummy sales.Sale
var saleDataDummyGetAll []sales.Sale

func setup() {
	saleUsecaseInterface = sales.NewSaleUsecase(&saleRepoInterfaceMock, time.Hour*1)

	//data mock hasil login
	saleDataDummy = sales.Sale{
		Id:               1,
		Percent:          30.5,
		MinimumPembelian: 2,
		StartDate:        time.Now(),
		EndDate:          time.Now(),
	}

	saleDataDummyGetAll = []sales.Sale{
		{
			Id:               1,
			Percent:          30.5,
			MinimumPembelian: 2,
			StartDate:        time.Now(),
			EndDate:          time.Now(),
		},
		{
			Id:               2,
			Percent:          10,
			MinimumPembelian: 2,
			StartDate:        time.Now(),
			EndDate:          time.Now(),
		},
	}
}

func TestInsertSale(t *testing.T) {
	setup()
	t.Run("Success insert", func(t *testing.T) {
		saleRepoInterfaceMock.On("InsertSale", mock.AnythingOfType("sales.Sale"), mock.Anything).Return(saleDataDummy, nil).Once()

		var requestInsertSale = sales.Sale{
			Percent:          30.5,
			MinimumPembelian: 2,
			StartDate:        time.Now(),
			EndDate:          time.Now(),
		}

		sale, err := saleUsecaseInterface.InsertSale(requestInsertSale, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, saleDataDummy, sale)
	})

	t.Run("Percent empty", func(t *testing.T) {
		saleRepoInterfaceMock.On("InsertSale", mock.AnythingOfType("sales.Sale"), mock.Anything).Return(sales.Sale{}, errors.New("Percent empty")).Once()

		var requestInsertSale = sales.Sale{
			Percent:          0,
			MinimumPembelian: 2,
			StartDate:        time.Now(),
			EndDate:          time.Now(),
		}
		sale, err := saleUsecaseInterface.InsertSale(requestInsertSale, context.Background())

		assert.Equal(t, errors.New("Percent empty"), err)
		assert.Equal(t, sales.Sale{}, sale)
	})

	t.Run("Minimum pembelian empty", func(t *testing.T) {
		saleRepoInterfaceMock.On("InsertSale", mock.AnythingOfType("sales.Sale"), mock.Anything).Return(sales.Sale{}, errors.New("Minimum pembelian empty")).Once()

		var requestInsertSale = sales.Sale{
			Percent:          30.5,
			MinimumPembelian: 0,
			StartDate:        time.Now(),
			EndDate:          time.Now(),
		}
		sale, err := saleUsecaseInterface.InsertSale(requestInsertSale, context.Background())

		assert.Equal(t, errors.New("Minimum pembelian empty"), err)
		assert.Equal(t, sales.Sale{}, sale)
	})

	t.Run("Error in database", func(t *testing.T) {
		saleRepoInterfaceMock.On("InsertSale", mock.AnythingOfType("sales.Sale"), mock.Anything).Return(sales.Sale{}, errors.New("Insert failed")).Once()

		var requestInsertSale = sales.Sale{
			Percent:          30.5,
			MinimumPembelian: 2,
			StartDate:        time.Now(),
			EndDate:          time.Now(),
		}
		sale, err := saleUsecaseInterface.InsertSale(requestInsertSale, context.Background())

		assert.Error(t, err)
		assert.Equal(t, sales.Sale{}, sale)
	})
}

func TestGetAllSale(t *testing.T) {
	setup()
	t.Run("Success Get All Sales", func(t *testing.T) {
		saleRepoInterfaceMock.On("GetAllSale", mock.Anything, mock.Anything).Return(saleDataDummyGetAll, nil).Once()

		sale, err := saleUsecaseInterface.GetAllSale(context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, saleDataDummyGetAll, sale)
	})

	t.Run("Error in database", func(t *testing.T) {
		saleRepoInterfaceMock.On("GetAllSale", mock.Anything, mock.Anything).Return([]sales.Sale{}, errors.New("Error in database")).Once()

		sale, err := saleUsecaseInterface.GetAllSale(context.Background())

		assert.Error(t, err)
		assert.Equal(t, []sales.Sale{}, sale)
	})
}

func TestGetSaleDetail(t *testing.T) {
	setup()
	t.Run("Success Get sale Detail", func(t *testing.T) {
		saleRepoInterfaceMock.On("GetSaleDetail", mock.Anything, mock.Anything).Return(saleDataDummy, nil).Once()

		sale, err := saleUsecaseInterface.GetSaleDetail(1, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, saleDataDummy, sale)
	})

	t.Run("Sale ID empty", func(t *testing.T) {
		saleRepoInterfaceMock.On("GetSaleDetail", mock.Anything, mock.Anything).Return(sales.Sale{}, errors.New("Sale ID empty")).Once()

		sale, err := saleUsecaseInterface.GetSaleDetail(0, context.Background())

		assert.Equal(t, errors.New("Sale ID empty"), err)
		assert.Equal(t, sales.Sale{}, sale)
	})

	t.Run("Sale not found in database", func(t *testing.T) {
		saleRepoInterfaceMock.On("GetSaleDetail", mock.Anything, mock.Anything).Return(sales.Sale{}, errors.New("Sale not found")).Once()

		sale, err := saleUsecaseInterface.GetSaleDetail(-1, context.Background())

		assert.Error(t, err)
		assert.Equal(t, sales.Sale{}, sale)
	})
}

func TestEditSale(t *testing.T) {
	setup()
	t.Run("Success Edit", func(t *testing.T) {
		saleRepoInterfaceMock.On("EditSale", mock.AnythingOfType("sales.Sale"), mock.Anything, mock.Anything).Return(saleDataDummy, nil).Once()

		var requestEditSale = sales.Sale{
			Percent:          30.5,
			MinimumPembelian: 2,
			StartDate:        time.Now(),
			EndDate:          time.Now(),
		}

		sale, err := saleUsecaseInterface.EditSale(requestEditSale, 1, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, saleDataDummy, sale)
	})

	t.Run("Sale ID empty", func(t *testing.T) {
		saleRepoInterfaceMock.On("EditSale", mock.AnythingOfType("sales.Sale"), mock.Anything, mock.Anything).Return(sales.Sale{}, errors.New("Sale ID empty")).Once()

		var requestEditSale = sales.Sale{
			Percent:          30.5,
			MinimumPembelian: 2,
			StartDate:        time.Now(),
			EndDate:          time.Now(),
		}

		sale, err := saleUsecaseInterface.EditSale(requestEditSale, 0, context.Background())

		assert.Equal(t, errors.New("Sale ID empty"), err)
		assert.Equal(t, sales.Sale{}, sale)
	})

	t.Run("Percent empty", func(t *testing.T) {
		saleRepoInterfaceMock.On("EditSale", mock.AnythingOfType("sales.Sale"), mock.Anything, mock.Anything).Return(sales.Sale{}, errors.New("Percent empty")).Once()

		var requestEditSale = sales.Sale{
			Percent:          0,
			MinimumPembelian: 2,
			StartDate:        time.Now(),
			EndDate:          time.Now(),
		}

		sale, err := saleUsecaseInterface.EditSale(requestEditSale, 1, context.Background())

		assert.Equal(t, errors.New("Percent empty"), err)
		assert.Equal(t, sales.Sale{}, sale)
	})

	t.Run("Minimum pembelian empty", func(t *testing.T) {
		saleRepoInterfaceMock.On("EditSale", mock.AnythingOfType("sales.Sale"), mock.Anything, mock.Anything).Return(sales.Sale{}, errors.New("Minimum pembelian empty")).Once()

		var requestEditSale = sales.Sale{
			Percent:          30.5,
			MinimumPembelian: 0,
			StartDate:        time.Now(),
			EndDate:          time.Now(),
		}

		sale, err := saleUsecaseInterface.EditSale(requestEditSale, 1, context.Background())

		assert.Equal(t, errors.New("Minimum pembelian empty"), err)
		assert.Equal(t, sales.Sale{}, sale)
	})

	t.Run("Sale not found", func(t *testing.T) {
		saleRepoInterfaceMock.On("EditSale", mock.AnythingOfType("sales.Sale"), mock.Anything, mock.Anything).Return(sales.Sale{}, errors.New("Sale not found")).Once()

		var requestEditSale = sales.Sale{
			Percent:          30.5,
			MinimumPembelian: 2,
			StartDate:        time.Now(),
			EndDate:          time.Now(),
		}

		sale, err := saleUsecaseInterface.EditSale(requestEditSale, 1, context.Background())

		assert.Error(t, err)
		assert.Equal(t, sales.Sale{}, sale)
	})
}

func TestDeleteSale(t *testing.T) {
	setup()
	t.Run("Success delete", func(t *testing.T) {
		saleRepoInterfaceMock.On("DeleteSale", mock.Anything, mock.Anything).Return(saleDataDummy, nil).Once()

		sale, err := saleUsecaseInterface.DeleteSale(1, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, saleDataDummy, sale)
	})

	t.Run("Sale ID empty", func(t *testing.T) {
		saleRepoInterfaceMock.On("DeleteSale", mock.Anything, mock.Anything).Return(sales.Sale{}, errors.New("Sale ID empty")).Once()

		sale, err := saleUsecaseInterface.DeleteSale(0, context.Background())

		assert.Equal(t, errors.New("Sale ID empty"), err)
		assert.Equal(t, sales.Sale{}, sale)
	})

	t.Run("Sale not found", func(t *testing.T) {
		saleRepoInterfaceMock.On("DeleteSale", mock.Anything, mock.Anything).Return(sales.Sale{}, errors.New("Sale not found")).Once()

		sale, err := saleUsecaseInterface.DeleteSale(-1, context.Background())

		assert.Error(t, err)
		assert.Equal(t, sales.Sale{}, sale)
	})
}
