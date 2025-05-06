package utils

import "github.com/zalando/go-keyring"

const (
	service = "stytch-cli"
	user    = "stytch-user"
)

// SaveToken persists the token securely
func SaveToken(tok string) error {
	return keyring.Set(service, user, tok)
}

// LoadToken retrieves the token
func LoadToken() (string, error) {
	return keyring.Get(service, user)
}

// DeleteToken logs out
func DeleteToken() error {
	return keyring.Delete(service, user)
}
