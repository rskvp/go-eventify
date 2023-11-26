package routers

import (
	"assalielmehdi/eventify/app/handlers"

	"github.com/gin-gonic/gin"
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

func (router *FlowRouter) Register(engine *gin.Engine) {
	group := engine.Group(router.BasePath)

	group.GET("/", router.Handler.HandleGetAll)

	group.GET("/:id", router.Handler.HandleGetOneById)

	group.POST("/", router.Handler.HandleAddOne)

	group.PUT("/", router.Handler.HandleUpdateOneById)

	group.DELETE("/:id", router.Handler.HandleDeleteOneById)

	group.POST("/run/:id", router.Handler.HandleRunById)
}
