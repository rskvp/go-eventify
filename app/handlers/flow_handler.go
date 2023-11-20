package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"assalielmehdi/eventify/app/services"
)

type FlowHandler struct {
	FlowService services.FlowService
}

func NewFlowHandler(flowService *services.FlowService) *FlowHandler {
	return &FlowHandler{
		FlowService: *flowService,
	}
}

func (handler *FlowHandler) HandleGetAll(ctx *gin.Context) {
	records := handler.FlowService.GetAll()

	ctx.JSON(http.StatusOK, records)
}

func (handler *FlowHandler) HandleGetOneById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "HandleGetOneById")
}

func (handler *FlowHandler) HandleAddOne(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "HandleAddOne")
}

func (handler *FlowHandler) HandleUpdateOneById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "HandleUpdateOneById")
}

func (handler *FlowHandler) HandleDeleteOneById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "HandleDeleteOneById")
}

func (handler *FlowHandler) HandleRunById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "HandleRunById")
}
