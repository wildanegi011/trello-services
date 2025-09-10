package main

import (
	"log"
	"net/http"
	"trello-services/infrastructure/config"
	"trello-services/infrastructure/db"
	delivery "trello-services/internal/delivery/http"
	"trello-services/internal/usecase"

	"github.com/gin-gonic/gin"
)

// global router to reuse between invocations (important for serverless)
var r *gin.Engine

func init() {
	// load environment variables
	config.LoadEnv()

	// Set Gin mode
	if config.GetEnv("MODE", "") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// build router once
	r = setupRouter()
}

func main() {
	// For environments where we run a normal server (not serverless)
	port := config.GetEnv("SERVER_PORT", "8080")
	if port[0] != ':' {
		port = ":" + port
	}

	log.Printf("Server running at %s\n", port)
	if err := http.ListenAndServe(port, r); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}

}

func setupRouter() *gin.Engine {
	database := db.NewPostgresDB()
	boardRepo := db.NewBoardRepo(database)
	boardUsecase := usecase.NewBoardUsecase(boardRepo)

	router := gin.Default()
	if err := router.SetTrustedProxies(nil); err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	api := router.Group("/api/v1")
	delivery.NewBoardHandler(api, boardUsecase)

	return router
}

// For serverless platforms (like Vercel/Railway handler entrypoint)
func Handler(w http.ResponseWriter, req *http.Request) {
	r.ServeHTTP(w, req)
}
