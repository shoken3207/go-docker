package sample

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SampleHandler struct{}

// @Summary サンプルAPI
// @Tags sample
// @Router /api/sample/helloWorld [get]
func (h *SampleHandler) PublicHelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!!!",
	})
}

// @Summary サンプルAPI
// @Tags sample
// @Security BearerAuth
// @Router /api/sample/protectedHelloWorld [get]
func (h *SampleHandler) ProtectedHelloWorld(c *gin.Context) {
	userId := c.GetString("userId")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!!!" + userId + "aa",
	})
}

func NewSampleHandler() *SampleHandler {
	return &SampleHandler{}
}