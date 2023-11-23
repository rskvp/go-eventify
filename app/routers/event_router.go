package routers

import (
	"assalielmehdi/eventify/app/handlers"
)

type EventRouter struct {
	BasePath string
	Handler  *handlers.EventHandler
}

func NewEventRouter(handler *handlers.EventHandler) *EventRouter {
	return &EventRouter{
		BasePath: "/api/events",
		Handler:  handler,
	}
}

func (router *EventRouter) Register(server *Server) {
	group := server.Router.Group(router.BasePath)

	// group.GET("/", router.Handler.HandleGetAll)

	// group.GET("/:id", router.Handler.HandleGetOneById)

	group.POST("/", router.Handler.HandleAddOne)

	// group.PUT("/", router.Handler.HandleUpdateOneById)

	// group.DELETE("/:id", router.Handler.HandleDeleteOneById)
}
