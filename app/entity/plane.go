package entity

import "github.com/gin-gonic/gin"

type (
	Plane struct {
		ID          string
		Model       string
		Description string
		Status      string
	}
	PlaneInput struct {
		Model       string `json:"model" :binding:"required"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}

	PlaneOutput struct {
		ID          string `json:"id"`
		Model       string `json:"model"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}

	PlaneRepository interface {
		Create(plane *Plane) error
		Find(id int) (Plane, error)
	}

	PlaneUseCase interface {
		Create(plane PlaneInput) (Plane, error)
		Find(id int) (PlaneOutput, error)
	}

	PlaneHandler interface {
		Create(context *gin.Context)
		Find(context *gin.Context)
	}

	PlanePresenter interface {
		Output(plane Plane) PlaneOutput
	}
)
