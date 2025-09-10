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
	// load env
	config.LoadEnv()

	port := config.GetEnv("SERVER_PORT", "8080")
	if port[0] != ':' {
		port = ":" + port
	}

	srv := &http.Server{
		Addr:    port,
		Handler: http.HandlerFunc(app.Handler), // use app.Handler
	}

	// run server in goroutine
	go func() {
		log.Printf("Server running at %s\n", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// graceful shutdown
	gracefulShutdown(srv, 5*time.Second)
}

func gracefulShutdown(srv *http.Server, timeout time.Duration) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited properly")
}
