package usecase

import (
	"time"

	"github.com/hmrkm/simple-auth/domain"
)

//go:generate mockgen -source=$GOFILE -self_package=github.com/hmrkm/simple-auth/$GOPACKAGE -package=$GOPACKAGE -destination=token_mock.go
type Token interface {
	Verify(token string, now time.Time) (domain.User, error)
}

type token struct {
	store domain.Store
}

func NewToken(s domain.Store) Token {
	return token{
		store: s,
	}
}

func (ta token) Verify(token string, now time.Time) (domain.User, error) {
	t := domain.Token{}
	if err := ta.store.First(&t, "token=?", token); err != nil {
		return domain.User{}, err
	}

	if !t.IsValid(now) {
		return domain.User{}, domain.ErrTokenWasExpired
	}

	u := domain.User{}
	if err := ta.store.First(&u, "id=?", t.UserId); err != nil {
		return domain.User{}, err
	}

	return u, nil
}
