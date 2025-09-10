package handler

import (
	"net/http"
	"trello-services/internal/app"
)

var router = app.NewRouter()

// Vercel entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
