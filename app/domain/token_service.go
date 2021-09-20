package domain

import (
	"time"

	"github.com/pkg/errors"
)

//go:generate mockgen -source=$GOFILE -self_package=github.com/hmrkm/simple-auth/$GOPACKAGE -package=$GOPACKAGE -destination=token_service_mock.go
type TokenService interface {
	Create(User, time.Time, int) (Token, error)
}

type tokenService struct {
	store Store
}

func NewTokenService(s Store) TokenService {
	return tokenService{
		store: s,
	}
}

func (ts tokenService) Create(u User, now time.Time, expireHour int) (Token, error) {
	expiredAt := now.Add(time.Duration(expireHour) * time.Hour)
	token := Token{
		Token:     CreateHash(now.String()),
		UserId:    u.Id,
		ExpiredAt: expiredAt,
	}

	if err := ts.store.Create(&token); err != nil {
		return Token{}, errors.WithStack(err)
	}

	return token, nil
}
