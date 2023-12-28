package service

import (
	"app/internal/storage"
	"github.com/mhchlib/go-kit/log"
)

type Service struct {
	store *storage.Store
	log   log.Logger
}

func New(store *storage.Store, log log.Logger) *Service {
	return &Service{
		store,
		log,
	}
}
