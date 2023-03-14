package request

import "github.com/afrizal423/my-shirt/business/kemejakeranjangs"

type KemejaKeranjangInsert struct {
	IdKemeja    int    `json:"idKemeja"`
	IdKeranjang int    `json:"idKeranjang"`
	Jumlah      int    `json:"jumlah"`
	Size        string `json:"size"`
}

type KemejaKeranjangEdit struct {
	Jumlah int    `json:"jumlah"`
	Size   string `json:"size"`
}

func (kk *KemejaKeranjangInsert) ToUsecase() *kemejakeranjangs.KemejaKeranjang {
	return &kemejakeranjangs.KemejaKeranjang{
		IdKemeja:    kk.IdKemeja,
		IdKeranjang: kk.IdKeranjang,
		Jumlah:      kk.Jumlah,
		Size:        kk.Size,
	}
}

func (kk *KemejaKeranjangEdit) ToUsecase() *kemejakeranjangs.KemejaKeranjang {
	return &kemejakeranjangs.KemejaKeranjang{
		Jumlah: kk.Jumlah,
		Size:   kk.Size,
	}
}
