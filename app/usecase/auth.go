package usecase

import (
	"time"

	"github.com/hmrkm/simple-auth/domain"
)

//go:generate mockgen -source=$GOFILE -self_package=github.com/hmrkm/simple-auth/$GOPACKAGE -package=$GOPACKAGE -destination=auth_mock.go
type AuthUsecase interface {
	Verify(email string, password string, now time.Time, tokenExpireHour int) (domain.Token, error)
}

type authUsecase struct {
	userService  domain.UserService
	tokenService domain.TokenService
}

func NewAuthUsecase(us domain.UserService, ts domain.TokenService) AuthUsecase {
	return authUsecase{
		userService:  us,
		tokenService: ts,
	}
}

func (a authUsecase) Verify(email string, password string, now time.Time, tokenExpireHour int) (domain.Token, error) {
	user, err := a.userService.Verify(email, password)
	if err != nil {
		return domain.Token{}, err
	}

	token, err := a.tokenService.Create(user, now, tokenExpireHour)
	if err != nil {
		return domain.Token{}, err
	}

	return token, nil
}
