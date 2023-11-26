package main

import (
	"assalielmehdi/eventify/app"
	"assalielmehdi/eventify/app/config"
	"assalielmehdi/eventify/app/explorer"
	"assalielmehdi/eventify/app/graph"
	"assalielmehdi/eventify/app/handlers"
	"assalielmehdi/eventify/app/repositories"
	"assalielmehdi/eventify/app/routers"
	"assalielmehdi/eventify/app/services"
)

func main() {
	// Config
	config.Init()
	serverConfig := config.GetEnvServerConfig()
	dbConfig := config.GetEnvDBConfig()

	// 0-Deps
	server := app.NewServer(serverConfig)
	db := app.NewDB(dbConfig)

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

	graphService := graph.NewGraphService(db)
	graphHandler := graph.NewGraphHandler(graphService)
	graphRouter := graph.NewGraphRouter(graphHandler)

	// 2-Deps

	// Run
	server.Register(flowRouter)
	server.Register(eventRouter)
	server.Register(exploterRouter)
	server.Register(graphRouter)

	server.Serve()
}
