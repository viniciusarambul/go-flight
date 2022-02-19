package pilot

import "github.com/viniciusarambul/go-flight/app/entity"

type PilotUseCase struct {
	pilotRepository entity.PilotRepository
	pilotPresenter  entity.PilotPresenter
}

func (pilotUseCase *PilotUseCase) Create(pilotInput entity.PilotInput) (entity.Pilot, error) {
	pilot := entity.Pilot{
		Name:     pilotInput.Name,
		LastName: pilotInput.LastName,
		Document: pilotInput.Document,
	}
	err := pilotUseCase.pilotRepository.Create(&pilot)
	return pilot, err
}

func (pilotUseCase *PilotUseCase) FindById(id int) (entity.PilotOutput, error) {
	pilot, err := pilotUseCase.pilotRepository.FindById(id)
	output := pilotUseCase.pilotPresenter.Output(pilot)
	return output, err
}

func (pilotUseCase *PilotUseCase) FindAll(pilotList *entity.Pilot) (entity.PilotList, error) {
	pilot, err := pilotUseCase.pilotRepository.FindAll(pilotList)
	output := pilotUseCase.pilotPresenter.ListOutput(pilot)
	return output, err
}

func NewPilotUseCase(repository entity.PilotRepository, presenter entity.PilotPresenter) entity.PilotUseCase {
	return &PilotUseCase{pilotRepository: repository, pilotPresenter: presenter}
}
