package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"playground/infrastructure/persistence"
	"playground/internal/app/handler"
	"playground/internal/app/model"
	"playground/web/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.File("./assets/ico/favicon.ico")
	})

	r.GET("/", handler.IndexView)

	apiGroup := r.Group("/api")
	apiGroup.Use(middleware.LoggerMiddleware())
	apiGroup.GET("/", handler.WelcomeHandler)
	apiGroup.GET("/hello", handler.HelloHandler)

	booksGroup := r.Group("/books")
	booksGroup.Use(middleware.LoggerMiddleware())

	booksGroup.POST("/", handler.CreateBooksHandler)
	booksGroup.GET("/", handler.CreateBooksHandler)

	// Get a single book by ID
	booksGroup.GET("/:id", func(c *gin.Context) {
		// Extract the book ID from the URL parameter
		bookID := c.Param("id")

		client, err := persistence.ConnectToMongoDB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to MongoDB"})
			return
		}
		defer client.Disconnect(context.Background())

		collection := client.Database("play_ground_go").Collection("books")

		var book model.Book // Replace YourBookType with the actual type

		// Query the book by its ID
		filter := bson.M{"_id": bookID}
		err = collection.FindOne(context.Background(), filter).Decode(&book)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}

		c.JSON(http.StatusOK, book)
	})

	// Update a book by ID
	booksGroup.PUT("/:id", func(c *gin.Context) {
		// Extract the book ID from the URL parameter
		bookID := c.Param("id")

		// Parse the request body into a Book struct
		var updatedBook model.Book // Replace YourBookType with the actual type
		if err := c.ShouldBindJSON(&updatedBook); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		client, err := persistence.ConnectToMongoDB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to MongoDB"})
			return
		}
		defer client.Disconnect(context.Background())

		collection := client.Database("play_ground_go").Collection("books")

		// Update the book by its ID
		filter := bson.M{"_id": bookID}
		update := bson.M{"$set": updatedBook}
		_, err = collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
			return
		}

		c.JSON(http.StatusOK, updatedBook)
	})

	// Delete a book by ID
	booksGroup.DELETE("/:id", func(c *gin.Context) {
		// Extract the book ID from the URL parameter
		bookID := c.Param("id")

		client, err := persistence.ConnectToMongoDB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to MongoDB"})
			return
		}
		defer client.Disconnect(context.Background())

		collection := client.Database("play_ground_go").Collection("books")

		// Delete the book by its ID
		filter := bson.M{"_id": bookID}
		_, err = collection.DeleteOne(context.Background(), filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
	})

	return r
}
