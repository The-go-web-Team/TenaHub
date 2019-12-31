package main

import (
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/delivery/http/handler"
	// "github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/entity"
	hcserviceRepository "github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/hcservice/repository"
	hcserviceService "github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/hcservice/service"
	"github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/user/repository"

	"github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/user/service"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
)

func main() {

	dbConn, err := gorm.Open("postgres", "postgres://postgres:password@localhost/tenahubdb?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbConn.Close()

	// errs := dbConn.CreateTable(&entity.Service{}).GetErrors()

	// if len(errs) > 0 {
	// 	panic(errs)
	// }

	userRepo := repository.NewUserGormRepo(dbConn)
	userServ := service.NewUserService(userRepo)

	serviceRepo := hcserviceRepository.NewServiceGormRepo(dbConn)
	serviceServ := hcserviceService.NewHcserviceService(serviceRepo)

	userHandl := handler.NewUserHander(userServ)
	hcservHandl := handler.NewServiceHandler(serviceServ)

	router := httprouter.New()

	router.GET("/v1/users", userHandl.GetUsers)
	router.GET("/v1/users/:id", userHandl.GetSingleUser)
	router.POST("/v1/user", userHandl.GetUser)
	router.PUT("/v1/users/:id", userHandl.PutUser)
	router.POST("/v1/users", userHandl.PostUser)
	router.DELETE("/v1/users/:id", userHandl.DeleteUser)

	router.GET("/v1/services/:id", hcservHandl.GetServices)
	router.GET("/v1/service/:id", hcservHandl.GetSingleService)
	router.PUT("/v1/services/:id", hcservHandl.PutService)
	router.DELETE("/v1/services/:id", hcservHandl.DeleteService)
	router.POST("/v1/services", hcservHandl.PostService)

	http.ListenAndServe(":8181", router)
}
