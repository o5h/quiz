package token

import (
	"crypto/rand"
	"encoding/base64"
)

var (
	//TODO: move to env
	JWT_KEY           = []byte("your_secret_key")
	SUBJECT_TOKEN_KEY = []byte("your_secret_key")
)

// GenerateRandomKey generates a random key of specified length and encodes it in base64.
func GenerateRandomKey(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	randomKey := base64.StdEncoding.EncodeToString(randomBytes)
	return randomKey, nil
}
