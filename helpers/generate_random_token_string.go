package helpers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func GenerateRandomTokenString(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)

	if err != nil {
		return " ", fmt.Errorf("unable to generate the token")
	}
	return hex.EncodeToString(bytes), nil
}
