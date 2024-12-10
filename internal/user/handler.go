package user

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type CreateUserRequest struct {
	Name  string `json:"name" binding:"required,min=3,max=100"`
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age" binding:"required"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}


// @Summary ユーザー情報取得
// @Description userIdからユーザーを1人取得
// @Tags user
// @Param id path integer true "userId"
// @Success 200 {object} User "ユーザー情報"
// @Failure 400 {object} ErrorResponse "リクエストエラー"
// @Failure 404 {object} ErrorResponse "ユーザーが見つかりません"
// @Router /api/user/{id} [get]
func (h *UserHandler) GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		log.Println(id)
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "idが不正です", Message: "idは1以上の整数である必要があります。"})
		return
	}

	// for _, user := range users {
	// 	if user.Id == id {
	// 		c.JSON(http.StatusOK, user)
	// 		return
	// 	}
	// }

	c.JSON(http.StatusNotFound, ErrorResponse{Error: "ユーザーが見つかりません。", Message: "idが同じユーザーが存在しません。"})
}


// @Summary ユーザー情報変更
// @description userIdが同じユーザーの情報を変更する
// @Tags user
// @param id path uint true "userId"
// @success 200 {object} User "ユーザー情報"
// failure 400 {object} ErrorResponse "リクエストエラー"
// failure 404 {object} ErrorResponse "ユーザーが見つかりません"
// @router /api/user/update/:id [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {

}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}
