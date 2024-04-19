package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func run() error {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		fmt.Println("no .env file found, trying to access data path elsewhere")
	}

	dataPath := os.Getenv("DATA_PATH")
	if dataPath == "" {
		return errors.New("no datapath found")
	}

	fmt.Println("Data path set to: ", dataPath)

	mux := http.NewServeMux()

	// all routes
	addRoutes(ctx, mux, dataPath)

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
