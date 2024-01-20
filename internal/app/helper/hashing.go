package helper

import (
	"crypto/sha256"
	"encoding/hex"
)

func GetHashing(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	hashedBytes := hash.Sum(nil)
	return hex.EncodeToString(hashedBytes)
}
