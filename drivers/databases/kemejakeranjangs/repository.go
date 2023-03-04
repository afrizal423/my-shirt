package kemejakeranjangs

import (
	"context"
	"kemejaku/business/kemejakeranjangs"

	"gorm.io/gorm"
)

type KemejaKeranjangRepository struct {
	db *gorm.DB
}

func NewKemejaKeranjangRepo(gormDb *gorm.DB) kemejakeranjangs.KemejaKeranjangRepoInterface {
	//yang direturn adalah interfacenya repo
	return &KemejaKeranjangRepository{
		db: gormDb,
	}
}

func (repo *KemejaKeranjangRepository) InsertKemejaKeranjang(kk kemejakeranjangs.KemejaKeranjang, ctx context.Context) (kemejakeranjangs.KemejaKeranjang, error) {
	kkDB := FromUsecase(kk)

	result := repo.db.Create(&kkDB)

	if result.Error != nil {
		return kemejakeranjangs.KemejaKeranjang{}, result.Error
	}
	return kkDB.ToUsecase(), nil
}

func (repo *KemejaKeranjangRepository) GetAllKemejaKeranjang(ctx context.Context) ([]kemejakeranjangs.KemejaKeranjang, error) {
	var kkDB []KemejaKeranjang

	result := repo.db.Find(&kkDB)

	if result.Error != nil {
		return []kemejakeranjangs.KemejaKeranjang{}, result.Error
	}
	return ToUsecaseList(kkDB), nil
}

func (repo *KemejaKeranjangRepository) GetKemejaKeranjangDetail(id int, ctx context.Context) (kemejakeranjangs.KemejaKeranjang, error) {
	var kkDB KemejaKeranjang

	result := repo.db.First(&kkDB, id)

	if result.Error != nil {
		return kemejakeranjangs.KemejaKeranjang{}, result.Error
	}
	return kkDB.ToUsecase(), nil
}

func (repo *KemejaKeranjangRepository) EditKemejaKeranjang(kk kemejakeranjangs.KemejaKeranjang, id int, ctx context.Context) (kemejakeranjangs.KemejaKeranjang, error) {
	kkDB := FromUsecase(kk)
	var newKK KemejaKeranjang

	result := repo.db.First(&newKK, id)

	if result.Error != nil {
		return kemejakeranjangs.KemejaKeranjang{}, result.Error
	}

	//ngecek kosong dan engga di mana?
	newKK.Jumlah = kkDB.Jumlah
	newKK.Size = kkDB.Size

	repo.db.Save(&newKK)
	return newKK.ToUsecase(), nil
}

func (repo *KemejaKeranjangRepository) DeleteKemejaKeranjang(id int, ctx context.Context) (kemejakeranjangs.KemejaKeranjang, error) {
	var kkDB KemejaKeranjang

	resultFind := repo.db.First(&kkDB, id)

	if resultFind.Error != nil {
		return kemejakeranjangs.KemejaKeranjang{}, resultFind.Error
	}

	result := repo.db.Delete(&kkDB, id)

	//kalo ngecek ga ada id kayak gitu pake result kah?
	if result.Error != nil {
		return kemejakeranjangs.KemejaKeranjang{}, result.Error
	}

	return kkDB.ToUsecase(), nil
}
