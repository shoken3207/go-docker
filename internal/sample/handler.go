package sample

import (
	"go-docker/internal/db"
	"go-docker/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SampleHandler struct{}

func (h *SampleHandler) HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!!!",
	})
}

func (h *SampleHandler)CreateUser(c *gin.Context) {
	var user models.User
	var request CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Name = request.Name
	user.Email = request.Email
	user.Pass_Hash = request.Password
	result := db.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *SampleHandler) GetUsers(c *gin.Context) {
	users := []models.User{}
	result := db.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func NewSampleHandler() *SampleHandler {
	return &SampleHandler{}
}