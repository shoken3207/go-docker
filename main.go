package main

import (
	"go-docker/internal/db"
	"go-docker/pkg/router"
	"go-docker/pkg/utils"
	"net/http"
	"os"
	"strings"

	// "github.com/swaggo/gin-swagger/swaggerFiles"
	_ "go-docker/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var swaggerUsername = os.Getenv("SWAGGER_USERNAME")
var swaggerPassword = os.Getenv("SWAGGER_PASSWORD")
func BasicAuthMiddleware() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		swaggerUsername: swaggerPassword,
	})
}

// @title ビジターゴーAPI
// @description このapiは、ビジターゴーのAPIで、ユーザー、スタジアム、遠征記録、などについて扱います。
// @version 1.0
// @accept json
// @produce json
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @host localhost:8080
// @BasePath /api
func main() {
	db.InitDB()
	ik := utils.NewImageKit()
	r := gin.Default()

    r.Use(func(c *gin.Context) {
		if strings.Contains(c.Request.Referer(), "/swagger/readonly/") {
			utils.ErrorResponse[any](c, http.StatusMethodNotAllowed, "読み取り専用モードでは実行できません。APIの実行には管理者モードでアクセスしてください。")
			c.Abort()
			return
		}
		c.Next()
	})

	router.SetupRouter(r, ik)

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, world!")
	})
    r.GET("/swagger/admin/*any", BasicAuthMiddleware(), ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/swagger/readonly/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
