package main

import (
	// "fmt"
	"net/http"

	// "github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/entity"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/delivery/http/handler"

	hcserviceRepository "github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/hcservice/repository"
	hcserviceService "github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/hcservice/service"

	commentRepository "github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/comment/repository"
	commentService "github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/comment/service"

	ratingRepository "github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/rating/repository"
	ratingService "github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/rating/service"

	hcRepository "github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/healthcenter/repository"
	hcService "github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/healthcenter/service"

	sesRepository "github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/session/repository"
	sesService "github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/session/service"

	"github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/user/repository"

	"github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/user/service"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	// "fmt"
)

func main() {

	dbConn, err := gorm.Open("postgres", "postgres://postgres:password@localhost/tenahubdb?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbConn.Close()

	// errs := dbConn.CreateTable(&entity.Session{}).GetErrors()

	// fmt.Println(errs)

	// if len(errs) > 0 {
	// 	panic(errs)
	// }

	userRepo := repository.NewUserGormRepo(dbConn)
	userServ := service.NewUserService(userRepo)

	serviceRepo := hcserviceRepository.NewServiceGormRepo(dbConn)
	serviceServ := hcserviceService.NewHcserviceService(serviceRepo)

	comRepo := commentRepository.NewCommentGormRepo(dbConn)
	comServ := commentService.NewCommentService(comRepo)

	hcRepo := hcRepository.NewHcRepository(dbConn)
	hcServ := hcService.NewHcService(hcRepo)

	ratingRepo := ratingRepository.NewGormRatingRepository(dbConn)
	ratingServ := ratingService.NewHcRatingService(ratingRepo)

	sessionRepo := sesRepository.NewSessionGormRepo(dbConn)
	sessionService := sesService.NewSessionService(sessionRepo)

	userHandl := handler.NewUserHander(userServ)
	hcservHandl := handler.NewServiceHandler(serviceServ)
	cmtHandl := handler.NewCommentHandler(comServ)
	hcHandl := handler.NewHcHandler(hcServ)
	ratingHandl := handler.NewRatingHandler(ratingServ)
	sesHandl := handler.NewSessionHandler(sessionService)

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

	router.GET("/v1/comments/:id", cmtHandl.GetComments)
	router.GET("/v1/comment/:id", cmtHandl.GetComment)
	router.PUT("/v1/comments/:id", cmtHandl.PutComment)
	router.DELETE("/v1/comments/:id", cmtHandl.DeleteComment)
	router.POST("/v1/comments", cmtHandl.PostComment)
	router.POST("/v1/comments/check", cmtHandl.Check)

	router.GET("/v1/healthcenters", hcHandl.GetHealthcenters)
	router.GET("/v1/healthcenter/:id", hcHandl.GetHealthcenter)
	router.GET("/v1/healthcenters/top/:amount", hcHandl.GetTop)

	router.GET("/v1/rating/:id", ratingHandl.GetRating)
	router.POST("/v1/rating", ratingHandl.PostRating)

	router.GET("/v1/session", sesHandl.GetSession)
	router.POST("/v1/session", sesHandl.PostSession)
	router.DELETE("/v1/session/:uuid", sesHandl.DeleteSession)

	http.ListenAndServe(":8181", router)
}
