package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"assalielmehdi/eventify/app/models"
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

	ctx.JSON(http.StatusOK, NewHandlerResponse(http.StatusOK, records))
}

func (handler *FlowHandler) HandleGetOneById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "HandleGetOneById")
}

func (handler *FlowHandler) HandleAddOne(ctx *gin.Context) {
	var payload models.Flow

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		log.Print(err)
		ctx.JSON(http.StatusBadRequest, NewHandlerResponse(http.StatusBadRequest, "Failed to bind JSON data"))
		return
	}

	record := handler.FlowService.AddOne(&payload)

	ctx.JSON(http.StatusCreated, NewHandlerResponse(http.StatusCreated, record))
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
