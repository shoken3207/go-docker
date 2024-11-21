package main

import (
	"fmt"
	"go-docker/backend/auth"
	"go-docker/backend/models"
	"go-docker/backend/user"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type CreateUserRequest struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func initDB() {
	dsn := "host=" + os.Getenv("DB_HOST") +
			"user=" + os.Getenv("DB_USER") +
			"password=" + os.Getenv("DB_PASSWORD") +
			"dbname=" + os.Getenv("DB_NAME") +
			"port=" + os.Getenv("DB_PORT") +
			"sslmode=disable"

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %V", err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %V", err)
	}
}

func main() {
	fmt.Println("hello world")
	initDB()
	r := gin.Default()
	r.GET("/helloworld", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World!!",
		})
	})
	r.POST("/users", func(c *gin.Context) {
		var user models.User
		var request CreateUserRequest
		if err := c.ShouldBindJSON(request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user.Name = request.Name
		user.Email = request.Email
		user.Pass_Hash = request.Password
		result := db.Create(&user)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, result)
	})

	r.GET("/users", func(c *gin.Context) {
		users := []models.User{}
		result := db.Find(&users)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, result)
	})


	{
		userRoute := r.Group("/api/user")
		userHandler := user.NewUserHandler()
		userRoute.GET("/", userHandler.GetUsers)
		userRoute.GET("/:id", userHandler.GetUserById)
	}
	{
		authRoute := r.Group("/api/auth")
		authHandler := auth.NewAuthHandler()
		authRoute.GET("/users", authHandler.GetUsers)
		authRoute.POST("/register", authHandler.Register)
		authRoute.POST("/login", authHandler.Login)
	}
	r.Run()
}
