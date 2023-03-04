package request

import "kemejaku/business/users"

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserEdit struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
	Street      string `json:"street"`
	Address     string `json:"address"`
	PostalCode  string `json:"postalCode"`
}

func (user *UserLogin) ToUsecase() *users.User {
	return &users.User{
		Email:    user.Email,
		Password: user.Password,
	}
}

func (user *UserEdit) ToUsecase() *users.User {
	return &users.User{
		Name:        user.Name,
		Email:       user.Email,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
		Street:      user.Street,
		Address:     user.Address,
		PostalCode:  user.PostalCode,
	}
}
