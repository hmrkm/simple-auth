package usecase

import "time"

type Token struct {
	Token     string    `json:"token" gorm:"column:token"`
	UserId    string    `json:"-" gorm:"column:user_id"`
	ExpiredAt time.Time `json:"expired_at" gorm:"column:expired_at"`
}

func (Token) TableName() string {
	return "tokens"
}

func (t Token) IsValid(now time.Time) bool {
	return !t.ExpiredAt.Before(now)
}

func (t Token) GetEpochExpiredAt() int {
	return int(t.ExpiredAt.UnixNano() / 1000)
}
