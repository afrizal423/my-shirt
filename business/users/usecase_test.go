package users_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/afrizal423/my-shirt/app/middleware"
	"github.com/afrizal423/my-shirt/business/users"
	"github.com/afrizal423/my-shirt/business/users/mocks"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// buat mock yang seolah-olah interface dari database
var userRepoInterfaceMock mocks.UserRepoInterface
var userUseCaseInterface users.UserUseCaseInterface
var userDataDummyLogin, userDataDummyEdit users.User
var userDataDummyGetAllUsers []users.User

func setup() {
	configJwt := middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}
	userUseCaseInterface = users.NewUseCase(&userRepoInterfaceMock, time.Hour*1, &configJwt)

	//data mock hasil login
	userDataDummyLogin = users.User{
		Id:       3,
		Name:     "Lee Mark",
		Email:    "leemark@gmail.com",
		Password: "1234",
		Token:    "",
	}

	userDataDummyGetAllUsers = []users.User{
		{
			Id:       1,
			Name:     "Lee Mark",
			Email:    "leemark@gmail.com",
			Password: "1234",
			Token:    "",
		},
		{
			Id:       2,
			Name:     "Jeong Jaehyun",
			Email:    "jeongjaehyun@gmail.com",
			Password: "1234",
			Token:    "",
		},
	}

	userDataDummyEdit = users.User{
		Id:          3,
		Name:        "Lee Mark",
		Email:       "leemark@gmail.com",
		Password:    "1234",
		PhoneNumber: "085723252648",
		Street:      "Jl Magelang 19",
		Address:     "Kota Magelang",
		PostalCode:  "56191",
	}
}

func TestLogin(t *testing.T) {
	setup()
	// t.Run("Success Login", func(t *testing.T) {
	// 	userRepoInterfaceMock.On("Login", mock.AnythingOfType("users.User"), mock.Anything).Return(userDataDummyLogin, nil).Once()

	// 	var requestLoginUser = users.User{
	// 		Email:    "leemark@gmail.com",
	// 		Password: "1234",
	// 	}

	// 	user, err := userUseCaseInterface.Login(requestLoginUser, context.Background())
	// 	assert.Equal(t, nil, err)
	// 	assert.Equal(t, userDataDummyLogin.Id, user.Id)
	// })

	t.Run("Email empty", func(t *testing.T) {
		userRepoInterfaceMock.On("Login", mock.AnythingOfType("users.User"), mock.Anything).Return(users.User{}, errors.New("Email empty")).Once()

		var requestLoginUser = users.User{
			Email:    "",
			Password: "123",
		}
		user, err := userUseCaseInterface.Login(requestLoginUser, context.Background())

		assert.Equal(t, errors.New("Email empty"), err)
		assert.Equal(t, users.User{}, user)
	})

	t.Run("Password empty", func(t *testing.T) {
		userRepoInterfaceMock.On("Login", mock.AnythingOfType("users.User"), mock.Anything).Return(users.User{}, errors.New("Password empty")).Once()

		var requestLoginUser = users.User{
			Email:    "leemark@gmail.com",
			Password: "",
		}
		user, err := userUseCaseInterface.Login(requestLoginUser, context.Background())

		assert.Equal(t, errors.New("Password empty"), err)
		assert.Equal(t, users.User{}, user)
	})

	t.Run("Error in database", func(t *testing.T) {
		userRepoInterfaceMock.On("Login", mock.AnythingOfType("users.User"), mock.Anything).Return(users.User{}, errors.New("User not found")).Once()

		var requestLoginUser = users.User{
			Email:    "leemark@gmail.com",
			Password: "1234",
		}

		user, err := userUseCaseInterface.Login(requestLoginUser, context.Background())
		assert.Error(t, err)
		assert.Equal(t, users.User{}, user)
	})
}

