package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"

	"playground/infrastructure/persistence"
	"playground/internal/app/model"
)

func InsertImageQr(qrImg model.QrImgDao) error {
	client, err := persistence.ConnectToMongoDB()
	if err != nil {
		return err
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {

		}
	}(client, context.Background())

	collection := client.Database(os.Getenv("DB_DATABASE")).Collection("qr-image")

	_, err = collection.InsertOne(context.Background(), qrImg)
	if err != nil {
		log.Println("Error inserting book:", err)
		return err
	}

	return nil
}
