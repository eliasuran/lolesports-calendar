package main

import (
	"net/http"
)

func addRoutes(
	mux *http.ServeMux,
	dataPath string,
) {
	mux.Handle("GET /", http.NotFoundHandler())
	mux.HandleFunc("GET /leagues", func(w http.ResponseWriter, r *http.Request) {
		get_active_leagues(w, r, dataPath)
	})
	mux.HandleFunc("GET /leagues/{id}", func(w http.ResponseWriter, r *http.Request) {
		get_league(w, r, dataPath)
	})
}
