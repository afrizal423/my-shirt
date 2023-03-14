package controllers

import (
	"net/http"
	"strconv"

	"github.com/afrizal423/my-shirt/app/middleware"
	"github.com/afrizal423/my-shirt/business/users"
	"github.com/afrizal423/my-shirt/controllers"
	"github.com/afrizal423/my-shirt/controllers/users/request"
	"github.com/afrizal423/my-shirt/controllers/users/response"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	usecase users.UserUseCaseInterface
}

// dipasangkan dengan routing
func NewUserController(uc users.UserUseCaseInterface) *UserController {
	return &UserController{
		usecase: uc,
	}
}

func (controller *UserController) Login(c echo.Context) error {
	ctx := c.Request().Context()

	//dari request
	var userLogin request.UserLogin

	err := c.Bind(&userLogin)

	userId := middleware.GetUserId(c)
	userLogin.ToUsecase().Id = userId

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "Error binding", err)
	}

	user, errRepo := controller.usecase.Login(*userLogin.ToUsecase(), ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(user))
}

func (controller *UserController) GetAllUsers(c echo.Context) error {
	ctx := c.Request().Context()

	user, errRepo := controller.usecase.GetAllUsers(ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "There is no user column", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecaseList(user))
}

func (controller *UserController) SignUp(c echo.Context) error {
	ctx := c.Request().Context()

	var userSignup request.UserLogin

	err := c.Bind(&userSignup)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "Error binding", err)
	}

	user, errRepo := controller.usecase.SignUp(*userSignup.ToUsecase(), ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "Error in repo", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(user))
}

func (controller *UserController) GetUserDetail(c echo.Context) error {
	ctx := c.Request().Context()

	userId, _ := strconv.Atoi(c.Param("userId"))

	user, errRepo := controller.usecase.GetUserDetail(userId, ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "User not found", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(user))
}

func (controller *UserController) EditUser(c echo.Context) error {
	ctx := c.Request().Context()

	var userEdit request.UserEdit
	userId, _ := strconv.Atoi(c.Param("userId"))

	err := c.Bind(&userEdit)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "Error binding", err)
	}

	user, errRepo := controller.usecase.EditUser(*userEdit.ToUsecase(), userId, ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "User not found", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(user))
}

func (controller *UserController) DeleteUser(c echo.Context) error {
	ctx := c.Request().Context()

	userId, _ := strconv.Atoi(c.Param("userId"))

	user, errRepo := controller.usecase.DeleteUser(userId, ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "User not found", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(user))
}
