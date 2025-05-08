package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/stytchauth/stytch-cli/cmd/internal"
	"github.com/stytchauth/stytch-cli/utils"

	"github.com/spf13/cobra"
)

const (
	Scopes        = "openid email profile admin:projects manage:project_settings manage:api_keys"
	serverTimeout = 5 * time.Minute
)

func NewAuthenticateCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "authenticate",
		Short: "Start authentication flow via Stytch",
		Run: func(cmd *cobra.Command, args []string) {
			stop := make(chan struct{})
			timer := time.NewTimer(serverTimeout)

			// Generate PKCE pair.
			verifier, challenge := internal.PKCECodePair()

			mux := http.NewServeMux()
			mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				handleCallback(w, r, verifier)
				stop <- struct{}{}
			})

			server := &http.Server{
				Addr:              utils.PortUrl,
				Handler:           mux,
				ReadHeaderTimeout: 1 * time.Second,
			}

			// Build the authentication URL
			u, _ := url.Parse("https://" + internal.BaseURI + "/oauth/authorize")
			params := u.Query()
			params.Add("response_type", "code")
			params.Add("client_id", utils.ClientId)
			params.Add("redirect_uri", fmt.Sprintf("http://%s", utils.PortUrl))
			params.Add("code_verifier", verifier)
			params.Add("code_challenge", challenge)
			params.Add("scope", Scopes)
			u.RawQuery = params.Encode()

			go func() {
				fmt.Println("Starting authentication flow...")
				fmt.Println("\n⚠️ If you do not have an active Stytch session you will be prompted to login ⚠️")
				fmt.Printf("If this happens, complete the login and then visit the following link to continue the authentication flow:\n\n%s\n\n", u.String())
				if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
					fmt.Printf("Server error: %v\n", err)
					panic(err)
				}
			}()

			// Give the user enough time to see the warning message about having an active session
			// before opening the browser window.
			time.Sleep(time.Second * 3)

			// Open browser
			utils.OpenBrowser(u.String())

			// Keep the program running
			select {
			case <-stop:
			case <-timer.C:
				log.Fatal("Timed out waiting for authentication flow to complete")
			}

			// shut down the server
			if err := server.Shutdown(context.Background()); err != nil {
				log.Fatalf("Server shutdown failed: %v", err)
			}
			fmt.Println("Server shutdown successfully")
		},
	}

	return command
}

func handleCallback(w http.ResponseWriter, r *http.Request, pkceVerifier string) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Missing code parameter", http.StatusBadRequest)
		return
	}
	fmt.Printf("✅ Received code: %s\n", code)

	// Exchange the code for an access token & save it
	getAccessTokenResp := utils.GetAccessTokenFromCode(code, pkceVerifier)
	if getAccessTokenResp.AccessToken == "" {
		log.Fatalf("Failed to get access token")
	}
	err := utils.SaveToken(getAccessTokenResp.AccessToken, utils.AccessToken)
	if err != nil {
		log.Fatalf("Failed to save access token: %v", err)
	}

	// Send 302 redirect to a friendly page (Stytch recommends redirecting away from localhost)
	http.Redirect(w, r, "https://stytch.com", http.StatusFound)
}
