package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"playground/internal/app/handler"
	register_access "playground/internal/app/handler/register-access"
	"playground/web/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	config := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	r.Use(cors.New(config))

	r.LoadHTMLGlob("templates/*")

	r.Static("/assets", "./assets")

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.File("./assets/ico/favicon.ico")
	})

	apiGroup := r.Group("api/v1")
	apiGroup.Use(middleware.LoggerMiddleware())

	apiGroup.GET("/", handler.WelcomeHandler)
	apiGroup.GET("/ping", handler.PingHandler)
	apiGroup.Any("/status/:code_number", handler.GetStatus)

	credentialGroup := apiGroup.Group("user")
	credentialGroup.POST("/login", handler.UserLogin)
	credentialGroup.POST("/register", handler.UserRegister)
	credentialGroup.GET("", handler.UserReadAll)
	credentialGroup.GET("/:id", handler.UserReadOneById)
	credentialGroup.DELETE("/:id", handler.UserDeleteById)

	genQrGroup := apiGroup.Group("gen-qr")
	genQrGroup.GET("/:key", handler.ReadQrCodeHandler)
	genQrGroup.POST("", handler.CreateQrCode)

	regQrGroup := apiGroup.Group("reg-qr")
	regQrGroup.POST("/", register_access.GenQr)
	regQrGroup.GET("/history", register_access.GetQrCode)
	regQrGroup.GET("/history-one/:key", register_access.GetQrCode)
	regQrGroup.GET("/history-between/:pages/:length", register_access.GetQrCode)

	parkingCal := apiGroup.Group("payment-parking")
	parkingCal.GET("/calculator-fee", register_access.GetQrCode)
	parkingCal.GET("/payment-fee", register_access.GetQrCode)

	token := r.Group("token")
	token.POST("/refresh", handler.RefreshHandler)

	tokenLock := apiGroup.Group("token/lock")
	tokenLock.Use(middleware.JwtMiddleware())
	tokenLock.GET("/check", handler.TokenCheckHandler)

	return r
}
