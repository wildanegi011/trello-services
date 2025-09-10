package app

import (
	"log"
	"trello-services/infrastructure/config"
	"trello-services/infrastructure/db"
	delivery "trello-services/internal/delivery/http"
	"trello-services/internal/usecase"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	// Load environment variables
	config.LoadEnv()

	// Set Gin mode
	if config.GetEnv("MODE", "") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Init DB + repo + usecase
	database := db.NewPostgresDB()
	boardRepo := db.NewBoardRepo(database)
	boardUsecase := usecase.NewBoardUsecase(boardRepo)

	// Router setup
	r := gin.Default()
	if err := r.SetTrustedProxies(nil); err != nil {
		log.Fatalf("failed to set trusted proxies: %v", err)
	}

	api := r.Group("/api/v1")
	delivery.NewBoardHandler(api, boardUsecase)

	return r
}
