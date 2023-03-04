package users

import (
	"context"
	"errors"
	"kemejaku/business/users"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(gormDb *gorm.DB) users.UserRepoInterface {
	//yang direturn adalah interfacenya repo
	return &UserRepository{
		db: gormDb,
	}
}

func (repo *UserRepository) Login(user users.User, ctx context.Context) (users.User, error) {
	userDB := FromUsecase(user)

	err := repo.db.Where("email = ?", userDB.Email).Preload("Keranjang").First(&userDB).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return users.User{}, errors.New("User not found")
		}
		return users.User{}, errors.New("Error in database")
	}

	return userDB.ToUsecase(), nil
}

func (repo *UserRepository) GetAllUsers(ctx context.Context) ([]users.User, error) {
	var usersDb []User

	result := repo.db.Preload("Keranjang").Preload("Keranjang").Find(&usersDb)

	if result.Error != nil {
		return []users.User{}, result.Error
	}

	// Kalo mau mengubah array
	return ToUsecaseList(usersDb), nil
}

func (repo *UserRepository) SignUp(user users.User, ctx context.Context) (users.User, error) {
	userDB := FromUsecase(user)

	result := repo.db.Preload("Keranjang").Create(&userDB)

	if result.Error != nil {
		return users.User{}, result.Error
	}
	return userDB.ToUsecase(), nil
}

func (repo *UserRepository) GetUserDetail(id int, ctx context.Context) (users.User, error) {
	var userDb User

	result := repo.db.Preload("Keranjang").First(&userDb, id)

	//kalo ga ketemu idnya gimana?
	if result.Error != nil {
		return users.User{}, errors.New("User not found")
	}

	return userDb.ToUsecase(), nil
}

func (repo *UserRepository) EditUser(user users.User, id int, ctx context.Context) (users.User, error) {
	userDB := FromUsecase(user)
	var newUser User

	result := repo.db.Preload("Keranjang").First(&newUser, id)

	if result.Error != nil {
		return users.User{}, errors.New("User not found")
	}

	//ngecek kosong dan engga di mana?
	newUser.Email = userDB.Email
	newUser.Password = userDB.Password
	newUser.Name = userDB.Name
	newUser.PhoneNumber = userDB.PhoneNumber
	newUser.Street = userDB.Street
	newUser.PostalCode = userDB.PostalCode

	repo.db.Save(&newUser)
	return newUser.ToUsecase(), nil
}

//Deletenya cuma ngubah deleted at
func (repo *UserRepository) DeleteUser(id int, ctx context.Context) (users.User, error) {
	var userDb User

	resultFind := repo.db.First(&userDb, id)

	if resultFind.Error != nil {
		return users.User{}, errors.New("User not found")
	}

	result := repo.db.Delete(&userDb, id)

	//kalo ngecek ga ada id kayak gitu pake result kah?
	if result.Error != nil {
		return users.User{}, errors.New("User not found")
	}

	return userDb.ToUsecase(), nil
}
