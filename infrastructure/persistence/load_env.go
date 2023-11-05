package persistence

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() error {
	ginMode := os.Getenv("GIN_MODE")
	fmt.Printf("GIN_MODE is set to: %s\n", ginMode)
	if ginMode == "debug" || ginMode == "" {
		if err := godotenv.Load(".env.development"); err != nil {
			return err
		}
	} else if ginMode == "release" {
		if err := godotenv.Load(".env.production"); err != nil {
			return err
		}
	} else {
		err := errors.New("cannot load envelopment please check it")
		if err != nil {
			return err
		}
	}

	return nil
}
