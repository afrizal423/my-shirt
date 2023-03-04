package kemejas

import (
	"context"
	"errors"
	"time"
)

//buat struct supaya bisa pake interface biar bisa dipasangkan dengan yang lainnya
type KemejaUsecase struct {
	// interface repo
	repo KemejaRepoInterface
	ctx  time.Duration
}

//generate usecase baru
func NewKemejaUsecase(kemejaRepo KemejaRepoInterface, contextTimeout time.Duration) KemejaUseCaseInterface {
	return &KemejaUsecase{
		repo: kemejaRepo,
		ctx:  contextTimeout,
	}
}

func (kemejaUseCase *KemejaUsecase) InsertKemeja(kemeja Kemeja, ctx context.Context) (Kemeja, error) {
	if kemeja.Nama == "" {
		return Kemeja{}, errors.New("Nama empty")
	}

	if kemeja.Deskripsi == "" {
		return Kemeja{}, errors.New("Deskripsi empty")
	}

	if kemeja.Harga == 0 {
		return Kemeja{}, errors.New("Harga empty")
	}

	if kemeja.Stock_L == 0 {
		return Kemeja{}, errors.New("Kemeja stock for size L empty")
	}

	if kemeja.Stock_M == 0 {
		return Kemeja{}, errors.New("Kemeja stock for size M empty")
	}

	if kemeja.Stock_S == 0 {
		return Kemeja{}, errors.New("Kemeja stock for size S empty")
	}

	kemejaRepo, err := kemejaUseCase.repo.InsertKemeja(kemeja, ctx)

	if err != nil {
		return Kemeja{}, err
	}

	return kemejaRepo, nil
}

func (kemejaUseCase *KemejaUsecase) GetAllKemeja(ctx context.Context) ([]Kemeja, error) {
	kemejaRepo, err := kemejaUseCase.repo.GetAllKemeja(ctx)

	if err != nil {
		return []Kemeja{}, err
	}

	return kemejaRepo, nil
}

func (kemejaUseCase *KemejaUsecase) GetKemejaDetail(id int, ctx context.Context) (Kemeja, error) {
	if id == 0 {
		return Kemeja{}, errors.New("Kemeja ID empty")
	}

	kemejaRepo, err := kemejaUseCase.repo.GetKemejaDetail(id, ctx)

	if err != nil {
		return Kemeja{}, err
	}

	return kemejaRepo, nil
}

func (kemejaUseCase *KemejaUsecase) EditKemeja(kemeja Kemeja, id int, ctx context.Context) (Kemeja, error) {
	if id == 0 {
		return Kemeja{}, errors.New("Kemeja ID empty")
	}

	if kemeja.Nama == "" {
		return Kemeja{}, errors.New("Nama empty")
	}

	if kemeja.Deskripsi == "" {
		return Kemeja{}, errors.New("Deskripsi empty")
	}

	if kemeja.Harga == 0 {
		return Kemeja{}, errors.New("Harga empty")
	}

	kemejaRepo, err := kemejaUseCase.repo.EditKemeja(kemeja, id, ctx)

	if err != nil {
		return Kemeja{}, err
	}

	return kemejaRepo, nil
}

func (kemejaUseCase *KemejaUsecase) DeleteKemeja(id int, ctx context.Context) (Kemeja, error) {
	if id == 0 {
		return Kemeja{}, errors.New("Kemeja ID empty")
	}

	kemejaRepo, err := kemejaUseCase.repo.DeleteKemeja(id, ctx)

	if err != nil {
		return Kemeja{}, err
	}

	return kemejaRepo, nil
}
