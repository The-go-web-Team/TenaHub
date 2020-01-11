package main

import (
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"


	serviceRepo "github.com/TenaHub/api/service/repository"
	serviceServ "github.com/TenaHub/api/service/service"
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
	"github.com/TenaHub/api/delivery/http/handler"
	"github.com/julienschmidt/httprouter"
)

func main()  {
	dbconn, err := gorm.Open("postgres", "postgres://postgres:0912345678@localhost/tenahub?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

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
	serviceRepository := serviceRepo.NewServiceGormRepo(dbconn)
	serviceService := serviceServ.NewServiceService(serviceRepository)
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
	router.POST("/v1/agent", agentHandler.GetAgent)
	router.POST("/v1/agents", agentHandler.PostAgent)
	router.OPTIONS("/v1/agent", agentHandler.PostAgent)
	router.DELETE("/v1/agent/:id", agentHandler.DeleteAgent)

	router.GET("/v1/healthcenter/:id", healthCenterHandler.GetSingleHealthCenter)
	router.GET("/v1/healthcenter/:id/agent", healthCenterHandler.GetHealthCentersByAgentId)
	router.POST("/v1/healthcenter", healthCenterHandler.GetHealthCenter)
	router.POST("/v1/healthcenters", healthCenterHandler.PostHealthCenter)
	router.PUT("/v1/healthcenter/:id", healthCenterHandler.PutHealthCenter)
	router.GET("/v1/healthcenter", healthCenterHandler.GetHealthCenter)
	router.GET("/v1/healthcenters", healthCenterHandler.GetHealthCenters)
	router.DELETE("/v1/healthcenter/:id", healthCenterHandler.DeleteHealthCenter)

	router.GET("/v1/user/:id", userHandler.GetSingleUser)
	router.GET("/v1/user", userHandler.GetUsers)
	router.DELETE("/v1/user/:id", userHandler.DeleteUser)

	//router.GET("/v1/service/:id", serviceHandler.GetSingleService)
	router.GET("/v1/service/:id", serviceHandler.GetServices)
	router.GET("/v1/services/pending/:id", serviceHandler.GetPendingServices)
	router.PUT("/v1/service/:id", serviceHandler.PutService)
	router.POST("/v1/services", serviceHandler.PostService)
	router.OPTIONS("/v1/services", serviceHandler.PostService)
	router.DELETE("/v1/service/:id", serviceHandler.DeleteService)

	//router.GET("/v1/feedback/:id", feedBackHandler.GetComment)
	router.GET("/v1/feedback/:id", feedBackHandler.GetComments)
	router.PUT("/v1/feedback/:id", feedBackHandler.PutComment)
	router.POST("/v1/feedback", feedBackHandler.PostComment)
	router.OPTIONS("/v1/feedback", feedBackHandler.PostComment)
	router.DELETE("/v1/feedback/:id", feedBackHandler.DeleteComment)


	http.ListenAndServe(":8181", router)
}

