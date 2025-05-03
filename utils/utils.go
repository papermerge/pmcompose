package utils

import (
	"crypto/rand"
	"math/big"
	"strings"
)

func GenerateSecretString(length int) (string, error) {
	// Define character set for the random string
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"

	// Convert string to rune slice for easier manipulation
	runes := []rune(charset)

	// Create a new string builder
	var sb strings.Builder
	sb.Grow(length) // Pre-allocate space for better performance

	// Generate random characters
	for i := 0; i < length; i++ {
		// Generate a random index
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(runes))))
		if err != nil {
			return "", err
		}

		// Append the character at the random index
		sb.WriteRune(runes[idx.Int64()])
	}

	return sb.String(), nil
}
