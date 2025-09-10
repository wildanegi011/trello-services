package handler

import (
	"net/http"

	"trello-services/internal/app"
)

// create the router once (cold start)
var router = app.NewRouter()

// Vercel entrypoint (must be exported "Handler")
func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
