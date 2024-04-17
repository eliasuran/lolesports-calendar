package main

import (
	"fmt"
	"net/http"
)

func addRoutes(
	mux *http.ServeMux,
) {
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "xdd!")
	})
}
