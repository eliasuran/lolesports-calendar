package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/eliasuran/lolesports-calendar-api/functions"
	"golang.org/x/oauth2"
)

func addRoutes(
	ctx context.Context,
	mux *http.ServeMux,
	pantryUrl string,
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
	mux.HandleFunc("POST /newCalendar", func(w http.ResponseWriter, r *http.Request) {
		client, err := functions.Validate(w, r)
		if err != nil {
			fmt.Fprintf(w, "Couldnt validate token: %v\n", err)
			return
		}
		calendar, err := functions.CreateCalendar(ctx, client)
		if err != nil {
			fmt.Fprintf(w, "Could not create calendar: %v\n", err)
			return
		}
		fmt.Println(calendar)
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
	// TODO: move this to auth functions
	mux.HandleFunc("GET /auth", func(w http.ResponseWriter, r *http.Request) {
		config, err := functions.GetConfig()
		if err != nil {
			fmt.Fprintf(w, "Error getting config: %v\n", err)
			return
		}

		code := r.Form["code"]

		if len(code) == 0 {
			url := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
			fmt.Fprintln(w, url)
			return
		}
	})
	mux.HandleFunc("GET /callback", func(w http.ResponseWriter, r *http.Request) {
		config, err := functions.GetConfig()
		if err != nil {
			fmt.Fprintf(w, "Error getting config: %v\n", err)
			return
		}
		code := r.URL.Query().Get("code")
		if code == "" {
			fmt.Fprintln(w, "No code in url")
		}
		token, err := functions.CreateToken(config, code)
		if err != nil {
			fmt.Fprintf(w, "Error getting token: %v\n", err)
			return
		}
		fmt.Fprintln(w, "Token:", token)
	})

	// data
	mux.HandleFunc("GET /leagues", func(w http.ResponseWriter, r *http.Request) {
		functions.Get_active_leagues(w, r, pantryUrl)
	})
	mux.HandleFunc("GET /leagues/{id}", func(w http.ResponseWriter, r *http.Request) {
		functions.Get_league(w, r, pantryUrl)
	})
}