func TestGetAllUsers(t *testing.T) {
	setup()
	t.Run("Success Get All Users", func(t *testing.T) {
		userRepoInterfaceMock.On("GetAllUsers", mock.Anything, mock.Anything).Return(userDataDummyGetAllUsers, nil).Once()

		user, err := userUseCaseInterface.GetAllUsers(context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, userDataDummyGetAllUsers, user)
	})

	t.Run("Error in database", func(t *testing.T) {
		userRepoInterfaceMock.On("GetAllUsers", mock.Anything, mock.Anything).Return([]users.User{}, errors.New("Error in database")).Once()

		user, err := userUseCaseInterface.GetAllUsers(context.Background())

		assert.Error(t, err)
		assert.Equal(t, []users.User{}, user)
	})
}

func TestSignUp(t *testing.T) {
	setup()
	t.Run("Success Signup", func(t *testing.T) {
		userRepoInterfaceMock.On("SignUp", mock.AnythingOfType("users.User"), mock.Anything).Return(userDataDummyLogin, nil).Once()

		var requestSignUpUser = users.User{
			Email:    "leemark@gmail.com",
			Password: "1234",
		}

		user, err := userUseCaseInterface.SignUp(requestSignUpUser, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, userDataDummyLogin, user)
	})

	t.Run("Email empty", func(t *testing.T) {
		userRepoInterfaceMock.On("SignUp", mock.AnythingOfType("users.User"), mock.Anything).Return(users.User{}, errors.New("Email empty")).Once()

		var requestLoginUser = users.User{
			Email:    "",
			Password: "123",
		}
		user, err := userUseCaseInterface.SignUp(requestLoginUser, context.Background())

		assert.Equal(t, errors.New("Email empty"), err)
		assert.Equal(t, users.User{}, user)
	})

	t.Run("Password empty", func(t *testing.T) {
		userRepoInterfaceMock.On("SignUp", mock.AnythingOfType("users.User"), mock.Anything).Return(users.User{}, errors.New("Password empty")).Once()

		var requestLoginUser = users.User{
			Email:    "leemark@gmail.com",
			Password: "",
		}
		user, err := userUseCaseInterface.SignUp(requestLoginUser, context.Background())

		assert.Equal(t, errors.New("Password empty"), err)
		assert.Equal(t, users.User{}, user)
	})

	t.Run("Email already exist", func(t *testing.T) {
		userRepoInterfaceMock.On("SignUp", mock.AnythingOfType("users.User"), mock.Anything).Return(users.User{}, errors.New("Email already exist")).Once()

		var requestSignUpUser = users.User{
			Email:    "leemark@gmail.com",
			Password: "123",
		}
		user, err := userUseCaseInterface.SignUp(requestSignUpUser, context.Background())

		assert.Error(t, err)
		assert.Equal(t, users.User{}, user)
	})
}

func TestGetUserDetail(t *testing.T) {
	setup()
	t.Run("Success Get User Detail", func(t *testing.T) {
		userRepoInterfaceMock.On("GetUserDetail", mock.Anything, mock.Anything).Return(userDataDummyLogin, nil).Once()

		user, err := userUseCaseInterface.GetUserDetail(1, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, userDataDummyLogin, user)
	})

	t.Run("User ID empty", func(t *testing.T) {
		userRepoInterfaceMock.On("GetUserDetail", mock.Anything, mock.Anything).Return(users.User{}, errors.New("User ID empty")).Once()

		user, err := userUseCaseInterface.GetUserDetail(0, context.Background())

		assert.Equal(t, errors.New("User ID empty"), err)
		assert.Equal(t, users.User{}, user)
	})

	t.Run("Users not found in database", func(t *testing.T) {
		userRepoInterfaceMock.On("GetUserDetail", mock.Anything, mock.Anything).Return(users.User{}, errors.New("User not found")).Once()

		user, err := userUseCaseInterface.GetUserDetail(-1, context.Background())

		assert.Error(t, err)
		assert.Equal(t, users.User{}, user)
	})
}

