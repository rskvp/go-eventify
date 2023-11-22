package handlers

import (
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
	record, err := handler.FlowService.GetOneById(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, NewHandlerResponse(http.StatusNotFound, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, NewHandlerResponse(http.StatusOK, record))
}

func (handler *FlowHandler) HandleAddOne(ctx *gin.Context) {
	var payload models.Flow

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, NewHandlerResponse(http.StatusBadRequest, err.Error()))
		return
	}

	record := handler.FlowService.AddOne(&payload)

	ctx.JSON(http.StatusCreated, NewHandlerResponse(http.StatusCreated, record))
}

func (handler *FlowHandler) HandleUpdateOneById(ctx *gin.Context) {
	var payload models.Flow

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, NewHandlerResponse(http.StatusBadRequest, err.Error()))
		return
	}

	record, err := handler.FlowService.UpdateOneById(&payload)
	if err != nil {
		ctx.JSON(http.StatusNotFound, NewHandlerResponse(http.StatusNotFound, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, NewHandlerResponse(http.StatusOK, record))
}

func (handler *FlowHandler) HandleDeleteOneById(ctx *gin.Context) {
	record, err := handler.FlowService.DeleteOneById(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, NewHandlerResponse(http.StatusNotFound, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, NewHandlerResponse(http.StatusOK, record))
}

func (handler *FlowHandler) HandleRunById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "HandleRunById")
}
