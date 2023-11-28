package graph

import (
	"github.com/gin-gonic/gin"
)

type GraphRouter struct {
	basePath string
	handler  *GraphHandler
}

func NewGraphRouter(handler *GraphHandler) *GraphRouter {
	return &GraphRouter{
		basePath: "/api/graph/",
		handler:  handler,
	}
}

func (router *GraphRouter) Register(engine *gin.Engine) {
	group := engine.Group(router.basePath)

	group.GET("flows/:flowId/", router.handler.HandleGetFlowGraph)

	group.PATCH("events/position/", router.handler.HandleUpdateEventPosition)
}
