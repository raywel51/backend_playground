package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	"playground/infrastructure/persistence"
	"playground/internal/app/model"
)

func GetAllBooks(model *model.Book) (*mongo.Cursor, error) {
	client, err := persistence.ConnectToMongoDB()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("play_ground_go").Collection("books")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	return cursor, nil
}

func InsertBooks(book *model.Book) error {
	client, err := persistence.ConnectToMongoDB()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(os.Getenv("DB_DATABASE")).Collection("books")

	_, err = collection.InsertOne(context.Background(), book)
	if err != nil {
		log.Println("Error inserting book:", err)
		return err
	}

	return nil
}
