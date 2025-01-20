package sample

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SampleHandler struct{}

// @Summary サンプルAPI
// @Description Hello Worldを返すだけのAPIです。
// @Tags sample
// @Router /api/admin/sample/helloWorld [get]
func (h *SampleHandler) PublicHelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!!!",
	})
}

// @Summary サンプルAPI
// @Description ログイン済みじゃないと実行できない、Hello Worldを返すだけのAPIです。
// @Tags sample
// @Security BearerAuth
// @Router /api/admin/sample/protectedHelloWorld [get]
func (h *SampleHandler) ProtectedHelloWorld(c *gin.Context) {
	userId := c.GetString("userId")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!!!" + userId,
	})
}

func NewSampleHandler() *SampleHandler {
	return &SampleHandler{}
}
