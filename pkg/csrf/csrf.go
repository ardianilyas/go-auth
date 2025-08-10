package csrf

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
)

func GenerateToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}

func ValidateToken(headerToken, cookieToken string) error {
	if headerToken == "" || cookieToken == "" {
		return errors.New("missing CSRF token")
	}

	if headerToken != cookieToken {
		return errors.New("CSRF token mismatch")
	}

	return nil
}