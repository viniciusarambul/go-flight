package repository

import (
	"github.com/viniciusarambul/go-flight/app/entity"
	"gorm.io/gorm"
)

type PilotMySQLRepository struct {
	DB *gorm.DB
}

func (p *PilotMySQLRepository) FindById(id int) (entity.Pilot, error) {
	var pilot entity.Pilot
	err := p.DB.First(&pilot, id)

	return pilot, err.Error
}

func (p *PilotMySQLRepository) Create(pilot *entity.Pilot) error {
	err := p.DB.Create(&pilot)
	return err.Error
}

func (p *PilotMySQLRepository) FindAll(*entity.Pilot) (entity.Pilot, error) {
	var pilot entity.Pilot
	err := p.DB.Find(&pilot)
	return pilot, err.Error
}

func NewPilotMySQLRepository(DB *gorm.DB) entity.PilotRepository {
	return &PilotMySQLRepository{DB}
}
