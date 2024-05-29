package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}

func WelcomeHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello from Gin!")
}

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": true})
}

func TokenCheckHandler(c *gin.Context) {
	userClaims, exists := c.Get("token")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User information not found"})
		return
	}

	exp, ok := userClaims.(jwt.MapClaims)["exp"].(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid expiration time"})
		return
	}

	// Convert the expiration time to a human-readable format
	expTime := time.Unix(int64(exp), 0)

	// Format time in a readable string
	normalTime := expTime.Format("2006-01-02 15:04:05 MST")

	c.JSON(http.StatusOK, gin.H{"status": true, "data": normalTime})
}

func GetStatus(c *gin.Context) {
	codeNumber := c.Param("code_number") // Get code_number parameter

	// Convert codeNumber to integer if needed
	codeNumberInt, err := strconv.Atoi(codeNumber)
	if err != nil {
		// Handle error if conversion fails (codeNumber is not a valid integer)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid code_number format"})
		return
	}

	// Respond with JSON (assuming codeNumberInt is the relevant status code)
	c.JSON(codeNumberInt, gin.H{"status": true})
}
