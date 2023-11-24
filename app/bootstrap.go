package app

import (
	"assalielmehdi/eventify/app/config"
	"assalielmehdi/eventify/app/explorer"
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
	eventRepository := repositories.NewEventRepository(db)

	flowService := services.NewFlowService(flowRepository)
	eventService := services.NewEventService(eventRepository)

	flowHandler := handlers.NewFlowHandler(flowService)
	eventHandler := handlers.NewEventHandler(eventService)

	flowRouter := routers.NewFlowRouter(flowHandler)
	eventRouter := routers.NewEventRouter(eventHandler)

	explorerService := explorer.NewExplorerService(db)
	explorerHandler := explorer.NewExplorerHandler(explorerService)
	exploterRouter := explorer.NewExplorerRouter(explorerHandler)

	// 2-Deps

	// Run
	flowRouter.Register(server)
	eventRouter.Register(server)
	exploterRouter.Register(server)

	server.Serve()
}
