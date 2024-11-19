package main

import (
	"go-docker/api/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/helloworld", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World!!!",
		})
	})

	userRoute := r.Group("/user")

	{
		userHandler := user.NewUserHandler()
		userRoute.GET("/", userHandler.GetUsers)
		userRoute.GET("/:id", userHandler.GetUserById)
	}
	r.Run()
}
