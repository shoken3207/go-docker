package user

import (
	"go-docker/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

var userService = NewUserService()

// @Summary userIdからユーザー情報取得
// @Description userIdからユーザーを1人取得
// @Tags user
// @Security BearerAuth
// @Param userId path integer true "userId"
// @Success 200 {object} utils.ApiResponse[UserResponse] "ユーザー情報"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 401 {object} utils.BasicResponse "認証エラー"
// @Failure 404 {object} utils.BasicResponse "not foundエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/user/userId/{userId} [get]
func (h *UserHandler) GetUserById(c *gin.Context) {
	request := GetUserByIdRequest{}
	if err := c.ShouldBindUri(&request); err != nil {
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります。")
		return
	}

	userResponse, err := userService.getUserByIdService(&request)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}
	utils.SuccessResponse[UserResponse](c, http.StatusOK, *userResponse, "ユーザー情報の取得に成功しました。")
}

// @Summary usernameからユーザー情報取得
// @Description usernameからユーザーを1人取得
// @Tags user
// @Security BearerAuth
// @Param username path string true "username"
// @Success 200 {object} utils.ApiResponse[UserResponse] "ユーザー情報"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 401 {object} utils.BasicResponse "認証エラー"
// @Failure 404 {object} utils.BasicResponse "not foundエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/user/username/{username} [get]
func (h *UserHandler) GetUserByUsername(c *gin.Context) {
	request := GetUserByUsernameRequest{}
	if err := c.ShouldBindUri(&request); err != nil {
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります。")
		return
	}

	userResponse, err := userService.getUserByUsernameService(&request)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}
	utils.SuccessResponse[UserResponse](c, http.StatusOK, *userResponse, "ユーザー情報の取得に成功しました。")
}

// @Summary ユーザーネームの重複チェック
// @Description リクエストと同じuserNameが登録済みかチェックする
// @Tags user
// @Param username path string true "username"
// @Success 200 {object} utils.ApiResponse[IsUniqueUsernameResponse] "一意かのフラグ"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/user/isUnique/{username} [get]
func (h *UserHandler) IsUniqueUsername(c *gin.Context) {
	request := IsUniqueUsernameRequest{}
	if err := c.ShouldBindUri(&request); err != nil {
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります。")
		return
	}

	isUnique, err := utils.CheckUsernameUnique(request.Username)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok && customErr.Code != http.StatusNotFound {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Message)
		}
		return
	}
	var message string
	if isUnique {
		message = "一意なユーザーネームです。"
	} else {
		message = "ユーザーネームが被っています。"
	}
	utils.SuccessResponse[IsUniqueUsernameResponse](c, http.StatusOK, IsUniqueUsernameResponse{IsUnique: isUnique}, message)
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
	userId, err := utils.StringToUint(c.GetString("userId"))
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}
	userResponse, err := userService.getUserByIdService(&GetUserByIdRequest{UserId: *userId})
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}
	utils.SuccessResponse[UserResponse](c, http.StatusOK, *userResponse, "ユーザー情報の取得に成功しました。")
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
	userId, requestBody, err := userService.validateUpdateUserRequest(c)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}

	userResponse, err := userService.updateUserService(userId, requestBody)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}

	utils.SuccessResponse[UserResponse](c, http.StatusOK, *userResponse, "ユーザー情報の更新に成功しました")
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}
