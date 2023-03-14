package main

import (
	"github.com/afrizal423/my-shirt/app/routes"

	userUsecase "github.com/afrizal423/my-shirt/business/users"
	userController "github.com/afrizal423/my-shirt/controllers/users"
	userRepo "github.com/afrizal423/my-shirt/drivers/databases/users"

	keranjangUsecase "github.com/afrizal423/my-shirt/business/keranjangs"
	keranjangController "github.com/afrizal423/my-shirt/controllers/keranjangs"
	keranjangRepo "github.com/afrizal423/my-shirt/drivers/databases/keranjangs"

	kemejaKeranjangUsecase "github.com/afrizal423/my-shirt/business/kemejakeranjangs"
	kemejaKeranjangController "github.com/afrizal423/my-shirt/controllers/kemejakeranjangs"
	kemejaKeranjangRepo "github.com/afrizal423/my-shirt/drivers/databases/kemejakeranjangs"

	kemejaUsecase "github.com/afrizal423/my-shirt/business/kemejas"
	kemejaController "github.com/afrizal423/my-shirt/controllers/kemejas"
	kemejaRepo "github.com/afrizal423/my-shirt/drivers/databases/kemejas"

	saleUsecase "github.com/afrizal423/my-shirt/business/sales"
	saleController "github.com/afrizal423/my-shirt/controllers/sales"
	saleRepo "github.com/afrizal423/my-shirt/drivers/databases/sales"

	"log"
	"time"

	_middleware "github.com/afrizal423/my-shirt/app/middleware"
	"github.com/afrizal423/my-shirt/drivers/databases/mysql"

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
