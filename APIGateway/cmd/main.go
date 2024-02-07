package main

import (
	"context"
	"errors"
	"github.com/layzy-wolf/goCommercyMessaging/APIGateway/internal/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router.NewRouter(),
	}

	idleConnClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		signal.Notify(sigint, syscall.SIGTERM)
		<-sigint

		log.Println("service interrupt received")

		log.Println("http server shutting down")
		time.Sleep(5 * time.Second)

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("http server shutdwon err: %v", err)
		}

		log.Println("shutdown complete")
		close(idleConnClosed)
	}()

	log.Printf("Starting server on port 8080")

	if err := srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Fatal http server failed to start: %v", err)
		}
	}
}
