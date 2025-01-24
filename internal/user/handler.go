package user

import (
	"go-docker/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imagekit-developer/imagekit-go"
)

type UserHandler struct{}

var userService = NewUserService()

// @Summary userIdからユーザー情報取得
// @Description userIdからユーザーを1人取得
// @Tags user
// @Security BearerAuth
// @Param userId path integer true "userId"
// @Success 200 {object} utils.ApiResponse[UserDetailResponse] "ユーザー情報"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 401 {object} utils.ErrorBasicResponse "認証エラー"
// @Failure 404 {object} utils.ErrorBasicResponse "not foundエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/user/userId/{userId} [get]
func (h *UserHandler) GetUserById(c *gin.Context) {
	var requestPath GetUserByIdRequestPath
	loginUserId, err, customErr := utils.ValidateRequest(c, &requestPath, nil, nil, true)
	if err != nil {
		if customErr, ok := customErr.(*utils.CustomError); ok {
			utils.HandleCustomError(c, customErr, err, requestPath)
			return
		}
	}

	userDetailResponse, err := userService.getUserByIdService(&requestPath, loginUserId)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, utils.CreateSingleMessage(customErr.Error()))
			return
		}
	}
	log.Printf("userDetailResponse: %v" ,userDetailResponse.FavoriteTeams)
	utils.SuccessResponse[UserDetailResponse](c, http.StatusOK, *userDetailResponse, utils.CreateSingleMessage("ユーザー情報の取得に成功しました。"))
}

// @Summary usernameからユーザー情報取得
// @Description usernameからユーザーを1人取得
// @Tags user
// @Security BearerAuth
// @Param username path string true "username"
// @Success 200 {object} utils.ApiResponse[UserDetailResponse] "ユーザー情報"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 401 {object} utils.ErrorBasicResponse "認証エラー"
// @Failure 404 {object} utils.ErrorBasicResponse "not foundエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/user/username/{username} [get]
func (h *UserHandler) GetUserByUsername(c *gin.Context) {
	var requestPath GetUserByUsernameRequestPath
	loginUserId, err, customErr := utils.ValidateRequest(c, &requestPath, nil, nil, true)
	if err != nil {
		if customErr, ok := customErr.(*utils.CustomError); ok {
			utils.HandleCustomError(c, customErr, err, requestPath)
			return
		}
	}

	userDetailResponse, err := userService.getUserByUsernameService(&requestPath, loginUserId)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, utils.CreateSingleMessage(customErr.Error()))
			return
		}
	}
	utils.SuccessResponse[UserDetailResponse](c, http.StatusOK, *userDetailResponse, utils.CreateSingleMessage("ユーザー情報の取得に成功しました。"))
}

// @Summary ユーザーネームの重複チェック
// @Description リクエストと同じuserNameが登録済みかチェックする
// @Tags user
// @Param username path string true "username"
// @Success 200 {object} utils.ApiResponse[IsUniqueUsernameResponse] "一意かのフラグ"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/user/isUnique/{username} [get]
func (h *UserHandler) IsUniqueUsername(c *gin.Context) {
	requestPath := IsUniqueUsernameRequestPath{}
	_, err, customErr := utils.ValidateRequest(c, &requestPath, nil, nil, false)
	if err != nil {
		if customErr, ok := customErr.(*utils.CustomError); ok {
			utils.HandleCustomError(c, customErr, err, requestPath)
			return
		}
	}

	isUnique, err := utils.CheckUsernameUnique(requestPath.Username)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok && customErr.Code != http.StatusNotFound {
			utils.ErrorResponse[any](c, customErr.Code, utils.CreateSingleMessage(customErr.Message))
		}
		return
	}
	var message string
	if isUnique {
		message = "使用できます"
	} else {
		message = "このユーザーネームは使用されています"
	}
	utils.SuccessResponse[IsUniqueUsernameResponse](c, http.StatusOK, IsUniqueUsernameResponse{IsUnique: isUnique, Message: message}, utils.CreateSingleMessage(message))
}

// @Summary ログイン済みの場合、ログインユーザーの情報を取得
// @Description ヘッダーのトークンからログイン済みのユーザーを取得する
// @Tags user
// @Security BearerAuth
// @Success 200 {object} utils.ApiResponse[UserDetailResponse] "成功"
// @Failure 401 {object} utils.ErrorBasicResponse "認証エラー"
// @Failure 404 {object} utils.ErrorBasicResponse "not foundエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/user/logined [get]
func (h *UserHandler) GetMyData(c *gin.Context) {
	loginUserId, err := utils.StringToUint(c.GetString("userId"))
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, utils.CreateSingleMessage(customErr.Error()))
			return
		}
	}
	userDetailResponse, err := userService.getUserByIdService(&GetUserByIdRequestPath{UserId: *loginUserId}, loginUserId)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, utils.CreateSingleMessage(customErr.Error()))
			return
		}
	}
	utils.SuccessResponse[UserDetailResponse](c, http.StatusOK, *userDetailResponse, utils.CreateSingleMessage("ユーザー情報の取得に成功しました。"))
}

// @Summary ユーザー情報変更
// @description ユーザーの情報を変更する
// @Tags user
// @Security BearerAuth
// @param request body UpdateUserRequestBody true "更新データ"
// @success 200 {object} utils.ApiResponse[UserResponse] "ユーザー情報"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 401 {object} utils.ErrorBasicResponse "認証エラー"
// @Failure 404 {object} utils.ErrorBasicResponse "not foundエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/user/update [put]
func (h *UserHandler) UpdateUser(c *gin.Context, ik *imagekit.ImageKit) {
	var requestBody UpdateUserRequestBody
	loginUserId, err, customErr := utils.ValidateRequest(c, nil, nil, &requestBody, true)
	if err != nil {
		if customErr, ok := customErr.(*utils.CustomError); ok {
			utils.HandleCustomError(c, customErr, err, requestBody)
			return
		}
	}
	userResponse, err := userService.updateUserService(ik, loginUserId, &requestBody)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, utils.CreateSingleMessage(customErr.Error()))
			return
		}
	}

	utils.SuccessResponse[UserResponse](c, http.StatusOK, *userResponse, utils.CreateSingleMessage("ユーザー情報の更新に成功しました"))
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}
