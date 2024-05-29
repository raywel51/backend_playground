package repository

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"os"

	"playground/infrastructure/persistence"
	"playground/internal/app/model/entity"
)

func GetVisitorInfo(visitorId int) (*entity.VisitorTypeDao, error) {
	client, err := persistence.ConnectToMongoDB()
	if err != nil {
		return nil, err
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {

		}
	}(client, context.Background())

	collection := client.Database(os.Getenv("DB_DATABASE")).Collection("visitor_type")

	var visitorTypeDao entity.VisitorTypeDao
	filter := bson.M{"visitor_type": visitorId}

	err = collection.FindOne(context.Background(), filter).Decode(&visitorTypeDao)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("visitor Type Not Found: %w", err)
		}
		return nil, fmt.Errorf("error fetching user: %w", err)
	}

	return &visitorTypeDao, nil
}
