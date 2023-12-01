package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
