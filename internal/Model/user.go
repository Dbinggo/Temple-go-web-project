package Model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func (User) TableName() string {
	return "users"
}
