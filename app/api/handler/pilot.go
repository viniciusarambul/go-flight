package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/viniciusarambul/go-flight/app/entity"
)

type PilotHandler struct {
	PilotUseCase entity.PilotUseCase
}

func (pilotHandler PilotHandler) FindAll(context *gin.Context) {
	var query *entity.Pilot
	pilot, err := pilotHandler.PilotUseCase.FindAll(query)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao chamar endpoint"})
		return
	}

	context.JSON(http.StatusOK, pilot)
}

func (pilotHandler PilotHandler) FindById(context *gin.Context) {
	ID, err := strconv.Atoi(context.Param("pilotId"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pilot ID"})
		return
	}

	pilot, err := pilotHandler.PilotUseCase.FindById(ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Pilot not found"})
		return
	}

	context.JSON(http.StatusOK, pilot)
}

func (pilotHandler PilotHandler) Create(context *gin.Context) {
	var pilotInput entity.PilotInput
	err := context.ShouldBindJSON(&pilotInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pilot, err := pilotHandler.PilotUseCase.Create(pilotInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, pilot)
}

func NewPilotHandler(engine *gin.Engine, PilotUseCase entity.PilotUseCase) entity.PilotHandler {
	handler := &PilotHandler{PilotUseCase}

	engine.GET("/pilot", handler.FindAll)
	engine.GET("/pilot/:pilotId", handler.FindById)
	engine.POST("/pilot", handler.Create)
	return handler
}
