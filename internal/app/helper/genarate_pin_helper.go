package helper

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

func GeneratePIN() (string, error) {
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(9000))
	if err != nil {
		return strconv.Itoa(0), err
	}
	return strconv.Itoa(int(randomNumber.Int64()) + 1000), nil
}
