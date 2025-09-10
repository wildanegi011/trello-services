package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"trello-services/infrastructure/config"
	"trello-services/internal/app"
)

func main() {
	r := app.NewRouter()

	// Get port from env (default 8080)
	port := config.GetEnv("SERVER_PORT", "8080")
	if port[0] != ':' {
		port = ":" + port
	}

	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	// Run server in goroutine
	go func() {
		log.Printf("Server running at %s\n", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	// Graceful shutdown
	gracefulShutdown(srv, 5*time.Second)
}

func gracefulShutdown(srv *http.Server, timeout time.Duration) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server forced to shutdown:", err)
	}

	log.Println("server exited properly")
}
