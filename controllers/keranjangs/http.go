package controllers

import (
	"net/http"
	"strconv"

	"github.com/afrizal423/my-shirt/business/keranjangs"
	"github.com/afrizal423/my-shirt/controllers"
	"github.com/afrizal423/my-shirt/controllers/keranjangs/request"
	"github.com/afrizal423/my-shirt/controllers/keranjangs/response"

	"github.com/labstack/echo/v4"
)

type KeranjangController struct {
	usecase keranjangs.KeranjangUseCaseInterface
}

func NewKeranjangController(uc keranjangs.KeranjangUseCaseInterface) *KeranjangController {
	return &KeranjangController{
		usecase: uc,
	}
}

func (controller *KeranjangController) InsertKeranjang(c echo.Context) error {
	ctx := c.Request().Context()

	var keranjang request.KeranjangInsert

	err := c.Bind(&keranjang)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "Error binding", err)
	}

	keranjangRepo, errRepo := controller.usecase.InsertKeranjang(*keranjang.ToUsecase(), ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "Internal error", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(keranjangRepo))
}

func (controller *KeranjangController) GetAllKeranjang(c echo.Context) error {
	ctx := c.Request().Context()

	keranjang, errRepo := controller.usecase.GetAllKeranjang(ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "Internal error", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecaseList(keranjang))
}

func (controller *KeranjangController) GetKeranjangDetail(c echo.Context) error {
	ctx := c.Request().Context()

	keranjangId, _ := strconv.Atoi(c.Param("keranjangId"))

	keranjang, errRepo := controller.usecase.GetKeranjangDetail(keranjangId, ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "Internal error", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(keranjang))
}

func (controller *KeranjangController) EditKeranjang(c echo.Context) error {
	ctx := c.Request().Context()

	var keranjang request.KeranjangEdit
	keranjangId, _ := strconv.Atoi(c.Param("keranjangId"))
	err := c.Bind(&keranjang)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "Error binding", err)
	}

	keranjangRepo, errRepo := controller.usecase.EditKeranjang(*keranjang.ToUsecase(), keranjangId, ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "Internal error", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(keranjangRepo))
}

func (controller *KeranjangController) DeleteKeranjang(c echo.Context) error {
	ctx := c.Request().Context()

	keranjangId, _ := strconv.Atoi(c.Param("keranjangId"))

	keranjang, errRepo := controller.usecase.DeleteKeranjang(keranjangId, ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "Internal error", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(keranjang))
}
