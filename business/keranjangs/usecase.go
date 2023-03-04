package keranjangs

import (
	"context"
	"errors"
	"time"
)

type KeranjangUsecase struct {
	repo KeranjangRepoInterface
	ctx  time.Duration
}

func NewKeranjangUcecase(kRepo KeranjangRepoInterface, contextTimeout time.Duration) KeranjangUseCaseInterface {
	return &KeranjangUsecase{
		repo: kRepo,
		ctx:  contextTimeout,
	}
}

//fungsi harus menempel pada struct
func (kUseCase *KeranjangUsecase) InsertKeranjang(kk Keranjang, ctx context.Context) (Keranjang, error) {
	if kk.IdUser == 0 {
		return Keranjang{}, errors.New("User ID empty")
	}

	kkRepo, err := kUseCase.repo.InsertKeranjang(kk, ctx)

	if err != nil {
		return Keranjang{}, err
	}

	return kkRepo, nil
}

func (kUseCase *KeranjangUsecase) GetAllKeranjang(ctx context.Context) ([]Keranjang, error) {
	kkRepo, err := kUseCase.repo.GetAllKeranjang(ctx)

	if err != nil {
		return []Keranjang{}, err
	}

	return kkRepo, nil
}

func (kUseCase *KeranjangUsecase) GetKeranjangDetail(id int, ctx context.Context) (Keranjang, error) {
	if id == 0 {
		return Keranjang{}, errors.New("Keranjang ID empty")
	}

	kkRepo, err := kUseCase.repo.GetKeranjangDetail(id, ctx)

	if err != nil {
		return Keranjang{}, err
	}

	return kkRepo, nil
}

func (kUseCase *KeranjangUsecase) EditKeranjang(kk Keranjang, id int, ctx context.Context) (Keranjang, error) {
	if id == 0 {
		return Keranjang{}, errors.New("Keranjang ID empty")
	}

	kkRepo, err := kUseCase.repo.EditKeranjang(kk, id, ctx)

	if err != nil {
		return Keranjang{}, err
	}

	return kkRepo, nil
}

func (kUseCase *KeranjangUsecase) DeleteKeranjang(id int, ctx context.Context) (Keranjang, error) {
	if id == 0 {
		return Keranjang{}, errors.New("Keranjang ID empty")
	}

	kkRepo, err := kUseCase.repo.DeleteKeranjang(id, ctx)

	if err != nil {
		return Keranjang{}, err
	}

	return kkRepo, nil
}
