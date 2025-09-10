package handler

import (
	"net/http"
	"trello-services/internal/app"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	app.Handler(w, r)
}
