package internal

import (
	"github.com/stytchauth/stytch-management-go/pkg/api"
)

var client *api.API

func SetDefaultMangoClient(c *api.API) {
	client = c
}

func GetDefaultMangoClient() *api.API {
	return client
}
