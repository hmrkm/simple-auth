package domain

//go:generate mockgen -source=$GOFILE -self_package=github.com/hmrkm/simple-auth/$GOPACKAGE -package=$GOPACKAGE -destination=user_service_mock.go
type UserService interface {
	Verify(email string, password string) (User, error)
}

type userService struct {
	store Store
}

func NewUserService(s Store) UserService {
	return userService{
		store: s,
	}
}

func (us userService) Verify(email string, password string) (User, error) {
	u := User{}
	if err := us.store.First(&u, "email=?", email); err != nil {
		return User{}, err
	}

	if u.Password != CreateHash(password) {
		return User{}, ErrInvalidPassword
	}

	return u, nil
}
