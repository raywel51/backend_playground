package main

import (
	"log"
	"net/http"
	"playground/routes"
)

func main() {
	address := "0.0.0.0:1661"
	log.Printf("Server is starting at http://%s\n", address)

	r := routes.SetupRouter()

	err := http.ListenAndServe(address, r)
	if err != nil {
		return
	}
}
