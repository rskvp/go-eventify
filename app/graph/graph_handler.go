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
	eventId := ctx.Param("eventId")
	var payload FlowGraphNodePosition

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, handlers.NewHandlerError(err))
		return
	}

	position, err := handler.graphService.UpdateEventPosition(eventId, &payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, handlers.NewHandlerError(err))
	}

	ctx.JSON(http.StatusOK, handlers.NewHandlerSuccess(position))
}
