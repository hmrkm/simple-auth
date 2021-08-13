package adapter

import (
	"time"

	"github.com/hmrkm/simple-auth/usecase"

	"github.com/pkg/errors"
)

//go:generate mockgen -source=$GOFILE -self_package=github.com/hmrkm/simple-auth/$GOPACKAGE -package=$GOPACKAGE -destination=token_mock.go
type TokenAdapter interface {
	Verify(GetV1VerifyParams, time.Time) (bool, error)
}

type tokenAdapter struct {
	store usecase.Store
}

func NewTokenAdapter(s usecase.Store) TokenAdapter {
	return tokenAdapter{
		store: s,
	}
}

func (ta tokenAdapter) Verify(p GetV1VerifyParams, now time.Time) (isValid bool, err error) {
	t := usecase.Token{}
	if err := ta.store.First(&t, "token=?", string(p.Token)); err != nil {
		return false, errors.WithStack(err)
	}

	if !t.IsValid(now) {
		return false, errors.WithStack(usecase.ErrTokenWasExpired)
	}

	return true, nil
}
