package entity

import "github.com/gin-gonic/gin"

type (
	Pilot struct {
		ID       int64 `gorm:"primary_key;auto_increment;not_null"`
		Name     string
		LastName string
		Document string
	}
	PilotInput struct {
		Name     string `json:"name" :binding:"required"`
		LastName string `json:"lastName"`
		Document string `json:"document"`
	}

	PilotOutput struct {
		ID       int64  `json:"id"`
		Name     string `json:"name"`
		LastName string `json:"lastName"`
		Document string `json:"document"`
	}

	PilotList struct {
		ID       int64  `json:"id"`
		Name     string `json:"name"`
		LastName string `json:"lastName"`
		Document string `json:"document"`
	}

	PilotRepository interface {
		Create(pilot *Pilot) error
		FindById(id int) (Pilot, error)
		FindAll(pilot *Pilot) (Pilot, error)
	}

	PilotUseCase interface {
		Create(pilot PilotInput) (Pilot, error)
		FindById(id int) (PilotOutput, error)
		FindAll(pilot *Pilot) (PilotList, error)
	}

	PilotHandler interface {
		Create(context *gin.Context)
		FindById(context *gin.Context)
		FindAll(context *gin.Context)
	}

	PilotPresenter interface {
		Output(pilot Pilot) PilotOutput
		ListOutput(pilot Pilot) PilotList
	}
)
