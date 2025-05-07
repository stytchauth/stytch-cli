package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	PortUrl   = "127.0.0.1:5001"
	ClientId  = "connected-app-live-d552cc32-a785-4371-bf85-0af85f5f7067"
	ProjectId = "project-live-0f74ccf8-79bd-4096-bd3f-5317c0e69a3b"
)

type GetAccessTokenResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func GetAccessTokenFromRefreshToken(tok string) string {
	// make request to stytch with refresh token to get access token
	// store the access token/refresh token locally
	tokenUrl := fmt.Sprintf("https://api.stytch.com/v1/public/%s/oauth2/token", ProjectId)
	requestBody := map[string]interface{}{
		"client_id":     ClientId,
		"grant_type":    "refresh_token",
		"refresh_token": tok,
	}

	bodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalf("Unable to marshal refresh token request to JSON: %v", err)
	}

	// Make the HTTP request
	req, _ := http.NewRequest("POST", tokenUrl, bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error exchanging refresh token: %v", err)
	}
	defer resp.Body.Close()

	// Read and print the response
	var getAccessTokenResp GetAccessTokenResp
	if err := json.NewDecoder(resp.Body).Decode(&getAccessTokenResp); err != nil {
		log.Fatalf("Error decoding access token response: %v", err)
	}

	// save the access token and refresh token
	saveTokensFromResp(getAccessTokenResp)

	return getAccessTokenResp.AccessToken
}

func GetAccessTokenFromCode(code string, pkceVerifier string) GetAccessTokenResp {
	// make request to stytch with code to get access token
	// store the access token/refresh token locally
	tokenUrl := fmt.Sprintf("https://api.stytch.com/v1/public/%s/oauth2/token", ProjectId)
	requestBody := map[string]interface{}{
		"client_id":     ClientId,
		"redirect_uri":  fmt.Sprintf("http://%s", PortUrl),
		"grant_type":    "authorization_code",
		"code":          code,
		"code_verifier": pkceVerifier,
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

	return getAccessTokenResp
}

func saveTokensFromResp(accessTokenResp GetAccessTokenResp) {
	if err := SaveToken(accessTokenResp.AccessToken, AccessToken); err != nil {
		log.Fatalf("Unable to save access token: %v", err)
	}
	if err := SaveToken(accessTokenResp.RefreshToken, RefreshToken); err != nil {
		log.Fatalf("Unable to save refresh token: %v", err)
	}
}
