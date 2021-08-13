package adapter

import (
	"time"

	"github.com/hmrkm/simple-auth/usecase"

	"github.com/pkg/errors"
)

//go:generate mockgen -source=$GOFILE -self_package=github.com/hmrkm/simple-auth/$GOPACKAGE -package=$GOPACKAGE -destination=auth_mock.go
type AuthAdapter interface {
	Verify(RequestPostAuth, time.Time, int) (ResponsePostAuth, error)
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

func (a authAdapter) Verify(req RequestPostAuth, now time.Time, tokenExpireHour int) (ResponsePostAuth, error) {
	isValid, user, err := a.userService.Verify(req.Email, req.Password)
	if err != nil {
		return ResponsePostAuth{}, err
	}
	if !isValid {
		return ResponsePostAuth{}, errors.WithStack(usecase.ErrInvalidVerify)
	}

	token, err := a.tokenService.Create(user, now, tokenExpireHour)
	if err != nil {
		return ResponsePostAuth{}, err
	}

	return ResponsePostAuth{
		Token:     token.Token,
		ExpiredAt: token.GetEpochExpiredAt(),
	}, nil
}
