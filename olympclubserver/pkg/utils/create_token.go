package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateSecureToken() string {
	b := make([]byte, 20)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
