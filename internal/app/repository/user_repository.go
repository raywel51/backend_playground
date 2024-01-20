package repository

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"

	"playground/infrastructure/persistence"
	"playground/internal/app/model/entity"
)

func InsertUser(userDao *entity.UserDao) error {
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

	_, err = collection.InsertOne(context.Background(), userDao)
	if err != nil {
		log.Println("Error inserting book:", err)
		return err
	}

	return nil
}

func SelectOneUserByUsername(username string) (*entity.UserDao, error) {
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

	var user entity.UserDao
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

func SelectOneUserByEmail(email string) (*entity.UserDao, error) {
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

	var user entity.UserDao
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

func DeleteUserById(id string) error {
	client, err := persistence.ConnectToMongoDB()
	if err != nil {
		return err
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			// Handle disconnect error if needed
		}
	}(client, context.Background())

	collection := client.Database(os.Getenv("DB_DATABASE")).Collection("users")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

func GetAllUsers() ([]entity.UserDao, error) {
	client, err := persistence.ConnectToMongoDB()
	if err != nil {
		return nil, err
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			// Handle disconnect error if needed
		}
	}(client, context.Background())

	collection := client.Database(os.Getenv("DB_DATABASE")).Collection("users")

	var users []entity.UserDao

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error fetching users: %w", err)
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, context.Background())

	if err := cursor.All(context.Background(), &users); err != nil {
		return nil, fmt.Errorf("error decoding users: %w", err)
	}

	return users, nil
}

func GetUserByID(userID string) (entity.UserDao, error) {
	client, err := persistence.ConnectToMongoDB()
	if err != nil {
		return entity.UserDao{}, err
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			// Handle disconnect error if needed
		}
	}(client, context.Background())

	collection := client.Database(os.Getenv("DB_DATABASE")).Collection("users")

	var user entity.UserDao

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return entity.UserDao{}, fmt.Errorf("invalid user ID: %s", userID)
	}

	filter := bson.M{"_id": objID}

	err = collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return entity.UserDao{}, fmt.Errorf("user not found with ID: %s", userID)
		}
		return entity.UserDao{}, fmt.Errorf("error fetching user: %w", err)
	}

	return user, nil
}
