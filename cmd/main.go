package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"os"
	"playground/infrastructure/persistence"
	"playground/router"
)

var mongoClient *mongo.Client

func init() {
	var err error
	mongoClient, err = persistence.ConnectToMongoDB()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %s", err.Error())
	}
}

func main() {
	defer func(mongoClient *mongo.Client, ctx context.Context) {
		err := mongoClient.Disconnect(ctx)
		if err != nil {

		}
	}(mongoClient, context.Background())

	err := persistence.LoadEnv()
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
