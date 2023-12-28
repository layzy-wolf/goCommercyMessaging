package service

import (
	"app/internal/storage"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func (s Service) LoginUser(email, passwd string) (token string, err error) {
	var hmacSampleSecret []byte

	user := s.store.DB.Take(storage.User{}, "Email = ? && Password = ?", email, passwd)
	if user.Error != nil {
		return "", user.Error
	}

	s.store.DB.Model(&user).Update("TTlToken", time.Now().AddDate(0, 1, 0))

	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
		//TODO see there
	})

	tokenString, err := tok.SignedString(hmacSampleSecret)

	return tokenString, err
}

func (s Service) RegisterUser(email, login, passwd string) (userId int32, err error) {
	user := storage.User{
		Login:    login,
		Email:    email,
		Password: passwd,
		TTLToken: time.Now(),
	}

	res := s.store.DB.Create(user)

	return int32(user.ID), res.Error
}
