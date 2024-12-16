package main

import (
	"go-docker/internal/db"
	"go-docker/pkg/router"
	"go-docker/pkg/utils"

	// "github.com/swaggo/gin-swagger/swaggerFiles"
	_ "go-docker/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

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
	r := router.SetupRouter(ik)
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, world!")
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
