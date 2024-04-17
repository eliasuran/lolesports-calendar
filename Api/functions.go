package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Data struct {
	Leagues Leagues
}

type Leagues struct {
	League []League
}

type League struct {
	ID       string
	Team1    string
	Team2    string
	DateTime string
}

func get_leagues(w http.ResponseWriter, r *http.Request, dataPath string) {
	jsonData, err := os.ReadFile(dataPath + "data.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not read file: %v\n", err)
		return
	}

	var leagues Data

	if err = json.Unmarshal(jsonData, &leagues); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not parse json: %v\n", err)
		return
	}

	w.WriteHeader(200)
	fmt.Fprintln(w, leagues)
}
