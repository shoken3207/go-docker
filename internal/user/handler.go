package user

import (
	"errors"
	"go-docker/pkg/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct{}

var userService = NewUserService()

// @Summary ユーザー情報取得
// @Description userIdからユーザーを1人取得
// @Tags user
// @Security BearerAuth
// @Param userId path integer true "userId"
// @Success 200 {object} utils.ApiResponse[UserResponse] "ユーザー情報"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 401 {object} utils.BasicResponse "認証エラー"
// @Failure 404 {object} utils.BasicResponse "not foundエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/user/{userId} [get]
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
	userResponse := UserResponse{
		Id:           user.ID,
		Email:        user.Email,
		Name:         user.Name,
		Description:  user.Description,
		ProfileImage: user.ProfileImage,
	}
	utils.SuccessResponse[UserResponse](c, http.StatusOK, userResponse, "ユーザー情報の取得に成功しました。")
}

// @Summary ログイン済みの場合、ログインユーザーの情報を取得
// @Description ヘッダーのトークンからユーザーを取得する
// @Tags user
// @Security BearerAuth
// @Success 200 {object} utils.ApiResponse[UserResponse] "成功"
// @Failure 401 {object} utils.BasicResponse "認証エラー"
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
	user, err := userService.findUserById(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.ErrorResponse[any](c, http.StatusNotFound, "ユーザーが見つかりません。")
		} else {
			utils.ErrorResponse[any](c, http.StatusInternalServerError, "内部エラーが発生しました。")
		}
		return
	}
	userResponse := UserResponse{
		Id:           user.ID,
		Email:        user.Email,
		Name:         user.Name,
		Description:  user.Description,
		ProfileImage: user.ProfileImage,
	}
	utils.SuccessResponse[UserResponse](c, http.StatusOK, userResponse, "ユーザー情報の取得に成功しました。")
}

// @Summary ユーザー情報変更
// @description userIdが同じユーザーの情報を変更する
// @Tags user
// @Security BearerAuth
// @param userId path uint true "userId"
// @param request body UpdateUserRequestBody true "userId"
// @success 200 {object} utils.ApiResponse[UserResponse] "ユーザー情報"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 401 {object} utils.BasicResponse "認証エラー"
// @Failure 404 {object} utils.BasicResponse "not foundエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @router /api/user/update/{userId} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}
