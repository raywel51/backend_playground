package handler

import (
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}

func WelcomeHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello from Gin!")
}

func IndexView(c *gin.Context) {
	osName := runtime.GOOS
	ginVersion := gin.Version

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":      "Hello Gin!",
		"ginVersion": ginVersion,
		"os":         osName,
	})
}
