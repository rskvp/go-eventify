package graph

import (
	"assalielmehdi/eventify/app/handlers"
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
		ctx.JSON(http.StatusInternalServerError, handlers.NewHandlerError(err))
	}

	ctx.JSON(http.StatusOK, handlers.NewHandlerSuccess(graph))
}

func (handler *GraphHandler) HandleUpdateEventPosition(ctx *gin.Context) {
	var payload []*FlowGraphNodePositionUpdate

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, handlers.NewHandlerError(err))
		return
	}

	err := handler.graphService.UpdateEventsPositions(payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, handlers.NewHandlerError(err))
	}

	ctx.JSON(http.StatusOK, handlers.NewHandlerSuccess("Batch update with success."))
}
