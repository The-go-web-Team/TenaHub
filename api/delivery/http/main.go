package main

import (
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	adminRepo "github.com/TenaHub/api/admin/repository"
	adminServ "github.com/TenaHub/api/admin/service"
	agentRepo "github.com/TenaHub/api/agent/repository"
	agentServ "github.com/TenaHub/api/agent/service"
	healthCenterRepo "github.com/TenaHub/api/healthcenter/repository"
	healthCenterServ "github.com/TenaHub/api/healthcenter/service"
	userRepo "github.com/TenaHub/api/user/repository"
	userServ "github.com/TenaHub/api/user/service"
	"github.com/TenaHub/api/delivery/http/handler"
	"github.com/julienschmidt/httprouter"
)

func main()  {
	dbconn, err := gorm.Open("postgres", "postgres://postgres:0912345678@localhost/tenahub?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()
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

	//errs := dbconn.CreateTable(&entity.Agent{},&entity.Admin{}, &entity.HealthCenter{}, &entity.User{}).GetErrors()
	//if len(errs)> 0 {
	//	panic(errs)
	//}else {
	//	fmt.Println("something is occurred")
	//}


	router := httprouter.New()

	router.GET("/v1/admin/:id", adminHandler.GetAdmin)
	router.GET("/v1/agent/:id", agentHandler.GetSingleAgent)
	router.GET("/v1/agent", agentHandler.GetAgents)
	router.PUT("/v1/agent/:id", agentHandler.PutAgent)
	router.POST("/v1/agent", agentHandler.PostAgent)
	router.DELETE("/v1/agent/:id", agentHandler.DeleteAgent)
	router.GET("/v1/healthcenter/:id", healthCenterHandler.GetSingleHealthCenter)
	router.GET("/v1/healthcenters", healthCenterHandler.GetHealthCenter)
	router.DELETE("/v1/healthcenter/:id", healthCenterHandler.DeleteHealthCenter)
	router.GET("/v1/user/:id", userHandler.GetSingleUser)
	router.GET("/v1/user", userHandler.GetUsers)
	router.DELETE("/v1/user/:id", userHandler.DeleteUser)
	http.ListenAndServe(":8181", router)
}

