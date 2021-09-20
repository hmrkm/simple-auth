package adapter

import (
	"time"

	"github.com/hmrkm/simple-auth/usecase"
)

type Auth interface {
	Auth(RequestAuth) (ResponseAuth, error)
	Verify(RequestVerify) (ResponseVerify, error)
}

type auth struct {
	authUsecase     usecase.Auth
	tokenUsecase    usecase.Token
	tokenExpireHour int
}

func NewAuth(
	au usecase.Auth,
	tu usecase.Token,
	teh int,
) Auth {
	return auth{
		authUsecase:     au,
		tokenUsecase:    tu,
		tokenExpireHour: teh,
	}
}

func (aa auth) Auth(req RequestAuth) (ResponseAuth, error) {
	t, err := aa.authUsecase.Verify(req.Email, req.Password, time.Now(), aa.tokenExpireHour)
	if err != nil {
		return ResponseAuth{}, err
	}

	return ResponseAuth{
		ExpiredAt: t.GetEpochExpiredAt(),
		Token:     t.Token,
	}, nil
}

func (aa auth) Verify(req RequestVerify) (ResponseVerify, error) {
	u, err := aa.tokenUsecase.Verify(req.Token, time.Now())
	if err != nil {
		return ResponseVerify{}, err
	}

	return ResponseVerify{
		User: VerifyUser{
			Email: u.Email,
			Id:    u.Id,
		},
	}, nil
}
