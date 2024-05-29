package credential_helper

import (
	"github.com/golang-jwt/jwt"
	"time"

	"playground/internal/app/model/request"
)

var SecretKey = []byte("Ur9ZMEoi9MmEGL")
var RefreshSecret = []byte("ebt8Cz8b2xnGvr")

func CreateToken(user request.CredentialLoginRequest, secretKey []byte, expires time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(expires).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
