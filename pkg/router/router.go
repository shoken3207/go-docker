package router

import (
	"go-docker/internal/adminTool"
	"go-docker/internal/auth"
	"go-docker/internal/expedition"
	"go-docker/internal/sample"
	"go-docker/internal/stadium"
	"go-docker/internal/team"
	"go-docker/internal/upload"
	"go-docker/internal/user"
	"go-docker/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/imagekit-developer/imagekit-go"
)

func SetupRouter(router *gin.Engine, ik *imagekit.ImageKit) *gin.Engine {
	api := router.Group("/api")

	sampleHandler := sample.NewSampleHandler()
	authHandler := auth.NewAuthHandler()
	userHandler := user.NewUserHandler()
	expeditionHandler := expedition.NewExpeditionHandler()
	uploadHandler := upload.NewUploadHandler()
	teamHandler := team.NewTeamHandler()
	stadiumHandler := stadium.NewStadiumHandler()
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

		publicTeamGroup := publicGroup.Group("/team")
		{
			publicTeamGroup.GET("/public", teamHandler.GetTeamsWithoutFavorites)
		}

		publicAdminGroup := publicGroup.Group("/admin")
		publicAdminStadiumGroup := publicAdminGroup.Group("/stadium")
		{
			publicAdminStadiumGroup.GET("/stadiums", adminToolHandler.GetStadiums)
			publicAdminStadiumGroup.GET("/idStadium/:id", adminToolHandler.GetIdStadiums)
			publicAdminStadiumGroup.POST("/stadiumAdd", adminToolHandler.StadiumAdd)
			publicAdminStadiumGroup.PUT("/update/:id", adminToolHandler.StadiumUpdate)
			publicAdminStadiumGroup.DELETE("/delete/:id", adminToolHandler.DeleteStadium)
		}

		publicAdminSportsGroup := publicAdminGroup.Group("/sports")
		{
			publicAdminSportsGroup.GET("/sports", adminToolHandler.GetSports)
			publicAdminSportsGroup.GET("/idSports/:id", adminToolHandler.GetIdSports)
			publicAdminSportsGroup.POST("/sportsAdd", adminToolHandler.SportsAdd)
			publicAdminSportsGroup.PUT("/update/:id", adminToolHandler.SportsUpdate)
			publicAdminSportsGroup.DELETE("/delete/:id", adminToolHandler.DeleteSports)
		}

		publicAdminLeagueGroup := publicAdminGroup.Group("/league")
		{
			publicAdminLeagueGroup.GET("/leagues", adminToolHandler.GetLeagues)
			publicAdminLeagueGroup.GET("/idLeague/:id", adminToolHandler.GetIdLeague)
			publicAdminLeagueGroup.POST("/leagueAdd", adminToolHandler.LeagueAdd)
			publicAdminLeagueGroup.PUT("/update/:id", adminToolHandler.LeagueUpdate)
			publicAdminLeagueGroup.DELETE("/delete/:id", adminToolHandler.DeleteLeague)
		}

		publicAdminTeamGroup := publicAdminGroup.Group("/team")
		{
			publicAdminTeamGroup.GET("/teams", adminToolHandler.GetTeams)
			publicAdminTeamGroup.GET("/idTeam/:id", adminToolHandler.GetIdTeam)
			publicAdminTeamGroup.POST("/teamAdd", adminToolHandler.TeamAdd)
			publicAdminTeamGroup.PUT("/update/:id", adminToolHandler.TeamUpdate)
			publicAdminTeamGroup.DELETE("/delete/:id", adminToolHandler.DeleteTeam)
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
			protectedUserGroup.PUT("/update/:userId", func(c *gin.Context) {
				userHandler.UpdateUser(c, ik)
			})
		}

		protectedExpeditionGroup := protectedGroup.Group("/expedition")
		{
			protectedExpeditionGroup.GET("/:expeditionId", expeditionHandler.GetExpeditionDetail)
			protectedExpeditionGroup.POST("/create", expeditionHandler.CreateExpedition)
			protectedExpeditionGroup.PUT("/update/:expeditionId", func(c *gin.Context) {
				expeditionHandler.UpdateExpedition(c, ik)
			})
			protectedExpeditionGroup.DELETE("/delete/:expeditionId", func(c *gin.Context) {
				expeditionHandler.DeleteExpedition(c, ik)
			})
			protectedExpeditionGroup.POST("/like/:expeditionId", expeditionHandler.LikeExpedition)
			protectedExpeditionGroup.GET("/list", expeditionHandler.GetExpeditionList)
			protectedExpeditionGroup.GET("/list/user", expeditionHandler.GetExpeditionListByUserId)
		}

		protectedTeamGroup := protectedGroup.Group("/team")
		{
			protectedTeamGroup.GET("/me", teamHandler.GetTeamsWithFavorites)
		}
		protectedStadiumGroup := protectedGroup.Group("/stadium")
		{
			protectedStadiumGroup.GET("/:stadiumId", stadiumHandler.GetStadium)
		}
	}

	return router
}
