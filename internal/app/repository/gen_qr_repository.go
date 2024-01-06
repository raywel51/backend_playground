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

func RegisterQrCode(genQrCodeDao *model.GenQrDao) error {
	client, err := persistence.ConnectToMongoDB()
	if err != nil {
		return err
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {

		}
	}(client, context.Background())

	collection := client.Database(os.Getenv("DB_DATABASE")).Collection("qr_code")

	_, err = collection.InsertOne(context.Background(), genQrCodeDao)
	if err != nil {
		log.Println("Error inserting book:", err)
		return err
	}

	return nil
}

func CheckPin(pinCode string) (*model.GenQrDao, error) {
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

	var qrCode model.GenQrDao
	filter := bson.M{"qr_key": pinCode}

	err = collection.FindOne(context.Background(), filter).Decode(&qrCode)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("PinCode Not Found: %w", err)
		}
		return nil, fmt.Errorf("error fetching user: %w", err)
	}

	return &qrCode, nil
}
