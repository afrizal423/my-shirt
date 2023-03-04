package users

import (
	"kemejaku/business/users"
	"kemejaku/drivers/databases/keranjangs"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id          int `gorm:"primaryKey;unique;autoIncrement:true"`
	Name        string
	Email       string `gorm:"unique"`
	Password    string
	PhoneNumber string
	Street      string
	Address     string
	PostalCode  string
	Keranjang   keranjangs.Keranjang `gorm:"foreignKey:IdUser"`
	CreatedAt   time.Time            `gorm:"autoCreateTime"`
	UpdatedAt   time.Time            `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt       `gorm:"index"`
}

func (user User) ToUsecase() users.User {
	newKeranjang := user.Keranjang.ToUsecase()

	return users.User{
		Id:          user.Id,
		Name:        user.Name,
		Email:       user.Email,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
		Street:      user.Street,
		Address:     user.Address,
		PostalCode:  user.PostalCode,
		Keranjang:   newKeranjang,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		DeletedAt:   user.DeletedAt,
	}
}

func ToUsecaseList(user []User) []users.User {
	var newUsers []users.User

	for _, v := range user {
		newUsers = append(newUsers, v.ToUsecase())
	}
	return newUsers
}

func FromUsecase(user users.User) User {
	newKeranjang := keranjangs.FromUsecase(user.Keranjang)

	return User{
		Id:          user.Id,
		Name:        user.Name,
		Email:       user.Email,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
		Street:      user.Street,
		Address:     user.Address,
		PostalCode:  user.PostalCode,
		Keranjang:   newKeranjang,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		DeletedAt:   user.DeletedAt,
	}
}
