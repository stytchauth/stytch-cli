package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/stytchauth/stytch-cli/utils"

	"github.com/spf13/cobra"
)

const (
	PortUrl       = "127.0.0.1:5001"
	ClientId      = "connected-app-live-d552cc32-a785-4371-bf85-0af85f5f7067"
	CodeChallenge = "dcvJOywUwb5HOWOyhmOI5dSc4_VHQU8Xkp9bXD-tGWI"
	CodeVerifier  = "00afa9f459ce29bd4f9cd89f3a26036c3a2a772abd929e3fe179cb41"
	ProjectId     = "project-live-0f74ccf8-79bd-4096-bd3f-5317c0e69a3b"
	Scope         = "openid email profile project"
)

func NewAuthenticateCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "authenticate",
		Short: "Start authentication flow via Stytch",
		Run: func(cmd *cobra.Command, args []string) {
			stop := make(chan struct{})

			mux := http.NewServeMux()
			mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				handleCallback(w, r)
				stop <- struct{}{}
			})

			server := &http.Server{
				Addr:    PortUrl,
				Handler: mux,
			}
			go func() {
				fmt.Printf("Listening on http://%s/\n", PortUrl)
				if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
					fmt.Printf("Server error: %v\n", err)
					panic(err)
				}
			}()

			// Build the authentication URL
			u, _ := url.Parse("https://ollie.dev.stytch.com/idp")

			params := u.Query()
			params.Add("response_type", "code")
			params.Add("client_id", ClientId)
			params.Add("redirect_uri", fmt.Sprintf("http://%s", PortUrl))
			params.Add("code_challenge", CodeChallenge)
			params.Add("code_verifier", CodeVerifier)
			params.Add("scope", Scope)
			u.RawQuery = params.Encode()

			// Open browser
			utils.OpenBrowser(u.String())

			// Keep the program running
			<-stop

			// shut down the server
			if err := server.Shutdown(context.Background()); err != nil {
				log.Fatalf("Server shutdown failed: %v", err)
			}
			fmt.Println("Server shutdown successfully")
		},
	}

	return command
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Missing code parameter", http.StatusBadRequest)
		return
	}
	fmt.Printf("✅ Received code: %s\n", code)

	accessToken := getAccessTokenFromCode(code)
	// Save the access token securely
	err := utils.SaveToken(accessToken)
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ Access token saved")

	// Send 302 redirect to a friendly page (Stytch recommends redirecting away from localhost)
	http.Redirect(w, r, "https://stytch.com", http.StatusFound)
}

type GetAccessTokenResp struct {
	AccessToken string `json:"access_token"`
}

func getAccessTokenFromCode(code string) string {
	// make request to stytch with code to get access token
	// store the access token/refresh token locally
	tokenUrl := fmt.Sprintf("https://api.ollie.dev.stytch.com/v1/public/%s/oauth2/token", ProjectId)
	requestBody := map[string]interface{}{
		"client_id":     ClientId,
		"redirect_uri":  fmt.Sprintf("http://%s", PortUrl),
		"grant_type":    "authorization_code",
		"code":          code,
		"code_verifier": CodeVerifier,
	}

	bodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalf("Unable to marshal auth code request to JSON: %v", err)
	}

	// Make the HTTP request
	req, _ := http.NewRequest("POST", tokenUrl, bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error exchanging auth code: %v", err)
	}
	defer resp.Body.Close()

	// Read and print the response
	var getAccessTokenResp GetAccessTokenResp
	if err := json.NewDecoder(resp.Body).Decode(&getAccessTokenResp); err != nil {
		log.Fatalf("Error decoding access token response: %v", err)
	}

	return getAccessTokenResp.AccessToken
}
