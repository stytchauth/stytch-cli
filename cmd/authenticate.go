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
	Scopes  = "openid email profile admin:projects manage:project_settings manage:api_keys"
	BaseURI = "stytch.com"
)

func NewAuthenticateCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "authenticate",
		Short: "Start authentication flow via Stytch",
		Run: func(cmd *cobra.Command, args []string) {
			stop := make(chan struct{})

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
			u, _ := url.Parse("https://" + BaseURI + "/oauth/authorize")
			params := u.Query()
			params.Add("response_type", "code")
			params.Add("client_id", utils.ClientId)
			params.Add("redirect_uri", fmt.Sprintf("http://%s", utils.PortUrl))
			params.Add("code_verifier", verifier)
			params.Add("code_challenge", challenge)
			params.Add("scope", Scopes)
			u.RawQuery = params.Encode()

			go func() {
				fmt.Printf("Once authenticated visit: " + "https://stytch.com/oauth/authorize?client_id=connected-app-live-c48152cf-8732-4981-8fd5-e52dd989d75f&code_challenge=YdGskCH87fSqeu6R3tpLdxjl3RTY_z3sEnUO0jGiVD8&code_verifier=mJKzE9M9OQm4J_PuUQaJlyoYLlAMsrXVqH4y0ME7Nic&redirect_uri=http://127.0.0.1:5001&response_type=code&scope=openid+email+profile+admin:projects+manage:project_settings+manage:api_keys" + "\n")
				if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
					fmt.Printf("Server error: %v\n", err)
					panic(err)
				}
			}()

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

func handleCallback(w http.ResponseWriter, r *http.Request, pkceVerifier string) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Missing code parameter", http.StatusBadRequest)
		return
	}
	fmt.Printf("✅ Received code: %s\n", code)

	// Exchange the code for an access token & save it
	utils.GetAccessTokenFromCode(code, pkceVerifier)

	// Send 302 redirect to a friendly page (Stytch recommends redirecting away from localhost)
	http.Redirect(w, r, "https://stytch.com", http.StatusFound)
}
