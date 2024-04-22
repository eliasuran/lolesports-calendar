package functions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

	for i := range leagues.Leagues {
		data = append(data, leagues.Leagues[i].ID)
	}

	w.WriteHeader(200)
	fmt.Fprintln(w, data)
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
