package main

import (
	"net/http"

	calendarapi "github.com/eliasuran/lolesports-calendar-api/CalendarApi"
)

func addRoutes(
	mux *http.ServeMux,
	dataPath string,
) {
	mux.Handle("GET /", http.NotFoundHandler())
	mux.HandleFunc("POST /auth", func(w http.ResponseWriter, r *http.Request) {
		calendarapi.Authorize()
	})
	mux.HandleFunc("GET /leagues", func(w http.ResponseWriter, r *http.Request) {
		get_active_leagues(w, r, dataPath)
	})
	mux.HandleFunc("GET /leagues/{id}", func(w http.ResponseWriter, r *http.Request) {
		get_league(w, r, dataPath)
	})
}
