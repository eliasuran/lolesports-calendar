package functions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Leagues struct {
	Active_leagues []ScheduleLeague
	All_leagues    []League
}

type ScheduleLeague struct {
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

type League struct {
	ID    string
	Name  string
	Logo  string
	Teams []Team
}

type Team struct {
	Name  string
	Short string
	Image string
}

func marshalJson() {}

func get_pantry_data(url string) ([]byte, error) {
	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return body, nil
}

func Get_active_leagues(w http.ResponseWriter, r *http.Request, url string) {
	jsonData, err := get_pantry_data(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could get data from pantry: %v\n", err)
		return
	}

	var leagues Leagues

	if err = json.Unmarshal(jsonData, &leagues); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not parse json: %v\n", err)
		return
	}

	var data []string

	for i := range leagues.Active_leagues {
		data = append(data, leagues.Active_leagues[i].ID)
	}

	w.WriteHeader(200)
	fmt.Fprintln(w, data)
}

func Get_schedule(w http.ResponseWriter, r *http.Request, url string) {
	id := r.PathValue("id")

	jsonData, err := get_pantry_data(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could get data from pantry: %v\n", err)
		return
	}

	var leagues Leagues

	if err = json.Unmarshal(jsonData, &leagues); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not parse json: %v\n", err)
		return
	}

	var league ScheduleLeague

	for i := range leagues.Active_leagues {
		if strings.ToLower(leagues.Active_leagues[i].ID) == strings.ToLower(id) {
			league = leagues.Active_leagues[i]
			break
		}
	}

	if league.Name == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%v not found\n", id)
		return
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

func Get_all_leagues(w http.ResponseWriter, r *http.Request, url string) {
	jsonData, err := get_pantry_data(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could get data from pantry: %v\n", err)
		return
	}

	var leagues Leagues

	if err = json.Unmarshal(jsonData, &leagues); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not parse json: %v\n", err)
		return
	}

	jsonLeagues, err := json.Marshal(leagues.All_leagues)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not marshal json data: %v\n", err)
		return
	}

	w.WriteHeader(200)
	fmt.Fprintln(w, string(jsonLeagues))
}

func Get_league(w http.ResponseWriter, r *http.Request, url string) {
	id := r.PathValue("id")

	jsonData, err := get_pantry_data(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could get data from pantry: %v\n", err)
		return
	}

	var leagues Leagues

	if err = json.Unmarshal(jsonData, &leagues); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not parse json: %v\n", err)
		return
	}

	var league League

	for i := range leagues.All_leagues {
		if strings.ToLower(leagues.All_leagues[i].ID) == strings.ToLower(id) {
			league = leagues.All_leagues[i]
			break
		}
	}

	if league.Name == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%v not found\n", id)
		return
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

func Get_team(w http.ResponseWriter, r *http.Request, url string) {
	id := r.PathValue("id")

	jsonData, err := get_pantry_data(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could get data from pantry: %v\n", err)
		return
	}

	var leagues Leagues

	if err = json.Unmarshal(jsonData, &leagues); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not parse json: %v\n", err)
		return
	}

	var team Team

	for i := range leagues.All_leagues {
		for y := range leagues.All_leagues[i].Teams {
			if strings.ToLower(leagues.All_leagues[i].Teams[y].Short) == strings.ToLower(id) {
				team = leagues.All_leagues[i].Teams[y]
				break
			}
		}
	}

	if team.Name == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%v not found\n", id)
		return
	}

	jsonTeam, err := json.Marshal(team)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not marshal json data: %v\n", err)
		return
	}

	w.WriteHeader(200)
	fmt.Fprintln(w, string(jsonTeam))
}
