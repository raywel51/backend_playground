package helper

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/google/uuid"
)

func GetHashing(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	hashedBytes := hash.Sum(nil)
	return hex.EncodeToString(hashedBytes)
}

func GetToken() string {
	randomString := uuid.New().String()
	return "Bearer " + randomString
}
