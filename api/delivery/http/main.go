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

	// serviceRepo "github.com/TenaHub/api/service/repository"
	// serviceServ "github.com/TenaHub/api/service/service"
	adminRepo "github.com/TenaHub/api/admin/repository"
	adminServ "github.com/TenaHub/api/admin/service"
	agentRepo "github.com/TenaHub/api/agent/repository"
	agentServ "github.com/TenaHub/api/agent/service"
	healthCenterRepo "github.com/TenaHub/api/healthcenter/repository"
	healthCenterServ "github.com/TenaHub/api/healthcenter/service"
	userRepo "github.com/TenaHub/api/user/repository"
	userServ "github.com/TenaHub/api/user/service"
	feedBackRepo "github.com/TenaHub/api/comment/repository"
	feedBackServ "github.com/TenaHub/api/comment/service"

	"github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/user/repository"

	"github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/user/service"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	// "fmt"
)


func main()  {
	// dbconn, err := gorm.Open("postgres", "postgres://postgres:0912345678@localhost/tenahub?sslmode=disable")
	dbconn, err := gorm.Open("postgres", "postgres://postgres:password@localhost/tenahubdb?sslmode=disable")
	
	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	// errs := dbconn.CreateTable(&entity.Session{}).GetErrors()

	// fmt.Println(errs)

	// if len(errs) > 0 {
	// 	panic(errs)
	// }

	userRepo := repository.NewUserGormRepo(dbconn)
	userServ := service.NewUserService(userRepo)

	serviceRepo := hcserviceRepository.NewServiceGormRepo(dbconn)
	serviceServ := hcserviceService.NewHcserviceService(serviceRepo)

	comRepo := commentRepository.NewCommentGormRepo(dbconn)
	comServ := commentService.NewCommentService(comRepo)

	hcRepo := hcRepository.NewHcRepository(dbconn)
	hcServ := hcService.NewHcService(hcRepo)

	ratingRepo := ratingRepository.NewGormRatingRepository(dbconn)
	ratingServ := ratingService.NewHcRatingService(ratingRepo)

	sessionRepo := sesRepository.NewSessionGormRepo(dbconn)
	sessionService := sesService.NewSessionService(sessionRepo)

	userHandl := handler.NewUserHander(userServ)
	hcservHandl := handler.NewServiceHandler(serviceServ)
	cmtHandl := handler.NewCommentHandler(comServ)
	hcHandl := handler.NewHcHandler(hcServ)
	ratingHandl := handler.NewRatingHandler(ratingServ)
	sesHandl := handler.NewSessionHandler(sessionService)


	// defer dbconn.Close()

	//errs := dbconn.CreateTable(&entity.HealthCenter{}).GetErrors()
	//if len(errs)> 0 {
	//	panic(errs)
	//}else {
	//	fmt.Println("something is occurred")
	//}


	adminRespository := adminRepo.NewAdminGormRepo(dbconn)
	adminService := adminServ.NewAdminService(adminRespository)
	adminHandler := handler.NewAdminHandler(adminService)

	////////////

	agentRespository := agentRepo.NewAgentGormRepo(dbconn)
	agentService := agentServ.NewAgentService(agentRespository)
	agentHandler := handler.NewAgentHandler(agentService)

	//////////////
	userRespository := userRepo.NewUserGormRepo(dbconn)
	userService := userServ.NewUserService(userRespository)
	userHandler := handler.NewUserHandler(userService)

	/////////////////

	healthCenterRespository := healthCenterRepo.NewHealthCenterGormRepo(dbconn)
	healthCenterService := healthCenterServ.NewHealthCenterService(healthCenterRespository)
	healthCenterHandler := handler.NewHealthCenterHandler(healthCenterService)


	/////////
	serviceRepository := hcserviceRepository.NewServiceGormRepo(dbconn)
	serviceService := hcserviceService.NewServiceService(serviceRepository)
	serviceHandler := handler.NewServiceHandler(serviceService)
	////////


	feedBackRepository := feedBackRepo.NewCommentGormRepo(dbconn)
	feedBackService := feedBackServ.NewCommentService(feedBackRepository)
	feedBackHandler := handler.NewCommentHandler(feedBackService)




	router := httprouter.New()

	router.GET("/v1/admin/:id", adminHandler.GetSingleAdmin)
	router.POST("/v1/admin", adminHandler.GetAdmin)
	router.PUT("/v1/admin/:id", adminHandler.PutAdmin)
	router.GET("/v1/agent/:id", agentHandler.GetSingleAgent)

	router.GET("/v1/agent", agentHandler.GetAgents)
	router.PUT("/v1/agent/:id", agentHandler.PutAgent)
	router.POST("/v1/agent", agentHandler.PostAgent)
	router.OPTIONS("/v1/agent", agentHandler.PostAgent)
	router.DELETE("/v1/agent/:id", agentHandler.DeleteAgent)

	// router.GET("/v1/healthcenter/:id", healthCenterHandler.GetSingleHealthCenter)
	// router.POST("/v1/healthcenter", healthCenterHandler.GetHealthCenter)
	// router.PUT("/v1/healthcenter/:id", healthCenterHandler.PutHealthCenter)
	// router.GET("/v1/healthcenter", healthCenterHandler.GetHealthCenter)
	// router.GET("/v1/healthcenters", healthCenterHandler.GetHealthCenters)
	// router.DELETE("/v1/healthcenter/:id", healthCenterHandler.DeleteHealthCenter)

	router.GET("/v1/user/:id", userHandler.GetSingleUser)
	router.GET("/v1/user", userHandler.GetUsers)
	router.DELETE("/v1/user/:id", userHandler.DeleteUser)

	//router.GET("/v1/service/:id", serviceHandler.GetSingleService)
	router.GET("/v1/service/:id", serviceHandler.GetServices)
	router.GET("/v1/pendingservice", serviceHandler.GetPendingServices)
	router.PUT("/v1/service/:id", serviceHandler.PutService)
	router.POST("/v1/service", serviceHandler.PostService)
	router.OPTIONS("/v1/service", serviceHandler.PostService)
	router.DELETE("/v1/service/:id", serviceHandler.DeleteService)

	//router.GET("/v1/feedback/:id", feedBackHandler.GetComment)
	router.GET("/v1/feedback/:id", feedBackHandler.GetComments)
	router.PUT("/v1/feedback/:id", feedBackHandler.PutComment)
	router.POST("/v1/feedback", feedBackHandler.PostComment)
	router.OPTIONS("/v1/feedback", feedBackHandler.PostComment)
	router.DELETE("/v1/feedback/:id", feedBackHandler.DeleteComment)

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

	// router.GET("/v1/healthcenter/:id", healthCenterHandler.GetSingleHealthCenter)
	router.POST("/v1/healthcenter", healthCenterHandler.GetHealthCenter)
	router.PUT("/v1/healthcenter/:id", healthCenterHandler.PutHealthCenter)
	router.GET("/v1/healthcenter", healthCenterHandler.GetHealthCenter)
	router.GET("/v1/healthcenters", healthCenterHandler.GetHealthCenters)
	router.DELETE("/v1/healthcenter/:id", healthCenterHandler.DeleteHealthCenter)

	router.GET("/v1/healthcenters/search", hcHandl.GetHealthcenters)
	router.GET("/v1/healthcenter/:id", hcHandl.GetHealthcenter)
	router.GET("/v1/healthcenters/top/:amount", hcHandl.GetTop)

	router.GET("/v1/rating/:id", ratingHandl.GetRating)
	router.POST("/v1/rating", ratingHandl.PostRating)

	router.GET("/v1/session", sesHandl.GetSession)
	router.POST("/v1/session", sesHandl.PostSession)
	router.DELETE("/v1/session/:uuid", sesHandl.DeleteSession)


	http.ListenAndServe(":8181", router)
}

