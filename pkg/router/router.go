package router

import (
	"go-docker/internal/auth"
	"go-docker/internal/sample"
	"go-docker/internal/user"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")

	sampleGroup := api.Group("/sample")
	sampleHandler := sample.NewSampleHandler()
	{
		sampleGroup.GET("/helloworld", sampleHandler.HelloWorld)
		sampleGroup.GET("/users", sampleHandler.GetUsers)
		sampleGroup.POST("/user", sampleHandler.CreateUser)
	}

	authGroup := api.Group("/auth")
	authHandler := auth.NewAuthHandler()
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
	}

	userGroup := api.Group("/user")
	userHandler := user.NewUserHandler()
	{
		userGroup.GET("/", userHandler.GetUsers)
		userGroup.GET("/:id", userHandler.GetUserById)
	}

	return router
}