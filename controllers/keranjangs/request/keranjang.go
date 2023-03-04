package request

import "kemejaku/business/keranjangs"

type KeranjangInsert struct {
	IdUser int `json:"idUser"`
}

type KeranjangEdit struct {
	Status bool `json:"status"`
}

func (keranjang *KeranjangInsert) ToUsecase() *keranjangs.Keranjang {
	return &keranjangs.Keranjang{
		IdUser: keranjang.IdUser,
	}
}

func (keranjang *KeranjangEdit) ToUsecase() *keranjangs.Keranjang {
	return &keranjangs.Keranjang{
		Status: keranjang.Status,
	}
}
