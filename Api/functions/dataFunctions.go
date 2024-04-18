package functions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Leagues struct {
	Leagues []League
}

type League struct {
	ID       string
	Name     string
	Schedule []Match
}

type Match struct {
	ID       string
	Team1    string
	Team2    string
	DateTime string
}

func Get_active_leagues(w http.ResponseWriter, r *http.Request, dataPath string) {
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
		data = append(data, leagues.Leagues[i].ID)
	}

	w.WriteHeader(200)
	fmt.Fprintln(w, data)
}

func Get_league(w http.ResponseWriter, r *http.Request, dataPath string) {
	id := r.PathValue("id")
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

	var league League

	for i := range leagues.Leagues {
		if strings.ToLower(leagues.Leagues[i].ID) == strings.ToLower(id) {
			league = leagues.Leagues[i]
			break
		}
	}

	jsonLeague, err := json.Marshal(league)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not marshal json data: %v\n", err)
		return
	}

	w.WriteHeader(200)
	fmt.Fprintln(w, string(jsonLeague))
}
