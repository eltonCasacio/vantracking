package passenger

import (
	"sync"

	f "github.com/eltoncasacio/vantracking/internal/domain/passenger/factory"
	repo "github.com/eltoncasacio/vantracking/internal/domain/passenger/repository"
)

type updateUseCase struct {
	repository repo.PassengerRepositoryInterface
}

func NewUseCase(repository repo.PassengerRepositoryInterface) *updateUseCase {
	return &updateUseCase{
		repository: repository,
	}
}

func (u *updateUseCase) FinalizeRoute(passengers PassengerInputDTO) error {
	wg := sync.WaitGroup{}
	wg.Add(len(passengers))
	for _, p := range passengers {
		go UpdatePassenger(
			f.PassengerInputDTO{
				ID:                p.ID,
				Name:              p.Name,
				Nickname:          p.Nickname,
				RouteCode:         p.RouteCode,
				Goes:              p.Goes,
				Comesback:         p.Comesback,
				RegisterConfirmed: p.RegisterConfirmed,
				SchoolName:        p.SchoolName,
				MonitorID:         p.MonitorID,
			}, u, &wg)
	}
	return nil
}

func UpdatePassenger(value f.PassengerInputDTO, u *updateUseCase, wg *sync.WaitGroup) error {
	passenger, err := f.PassengerFactory().Instance(value)
	if err != nil {
		wg.Done()
		return err
	}

	err = u.repository.Update(passenger)
	if err != nil {
		wg.Done()
		return err
	}
	wg.Done()
	return nil
}
