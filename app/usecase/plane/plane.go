package plane

import "github.com/viniciusarambul/go-flight/app/entity"

type PlaneUseCase struct {
	planeRepository entity.PlaneRepository
	planePresenter  entity.PlanePresenter
}

func (planeUseCase *PlaneUseCase) Create(planeInput entity.PlaneInput) (entity.Plane, error) {
	plane := entity.Plane{
		Model:       planeInput.Model,
		Description: planeInput.Description,
		Status:      planeInput.Status,
	}
	err := planeUseCase.planeRepository.Create(&plane)
	return plane, err
}

func (planeUseCase *PlaneUseCase) FindById(id int) (entity.PlaneOutput, error) {
	plane, err := planeUseCase.planeRepository.FindById(id)
	output := planeUseCase.planePresenter.Output(plane)
	return output, err
}

func NewPlaneUseCase(repository entity.PlaneRepository, presenter entity.PlanePresenter) entity.PlaneUseCase {
	return &PlaneUseCase{planeRepository: repository, planePresenter: presenter}
}
