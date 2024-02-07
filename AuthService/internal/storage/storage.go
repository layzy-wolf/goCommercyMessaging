package storage

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Store struct {
	DB *gorm.DB
}

func New(
	database string,
	port string,
	login string,
	passwd string) *Store {
	dsn := "host=" + database + " user=" + login + " password=" + passwd + " dbname=srvices" + " port=" + port + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect with db")
	}

	err = db.AutoMigrate(User{})
	if err != nil {
		log.Println("Error", "migration failed")
	}

	return &Store{
		db,
	}
}
