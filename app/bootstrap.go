package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"assalielmehdi/eventify/app/config"
	"assalielmehdi/eventify/app/explorer"
	"assalielmehdi/eventify/app/graph"
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

	graphService := graph.NewGraphService(db)
	graphHandler := graph.NewGraphHandler(graphService)
	graphRouter := graph.NewGraphRouter(graphHandler)

	// 2-Deps

	// Run

	flowRouter.Register(server)
	eventRouter.Register(server)
	exploterRouter.Register(server)
	graphRouter.Register(server)

	server.Router.Static("/static/", "static/")
	server.Router.LoadHTMLGlob("templates/*")
	server.Router.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})

	server.Serve()
}
