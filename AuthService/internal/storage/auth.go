package storage

import (
	"gorm.io/gorm"
	"time"
)

type Auth interface {
	LoginUser(email, passwd string) (token string, err error)
	RegisterUser(email, login, passwd string) (userId int32, err error)
}

type User struct {
	gorm.Model
	Login    string
	Email    string `gorm:"unique"`
	Password string `gorm:"unique"`
	TTLToken time.Time
}
