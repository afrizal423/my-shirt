package response

import (
	"time"

	"github.com/afrizal423/my-shirt/business/keranjangs"
	"github.com/afrizal423/my-shirt/business/users"

	"gorm.io/gorm"
)

type UserResponse struct {
	Id          int                  `json:"id"`
	Email       string               `json:"email"`
	Name        string               `json:"name"`
	Password    string               `json:"password"`
	Token       string               `json:"token"`
	PhoneNumber string               `json:"phoneNumber"`
	Street      string               `json:"street"`
	Address     string               `json:"address"`
	PostalCode  string               `json:"postalCode"`
	Keranjang   keranjangs.Keranjang `json:"keranjang"`
	CreatedAt   time.Time            `json:"createdAt"`
	UpdatedAt   time.Time            `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt       `json:"deletedAt"`
}

func FromUsecase(user users.User) UserResponse {
	return UserResponse{
		Id:          user.Id,
		Name:        user.Name,
		Email:       user.Email,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
		Street:      user.Street,
		Address:     user.Address,
		PostalCode:  user.PostalCode,
		Keranjang:   user.Keranjang,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		DeletedAt:   user.DeletedAt,
		Token:       user.Token,
	}
}

func FromUsecaseList(user []users.User) []UserResponse {
	var userResponse []UserResponse

	for _, v := range user {
		userResponse = append(userResponse, FromUsecase(v))
	}

	return userResponse
}
