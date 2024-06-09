package main

import (
	"ApiGateway/config"
	"ApiGateway/internal/transport"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// инициализация конфига
	cfg := config.Load()

	// инициализация маршрутизатора
	r := transport.Handler(cfg)

	// Создание сервера
	srv := &http.Server{
		Addr: fmt.Sprintf(":%v", cfg.GatewayPort),
		//Handler: router.NewRouter(&cfg),
		Handler: r,
	}

	log.Printf("Starting server on port: %v", srv.Addr)

	// Запуск сервера
	if err := srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Fatal http server failed to start: %v", err)
		}
	}
}
