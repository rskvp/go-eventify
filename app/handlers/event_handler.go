package handlers

import (
	"assalielmehdi/eventify/app/models"
	"assalielmehdi/eventify/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	EventService services.EventService
}

func NewEventHandler(eventService *services.EventService) *EventHandler {
	return &EventHandler{
		EventService: *eventService,
	}
}

func (handler *EventHandler) HandleAddOne(ctx *gin.Context) {
	var payload models.Event

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, NewHandlerError(err))
		return
	}

	record := handler.EventService.AddOne(&payload)

	ctx.JSON(http.StatusCreated, NewHandlerSuccess(record))
}
