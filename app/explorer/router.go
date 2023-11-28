package explorer

import (
	"github.com/gin-gonic/gin"
)

type ExplorerRouter struct {
	basePath string
	handler  *ExplorerHandler
}

func NewExplorerRouter(handler *ExplorerHandler) *ExplorerRouter {
	return &ExplorerRouter{
		basePath: "/api/explorer/",
		handler:  handler,
	}
}

func (router *ExplorerRouter) Register(engine *gin.Engine) {
	group := engine.Group(router.basePath)

	group.GET("tree/", router.handler.HandleGetTree)

	group.POST("flows/", router.handler.HandleAddFlow)

	group.POST("events/", router.handler.HandleAddEvent)
}
