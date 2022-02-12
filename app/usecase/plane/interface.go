package plane

import (
	"github.com/viniciusarambul/go-flight/app/entity"
)

type PlaneRepository interface {
	Insert(plane entity.Plane) error
	Select(plane entity.Plane) error
}
