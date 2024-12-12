package user

import (
	"errors"
	"go-docker/models"
	"go-docker/pkg/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct{}

var userService = NewUserService()

// @Summary ログイン済みの場合、ログインユーザーの情報を取得
// @Description ヘッダーのトークンからユーザーを取得する
// @Tags user
// @Success 200 {object} utils.ApiResponse[LoginResponse] "成功"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 404 {object} utils.BasicResponse "not foundエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/user/logined [get]
func (h *UserHandler) GetMyData(c *gin.Context) {
	strUserId := c.GetString("userId")
	u, err := strconv.ParseUint(strUserId, 10, 64)
	if err != nil {
		log.Printf("変換エラー:", err)
		return
	}
	userId := uint(u)
	// userIDUint := uint(userId)
	// userId, err := utils.GetUserIdFromJWT(token)
	// if err != nil {
	// 	utils.ErrorResponse[any](c, http.StatusInternalServerError, "内部エラーが発生しました。")
	// 	return
	// }
	user, err := userService.findUserById(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.ErrorResponse[any](c, http.StatusNotFound, "ユーザーが見つかりません。")
		} else {
			utils.ErrorResponse[any](c, http.StatusInternalServerError, "内部エラーが発生しました。")
		}
		return
	}
	utils.SuccessResponse[models.User](c, http.StatusOK, *user, "")
}

// @Summary ユーザー情報取得
// @Description pathのuserIdからユーザーを1人取得
// @Tags user
// @Param id path integer true "userId"
// @Success 200 {object} utils.ApiResponse[LoginResponse] "成功"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 404 {object} utils.BasicResponse "not foundエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/user/{id} [get]
func (h *UserHandler) GetUserById(c *gin.Context) {
	request := GetUserByIdRequest{}
	if err := c.ShouldBindUri(&request); err != nil {
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります。")
		return
	}
	user, err := userService.findUserById(request.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.ErrorResponse[any](c, http.StatusNotFound, "認証に失敗しました。")
		} else {
			utils.ErrorResponse[any](c, http.StatusInternalServerError, "内部エラーが発生しました。")
		}
		return
	}
	utils.SuccessResponse[models.User](c, http.StatusOK, *user, "")
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}
