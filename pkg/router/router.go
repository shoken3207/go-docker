package router

import (
	"go-docker/internal/auth"
	"go-docker/internal/expedition"
	"go-docker/internal/sample"
	"go-docker/internal/user"
	"go-docker/pkg/utils"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")

	sampleHandler := sample.NewSampleHandler()
	authHandler := auth.NewAuthHandler()
	userHandler := user.NewUserHandler()
	expeditionHandler := expedition.NewExpeditionHandler()

	publicGroup := api.Group("")
	{
		publicSampleGroup := publicGroup.Group("/sample")
		{
			publicSampleGroup.GET("/helloWorld", sampleHandler.ProtectedHelloWorld)
		}

		publicAuthGroup := publicGroup.Group("/auth")
		{
			publicAuthGroup.GET("/emailVerified/:email", authHandler.EmailVerification)
			publicAuthGroup.POST("/register", authHandler.Register)
			publicAuthGroup.POST("/login", authHandler.Login)
			publicAuthGroup.PUT("/resetPass", authHandler.ResetPass)

		}
	}


	protectedGroup := api.Group("")
	protectedGroup.Use(utils.AuthMiddleware())
	{
		protectedSampleGroup := protectedGroup.Group("/sample")
		{
			protectedSampleGroup.GET("/protectedHelloWorld", sampleHandler.ProtectedHelloWorld)
		}

		protectedAuthGroup := protectedGroup.Group("/auth")
		{
			protectedAuthGroup.PUT("/updatePass", authHandler.UpdatePass)
		}

		protectedUserGroup := protectedGroup.Group("/user")
		{
			protectedUserGroup.GET("/:id", userHandler.GetUserById)
		}

		protectedExpeditionGroup := protectedGroup.Group("/expedition")
		{
			protectedExpeditionGroup.POST("/create", expeditionHandler.CreateExpedition)
		}
	}




	return router
}