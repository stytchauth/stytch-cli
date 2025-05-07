package internal

import (
	"log"
	"sync"

	"github.com/stytchauth/stytch-management-go/v2/pkg/api"

	"github.com/stytchauth/stytch-cli/utils"
)

var BaseUri = "ollie.dev.stytch.com"
var MangoClient = sync.OnceValue(func() *api.API {
	token, err := utils.LoadToken()
	if err != nil {
		log.Fatal("Unable to load access token: ", err)
	}
<<<<<<< HEAD
	return api.NewAccessTokenClient(token, api.WithBaseURI("https://management." + BaseUri))
=======
	return api.NewAccessTokenClient(token, api.WithBaseURI("https://management.staging.stytch.com"))
>>>>>>> main
})
