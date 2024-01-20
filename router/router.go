package router

import (
	"github.com/gin-gonic/gin"

	"playground/internal/app/handler"
	"playground/internal/app/handler/register-access"
	"playground/web/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.Static("/assets", "./assets")

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
	genQrGroup.GET("/:key", handler.ReadQrCodeHandler)
	genQrGroup.POST("", handler.CreateQrCode)

	regQrGroup := r.Group("v2/reg-qr")
	regQrGroup.Use(middleware.LoggerMiddleware())
	regQrGroup.POST("", register_access.GenQr)
	regQrGroup.GET("/history", register_access.GetQrCode)
	regQrGroup.GET("/history-one/:key", register_access.GetQrCode)
	regQrGroup.GET("/history-between/:pages/:length", register_access.GetQrCode)

	parkingCal := r.Group("v2/payment-parking")
	parkingCal.Use(middleware.LoggerMiddleware())
	regQrGroup.GET("/calculator-fee", register_access.GetQrCode)
	regQrGroup.GET("/payment-fee", register_access.GetQrCode)

	token := r.Group("v2/token")
	token.Use(middleware.LoggerMiddleware())
	token.POST("/refresh", handler.RefreshHandler)

	tokenLock := r.Group("v2/token/lock")
	tokenLock.Use(middleware.LoggerMiddleware())
	tokenLock.Use(middleware.JwtMiddleware())
	tokenLock.GET("/check", handler.TokenCheckHandler)

	return r
}
