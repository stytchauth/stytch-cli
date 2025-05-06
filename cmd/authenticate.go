package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/stytchauth/stytch-cli/utils"

	"github.com/spf13/cobra"
)

const PortUrl = "127.0.0.1:5001"
const ClientId = "connected-app-live-d552cc32-a785-4371-bf85-0af85f5f7067"
const CodeChallenge = "dcvJOywUwb5HOWOyhmOI5dSc4_VHQU8Xkp9bXD-tGWI"
const CodeVerifier = "00afa9f459ce29bd4f9cd89f3a26036c3a2a772abd929e3fe179cb41"
const ProjectId = "project-live-0f74ccf8-79bd-4096-bd3f-5317c0e69a3b"

func NewAuthenticateCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "authenticate",
		Short: "Start authentication flow via Stytch",
		Run: func(cmd *cobra.Command, args []string) {
			http.HandleFunc("/", handleCallback)

			go func() {
				fmt.Printf("Listening on http://%s/\n", PortUrl)
				if err := http.ListenAndServe(PortUrl, nil); err != nil {
					log.Fatalf("Failed to start server: %v", err)
				}
			}()

			// Build the authentication URL
			authURL := fmt.Sprintf("https://ollie.dev.stytch.com/idp?response_type=code&client_id=%s&redirect_uri=http://%s&code_challenge=%s&code_verifier=%s", ClientId, PortUrl, CodeChallenge, CodeVerifier)

			// Open browser
			utils.OpenBrowser(authURL)

			// Keep the program running
			select {}
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
	fmt.Printf("access token: %s\n", accessToken)

	// Save the access token securely
	err := utils.SaveToken(accessToken)
	if err != nil {
		panic(err)
	}

	// Send 302 redirect to a friendly page (Stytch recommends redirecting away from localhost)
	http.Redirect(w, r, "https://stytch.com", http.StatusFound)

	// Optionally: shut down CLI here
	go func() {
		log.Println("Shutting down CLI after receiving code")
		os.Exit(0) // Uncomment if you want it to exit
	}()
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
		panic(err)
	}
	fmt.Printf("Request body: %s\n", string(bodyBytes))

	// Make the HTTP request
	req, err := http.NewRequest("POST", tokenUrl, bytes.NewBuffer(bodyBytes))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	// Read and print the response
	getAccessTokenResp := &GetAccessTokenResp{}
	err = json.NewDecoder(resp.Body).Decode(getAccessTokenResp)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Status: %s\n", resp.Status)
	return getAccessTokenResp.AccessToken
}
