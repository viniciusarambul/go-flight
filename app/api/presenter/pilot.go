package presenter

import "github.com/viniciusarambul/go-flight/app/entity"

type pilotPresenter struct{}

func (pilotPresenter pilotPresenter) Output(pilot entity.Pilot) entity.PilotOutput {
	return entity.PilotOutput{
		ID:       pilot.ID,
		Name:     pilot.Name,
		LastName: pilot.LastName,
		Document: pilot.Document,
	}
}

func (pilotPresenter pilotPresenter) ListOutput(pilot entity.Pilot) entity.PilotList {
	return entity.PilotList{
		ID:       pilot.ID,
		Name:     pilot.LastName,
		LastName: pilot.LastName,
		Document: pilot.Document,
	}

}

func NewPilotPresenter() entity.PilotPresenter {
	return &pilotPresenter{}
}
