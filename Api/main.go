package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/eliasuran/lolesports-calendar-api/middleware"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func run() error {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		fmt.Println("no .env file found, trying to access data path elsewhere")
	}

	pantryUrl := os.Getenv("PANTRY_URL")

	mux := http.NewServeMux()

	// all routes
	addRoutes(ctx, mux, pantryUrl)

	// defaults to port 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	stack := middleware.CreateStack(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":" + port,
		Handler: stack(cors.Default().Handler(mux)),
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