func TestEditUser(t *testing.T) {
	setup()
	t.Run("Success Edit", func(t *testing.T) {
		userRepoInterfaceMock.On("EditUser", mock.AnythingOfType("users.User"), mock.Anything, mock.Anything).Return(userDataDummyEdit, nil).Once()

		var requestEditUser = users.User{
			Name:     "Lee Mark",
			Email:    "leemark@gmail.com",
			Password: "123",
		}

		user, err := userUseCaseInterface.EditUser(requestEditUser, 1, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, userDataDummyEdit, user)
	})

	t.Run("User ID empty", func(t *testing.T) {
		userRepoInterfaceMock.On("EditUser", mock.AnythingOfType("users.User"), mock.Anything, mock.Anything).Return(users.User{}, errors.New("User ID empty")).Once()

		var requestLoginUser = users.User{
			Email:    "leemark@gmail.com",
			Password: "123",
		}

		user, err := userUseCaseInterface.EditUser(requestLoginUser, 0, context.Background())

		assert.Equal(t, errors.New("User ID empty"), err)
		assert.Equal(t, users.User{}, user)
	})

	t.Run("Email empty", func(t *testing.T) {
		userRepoInterfaceMock.On("EditUser", mock.AnythingOfType("users.User"), mock.Anything, mock.Anything).Return(users.User{}, errors.New("Email empty")).Once()

		var requestLoginUser = users.User{
			Email:    "",
			Password: "123",
		}
		user, err := userUseCaseInterface.EditUser(requestLoginUser, 1, context.Background())

		assert.Equal(t, errors.New("Email empty"), err)
		assert.Equal(t, users.User{}, user)
	})

	t.Run("Password empty", func(t *testing.T) {
		userRepoInterfaceMock.On("EditUser", mock.AnythingOfType("users.User"), mock.Anything, mock.Anything).Return(users.User{}, errors.New("Password empty")).Once()

		var requestLoginUser = users.User{
			Email:    "leemark@gmail.com",
			Password: "",
		}
		user, err := userUseCaseInterface.EditUser(requestLoginUser, 1, context.Background())

		assert.Equal(t, errors.New("Password empty"), err)
		assert.Equal(t, users.User{}, user)
	})

	t.Run("User not found", func(t *testing.T) {
		userRepoInterfaceMock.On("EditUser", mock.AnythingOfType("users.User"), mock.Anything, mock.Anything).Return(users.User{}, errors.New("User not found")).Once()

		var requestEditUser = users.User{
			Email:    "leemark@gmail.com",
			Password: "123",
		}
		user, err := userUseCaseInterface.EditUser(requestEditUser, 1, context.Background())

		assert.Error(t, err)
		assert.Equal(t, users.User{}, user)
	})
}

func TestDeleteUser(t *testing.T) {
	setup()
	t.Run("Success delete", func(t *testing.T) {
		userRepoInterfaceMock.On("DeleteUser", mock.Anything, mock.Anything).Return(userDataDummyLogin, nil).Once()

		user, err := userUseCaseInterface.DeleteUser(1, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, userDataDummyLogin, user)
	})

	t.Run("User ID empty", func(t *testing.T) {
		userRepoInterfaceMock.On("DeleteUser", mock.Anything, mock.Anything).Return(users.User{}, errors.New("User ID empty")).Once()

		user, err := userUseCaseInterface.DeleteUser(0, context.Background())

		assert.Equal(t, errors.New("User ID empty"), err)
		assert.Equal(t, users.User{}, user)
	})

	t.Run("Users not found", func(t *testing.T) {
		userRepoInterfaceMock.On("DeleteUser", mock.Anything, mock.Anything).Return(users.User{}, errors.New("User not found")).Once()

		user, err := userUseCaseInterface.DeleteUser(-1, context.Background())

		assert.Error(t, err)
		assert.Equal(t, users.User{}, user)
	})
}
