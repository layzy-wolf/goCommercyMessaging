package storage

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Login    string
	Email    string `gorm:"unique"`
	Password string `gorm:"unique"`
	TTLToken time.Time
}
