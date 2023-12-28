package storage

import (
	"github.com/mhchlib/go-kit/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store struct {
	DB  *gorm.DB
	log log.Logger
}

func New(log log.Logger,
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
		_ = log.Log("Error", "migration failed")
	}

	return &Store{
		db,
		log,
	}
}
