package players

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Username string
	Password string
	CharName string
}

type APIPlayer struct {
	Username string
	CharName string
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
