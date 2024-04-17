package main

import (
	"net/http"
)

func addRoutes(
	mux *http.ServeMux,
	dataPath string,
) {
	mux.Handle("GET /", http.NotFoundHandler())
	mux.HandleFunc("GET /active_leagues", func(w http.ResponseWriter, r *http.Request) {
		get_active_leagues(w, r, dataPath)
	})
}
