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
	"playground/internal/app/model/entity"
)

func InsertOneQrCode(genQrCodeDao *entity.GenQrDao) error {
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

func FindHistoryQrCode() ([]*entity.GenQrDao, error) {
	client, err := persistence.ConnectToMongoDB()
	if err != nil {
		return nil, err
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Println("Error disconnecting from MongoDB:", err)
		}
	}(client, context.Background())

	collection := client.Database(os.Getenv("DB_DATABASE")).Collection("qr_code")

	// Find all documents in the collection
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error finding documents:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Decode the results into a slice of GenQrDao
	var result []*entity.GenQrDao
	for cursor.Next(context.Background()) {
		var document entity.GenQrDao
		err := cursor.Decode(&document)
		if err != nil {
			log.Println("Error decoding document:", err)
			return nil, err
		}
		result = append(result, &document)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Error iterating through cursor:", err)
		return nil, err
	}

	return result, nil
}

func CheckPin(pinCode string) (*entity.GenQrDao, error) {
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

	var qrCode entity.GenQrDao
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
