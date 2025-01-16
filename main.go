package main

import (
	"go-docker/internal/db"
	"go-docker/pkg/router"
	"go-docker/pkg/utils"
	"log"
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

func customCorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			allowedOrigins := []string{
				os.Getenv("BASE_URL"),
				"http://localhost:3000",
				"http://localhost:8080",
				"capacitor://localhost",
				"ionic://localhost",
			}

			allowed := false
			for _, allowedOrigin := range allowedOrigins {
				if origin == allowedOrigin {
					allowed = true
					c.Header("Access-Control-Allow-Origin", origin)
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

    r.Use(func(c *gin.Context) {
        log.Println(c.Request.Referer())
		log.Println(c.Request.URL.Path)
		log.Println(strings.HasSuffix(c.Request.URL.Path, ".html"))
		log.Println(strings.HasSuffix(c.Request.URL.Path, ".css"))
		log.Println(strings.HasSuffix(c.Request.URL.Path, ".js"))
		log.Println(strings.HasSuffix(c.Request.URL.Path, ".json"))
		log.Println(strings.HasSuffix(c.Request.URL.Path, ".png"))
        if strings.HasPrefix(c.Request.URL.Path, "/swagger/readonly/") && 
           (strings.HasSuffix(c.Request.URL.Path, ".html") ||
            strings.HasSuffix(c.Request.URL.Path, ".css") ||
            strings.HasSuffix(c.Request.URL.Path, ".js") ||
            strings.HasSuffix(c.Request.URL.Path, ".json") ||
            strings.HasSuffix(c.Request.URL.Path, ".png")) {
            c.Next()
            return
        }
		log.Println("APIリクエストのみをブロック")
        if strings.Contains(c.Request.Referer(), "/swagger/readonly/") {
            utils.ErrorResponse[any](c, http.StatusMethodNotAllowed, "読み取り専用モードでは実行できません。")
            c.Abort()
            return
        }
        c.Next()
    })

	r.Use(customCorsMiddleware())

	router.SetupRouter(r, ik)

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, world!")
	})
    r.GET("/swagger/admin/*any", BasicAuthMiddleware(), ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/swagger/readonly/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
