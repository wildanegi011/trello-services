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
	"trello-services/infrastructure/db"
	delivery "trello-services/internal/delivery/http"
	"trello-services/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	// load environment variables
	config.LoadEnv()

	// Set Gin mode
	if config.GetEnv("MODE", "") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// build router
	r := setupRouter()

	// get port form env (default : 8080)
	port := config.GetEnv("SERVER_PORT", ":8080")

	// create server
	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	// run server in gorountine
	go func() {
		log.Printf("server running at %s\n", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	// graceful shutdown
	gracefulShutdown(srv, 5*time.Second)

}

func setupRouter() *gin.Engine {
	database := db.NewPostgresDB()
	boardRepo := db.NewBoardRepo(database)
	boardUsecase := usecase.NewBoardUsecase(boardRepo)

	r := gin.Default()
	if err := r.SetTrustedProxies(nil); err != nil {
		log.Fatalf("failed to set trusted proxies: %v", err)
	}

	api := r.Group("/api/v1")
	delivery.NewBoardHandler(api, boardUsecase)

	return r
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
