package routes

import (
	kemejaKeranjangController "kemejaku/controllers/kemejakeranjangs"
	kemejaController "kemejaku/controllers/kemejas"
	keranjangController "kemejaku/controllers/keranjangs"
	saleController "kemejaku/controllers/sales"
	userController "kemejaku/controllers/users"

	_middleware "kemejaku/app/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouteControllerList struct {
	UserController            userController.UserController
	KeranjangController       keranjangController.KeranjangController
	KemejaController          kemejaController.KemejaController
	KemejaKeranjangController kemejaKeranjangController.KemejaKeranjangController
	SaleController            saleController.SaleController
	JWTConfig                 *_middleware.ConfigJWT
}

func (controller RouteControllerList) RouteRegister(e *echo.Echo) {

	eUser := e.Group("/user")
	eUser.Use(middleware.JWTWithConfig(controller.JWTConfig.Init()))
	e.POST("/login", controller.UserController.Login)
	e.POST("/signup", controller.UserController.SignUp)
	// e.POST("/logout", users.LogoutContorller)
	eUser.PUT("/:userId", controller.UserController.EditUser)
	eUser.DELETE("/:userId", controller.UserController.DeleteUser)
	eUser.GET("/:userId", controller.UserController.GetUserDetail)
	e.GET("/users", controller.UserController.GetAllUsers)
	// users.GET("/", controller.UserController.Login, middleware.JWTWithConfig(controller.JWTConfig))

	eKeranjang := e.Group("/keranjang")

	eKeranjang.GET("/:keranjangId", controller.KeranjangController.GetKeranjangDetail)
	e.GET("/keranjangs", controller.KeranjangController.GetAllKeranjang)
	e.POST("/keranjangs", controller.KeranjangController.InsertKeranjang)
	eKeranjang.PUT("/:keranjangId", controller.KeranjangController.EditKeranjang)
	eKeranjang.DELETE("/:keranjangId", controller.KeranjangController.DeleteKeranjang)

	eKemejaKeranjang := e.Group("/kemejakeranjang")

	eKemejaKeranjang.GET("/:kemejaKeranjangId", controller.KemejaKeranjangController.GetKemejaKeranjangDetail)
	e.GET("/kemejakeranjangs", controller.KemejaKeranjangController.GetAllKemejaKeranjang)
	e.POST("/kemejakeranjangs", controller.KemejaKeranjangController.InsertKemejaKeranjang)
	eKemejaKeranjang.PUT("/:kemejaKeranjangId", controller.KemejaKeranjangController.EditKemejaKeranjang)
	eKemejaKeranjang.DELETE("/:kemejaKeranjangId", controller.KemejaKeranjangController.DeleteKemejaKeranjang)

	eKemeja := e.Group("/kemeja")

	eKemeja.GET("/:kemejaId", controller.KemejaController.GetKemejaDetail)
	e.GET("/kemejas", controller.KemejaController.GetAllKemeja)
	e.POST("/kemejas", controller.KemejaController.InsertKemeja)
	eKemeja.PUT("/:kemejaId", controller.KemejaController.EditKemeja)
	eKemeja.DELETE("/:kemejaId", controller.KemejaController.DeleteKemeja)

	eSale := e.Group("/sale")

	eSale.GET("/:saleId", controller.SaleController.GetSaleDetail)
	e.GET("/sales", controller.SaleController.GetAllSale)
	e.POST("/sales", controller.SaleController.InsertSale)
	eSale.PUT("/:saleId", controller.SaleController.EditSale)
	eSale.DELETE("/:saleId", controller.SaleController.DeleteSale)
}
