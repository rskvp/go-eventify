package routers

import (
	"assalielmehdi/eventify/app/handlers"
)

type FlowRouter struct {
	BasePath string
	Handler  *handlers.FlowHandler
}

func NewFlowRouter(handler *handlers.FlowHandler) *FlowRouter {
	return &FlowRouter{
		BasePath: "/api/flows",
		Handler:  handler,
	}
}

func (router *FlowRouter) Register(server *Server) {
	group := server.Router.Group(router.BasePath)

	group.GET("/", router.Handler.HandleGetAll)

	group.GET("/:id", router.Handler.HandleGetOneById)

	group.POST("/", router.Handler.HandleAddOne)

	group.PUT("/:id", router.Handler.HandleUpdateOneById)

	group.DELETE("/:id", router.Handler.HandleDeleteOneById)

	group.POST("/run/:id", router.Handler.HandleRunById)
}
