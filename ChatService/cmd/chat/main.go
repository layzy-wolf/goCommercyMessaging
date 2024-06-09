package main

import (
	"ChatService/config"
	"ChatService/internal/server"
	"ChatService/internal/service"
	"log"
)

func main() {
	// инициализация конфиг файла
	cfg := config.Load()

	// подключение к сервису-хранилищу
	client, conn := service.StoreClient(cfg)

	// после завершения работы отключение от сервиса
	defer conn.Close()

	// инициализация сервиса
	serv := service.New(client)

	// запуск локального сервера
	if err := server.Engine(serv, &cfg); err != nil {
		log.Panicf("failed to serve: %v", err)
	}
}
