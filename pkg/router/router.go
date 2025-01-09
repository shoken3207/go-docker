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
			publicAuthGroup.GET("/emailVerification", authHandler.EmailVerification)
			publicAuthGroup.POST("/register", authHandler.Register)
			publicAuthGroup.POST("/login", authHandler.Login)
			publicAuthGroup.PUT("/resetPass", authHandler.ResetPass)

		}

		publicUploadGroup := publicGroup.Group("/upload")
		{
			publicUploadGroup.POST("/images", func(c *gin.Context) {
				uploadHandler.UploadImages(c, ik)
			})
		}

		// publicTeamGroup := publicGroup.Group("/teams")
		// {
		// 	publicTeamGroup.POST("/teamAdd", adminToolHandler.teamAdd)
		// }

		publicStadiumGroup := publicGroup.Group("/stadium")
		{
			publicStadiumGroup.GET("/stadiums", adminToolHandler.GetStadiums)
			publicStadiumGroup.POST("/stadiumAdd", adminToolHandler.StadiumAdd)
			publicStadiumGroup.PUT("/update", adminToolHandler.StadiumUpdate)
			publicStadiumGroup.DELETE("/delete", adminToolHandler.DeleteStadium)
		}

		publicSportsGroup := publicGroup.Group("/sports")
		{
			publicSportsGroup.GET("/sports", adminToolHandler.GetSports)
			publicSportsGroup.POST("/sportsAdd", adminToolHandler.SportsAdd)
			publicSportsGroup.PUT("/update", adminToolHandler.SportsUpdate)
			publicSportsGroup.DELETE("/delete", adminToolHandler.DeleteSports)
		}

		publicLeagueGroup := publicGroup.Group("/league")
		{
			publicLeagueGroup.GET("/leagues", adminToolHandler.GetLeagues)
			publicLeagueGroup.POST("/leagueAdd", adminToolHandler.LeagueAdd)
			publicLeagueGroup.PUT("/update", adminToolHandler.LeagueUpdate)
			publicLeagueGroup.DELETE("/delete", adminToolHandler.DeleteLeague)
		}

		publicTeamGroup := publicGroup.Group("/team")
		{
			publicTeamGroup.GET("/teams", adminToolHandler.GetTeams)
			publicTeamGroup.POST("/teamAdd", adminToolHandler.TeamAdd)
			publicTeamGroup.PUT("/update", adminToolHandler.TeamUpdate)
			publicTeamGroup.DELETE("/delete", adminToolHandler.DeleteTeam)
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
			protectedExpeditionGroup.PUT("/update/:expeditionId", func(c *gin.Context) {
				expeditionHandler.UpdateExpedition(c, ik)
			})
			protectedExpeditionGroup.DELETE("/delete/:expeditionId", func(c *gin.Context) {
				expeditionHandler.DeleteExpedition(c, ik)
			})
			protectedExpeditionGroup.POST("/like/:expeditionId", func(c *gin.Context) {
				expeditionHandler.LikeExpedition(c)
			})
			protectedExpeditionGroup.DELETE("/unlike/:expeditionId", func(c *gin.Context) {
				expeditionHandler.UnlikeExpedition(c)
			})
		}
	}

	return router
}
