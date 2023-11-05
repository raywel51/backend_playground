package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"playground/internal/app/model"
	"playground/internal/app/repository"
)

func CreateBooksHandler(c *gin.Context) {
	var book model.Book

	if c.ContentType() == "application/x-www-form-urlencoded" {
		if err := c.ShouldBindWith(&book, binding.Form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
			return
		}
	} else {
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
			return
		}
	}

	// Input validation (assuming Title is required and Publication is a positive integer)
	if book.Title == "" || book.Publication <= 0 {
		fmt.Print(book)
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "Title is required, and Publication must be a positive integer"})
		return
	}

	// Create an ObjectID for the book
	book.ID = primitive.NewObjectID()

	err := repository.InsertBooks(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book created successfully", "book": book})
}

//
//func GetBooksHandler(c *gin.Context) {
//	client, err := datastore.ConnectToMongoDB()
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to MongoDB", "err": err})
//		return
//	}
//	defer func(client *mongo.Client, ctx context.Context) {
//		err := client.Disconnect(ctx)
//		if err != nil {
//
//		}
//	}(client, context.Background())
//
//	var books []models.Book // Replace YourBookType with the actual type
//
//	// Query all books from the MongoDB collection
//	cursor := repository.GetAllBooks(c)
//
//	// Iterate through the cursor and decode documents into books
//	for cursor.Next(context.Background()) {
//		var book models.Book
//		if err := cursor.Decode(&book); err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode book"})
//			return
//		}
//		books = append(books, book)
//	}
//
//	c.JSON(http.StatusOK, books)
//}
