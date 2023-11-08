package http

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"io/ioutil"
	"net/http"
)

func (a *Api) getDriveOauthToken(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	b, err := ioutil.ReadFile("/Users/manginav/awesomeProject1/go_app/configs/credentials.json")
	if err != nil {
		a.WriteErrorResponse(r, w, fmt.Errorf("unable to read client secret file: %v", err))
		return
	}

	config, err := google.ConfigFromJSON(b, drive.DriveFileScope)
	if err != nil {
		a.WriteErrorResponse(r, w, fmt.Errorf("unable to parse client secret file to config: %v", err))
		return
	}

	client := getClient(ctx, config)

	// Create a Drive service client.
	srv, err := drive.New(client)
	if err != nil {
		a.WriteErrorResponse(r, w, fmt.Errorf("unable to retrieve Drive client: %v", err))
		return
	}

	// Your existing code here...

	a.WriteJsonBodyResponse(r, w, srv, http.StatusOK)
}

func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
	token := getTokenFromWeb(ctx, config)
	return config.Client(ctx, token)
}

func getTokenFromWeb(ctx context.Context, config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	fmt.Scan(&authCode)

	tok, _ := config.Exchange(ctx, authCode)
	return tok
}
