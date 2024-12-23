package router

import (
	"go-docker/internal/adminTool"
	"go-docker/internal/auth"
	"go-docker/internal/expedition"
	"go-docker/internal/sample"
	"go-docker/internal/upload"
	"go-docker/internal/user"
	"go-docker/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/imagekit-developer/imagekit-go"
)

func SetupRouter(ik *imagekit.ImageKit) *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")

	sampleHandler := sample.NewSampleHandler()
	authHandler := auth.NewAuthHandler()
	userHandler := user.NewUserHandler()
	expeditionHandler := expedition.NewExpeditionHandler()
	uploadHandler := upload.NewUploadHandler()
	adminToolHandler := adminTool.NewAdminToolHandler()

	publicGroup := api.Group("")
	{
		publicSampleGroup := publicGroup.Group("/sample")
		{
			publicSampleGroup.GET("/helloWorld", sampleHandler.ProtectedHelloWorld)
		}

		publicUserGroup := publicGroup.Group("/user")
		{
			publicUserGroup.GET("/isUnique/:username", userHandler.IsUniqueUsername)
		}

		publicAuthGroup := publicGroup.Group("/auth")
		{
			publicAuthGroup.GET("/emailVerification/:email", authHandler.EmailVerification)
			publicAuthGroup.POST("/register", authHandler.Register)
			publicAuthGroup.POST("/login", authHandler.Login)
			publicAuthGroup.PUT("/resetPass", authHandler.ResetPass)

		}

		// publicTeamGroup := publicGroup.Group("/teams")
		// {
		// 	publicTeamGroup.POST("/teamAdd", adminToolHandler.teamAdd)
		// }

		publicStadiumGroup := publicGroup.Group("/stadium")
		{
			publicStadiumGroup.POST("/stadiumAdd", adminToolHandler.StadiumAdd)
			publicStadiumGroup.PUT("/update", adminToolHandler.StadiumUpdate)
			publicStadiumGroup.DELETE("/delete", adminToolHandler.DeleteStadium)
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
			protectedAuthGroup.PUT("/updatePass/:userId", authHandler.UpdatePass)
		}

		protectedUserGroup := protectedGroup.Group("/user")
		{
			protectedUserGroup.GET("/userId/:userId", userHandler.GetUserById)
			protectedUserGroup.GET("/username/:username", userHandler.GetUserByUsername)
			protectedUserGroup.GET("/logined", userHandler.GetMyData)
			protectedUserGroup.PUT("/update/:userId", userHandler.UpdateUser)
		}

		protectedExpeditionGroup := protectedGroup.Group("/expedition")
		{
			protectedExpeditionGroup.POST("/create", expeditionHandler.CreateExpedition)
		}

		protectedUploadGroup := protectedGroup.Group("/upload")
		{
			protectedUploadGroup.POST("/images", func(c *gin.Context) {
				uploadHandler.UploadImages(c, ik)
			})
		}
	}

	return router
}
