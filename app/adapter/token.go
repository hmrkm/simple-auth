package adapter

import (
	"time"

	"github.com/hmrkm/simple-auth/usecase"

	"github.com/pkg/errors"
)

//go:generate mockgen -source=$GOFILE -self_package=github.com/hmrkm/simple-auth/$GOPACKAGE -package=$GOPACKAGE -destination=token_mock.go
type TokenAdapter interface {
	Verify(RequestVerify, time.Time) (ResponseVerify, error)
}

type tokenAdapter struct {
	store usecase.Store
}

func NewTokenAdapter(s usecase.Store) TokenAdapter {
	return tokenAdapter{
		store: s,
	}
}

func (ta tokenAdapter) Verify(p RequestVerify, now time.Time) (res ResponseVerify, err error) {
	t := usecase.Token{}
	if err := ta.store.First(&t, "token=?", string(p.Token)); err != nil {
		if ta.store.IsNotFoundError(err) {
			return ResponseVerify{}, errors.WithStack(usecase.ErrNotFound)
		}
		return ResponseVerify{}, errors.WithStack(err)
	}

	if !t.IsValid(now) {
		return ResponseVerify{}, errors.WithStack(usecase.ErrTokenWasExpired)
	}

	u := usecase.User{}
	if err := ta.store.First(&u, "id=?", t.UserId); err != nil {
		return ResponseVerify{}, errors.WithStack(err)
	}

	return ResponseVerify{
		User: VerifyUser{
			Id:    u.Id,
			Email: u.Email,
		},
	}, nil
}
