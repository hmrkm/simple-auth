package domain

//go:generate mockgen -source=$GOFILE -self_package=github.com/hmrkm/simple-auth/$GOPACKAGE -package=$GOPACKAGE -destination=store_mock.go
type Store interface {
	Find(destAddr interface{}, cond string, params ...interface{}) error
	First(destAddr interface{}, cond string, params ...interface{}) error
	Create(value interface{}) error
}
