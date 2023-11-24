package explorer

import "assalielmehdi/eventify/app/routers"

type ExplorerRouter struct {
	basePath string
	handler  *ExplorerHandler
}

func NewExplorerRouter(handler *ExplorerHandler) *ExplorerRouter {
	return &ExplorerRouter{
		basePath: "/api/explorer",
		handler:  handler,
	}
}

func (router *ExplorerRouter) Register(server *routers.Server) {
	group := server.Router.Group(router.basePath)

	group.GET("/tree", router.handler.HandleGetTree)
}
