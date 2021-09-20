package adapter

import (
	"time"

	"github.com/hmrkm/simple-auth/usecase"
)

type AuthAdapter interface {
	Auth(RequestAuth) (ResponseAuth, error)
	Verify(RequestVerify) (ResponseVerify, error)
}

type authAdapter struct {
	AuthUsecase     usecase.AuthUsecase
	TokenUsecase    usecase.TokenUsecase
	TokenExpireHour int
}

func NewAuthAdapter(
	au usecase.AuthUsecase,
	tu usecase.TokenUsecase,
	teh int,
) AuthAdapter {
	return authAdapter{
		AuthUsecase:     au,
		TokenUsecase:    tu,
		TokenExpireHour: teh,
	}
}

func (aa authAdapter) Auth(req RequestAuth) (ResponseAuth, error) {
	t, err := aa.AuthUsecase.Verify(req.Email, req.Password, time.Now(), aa.TokenExpireHour)
	if err != nil {
		return ResponseAuth{}, err
	}

	return ResponseAuth{
		ExpiredAt: t.GetEpochExpiredAt(),
		Token:     t.Token,
	}, nil
}

func (aa authAdapter) Verify(req RequestVerify) (ResponseVerify, error) {
	u, err := aa.TokenUsecase.Verify(req.Token, time.Now())
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
