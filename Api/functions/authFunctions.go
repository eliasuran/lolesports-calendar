package functions

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

func GetConfig(w http.ResponseWriter) (*oauth2.Config, error) {
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		w.WriteHeader(500)
		return nil, err
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		w.WriteHeader(500)
		return nil, err
	}
	return config, nil
}

// Request a token from the web, then returns the retrieved token.
func CreateToken(w http.ResponseWriter, config *oauth2.Config, authCode string) (*oauth2.Token, error) {
	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil, err
	}
	return tok, nil
}

// validate user's token
func Validate(w http.ResponseWriter, r *http.Request) (*http.Client, error) {
	// getting token data from body
	r.ParseForm()

	var token *oauth2.Token

	auth_header := r.Header.Get("Authorization")
	if auth_header == "" {
		w.WriteHeader(http.StatusBadRequest)
		return nil, errors.New("No access token provided")
	}

	split_auth_header := strings.Split(auth_header, " ")
	if len(split_auth_header) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		return nil, errors.New("Auth header in invalid format")
	}

	token_type := strings.Split(auth_header, " ")[0]
	access_token := strings.Split(auth_header, " ")[1]

	token = &oauth2.Token{
		TokenType:   token_type,
		AccessToken: access_token,
	}

	b, err := os.ReadFile("credentials.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, err
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, err
	}
	return config.Client(context.Background(), token), nil
}
