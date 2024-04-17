package main

import (
	"fmt"
	"net/http"
	"os"
)

func run() error {
	mux := http.NewServeMux()

	// all routes
	addRoutes(mux)

	// defaults to port 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Printf("Listening on port %s\n", port)
	server.ListenAndServe()
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
