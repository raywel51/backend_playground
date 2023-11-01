package routes

import (
	"playground/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.File("./assets/ico/favicon.ico")
	})

	apiGroup := r.Group("/api")
	apiGroup.GET("/", handlers.WelcomeHandler)
	apiGroup.GET("/hello", handlers.HelloHandler)

	return r
}
