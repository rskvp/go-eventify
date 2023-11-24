package explorer

import (
	"assalielmehdi/eventify/app/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExplorerHandler struct {
	explorerService *ExplorerService
}

func NewExplorerHandler(explorerService *ExplorerService) *ExplorerHandler {
	return &ExplorerHandler{
		explorerService: explorerService,
	}
}

func (handler *ExplorerHandler) HandleGetTree(ctx *gin.Context) {
	tree, err := handler.explorerService.GetTree()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, handlers.NewHandlerError(err))
	}

	ctx.JSON(http.StatusOK, handlers.NewHandlerSuccess(tree))
}
