package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Status   int         `json:"status"`
	Message  string      `json:"message"`
	Messages []string    `json:"messages"`
	Data     interface{} `json:"data"`
}

type BaseWithoutMessages struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ErrorResponseWithoutMessages(c echo.Context, status int, message string) error {
	response := BaseWithoutMessages{}
	response.Status = status
	response.Message = message
	return c.JSON(status, response)
}

func SuccessResponse(c echo.Context, data interface{}) error {
	response := BaseResponse{}
	response.Status = http.StatusOK
	response.Message = "success"
	response.Data = data
	return c.JSON(http.StatusOK, response)
}

func ErrorResponse(c echo.Context, status int, message string, errs error) error {
	response := BaseResponse{}
	response.Status = status
	response.Message = message
	response.Messages = []string{errs.Error()}
	return c.JSON(status, response)
}
