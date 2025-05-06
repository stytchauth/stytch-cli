package internal

import (
	"log"

	"github.com/stytchauth/stytch-management-go/v2/pkg/api"

	"github.com/stytchauth/stytch-cli/utils"
)

var client *api.API

func SetDefaultMangoClient(c *api.API) {
	client = c
}

func GetDefaultMangoClient() *api.API {
	token, err := utils.LoadToken()
	if err != nil {
		log.Fatal("Unable to load access token: ", err)
	}
	return api.NewAccessTokenClient(token)
}
