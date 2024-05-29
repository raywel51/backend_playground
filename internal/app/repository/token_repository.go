package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"

	"playground/infrastructure/persistence"
	"playground/internal/app/model/entity"
)

func InsertOrUpdateToken(tokenDao *entity.TokenDao) error {
	client, err := persistence.ConnectToMongoDB()
	if err != nil {
		return err
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Println("Error disconnecting from MongoDB:", err)
		}
	}(client, context.Background())

	collection := client.Database(os.Getenv("DB_DATABASE")).Collection("token")

	// Check if the document with the given username already exists
	_, err = collection.FindOne(context.Background(), bson.M{"username": tokenDao.Username}).DecodeBytes()
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Println("Error querying database:", err)
			return err
		}
		// Document doesn't exist, insert a new one
		_, err = collection.InsertOne(context.Background(), tokenDao)
		if err != nil {
			log.Println("Error inserting token:", err)
			return err
		}
	} else {
		// Document exists, update the existing one
		update := bson.M{"$set": bson.M{"token": tokenDao.Token, "expiry": tokenDao.Expiry}}
		_, err := collection.UpdateOne(context.Background(), bson.M{"username": tokenDao.Username}, update)
		if err != nil {
			log.Println("Error updating token:", err)
			return err
		}
	}

	return nil
}
