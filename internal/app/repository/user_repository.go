package repository

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"

	"playground/infrastructure/persistence"
	"playground/internal/app/model"
)

func InsertUser(book *model.User) error {
	client, err := persistence.ConnectToMongoDB()
	if err != nil {
		return err
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {

		}
	}(client, context.Background())

	collection := client.Database(os.Getenv("DB_DATABASE")).Collection("users")

	_, err = collection.InsertOne(context.Background(), book)
	if err != nil {
		log.Println("Error inserting book:", err)
		return err
	}

	return nil
}

func SelectOneUserByUsername(username string) (*model.User, error) {
	client, err := persistence.ConnectToMongoDB()
	if err != nil {
		return nil, err
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {

		}
	}(client, context.Background())

	collection := client.Database(os.Getenv("DB_DATABASE")).Collection("users")

	var user model.User
	filter := bson.M{"username": username}

	err = collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("error fetching user: %w", err)
	}

	return &user, nil
}

func SelectOneUserByEmail(email string) (*model.User, error) {
	client, err := persistence.ConnectToMongoDB()
	if err != nil {
		return nil, err
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {

		}
	}(client, context.Background())

	collection := client.Database(os.Getenv("DB_DATABASE")).Collection("users")

	var user model.User
	filter := bson.M{"email": email}

	err = collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("error fetching user: %w", err)
	}

	return &user, nil
}
