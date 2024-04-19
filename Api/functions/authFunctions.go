package functions

import (
	"context"
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
	token := createToken(config)
	return token, nil
}

// Request a token from the web, then returns the retrieved token.
// TODO: replace log fatals
func createToken(config *oauth2.Config) *oauth2.Token {
	authCode := getAuthCode(config)

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
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

// validate users token
// TODO: return error from this function
func Validate(w http.ResponseWriter, r *http.Request) *http.Client {
	// getting token data from body
	r.ParseForm()
	access_token := r.Form["access_token"][0]
	refresh_token := r.Form["refresh_token"][0]
	expiryString := r.Form["expiry"][0]

	if access_token == "" || refresh_token == "" || expiryString == "" {
		fmt.Fprintln(w, "Token couldnt be made because not all parts were provided body. Needs access_token, refresh_token and expiry")
	}

	timeFormat := "2024-04-15T11:55:11.215558+02:00"
	expiry, err := time.Parse(timeFormat, expiryString)
	if err != nil {
		fmt.Fprintf(w, "Couldnt parse expiry string: %v\n", err)
		return nil
	}

	token := &oauth2.Token{
		AccessToken:  access_token,
		TokenType:    "Bearer",
		RefreshToken: refresh_token,
		Expiry:       expiry,
	}

	b, err := os.ReadFile("credentials.json")
	if err != nil {
		fmt.Fprintf(w, "Unable to read client secret file: %v", err)
		return nil
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		fmt.Fprintf(w, "Unable to parse client secret file to config: %v", err)
		return nil
	}
	return config.Client(context.Background(), token)
}
