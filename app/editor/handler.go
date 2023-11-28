package editor

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"assalielmehdi/eventify/app"
	"assalielmehdi/eventify/app/models"
)

type EditorHandler struct {
	editorService *EditorService
}

func NewEditorHandler(editorService *EditorService) *EditorHandler {
	return &EditorHandler{
		editorService: editorService,
	}
}

func (handler *EditorHandler) HandleGeFlowById(ctx *gin.Context) {
	record, err := handler.editorService.GetFlowById(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, app.NewServerResponseError(err))
		return
	}

	ctx.JSON(http.StatusOK, app.NewServerResponseSuccess(record))
}

func (handler *EditorHandler) HandleUpdateFlowById(ctx *gin.Context) {
	var payload models.Flow

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, app.NewServerResponseError(err))
		return
	}

	err := handler.editorService.UpdateFlowById(&payload)
	if err != nil {
		ctx.JSON(http.StatusNotFound, app.NewServerResponseError(err))
		return
	}

	ctx.JSON(http.StatusOK, app.NewServerEmptyResponseSuccess())
}

func (handler *EditorHandler) HandleDeleteFlowById(ctx *gin.Context) {
	err := handler.editorService.DeleteFlowById(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, app.NewServerResponseError(err))
		return
	}

	ctx.JSON(http.StatusOK, app.NewServerEmptyResponseSuccess())
}

func (handler *EditorHandler) HandleGetEventById(ctx *gin.Context) {
	record, err := handler.editorService.GetEventById(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, app.NewServerResponseError(err))
		return
	}

	ctx.JSON(http.StatusOK, app.NewServerResponseSuccess(record))
}

func (handler *EditorHandler) HandleUpdateEventById(ctx *gin.Context) {
	var payload models.Event

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, app.NewServerResponseError(err))
		return
	}

	err := handler.editorService.UpdateEventById(&payload)
	if err != nil {
		ctx.JSON(http.StatusNotFound, app.NewServerResponseError(err))
		return
	}

	ctx.JSON(http.StatusOK, app.NewServerEmptyResponseSuccess())
}

func (handler *EditorHandler) HandleDeleteEventById(ctx *gin.Context) {
	err := handler.editorService.DeleteEventById(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, app.NewServerResponseError(err))
		return
	}

	ctx.JSON(http.StatusOK, app.NewServerEmptyResponseSuccess())
}
