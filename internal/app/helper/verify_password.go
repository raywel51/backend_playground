package helper

import "fmt"

func VerifyPassword(inputPassword, storedPassword string) bool {
	key := GetHashing(inputPassword)
	fmt.Println(key)
	return key == storedPassword
}
