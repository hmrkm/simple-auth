package adapter

import (
	"time"

	"github.com/hmrkm/simple-auth/usecase"
)

//go:generate mockgen -source=$GOFILE -self_package=github.com/hmrkm/simple-auth/$GOPACKAGE -package=$GOPACKAGE -destination=auth_mock.go
type AuthAdapter interface {
	Verify(RequestAuth, time.Time, int) (ResponseAuth, error)
}

type authAdapter struct {
	userService  usecase.UserService
	tokenService usecase.TokenService
}

func NewAuthAdapter(us usecase.UserService, ts usecase.TokenService) AuthAdapter {
	return authAdapter{
		userService:  us,
		tokenService: ts,
	}
}

func (a authAdapter) Verify(req RequestAuth, now time.Time, tokenExpireHour int) (ResponseAuth, error) {
	user, err := a.userService.Verify(req.Email, req.Password)
	if err != nil {
		return ResponseAuth{}, err
	}

	token, err := a.tokenService.Create(user, now, tokenExpireHour)
	if err != nil {
		return ResponseAuth{}, err
	}

	return ResponseAuth{
		Token:     token.Token,
		ExpiredAt: token.GetEpochExpiredAt(),
	}, nil
}
