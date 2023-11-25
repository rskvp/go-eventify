package graph

import "assalielmehdi/eventify/app/routers"

type GraphRouter struct {
	basePath string
	handler  *GraphHandler
}

func NewGraphRouter(handler *GraphHandler) *GraphRouter {
	return &GraphRouter{
		basePath: "/api/graph",
		handler:  handler,
	}
}

func (router *GraphRouter) Register(server *routers.Server) {
	group := server.Router.Group(router.basePath)

	group.GET("/flows/:flowId", router.handler.HandleGetFlowGraph)

	group.PATCH("/events/position", router.handler.HandleUpdateEventPosition)
}
