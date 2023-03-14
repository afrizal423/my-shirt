package kemejas_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/afrizal423/my-shirt/business/kemejas"
	"github.com/afrizal423/my-shirt/business/kemejas/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// buat mock yang seolah-olah interface dari database
var kemejaRepoInterfaceMock mocks.KemejaRepoInterface
var kemejaUseCaseInterface kemejas.KemejaUseCaseInterface
var kemejaDataDummy, kemejaDataDummyEdit kemejas.Kemeja
var kemejaDataDummyGetAllKemejas []kemejas.Kemeja

func setup() {
	kemejaUseCaseInterface = kemejas.NewKemejaUsecase(&kemejaRepoInterfaceMock, time.Hour*1)

	//data mock hasil login
	kemejaDataDummy = kemejas.Kemeja{
		Id:        1,
		Nama:      "Basic shirt",
		Deskripsi: "Bahan rayon",
		Harga:     150000,
		Stock_L:   10,
		Stock_M:   10,
		Stock_S:   10,
		IdSale:    1,
	}

	kemejaDataDummyGetAllKemejas = []kemejas.Kemeja{
		{
			Id:        1,
			Nama:      "Basic shirt",
			Deskripsi: "Bahan rayon",
			Harga:     150000,
			Stock_L:   10,
			Stock_M:   10,
			Stock_S:   10,
			IdSale:    1,
		},
		{
			Id:        2,
			Nama:      "Basic shirt",
			Deskripsi: "Bahan rayon",
			Harga:     150000,
			Stock_L:   10,
			Stock_M:   10,
			Stock_S:   10,
			IdSale:    1,
		},
	}
}

