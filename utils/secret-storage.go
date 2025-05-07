package utils

import "github.com/zalando/go-keyring"

const (
	service = "stytch-cli"
	user    = "stytch-user"
)

type TokenType string

const (
	AccessToken  TokenType = "access_token"
	RefreshToken TokenType = "refresh_token"
)

// SaveToken persists the token securely
func SaveToken(tok string, typ TokenType) error {
	return keyring.Set(service, user+string(typ), tok)
}

// LoadToken retrieves the token
func LoadToken(typ TokenType) (string, error) {
	return keyring.Get(service, user+string(typ))
}

// DeleteToken logs out
func DeleteToken(typ TokenType) error {
	return keyring.Delete(service, user+string(typ))
}
