package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// GenerateAPIToken generates a random API token and returns (plaintext, hash).
// The plaintext is shown to the user only once; the hash is stored in the database.
func GenerateAPIToken() (plaintext string, hash string, err error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", "", fmt.Errorf("generate random bytes: %w", err)
	}

	plaintext = fmt.Sprintf("goipam_%s", hex.EncodeToString(bytes))
	hash = HashAPIToken(plaintext)
	return plaintext, hash, nil
}

// HashAPIToken hashes an API token using SHA-256.
func HashAPIToken(token string) string {
	h := sha256.Sum256([]byte(token))
	return hex.EncodeToString(h[:])
}
