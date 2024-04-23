package functions

import (
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

func GetConfig() (*oauth2.Config, error) {
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		return nil, err
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// Request a token from the web, then returns the retrieved token.
func CreateToken(config *oauth2.Config, authCode string) (*oauth2.Token, error) {
	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		return nil, err
	}
	return tok, nil
}

// validate user's token
func Validate(w http.ResponseWriter, r *http.Request) (*http.Client, error) {
	// getting token data from body
	r.ParseForm()

	token_type := r.Form["type"]
	if len(token_type) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return nil, errors.New("No token type provided")
	}

	var token *oauth2.Token

	if token_type[0] == "refresh_token" && len(r.Form["refresh_token"]) > 0 {
		refresh_token := r.Form["refresh_token"][0]
		token = &oauth2.Token{
			RefreshToken: refresh_token,
		}
	} else if token_type[0] == "access_token" && len(r.Form["access_token"]) > 0 && len(r.Form["expiry"]) > 0 {
		access_token := r.Form["access_token"][0]
		expiryString := r.Form["expiry"][0]
		timeFormat := "2006-01-02 15:04:05.999999999 -0700 MST"
		expiry, err := time.Parse(timeFormat, expiryString)
		if err != nil {
			return nil, err
		}
		token = &oauth2.Token{
			AccessToken: access_token,
			TokenType:   "Bearer",
			Expiry:      expiry,
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return nil, errors.New("Bad request, invalid token type or not all required arguments passed in body")
	}

	b, err := os.ReadFile("credentials.json")
	if err != nil {
		return nil, err
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		return nil, err
	}
	return config.Client(context.Background(), token), nil
}
