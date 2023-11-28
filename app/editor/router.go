package editor

import (
	"github.com/gin-gonic/gin"
)

type EditorRouter struct {
	basePath string
	handler  *EditorHandler
}

func NewEditorRouter(handler *EditorHandler) *EditorRouter {
	return &EditorRouter{
		basePath: "/api/editor/",
		handler:  handler,
	}
}

func (router *EditorRouter) Register(engine *gin.Engine) {
	group := engine.Group(router.basePath)

	group.GET("flows/:id/", router.handler.HandleGeFlowById)

	group.PUT("flows/", router.handler.HandleUpdateFlowById)

	group.DELETE("flows/:id/", router.handler.HandleDeleteFlowById)

	group.GET("events/:id/", router.handler.HandleGetEventById)

	group.PUT("events/", router.handler.HandleUpdateEventById)

	group.DELETE("events/:id/", router.handler.HandleDeleteEventById)
}
