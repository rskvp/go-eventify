package app

import (
	"assalielmehdi/eventify/app/config"
	"assalielmehdi/eventify/app/handlers"
	"assalielmehdi/eventify/app/repositories"
	"assalielmehdi/eventify/app/routers"
	"assalielmehdi/eventify/app/services"
)

func Bootstrap() {
	// Init
	config.Init()

	// 0-Deps
	server := routers.NewServer()
	db := repositories.NewDB(repositories.DBTypeSqlite)

	// 1-Deps
	flowRepository := repositories.NewFlowRepository(db)

	flowService := services.NewFlowService(flowRepository)

	flowHandler := handlers.NewFlowHandler(flowService)

	flowRouter := routers.NewFlowRouter(flowHandler)

	// 2-Deps

	// Run
	flowRouter.Register(server)

	server.Serve()
}
