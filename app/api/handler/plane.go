package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/viniciusarambul/go-flight/app/entity"
)

type PlaneHandler struct {
	PlaneUseCase entity.PlaneUseCase
}

func (planeHundler PlaneHandler) Find(context *gin.Context) {
	ID, err := strconv.Atoi(context.Param("planeId"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid plane ID"})
		return
	}

	plane, err := planeHundler.PlaneUseCase.Find(ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Plane not found"})
		return
	}

	context.JSON(http.StatusOK, plane)
}

func (planeHandler PlaneHandler) Create(context *gin.Context) {
	var planeInput entity.PlaneInput
	err := context.ShouldBindJSON(&planeInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	plane, err := planeHandler.PlaneUseCase.Create(planeInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, plane)
}

func NewPlaneHandler(engine *gin.Engine, PlaneUseCase entity.PlaneUseCase) entity.PlaneHandler {
	handler := &PlaneHandler{PlaneUseCase}
	engine.GET("/plane/:planeId", handler.Find)
	engine.POST("/plane", handler.Create)
	return handler
}
