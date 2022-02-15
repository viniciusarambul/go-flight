package repository

import (
	"github.com/viniciusarambul/go-flight/app/entity"
	"gorm.io/gorm"
)

type PlaneMySQLRepository struct {
	DB *gorm.DB
}

func (p *PlaneMySQLRepository) Find(id int) (entity.Plane, error) {
	var plane entity.Plane
	err := p.DB.First(&plane, id)

	return plane, err.Error
}

func (p *PlaneMySQLRepository) Create(plane *entity.Plane) error {
	err := p.DB.Create(&plane)
	return err.Error
}

func NewPlaneMySQLRepository(DB *gorm.DB) entity.PlaneRepository {
	return &PlaneMySQLRepository{DB}
}