func TestInsertKemeja(t *testing.T) {
	setup()
	t.Run("Success insert", func(t *testing.T) {
		kemejaRepoInterfaceMock.On("InsertKemeja", mock.AnythingOfType("kemejas.Kemeja"), mock.Anything).Return(kemejaDataDummy, nil).Once()

		var requestInsertKemeja = kemejas.Kemeja{
			Nama:      "Basic shirt",
			Deskripsi: "Bahan rayon",
			Harga:     150000,
			Stock_L:   10,
			Stock_M:   10,
			Stock_S:   10,
			IdSale:    1,
		}

		kemeja, err := kemejaUseCaseInterface.InsertKemeja(requestInsertKemeja, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, kemejaDataDummy, kemeja)
	})

	t.Run("Nama empty", func(t *testing.T) {
		kemejaRepoInterfaceMock.On("InsertKemeja", mock.AnythingOfType("keranjangs.Keranjang"), mock.Anything).Return(kemejas.Kemeja{}, errors.New("Nama empty")).Once()

		var requestInsertKemeja = kemejas.Kemeja{
			Nama:      "",
			Deskripsi: "Bahan rayon",
			Harga:     150000,
			Stock_L:   10,
			Stock_M:   10,
			Stock_S:   10,
			IdSale:    1,
		}

		user, err := kemejaUseCaseInterface.InsertKemeja(requestInsertKemeja, context.Background())

		assert.Equal(t, errors.New("Nama empty"), err)
		assert.Equal(t, kemejas.Kemeja{}, user)
	})

	t.Run("Deskripsi empty", func(t *testing.T) {
		kemejaRepoInterfaceMock.On("InsertKemeja", mock.AnythingOfType("keranjangs.Keranjang"), mock.Anything).Return(kemejas.Kemeja{}, errors.New("Deskripsi empty")).Once()

		var requestInsertKemeja = kemejas.Kemeja{
			Nama:      "Flanel shirt",
			Deskripsi: "",
			Harga:     150000,
			Stock_L:   10,
			Stock_M:   10,
			Stock_S:   10,
			IdSale:    1,
		}

		user, err := kemejaUseCaseInterface.InsertKemeja(requestInsertKemeja, context.Background())

		assert.Equal(t, errors.New("Deskripsi empty"), err)
		assert.Equal(t, kemejas.Kemeja{}, user)
	})

	t.Run("Harga empty", func(t *testing.T) {
		kemejaRepoInterfaceMock.On("InsertKemeja", mock.AnythingOfType("keranjangs.Keranjang"), mock.Anything).Return(kemejas.Kemeja{}, errors.New("Harga empty")).Once()

		var requestInsertKemeja = kemejas.Kemeja{
			Nama:      "Flanel shirt",
			Deskripsi: "Bahan flanel",
			Harga:     0,
			Stock_L:   10,
			Stock_M:   10,
			Stock_S:   10,
			IdSale:    1,
		}

		user, err := kemejaUseCaseInterface.InsertKemeja(requestInsertKemeja, context.Background())

		assert.Equal(t, errors.New("Harga empty"), err)
		assert.Equal(t, kemejas.Kemeja{}, user)
	})

	t.Run("Size L empty", func(t *testing.T) {
		kemejaRepoInterfaceMock.On("InsertKemeja", mock.AnythingOfType("keranjangs.Keranjang"), mock.Anything).Return(kemejas.Kemeja{}, errors.New("Kemeja stock for size L empty")).Once()

		var requestInsertKemeja = kemejas.Kemeja{
			Nama:      "Flanel shirt",
			Deskripsi: "Bahan flanel",
			Harga:     150000,
			Stock_L:   0,
			Stock_M:   10,
			Stock_S:   10,
			IdSale:    1,
		}

		user, err := kemejaUseCaseInterface.InsertKemeja(requestInsertKemeja, context.Background())

		assert.Equal(t, errors.New("Kemeja stock for size L empty"), err)
		assert.Equal(t, kemejas.Kemeja{}, user)
	})

	t.Run("Size M empty", func(t *testing.T) {
		kemejaRepoInterfaceMock.On("InsertKemeja", mock.AnythingOfType("keranjangs.Keranjang"), mock.Anything).Return(kemejas.Kemeja{}, errors.New("Kemeja stock for size M empty")).Once()

		var requestInsertKemeja = kemejas.Kemeja{
			Nama:      "Flanel shirt",
			Deskripsi: "Bahan flanel",
			Harga:     150000,
			Stock_L:   10,
			Stock_M:   0,
			Stock_S:   10,
			IdSale:    1,
		}

		user, err := kemejaUseCaseInterface.InsertKemeja(requestInsertKemeja, context.Background())

		assert.Equal(t, errors.New("Kemeja stock for size M empty"), err)
		assert.Equal(t, kemejas.Kemeja{}, user)
	})

	t.Run("Size S empty", func(t *testing.T) {
		kemejaRepoInterfaceMock.On("InsertKemeja", mock.AnythingOfType("keranjangs.Keranjang"), mock.Anything).Return(kemejas.Kemeja{}, errors.New("Kemeja stock for size S empty")).Once()

		var requestInsertKemeja = kemejas.Kemeja{
			Nama:      "Flanel shirt",
			Deskripsi: "Bahan flanel",
			Harga:     150000,
			Stock_L:   10,
			Stock_M:   10,
			Stock_S:   0,
			IdSale:    1,
		}

		user, err := kemejaUseCaseInterface.InsertKemeja(requestInsertKemeja, context.Background())

		assert.Equal(t, errors.New("Kemeja stock for size S empty"), err)
		assert.Equal(t, kemejas.Kemeja{}, user)
	})

	t.Run("Error in database", func(t *testing.T) {
		kemejaRepoInterfaceMock.On("InsertKemeja", mock.AnythingOfType("kemejas.Kemeja"), mock.Anything).Return(kemejas.Kemeja{}, errors.New("Insert failed")).Once()

		var requestInsertKemeja = kemejas.Kemeja{
			Nama:      "Basic shirt",
			Deskripsi: "Bahan rayon",
			Harga:     150000,
			Stock_L:   10,
			Stock_M:   10,
			Stock_S:   10,
			IdSale:    1,
		}
		kemeja, err := kemejaUseCaseInterface.InsertKemeja(requestInsertKemeja, context.Background())

		assert.Error(t, err)
		assert.Equal(t, kemejas.Kemeja{}, kemeja)
	})
}

