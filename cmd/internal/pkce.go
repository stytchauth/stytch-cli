package internal

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

const LengthDefault = 32

func codeVerifier() string {
	b := make([]byte, LengthDefault)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return base64.RawURLEncoding.EncodeToString(b)
}

func codeChallenge(v string) string {
	h := sha256.Sum256([]byte(v))
	return base64.RawURLEncoding.EncodeToString(h[:])
}

func PKCECodePair() (verifier string, challenge string) {
	verifier = codeVerifier()
	challenge = codeChallenge(verifier)
	return verifier, challenge
}
