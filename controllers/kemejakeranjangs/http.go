package controllers

import (
	"kemejaku/business/kemejakeranjangs"
	"kemejaku/controllers"
	"kemejaku/controllers/kemejakeranjangs/request"
	"kemejaku/controllers/kemejakeranjangs/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type KemejaKeranjangController struct {
	usecase kemejakeranjangs.KemejaKeranjangUseCaseInterface
}

func NewKemejaKeranjangController(uc kemejakeranjangs.KemejaKeranjangUseCaseInterface) *KemejaKeranjangController {
	return &KemejaKeranjangController{
		usecase: uc,
	}
}

func (controller *KemejaKeranjangController) InsertKemejaKeranjang(c echo.Context) error {
	ctx := c.Request().Context()

	var kk request.KemejaKeranjangInsert

	err := c.Bind(&kk)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "Error binding", err)
	}

	user, errRepo := controller.usecase.InsertKemejaKeranjang(*kk.ToUsecase(), ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "Failed to insert new kemeja keranjang", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(user))
}

func (controller *KemejaKeranjangController) GetAllKemejaKeranjang(c echo.Context) error {
	ctx := c.Request().Context()

	kk, errRepo := controller.usecase.GetAllKemejaKeranjang(ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "There is no kemeja keranjang column", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecaseList(kk))
}

func (controller *KemejaKeranjangController) GetKemejaKeranjangDetail(c echo.Context) error {
	ctx := c.Request().Context()

	kemejaKeranjangId, _ := strconv.Atoi(c.Param("kemejaKeranjangId"))

	kk, errRepo := controller.usecase.GetKemejaKeranjangDetail(kemejaKeranjangId, ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "Kemeja keranjang not found", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(kk))
}

func (controller *KemejaKeranjangController) EditKemejaKeranjang(c echo.Context) error {
	ctx := c.Request().Context()

	var kk request.KemejaKeranjangEdit
	kemejaKeranjangId, _ := strconv.Atoi(c.Param("kemejaKeranjangId"))
	err := c.Bind(&kk)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "Error binding", err)
	}

	kkRepo, errRepo := controller.usecase.EditKemejaKeranjang(*kk.ToUsecase(), kemejaKeranjangId, ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "Kemeja keranjang not found", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(kkRepo))
}

func (controller *KemejaKeranjangController) DeleteKemejaKeranjang(c echo.Context) error {
	ctx := c.Request().Context()

	kemejaKeranjangId, _ := strconv.Atoi(c.Param("kemejaKeranjangId"))

	kk, errRepo := controller.usecase.DeleteKemejaKeranjang(kemejaKeranjangId, ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "Kemeja keranjang not found", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(kk))
}
