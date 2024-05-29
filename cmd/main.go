package main

import (
	"fmt"
	"net/http"
	"os"

	"playground/infrastructure/persistence"
	"playground/router"
)

func main() {

	var err error
	err = persistence.LoadEnv()
	if err != nil {
		return
	}

	address := os.Getenv("HOST_PORT")
	fmt.Printf("Server is starting at http://%s\n\n", address)

	r := router.SetupRouter()

	err = http.ListenAndServe(address, r)
	if err != nil {
		return
	}
}
