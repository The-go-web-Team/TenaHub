package main

import (
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/delivery/http/handler"
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

	// errs := dbConn.CreateTable(&entity.User{}).GetErrors()

	// if len(errs) > 0 {
	// 	panic(errs)
	// }

	userRepo := repository.NewUserGormRepo(dbConn)
	userServ := service.NewUserService(userRepo)

	userHandl := handler.NewUserHander(userServ)

	router := httprouter.New()

	router.GET("/v1/users", userHandl.GetUsers)
	router.GET("/v1/users/:id", userHandl.GetSingleUser)
	router.POST("/v1/user", userHandl.GetUser)
	router.PUT("/v1/users/:id", userHandl.PutUser)
	router.POST("/v1/users", userHandl.PostUser)
	router.DELETE("/v1/users/:id", userHandl.DeleteUser)

	http.ListenAndServe(":8181", router)
}
