package main

import (
	"auth-jwt/controllers"
	"auth-jwt/database"
	"auth-jwt/middlewares"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

var APPLICATION_NAME="My Simple JWT App"
var LOGIN_EXPIRATION_DURATION= time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD=jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY=[]byte("the secret of kalimdor")

func main()  {
	database.Connect("root:@tcp(localhost:3306)/otentikasi?parseTime=true")
	database.Migrate()
	router := initRouter()
	router.Run(":8099")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}