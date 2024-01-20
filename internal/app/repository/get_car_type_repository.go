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

func GetCarTypeInfo(carType int) (*entity.CarTypeDao, error) {
	client, err := persistence.ConnectToMongoDB()
	if err != nil {
		return nil, err
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {

		}
	}(client, context.Background())

	collection := client.Database(os.Getenv("DB_DATABASE")).Collection("car_type")

	var carTypeDao entity.CarTypeDao
	filter := bson.M{"car_type": carType}

	err = collection.FindOne(context.Background(), filter).Decode(&carTypeDao)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("visitor Type Not Found: %w", err)
		}
		return nil, fmt.Errorf("error fetching user: %w", err)
	}

	return &carTypeDao, nil
}
