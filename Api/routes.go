package main

import (
	"context"
	"encoding/json"
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

	// calendar api
	mux.HandleFunc("GET /calendars", func(w http.ResponseWriter, r *http.Request) {
		client, err := functions.Validate(w, r)
		if err != nil {
			fmt.Fprintf(w, "Couldnt validate token: %v\n", err)
			return
		}
		calendars, err := functions.MyCalendars(ctx, client)
		fmt.Fprintln(w, calendars.Items)
	})
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
		_, err = functions.CreateCalendar(ctx, client)
		if err != nil {
			fmt.Fprintf(w, "Could not create calendar: %v\n", err)
			return
		}
		fmt.Fprintln(w, "Calendar created!")
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
		config, err := functions.GetConfig(w)
		if err != nil {
			fmt.Printf("Error getting config: %v\n", err)
			return
		}

		url := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
		json, err := json.Marshal(url)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error marshaling json: %v\n", err)
			return
		}
		fmt.Fprintln(w, string(json))
	})
	mux.HandleFunc("GET /token", func(w http.ResponseWriter, r *http.Request) {
		config, err := functions.GetConfig(w)
		if err != nil {
			fmt.Printf("Error getting config: %v\n", err)
			return
		}
		code := r.URL.Query().Get("code")
		if code == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("No token found")
			return
		}
		token, err := functions.CreateToken(w, config, code)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("Error getting token: %v\n", err)
			return
		}

		json, err := json.Marshal(token)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error marshaling token: %v\n", err)
			return
		}
		fmt.Fprintln(w, string(json))
	})

	// data
	mux.HandleFunc("GET /active_leagues", func(w http.ResponseWriter, r *http.Request) {
		functions.Get_active_leagues(w, r, pantryUrl)
	})
	mux.HandleFunc("GET /schedule/{id}", func(w http.ResponseWriter, r *http.Request) {
		functions.Get_schedule(w, r, pantryUrl)
	})
	mux.HandleFunc("GET /all_leagues", func(w http.ResponseWriter, r *http.Request) {
		functions.Get_all_leagues(w, r, pantryUrl)
	})
	mux.HandleFunc("GET /league/{id}", func(w http.ResponseWriter, r *http.Request) {
		functions.Get_league(w, r, pantryUrl)
	})
	mux.HandleFunc("GET /team/{id}", func(w http.ResponseWriter, r *http.Request) {
		functions.Get_team(w, r, pantryUrl)
	})

	// 404
	mux.Handle("GET /", http.NotFoundHandler())
}
