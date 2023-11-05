package persistence

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongoDB() (*mongo.Client, error) {
	err := LoadEnv()
	if err != nil {
		return nil, err
	}
	clientOptions := options.Client().ApplyURI("mongodb://root:GjvwP6EFwWx2W7@raywel.ddns.net:27017/")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB")
	return client, nil
}
