package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/eliasuran/lolesports-calendar-api/functions"
)

func addRoutes(
	ctx context.Context,
	mux *http.ServeMux,
	dataPath string,
) {
	mux.Handle("GET /", http.NotFoundHandler())

	// calendar api
	mux.HandleFunc("POST /newEvent", func(w http.ResponseWriter, r *http.Request) {
		client, err := functions.Validate(w, r)
		if err != nil {
			fmt.Fprintf(w, "Couldnt validate token: %v\n", err)
			return
		}
		eventLink := functions.CreateEvent(ctx, client)
		fmt.Fprintln(w, eventLink)
	})

	// auth
	mux.HandleFunc("POST /validate", func(w http.ResponseWriter, r *http.Request) {
		client, err := functions.Validate(w, r)
		if err != nil {
			fmt.Fprintf(w, "Couldnt validate token: %v\n", err)
			return
		}
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
