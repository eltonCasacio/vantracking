package driver

import (
	"time"

	repo "github.com/eltoncasacio/vantracking/internal/domain/driver/repository"
	"github.com/go-chi/jwtauth"
)

type AuthenticateUseCase struct {
	repository    repo.DriverRepositoryInterface
	JWT           *jwtauth.JWTAuth
	JwtExpiriesIn int
}

func NewAuthenticateUseCase(repository repo.DriverRepositoryInterface, jwt *jwtauth.JWTAuth, jwtExpiriesIn int) *AuthenticateUseCase {
	return &AuthenticateUseCase{
		repository:    repository,
		JWT:           jwt,
		JwtExpiriesIn: jwtExpiriesIn,
	}
}

func (u *AuthenticateUseCase) Authenticate(cpf string) (OutputDTO, error) {
	driver, err := u.repository.FindByCPF(cpf)
	if err != nil {
		return OutputDTO{}, err
	}

	_, token, err := u.JWT.Encode(map[string]interface{}{
		"sub": driver.ID,
		"exp": time.Now().Add(time.Second * time.Duration(u.JwtExpiriesIn)).Unix(),
	})

	accessToken := OutputDTO{
		AccessToken: token,
		User: User{
			ID:   driver.ID.String(),
			Name: driver.Name,
		},
	}
	return accessToken, nil
}
