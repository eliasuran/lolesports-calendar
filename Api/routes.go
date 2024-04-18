package main

import (
	"net/http"

	"github.com/eliasuran/lolesports-calendar-api/functions"
	"golang.org/x/oauth2"
)

func addRoutes(
	mux *http.ServeMux,
	dataPath string,
) {
	mux.Handle("GET /", http.NotFoundHandler())

	// auth
	mux.HandleFunc("POST /validate", func(w http.ResponseWriter, r *http.Request) {
		// move this to functions
		r.ParseForm()
		access_token := r.Form["token"][0]

		token := &oauth2.Token{
			AccessToken: access_token,
			// TODO: add rest of the fields
		}

		client := functions.Validate(token)
	})
	mux.HandleFunc("POST /auth", func(w http.ResponseWriter, r *http.Request) {
		functions.GetToken()
	})
	mux.HandleFunc("GET /callback", func(w http.ResponseWriter, r *http.Request) {
		functions.Auth_callback(w, r)
	})

	// data
	mux.HandleFunc("GET /leagues", func(w http.ResponseWriter, r *http.Request) {
		functions.Get_active_leagues(w, r, dataPath)
	})
	mux.HandleFunc("GET /leagues/{id}", func(w http.ResponseWriter, r *http.Request) {
		functions.Get_league(w, r, dataPath)
	})
}
