package explorer

import (
	"assalielmehdi/eventify/app"
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
		ctx.JSON(http.StatusInternalServerError, app.NewServerResponseError(err))
		return
	}

	ctx.JSON(http.StatusOK, app.NewServerResponseSuccess(tree))
}

func (handler *ExplorerHandler) HandleAddFlow(ctx *gin.Context) {
	var payload AddFlowRequest

	if err := ctx.BindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, app.NewServerResponseError(err))
		return
	}

	record, err := handler.explorerService.AddFlow(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, app.NewServerResponseError(err))
		return
	}

	ctx.JSON(http.StatusCreated, app.NewServerResponseSuccess(record))
}

func (handler *ExplorerHandler) HandleAddEvent(ctx *gin.Context) {
	var payload AddEventRequest

	if err := ctx.BindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, app.NewServerResponseError(err))
		return
	}

	record, err := handler.explorerService.AddEvent(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, app.NewServerResponseError(err))
		return
	}

	ctx.JSON(http.StatusCreated, app.NewServerResponseSuccess(record))
}
