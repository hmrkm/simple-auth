package domain

type User struct {
	Id       string `json:"id" gorm:"column:id"`
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"-" gorm:"column:password"`
}

func (User) TableName() string {
	return "users"
}
