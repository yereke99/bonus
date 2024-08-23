package traits

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"
)

func Generate() int {
	rand.Seed(time.Now().UnixNano())

	number := rand.Intn(9000) + 1000
	return number
}

// generateQRToken generates a SHA-256 hash from the given user details
func GenerateQRToken(userName, userLastName, email string) string {
	// Concatenate the user details to form a unique string
	data := userName + userLastName + email

	// Create a new SHA-256 hash
	hash := sha256.New()
	hash.Write([]byte(data))

	// Convert the hash to a hexadecimal string
	return hex.EncodeToString(hash.Sum(nil))
}
