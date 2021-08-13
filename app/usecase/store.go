package usecase

//go:generate mockgen -source=$GOFILE -self_package=github.com/hmrkm/simple-auth/$GOPACKAGE -package=$GOPACKAGE -destination=store_mock.go
type Store interface {
	Find(interface{}, string, ...interface{}) error
	First(interface{}, string, ...interface{}) error
	Create(interface{}) error
}
