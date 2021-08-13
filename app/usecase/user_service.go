package usecase

import (
	"github.com/pkg/errors"
)

//go:generate mockgen -source=$GOFILE -self_package=github.com/hmrkm/simple-auth/$GOPACKAGE -package=$GOPACKAGE -destination=user_service_mock.go
type UserService interface {
	Verify(string, string) (bool, User, error)
}

type userService struct {
	store Store
}

func NewUserService(s Store) UserService {
	return userService{
		store: s,
	}
}

func (us userService) Verify(email string, password string) (isValid bool, user User, err error) {
	u := User{}
	if err := us.store.First(&u, "email=?", email); err != nil {
		return false, User{}, errors.WithStack(err)
	}

	if u.Password != CreateHash(password) {
		return false, User{}, errors.WithStack(ErrInvalidPassword)
	}

	return true, u, nil
}