func TestGetAllkemeja(t *testing.T) {
	setup()
	t.Run("Success Get All kemejas", func(t *testing.T) {
		kemejaRepoInterfaceMock.On("GetAllKemeja", mock.Anything, mock.Anything).Return(kemejaDataDummyGetAllKemejas, nil).Once()

		kemeja, err := kemejaUseCaseInterface.GetAllKemeja(context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, kemejaDataDummyGetAllKemejas, kemeja)
	})

	t.Run("Error in database", func(t *testing.T) {
		kemejaRepoInterfaceMock.On("GetAllKemeja", mock.Anything, mock.Anything).Return([]kemejas.Kemeja{}, errors.New("Error in database")).Once()

		kemeja, err := kemejaUseCaseInterface.GetAllKemeja(context.Background())

		assert.Error(t, err)
		assert.Equal(t, []kemejas.Kemeja{}, kemeja)
	})
}

func TestGetkemejaDetail(t *testing.T) {
	setup()
	t.Run("Success Get kemeja Detail", func(t *testing.T) {
		kemejaRepoInterfaceMock.On("GetKemejaDetail", mock.Anything, mock.Anything).Return(kemejaDataDummy, nil).Once()

		kemeja, err := kemejaUseCaseInterface.GetKemejaDetail(1, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, kemejaDataDummy, kemeja)
	})

	t.Run("Kemeja ID empty", func(t *testing.T) {
		kemejaRepoInterfaceMock.On("GetKemejaDetail", mock.Anything, mock.Anything).Return(kemejas.Kemeja{}, errors.New("Kemeja ID empty")).Once()

		user, err := kemejaUseCaseInterface.GetKemejaDetail(0, context.Background())

		assert.Equal(t, errors.New("Kemeja ID empty"), err)
		assert.Equal(t, kemejas.Kemeja{}, user)
	})

	t.Run("Kemejas not found in database", func(t *testing.T) {
		kemejaRepoInterfaceMock.On("GetKemejaDetail", mock.Anything, mock.Anything).Return(kemejas.Kemeja{}, errors.New("Kemeja not found")).Once()

		kemeja, err := kemejaUseCaseInterface.GetKemejaDetail(-1, context.Background())

		assert.Error(t, err)
		assert.Equal(t, kemejas.Kemeja{}, kemeja)
	})
}

