package keranjangs

import (
	"context"
	"kemejaku/business/keranjangs"

	"gorm.io/gorm"
)

type KeranjangRepository struct {
	db *gorm.DB
}

func NewKeranjangRepo(gormDb *gorm.DB) keranjangs.KeranjangRepoInterface {
	//yang direturn adalah interfacenya repo
	return &KeranjangRepository{
		db: gormDb,
	}
}

func (repo *KeranjangRepository) InsertKeranjang(kk keranjangs.Keranjang, ctx context.Context) (keranjangs.Keranjang, error) {
	keranjangDB := FromUsecase(kk)

	result := repo.db.Preload("KemejaKeranjang").Create(&keranjangDB)

	if result.Error != nil {
		return keranjangs.Keranjang{}, result.Error
	}
	return keranjangDB.ToUsecase(), nil
}

func (repo *KeranjangRepository) GetAllKeranjang(ctx context.Context) ([]keranjangs.Keranjang, error) {
	var keranjangDB []Keranjang

	result := repo.db.Preload("KemejaKeranjang").Find(&keranjangDB)

	if result.Error != nil {
		return []keranjangs.Keranjang{}, result.Error
	}
	return ToUsecaseList(keranjangDB), nil
}

func (repo *KeranjangRepository) GetKeranjangDetail(id int, ctx context.Context) (keranjangs.Keranjang, error) {
	var keranjangDB Keranjang

	result := repo.db.Preload("KemejaKeranjang").First(&keranjangDB, id)

	if result.Error != nil {
		return keranjangs.Keranjang{}, result.Error
	}
	return keranjangDB.ToUsecase(), nil
}

func (repo *KeranjangRepository) EditKeranjang(kk keranjangs.Keranjang, id int, ctx context.Context) (keranjangs.Keranjang, error) {
	keranjangDB := FromUsecase(kk)
	var newKeranjang Keranjang

	result := repo.db.Preload("KemejaKeranjang").First(&newKeranjang, id)

	if result.Error != nil {
		return keranjangs.Keranjang{}, result.Error
	}

	//ngecek kosong dan engga di mana?
	newKeranjang.Status = keranjangDB.Status

	repo.db.Save(&newKeranjang)
	return newKeranjang.ToUsecase(), nil
}

func (repo *KeranjangRepository) DeleteKeranjang(id int, ctx context.Context) (keranjangs.Keranjang, error) {
	var keranjangDB Keranjang

	resultFind := repo.db.First(&keranjangDB, id)

	if resultFind.Error != nil {
		return keranjangs.Keranjang{}, resultFind.Error
	}

	result := repo.db.Delete(&keranjangDB, id)

	//kalo ngecek ga ada id kayak gitu pake result kah?
	if result.Error != nil {
		return keranjangs.Keranjang{}, result.Error
	}

	return keranjangDB.ToUsecase(), nil
}
