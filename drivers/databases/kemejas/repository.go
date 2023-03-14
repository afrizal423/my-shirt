package kemejas

import (
	"context"

	"github.com/afrizal423/my-shirt/business/kemejas"

	"gorm.io/gorm"
)

type KemejaRepository struct {
	db *gorm.DB
}

func NewKemejaRepo(gormDb *gorm.DB) kemejas.KemejaRepoInterface {
	//yang direturn adalah interfacenya repo
	return &KemejaRepository{
		db: gormDb,
	}
}

func (repo *KemejaRepository) InsertKemeja(kemeja kemejas.Kemeja, ctx context.Context) (kemejas.Kemeja, error) {
	kemejaDB := FromUsecase(kemeja)

	result := repo.db.Create(&kemejaDB)

	if result.Error != nil {
		return kemejas.Kemeja{}, result.Error
	}
	return kemejaDB.ToUsecase(), nil
}

func (repo *KemejaRepository) GetAllKemeja(ctx context.Context) ([]kemejas.Kemeja, error) {
	var kemejaDB []Kemeja

	result := repo.db.Find(&kemejaDB)

	if result.Error != nil {
		return []kemejas.Kemeja{}, result.Error
	}
	return ToUsecaseList(kemejaDB), nil
}

func (repo *KemejaRepository) GetKemejaDetail(id int, ctx context.Context) (kemejas.Kemeja, error) {
	var kemejaDB Kemeja

	result := repo.db.First(&kemejaDB, id)

	if result.Error != nil {
		return kemejas.Kemeja{}, result.Error
	}
	return kemejaDB.ToUsecase(), nil
}

func (repo *KemejaRepository) EditKemeja(kemeja kemejas.Kemeja, id int, ctx context.Context) (kemejas.Kemeja, error) {
	kemejaDB := FromUsecase(kemeja)
	var newKemeja Kemeja

	result := repo.db.First(&newKemeja, id)

	if result.Error != nil {
		return kemejas.Kemeja{}, result.Error
	}

	//ngecek kosong dan engga di mana?
	newKemeja.Nama = kemejaDB.Nama
	newKemeja.Deskripsi = kemejaDB.Deskripsi
	newKemeja.Harga = kemejaDB.Harga
	newKemeja.Stock_L = kemejaDB.Stock_L
	newKemeja.Stock_M = kemejaDB.Stock_M
	newKemeja.Stock_S = kemejaDB.Stock_S

	repo.db.Save(&newKemeja)
	return newKemeja.ToUsecase(), nil
}

func (repo *KemejaRepository) DeleteKemeja(id int, ctx context.Context) (kemejas.Kemeja, error) {
	var kemejaDB Kemeja

	resultFind := repo.db.First(&kemejaDB, id)

	if resultFind.Error != nil {
		return kemejas.Kemeja{}, resultFind.Error
	}

	result := repo.db.Delete(&kemejaDB, id)

	//kalo ngecek ga ada id kayak gitu pake result kah?
	if result.Error != nil {
		return kemejas.Kemeja{}, result.Error
	}

	return kemejaDB.ToUsecase(), nil
}
