package plane

import (
	"fmt"

	"github.com/viniciusarambul/go-flight/app/entity"
)

type ListPlanes struct {
	Repository PlaneRepository
}

func (l ListPlanes) Exec() (ShowPlaneOutputDto, error) {
	plane := entity.Plane{}

	err := l.Repository.Select(plane)
	fmt.Println("dentro do exec", err)
	if err != nil {
		return ShowPlaneOutputDto{}, err
	}

	output := ShowPlaneOutputDto{}
	output.ID = plane.ID
	output.Description = plane.Description
	output.Model = plane.Model
	output.Status = plane.Status
	fmt.Println("dentro do output", output)
	return output, nil
}
