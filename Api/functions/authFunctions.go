package functions

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

func Auth_callback(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query().Get("code"))
	fmt.Fprintln(w, "Nice! You can return to the terminal")
}

func GetToken() *oauth2.Token {
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		fmt.Printf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		fmt.Printf("Unable to parse client secret file to config: %v", err)
	}
	token := createToken(config)
	return token
}

// authorize user (NOT DONE CHANGE THIS TO ONLY AUTHORIZE AND RETURN A TOKEN AND NOT RETURN A CLIENT)
func Validate(token *oauth2.Token) *http.Client {
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		fmt.Printf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		fmt.Printf("Unable to parse client secret file to config: %v", err)
	}
	client, token := GetClient(config, token)
	return client
}

func GetClient(config *oauth2.Config, token *oauth2.Token) (*http.Client, *oauth2.Token) {
	tok := token
	if tok.AccessToken == "" {
		tok = createToken(config)
	}
	return config.Client(context.Background(), tok), tok
}

// generating url to visit to login user
func authUrl(config *oauth2.Config) string {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return string(authURL)
}

// Request a token from the web, then returns the retrieved token.
func createToken(config *oauth2.Config) *oauth2.Token {
	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}
