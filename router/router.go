package router

import (
	"github.com/gin-gonic/gin"

	"playground/internal/app/handler"
	"playground/web/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.File("./assets/ico/favicon.ico")
	})

	r.GET("/", handler.WelcomeHandler)
	r.GET("/ping", handler.PingHandler)

	credentialGroup := r.Group("v2/user")
	credentialGroup.Use(middleware.LoggerMiddleware())
	credentialGroup.POST("/login", handler.UserLogin)
	credentialGroup.POST("/register", handler.UserRegister)
	credentialGroup.GET("", handler.UserReadAll)
	credentialGroup.GET("/:id", handler.UserReadOneById)
	credentialGroup.DELETE("/:id", handler.UserDeleteById)

	genQrGroup := r.Group("v2/gen-qr")
	genQrGroup.Use(middleware.LoggerMiddleware())
	genQrGroup.GET("/:key", handler.ReadQrCodeHandler)
	genQrGroup.POST("", handler.CreateQrCode)

	regQrGroup := r.Group("v2/reg-qr")
	regQrGroup.Use(middleware.LoggerMiddleware())
	regQrGroup.POST("", handler.GenQr)

	return r
}
