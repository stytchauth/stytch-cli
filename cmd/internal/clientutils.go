package internal

import (
	"log"
	"sync"

	"github.com/stytchauth/stytch-management-go/v2/pkg/api"

	"github.com/stytchauth/stytch-cli/utils"
)

var BaseURI   = "stytch.com"
var MangoClient = sync.OnceValue(func() *api.API {
	token, err := utils.LoadToken()
	if err != nil {
		log.Fatal("Unable to load access token: ", err)
	}
	return api.NewAccessTokenClient(token, api.WithBaseURI("https://management."+BaseURI))
})
