package main

import (
	"net/http"

	"github.com/eliasuran/lolesports-calendar-api/functions"
)

func addRoutes(
	mux *http.ServeMux,
	dataPath string,
) {
	mux.Handle("GET /", http.NotFoundHandler())

	// auth
	mux.HandleFunc("POST /auth", func(w http.ResponseWriter, r *http.Request) {
		functions.Authorize()
	})
	mux.HandleFunc("POST /callback", func(w http.ResponseWriter, r *http.Request) {
		token := functions.Token{}
		functions.Auth_callback(token)
	})

	// data
	mux.HandleFunc("GET /leagues", func(w http.ResponseWriter, r *http.Request) {
		functions.Get_active_leagues(w, r, dataPath)
	})
	mux.HandleFunc("GET /leagues/{id}", func(w http.ResponseWriter, r *http.Request) {
		functions.Get_league(w, r, dataPath)
	})
}
