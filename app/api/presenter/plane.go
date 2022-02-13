package presenter

import "github.com/viniciusarambul/go-flight/app/entity"

type planePresenter struct{}

func (planePresenter planePresenter) Output(plane entity.Plane) entity.PlaneOutput {
	return entity.PlaneOutput{
		ID:          plane.ID,
		Model:       plane.Model,
		Description: plane.Description,
		Status:      plane.Status,
	}
}

func NewPlanePresenter() entity.PlanePresenter {
	return &planePresenter{}
}
