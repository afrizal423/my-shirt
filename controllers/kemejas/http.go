package controllers

import (
	"net/http"
	"strconv"

	"github.com/afrizal423/my-shirt/business/kemejas"
	"github.com/afrizal423/my-shirt/controllers"
	"github.com/afrizal423/my-shirt/controllers/kemejas/request"
	"github.com/afrizal423/my-shirt/controllers/kemejas/response"

	"github.com/labstack/echo/v4"
)

type KemejaController struct {
	usecase kemejas.KemejaUseCaseInterface
}

func NewKemejaController(uc kemejas.KemejaUseCaseInterface) *KemejaController {
	return &KemejaController{
		usecase: uc,
	}
}

func (controller *KemejaController) InsertKemeja(c echo.Context) error {
	ctx := c.Request().Context()

	var kemeja request.Kemeja

	err := c.Bind(&kemeja)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "Error binding", err)
	}

	kemejaRepo, errRepo := controller.usecase.InsertKemeja(*kemeja.ToUsecase(), ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "Failed to insert new kemeja", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(kemejaRepo))
}

func (controller *KemejaController) GetAllKemeja(c echo.Context) error {
	ctx := c.Request().Context()

	kemeja, errRepo := controller.usecase.GetAllKemeja(ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "There is no kemeja column", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecaseList(kemeja))
}

func (controller *KemejaController) GetKemejaDetail(c echo.Context) error {
	ctx := c.Request().Context()

	kemejaId, _ := strconv.Atoi(c.Param("kemejaId"))

	kemeja, errRepo := controller.usecase.GetKemejaDetail(kemejaId, ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "Kemeja not found", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(kemeja))
}

func (controller *KemejaController) EditKemeja(c echo.Context) error {
	ctx := c.Request().Context()

	var kemeja request.Kemeja
	kemejaId, _ := strconv.Atoi(c.Param("kemejaId"))
	err := c.Bind(&kemeja)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "Error binding", err)
	}

	kemejaRepo, errRepo := controller.usecase.EditKemeja(*kemeja.ToUsecase(), kemejaId, ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "Kemeja not found", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(kemejaRepo))
}

func (controller *KemejaController) DeleteKemeja(c echo.Context) error {
	ctx := c.Request().Context()

	kemejaId, _ := strconv.Atoi(c.Param("kemejaId"))

	if kemejaId == 0 {
		return controllers.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Kemeja ID empty")
	}

	kemeja, errRepo := controller.usecase.DeleteKemeja(kemejaId, ctx)

	if errRepo != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "Kemeja not found", errRepo)
	}

	return controllers.SuccessResponse(c, response.FromUsecase(kemeja))
}
