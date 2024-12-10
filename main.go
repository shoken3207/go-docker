package main

import (
	"go-docker/internal/db"
	"go-docker/pkg/router"
	"log"

	"github.com/joho/godotenv"

	// "github.com/swaggo/gin-swagger/swaggerFiles"
	_ "go-docker/docs"

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
	if err := godotenv.Load(".env"); err != nil {
        log.Printf("環境変数をロードできませんでした: %v", err)
    }
	db.InitDB()
	r := router.SetupRouter()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
