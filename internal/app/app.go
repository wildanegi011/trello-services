package app

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
	// load env
	config.LoadEnv()

	// set Gin mode
	if config.GetEnv("MODE", "") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// build router once
	r = setupRouter()
}

// Setup router and dependencies
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

// Expose HTTP handler for both server and serverless
func Handler(w http.ResponseWriter, req *http.Request) {
	r.ServeHTTP(w, req)
}
