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

	cfg := config.Load()
	r := transport.Handler(cfg)

	srv := &http.Server{
		Addr: fmt.Sprintf(":%v", cfg.GatewayPort),
		//Handler: router.NewRouter(&cfg),
		Handler: r,
	}

	//idleConnClosed := make(chan struct{})
	//
	//go shutdown.Graceful(idleConnClosed, srv)
	//
	log.Printf("Starting server on port: %v", srv.Addr)

	if err := srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Fatal http server failed to start: %v", err)
		}
	}
}