func TestEditkemeja(t *testing.T) {
	setup()
	t.Run("Success Edit", func(t *testing.T) {
		kemejaRepoInterfaceMock.On("EditKemeja", mock.AnythingOfType("kemejas.Kemeja"), mock.Anything, mock.Anything).Return(kemejaDataDummyEdit, nil).Once()

		var requestEditKemeja = kemejas.Kemeja{
			Nama:      "Basic shirt",
			Deskripsi: "Bahan rayon",
			Harga:     150000,
			Stock_L:   10,
			Stock_M:   10,
			Stock_S:   10,
			IdSale:    1,
		}

		kemeja, err := kemejaUseCaseInterface.EditKemeja(requestEditKemeja, 1, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, kemejaDataDummyEdit, kemeja)
	})

	t.Run("Kemeja ID empty", func(t *testing.T) {
		kemejaRepoInterfaceMock.On("EditKemeja", mock.AnythingOfType("keranjangs.Keranjang"), mock.Anything, mock.Anything).Return(kemejas.Kemeja{}, errors.New("Kemeja ID empty")).Once()

		var requestEditKemeja = kemejas.Kemeja{
			Nama:      "Flanel shirt",
			Deskripsi: "Bahan flanel",
			Harga:     150000,
			Stock_L:   10,
			Stock_M:   10,
			Stock_S:   0,
			IdSale:    1,
		}

		user, err := kemejaUseCaseInterface.EditKemeja(requestEditKemeja, 0, context.Background())

		assert.Equal(t, errors.New("Kemeja ID empty"), err)
		assert.Equal(t, kemejas.Kemeja{}, user)
	})

	t.Run("Nama empty", func(t *testing.T) {
		kemejaRepoInterfaceMock.On("EditKemeja", mock.AnythingOfType("keranjangs.Keranjang"), mock.Anything, mock.Anything).Return(kemejas.Kemeja{}, errors.New("Nama empty")).Once()

		var requestEditKemeja = kemejas.Kemeja{
			Nama:      "",
			Deskripsi: "Bahan flanel",
			Harga:     150000,
			Stock_L:   10,
			Stock_M:   10,
			Stock_S:   10,
			IdSale:    1,
		}

		user, err := kemejaUseCaseInterface.EditKemeja(requestEditKemeja, 1, context.Background())

		assert.Equal(t, errors.New("Nama empty"), err)
		assert.Equal(t, kemejas.Kemeja{}, user)
	})

	t.Run("Deskripsi empty", func(t *testing.T) {
		kemejaRepoInterfaceMock.On("EditKemeja", mock.AnythingOfType("keranjangs.Keranjang"), mock.Anything, mock.Anything).Return(kemejas.Kemeja{}, errors.New("Deskripsi empty")).Once()

		var requestEditKemeja = kemejas.Kemeja{
			Nama:      "Flanel shirt",
			Deskripsi: "",
			Harga:     150000,
			Stock_L:   10,
			Stock_M:   10,
			Stock_S:   10,
			IdSale:    1,
		}

		user, err := kemejaUseCaseInterface.EditKemeja(requestEditKemeja, 1, context.Background())

		assert.Equal(t, errors.New("Deskripsi empty"), err)
		assert.Equal(t, kemejas.Kemeja{}, user)
	})

	t.Run("Harga empty", func(t *testing.T) {
		kemejaRepoInterfaceMock.On("EditKemeja", mock.AnythingOfType("keranjangs.Keranjang"), mock.Anything, mock.Anything).Return(kemejas.Kemeja{}, errors.New("Harga empty")).Once()

		var requestEditKemeja = kemejas.Kemeja{
			Nama:      "Flanel shirt",
			Deskripsi: "Bahan flanel",
			Harga:     0,
			Stock_L:   10,
			Stock_M:   10,
			Stock_S:   10,
			IdSale:    1,
		}

		user, err := kemejaUseCaseInterface.EditKemeja(requestEditKemeja, 1, context.Background())

		assert.Equal(t, errors.New("Harga empty"), err)
		assert.Equal(t, kemejas.Kemeja{}, user)
	})

	t.Run("Kemeja not found", func(t *testing.T) {
		kemejaRepoInterfaceMock.On("EditKemeja", mock.AnythingOfType("kemejas.Kemeja"), mock.Anything, mock.Anything).Return(kemejas.Kemeja{}, errors.New("Kemeja not found")).Once()

		var requestEditKemeja = kemejas.Kemeja{
			Nama:      "Basic shirt",
			Deskripsi: "Bahan rayon",
			Harga:     150000,
			Stock_L:   10,
			Stock_M:   10,
			Stock_S:   10,
			IdSale:    1,
		}
		kemeja, err := kemejaUseCaseInterface.EditKemeja(requestEditKemeja, 1, context.Background())

		assert.Error(t, err)
		assert.Equal(t, kemejas.Kemeja{}, kemeja)
	})
}

func TestDeletekemeja(t *testing.T) {
	setup()
	t.Run("Success delete", func(t *testing.T) {
		kemejaRepoInterfaceMock.On("DeleteKemeja", mock.Anything, mock.Anything).Return(kemejaDataDummy, nil).Once()

		kemeja, err := kemejaUseCaseInterface.DeleteKemeja(1, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, kemejaDataDummy, kemeja)
	})

	t.Run("Kemeja ID empty", func(t *testing.T) {
		kemejaRepoInterfaceMock.On("DeleteKemeja", mock.Anything, mock.Anything).Return(kemejas.Kemeja{}, errors.New("Kemeja ID empty")).Once()

		user, err := kemejaUseCaseInterface.DeleteKemeja(0, context.Background())

		assert.Equal(t, errors.New("Kemeja ID empty"), err)
		assert.Equal(t, kemejas.Kemeja{}, user)
	})

	t.Run("Kemeja not found", func(t *testing.T) {
		kemejaRepoInterfaceMock.On("DeleteKemeja", mock.Anything, mock.Anything).Return(kemejas.Kemeja{}, errors.New("Kemeja not found")).Once()

		kemeja, err := kemejaUseCaseInterface.DeleteKemeja(-1, context.Background())

		assert.Error(t, err)
		assert.Equal(t, kemejas.Kemeja{}, kemeja)
	})
}
