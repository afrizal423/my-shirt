package main

import (
	"kemejaku/app/routes"

	userUsecase "kemejaku/business/users"
	userController "kemejaku/controllers/users"
	userRepo "kemejaku/drivers/databases/users"

	keranjangUsecase "kemejaku/business/keranjangs"
	keranjangController "kemejaku/controllers/keranjangs"
	keranjangRepo "kemejaku/drivers/databases/keranjangs"

	kemejaKeranjangUsecase "kemejaku/business/kemejakeranjangs"
	kemejaKeranjangController "kemejaku/controllers/kemejakeranjangs"
	kemejaKeranjangRepo "kemejaku/drivers/databases/kemejakeranjangs"

	kemejaUsecase "kemejaku/business/kemejas"
	kemejaController "kemejaku/controllers/kemejas"
	kemejaRepo "kemejaku/drivers/databases/kemejas"

	saleUsecase "kemejaku/business/sales"
	saleController "kemejaku/controllers/sales"
	saleRepo "kemejaku/drivers/databases/sales"

	_middleware "kemejaku/app/middleware"
	"kemejaku/drivers/databases/mysql"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	//viper digunakan untuk membaca file config.json
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service Run and Debug mode")
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&userRepo.User{},
		&keranjangRepo.Keranjang{},
		&kemejaKeranjangRepo.KemejaKeranjang{},
		&saleRepo.Sale{},
		&kemejaRepo.Kemeja{},
	)
}

func main() {
	configDb := mysql.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	db := configDb.InitialDB()
	dbMigrate(db)

	jwt := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	e := echo.New()

	userRepoInterface := userRepo.NewUserRepository(db)
	userUseCaseInterface := userUsecase.NewUseCase(userRepoInterface, timeoutContext, &jwt)
	userControllerInterface := userController.NewUserController(userUseCaseInterface)

	kemejaKeranjangRepoInterface := kemejaKeranjangRepo.NewKemejaKeranjangRepo(db)
	kemejaKeranjangUseCaseInterface := kemejaKeranjangUsecase.NewKemejaKeranjangUsecase(kemejaKeranjangRepoInterface, timeoutContext)
	kemejaKeranjangControllerInterface := kemejaKeranjangController.NewKemejaKeranjangController(kemejaKeranjangUseCaseInterface)

	keranjangRepoInterface := keranjangRepo.NewKeranjangRepo(db)
	keranjangUseCaseInterface := keranjangUsecase.NewKeranjangUcecase(keranjangRepoInterface, timeoutContext)
	keranjangControllerInterface := keranjangController.NewKeranjangController(keranjangUseCaseInterface)

	saleRepoInterface := saleRepo.NewSaleRepo(db)
	saleUseCaseInterface := saleUsecase.NewSaleUsecase(saleRepoInterface, timeoutContext)
	saleControllerInterface := saleController.NewSaleController(saleUseCaseInterface)

	kemejaRepoInterface := kemejaRepo.NewKemejaRepo(db)
	kemejaUseCaseInterface := kemejaUsecase.NewKemejaUsecase(kemejaRepoInterface, timeoutContext)
	kemejaControllerInterface := kemejaController.NewKemejaController(kemejaUseCaseInterface)

	routesInit := routes.RouteControllerList{
		UserController:            *userControllerInterface,
		KeranjangController:       *keranjangControllerInterface,
		SaleController:            *saleControllerInterface,
		KemejaController:          *kemejaControllerInterface,
		KemejaKeranjangController: *kemejaKeranjangControllerInterface,
		JWTConfig:                 &jwt,
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
