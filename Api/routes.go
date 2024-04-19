package main

import (
	"fmt"
	"net/http"

	"github.com/eliasuran/lolesports-calendar-api/functions"
)

func addRoutes(
	mux *http.ServeMux,
	dataPath string,
) {
	mux.Handle("GET /", http.NotFoundHandler())

	// auth
	mux.HandleFunc("POST /validate", func(w http.ResponseWriter, r *http.Request) {
		client := functions.Validate(w, r)
		fmt.Fprintln(w, client)
	})
	mux.HandleFunc("POST /auth", func(w http.ResponseWriter, r *http.Request) {
		token, err := functions.GetToken()
		if err != nil {
			fmt.Fprintf(w, "Error getting token: %v\n", err)
			return
		}
		fmt.Fprintln(w, "Token:", token)
	})

	// data
	mux.HandleFunc("GET /leagues", func(w http.ResponseWriter, r *http.Request) {
		functions.Get_active_leagues(w, r, dataPath)
	})
	mux.HandleFunc("GET /leagues/{id}", func(w http.ResponseWriter, r *http.Request) {
		functions.Get_league(w, r, dataPath)
	})
}
