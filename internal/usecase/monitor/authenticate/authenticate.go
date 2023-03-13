package monitor

import (
	"time"

	repo "github.com/eltoncasacio/vantracking/internal/domain/monitor/repository"
	"github.com/go-chi/jwtauth"
)

type authenticateUseCase struct {
	repository    repo.MonitorRepositoryInterface
	JWT           *jwtauth.JWTAuth
	JwtExpiriesIn int
}

func NewUseCase(repository repo.MonitorRepositoryInterface, jwt *jwtauth.JWTAuth, jwtExpiriesIn int) *authenticateUseCase {
	return &authenticateUseCase{
		repository:    repository,
		JWT:           jwt,
		JwtExpiriesIn: jwtExpiriesIn,
	}
}

func (u *authenticateUseCase) Authenticate(cpf string) (OutputDTO, error) {
	monitor, err := u.repository.FindByCPF(cpf)
	if err != nil {
		return OutputDTO{}, err
	}

	_, token, err := u.JWT.Encode(map[string]interface{}{
		"sub": monitor.ID,
		"exp": time.Now().Add(time.Second * time.Duration(u.JwtExpiriesIn)).Unix(),
	})

	accessToken := OutputDTO{
		AccessToken: token,
		User: User{
			ID:   monitor.ID.String(),
			Name: monitor.Name,
		},
	}
	return accessToken, nil
}
