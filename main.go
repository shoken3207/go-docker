package main

import (
	"go-docker/internal/db"
	"go-docker/pkg/router"
	"go-docker/pkg/utils"
	"log"
	"net/http"
	"net/url"
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

func checkAdminPermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		referer := c.Request.Referer()
		requestPath := c.Request.URL.Path
		readonlyPath := "/swagger/readonly/"
		if strings.Contains(requestPath, readonlyPath) {
			c.Next()
			return
		}

		if origin != "" && strings.Contains(origin, readonlyPath) {
			utils.ErrorResponse[any](c, http.StatusForbidden, "読み取り専用モードではAPIの実行はできません。")
			log.Printf("読み取り専用UIからのAPIリクエストをOriginでブロック: Origin=%s", origin)
			c.Abort()
			return
		}
        if origin == "" && strings.Contains(referer, readonlyPath) {
			utils.ErrorResponse[any](c, http.StatusForbidden, "読み取り専用モードではAPIの実行はできません。")
			log.Printf("読み取り専用UIからのAPIリクエストをRefererでブロック: Referer=%s", referer)
			c.Abort()
			return
		}
        c.Next()
	}
}

func customCorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Referer()
		if origin != "" {
			allowedOrigins := []string{
				os.Getenv("BASE_URL"),
				"http://localhost:3000",
				"http://127.0.0.1:5050",
				"http://localhost:8080",
				"capacitor://localhost",
				"ionic://localhost",
			}

			allowed := false

			parsedURL, err := url.Parse(origin)
			if err != nil {
				utils.ErrorResponse[any](c, http.StatusForbidden, "無効なオリジンです")
				c.Abort()
				return
			}
			originHost := parsedURL.Scheme + "://" + parsedURL.Host

			for _, allowedOrigin := range allowedOrigins {
				if originHost == allowedOrigin {
					allowed = true
					c.Header("Access-Control-Allow-Origin", originHost)
					c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
					c.Header("Access-Control-Allow-Headers", "*")
					c.Header("Access-Control-Allow-Credentials", "true")
					c.Header("Access-Control-Expose-Headers", "Content-Length, Authorization")
					c.Header("Access-Control-Max-Age", "43200")
					break
				}
			}

			if !allowed {
				utils.ErrorResponse[any](c, http.StatusForbidden, "このオリジンからのアクセスは許可されていません")
				c.Abort()
				return
			}
		}

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
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

    r.Use(checkAdminPermission())

	r.Use(customCorsMiddleware())

	router.SetupRouter(r, ik)

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, world!")
	})
    r.GET("/swagger/admin/*any", BasicAuthMiddleware(), ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/swagger/readonly/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
