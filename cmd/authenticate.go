package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

const PORT = ":5001"

var authenticateCmd = &cobra.Command{
	Use:   "authenticate",
	Short: "Start authentication flow via Stytch",
	Run: func(cmd *cobra.Command, args []string) {
		// Start local server
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/" {
				handleCallback(w, r)
			} else {
				http.NotFound(w, r)
			}
		})

		go func() {
			fmt.Printf("Listening on http://127.0.0.1%s/\n", PORT)	
			if err := http.ListenAndServe("127.0.0.1"+PORT, nil); err != nil {
				log.Fatalf("Failed to start server: %v", err)
			}
		}()

		// Build the authentication URL (replace this with your actual URL from Stytch)
		authURL := "https://ollie.dev.stytch.com/idp?response_type=code&client_id=connected-app-live-d552cc32-a785-4371-bf85-0af85f5f7067&redirect_uri=http://127.0.0.1:5001&code_challenge=47DEQpj8HBSa-_TImW-5JCeuQeRkm5NMpJWZG3hSuFU&code_verifier=cbQ0SS1rHryO9V_xKa26myrXwamsWnBQw7NPp6SIu_Y"

		// Open browser
		openBrowser(authURL)
		fmt.Println("OPENING BROWSER")

		// Keep the program running
		select {}
	},
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handling callback for %s\n", r.URL.Path)
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Missing code parameter", http.StatusBadRequest)
		return
	}

	fmt.Printf("✅ Received code: %s\n", code)

	// make request to stytch with code to get access token
	// store the access token/refresh token locally

	// Send 302 redirect to a friendly page (Stytch recommends redirecting away from localhost)
	http.Redirect(w, r, "https://stytch.com/success", http.StatusFound)

	// Optionally: shut down CLI here
	go func() {
		log.Println("Shutting down CLI after receiving code")
		// os.Exit(0) // Uncomment if you want it to exit
	}()
}

func openBrowser(url string) {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "darwin":
		cmd = "open"
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", url}
	default: // Linux and others
		cmd = "xdg-open"
	}
	if cmd != "" {
		args = append([]string{url}, args...)
		exec.Command(cmd, args...).Start()
	}
}

func init() {
	rootCmd.AddCommand(authenticateCmd)
}
