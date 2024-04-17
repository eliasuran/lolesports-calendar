package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Leagues struct {
	Leagues []League
}

type League struct {
	Name     string
	Schedule []Match
}

type Match struct {
	ID       string
	Team1    string
	Team2    string
	DateTime string
}

func get_active_leagues(w http.ResponseWriter, r *http.Request, dataPath string) {
	jsonData, err := os.ReadFile(dataPath + "data.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not read file: %v\n", err)
		return
	}

	var leagues Leagues

	if err = json.Unmarshal(jsonData, &leagues); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not parse json: %v\n", err)
		return
	}

	var data []string

	for i := range leagues.Leagues {
		data = append(data, leagues.Leagues[i].Name)
	}

	w.WriteHeader(200)
	fmt.Fprintln(w, data)
}
