package internal

import (
	"fmt"
	"log"
	"sync"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/stytchauth/stytch-management-go/v2/pkg/api"

	"github.com/stytchauth/stytch-cli/utils"
)

var BaseURI   = "stytch.com"
var MangoClient = sync.OnceValue(func() *api.API {
	token, err := utils.LoadToken(utils.AccessToken)
	if err != nil {
		log.Fatal("Unable to load access token: ", err)
	}
	if tokenIsExpired(token) {
		token = utils.GetAccessTokenFromRefreshToken(token)
	}
	return api.NewAccessTokenClient(token, api.WithBaseURI("https://management"+BaseURI))
})

func tokenIsExpired(tok string) bool {
	token, _, err := new(jwt.Parser).ParseUnverified(tok, jwt.MapClaims{})
	if err != nil {
		panic(err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		panic("could not parse claims")
	}

	expFloat, ok := claims["exp"].(float64)
	if !ok {
		panic("exp claim is missing or not a number")
	}

	expTime := time.Unix(int64(expFloat), 0)
	fmt.Println("Token expires at:", expTime)

	return time.Now().After(expTime)
}
