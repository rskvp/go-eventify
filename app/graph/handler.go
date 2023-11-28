package graph

import (
	"assalielmehdi/eventify/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GraphHandler struct {
	graphService *GraphService
}

func NewGraphHandler(graphService *GraphService) *GraphHandler {
	return &GraphHandler{
		graphService: graphService,
	}
}

func (handler *GraphHandler) HandleGetFlowGraph(ctx *gin.Context) {
	flowId := ctx.Param("flowId")
	graph, err := handler.graphService.GetFlowGraph(flowId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, app.NewServerResponseError(err))
		return
	}

	ctx.JSON(http.StatusOK, app.NewServerResponseSuccess(graph))
}

func (handler *GraphHandler) HandleUpdateEventPosition(ctx *gin.Context) {
	var payload []*FlowGraphNodePositionUpdate

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, app.NewServerResponseError(err))
		return
	}

	err := handler.graphService.UpdateEventsPositions(payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, app.NewServerResponseError(err))
		return
	}

	ctx.JSON(http.StatusOK, app.NewServerResponseSuccess("Batch update with success."))
}
