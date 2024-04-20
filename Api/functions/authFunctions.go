package functions

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

func GetToken() (*oauth2.Token, error) {
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		return nil, err
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		return nil, err
	}
	token, err := createToken(config)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// Request a token from the web, then returns the retrieved token.
func createToken(config *oauth2.Config) (*oauth2.Token, error) {
	authCode := getAuthCode(config)

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		return nil, err
	}
	return tok, nil
}

// making chan from getting auth code
var authCodeChan = make(chan string)

// generating and visiting url
// then getting the auth code search param from the callback endpoint
func getAuthCode(config *oauth2.Config) string {
	url := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/callback", handleCallback)
	go http.ListenAndServe(":8888", nil)

	for {
		code := <-authCodeChan
		if code != "" {
			return code
		}
		time.Sleep(1 * time.Second)
	}
}

// handling callback
func handleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Auth code not found", http.StatusBadRequest)
	}
	authCodeChan <- code
	fmt.Fprintln(w, "Auth successful! You can close this window")
}

// validate user's token
func Validate(w http.ResponseWriter, r *http.Request) (*http.Client, error) {
	// getting token data from body
	r.ParseForm()

	token_type := r.Form["type"]
	if len(token_type) == 0 {
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
		return nil, errors.New("Malformed request, invalid token type or not all required arguments passed in body")
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
