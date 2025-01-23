package user

import (
	"go-docker/pkg/utils"
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
	loginUserId, err := utils.StringToUint(c.GetString("userId"))
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}
	request := GetUserByIdRequest{}
	if err := c.ShouldBindUri(&request); err != nil {
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります。")
		return
	}

	userDetailResponse, err := userService.getUserByIdService(&request, loginUserId)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}
	utils.SuccessResponse[UserDetailResponse](c, http.StatusOK, *userDetailResponse, "ユーザー情報の取得に成功しました。")
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
	loginUserId, err := utils.StringToUint(c.GetString("userId"))
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}
	request := GetUserByUsernameRequest{}
	if err := c.ShouldBindUri(&request); err != nil {
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります。")
		return
	}

	userDetailResponse, err := userService.getUserByUsernameService(&request, loginUserId)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}
	utils.SuccessResponse[UserDetailResponse](c, http.StatusOK, *userDetailResponse, "ユーザー情報の取得に成功しました。")
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
	request := IsUniqueUsernameRequest{}
	if err := c.ShouldBindUri(&request); err != nil {
		errorMessage := "リクエストに不備があります。"
		// if validationErrors, ok := err.(validator.ValidationErrors); ok {
		// 	log.Printf("validationErrors: %v", validationErrors)
		// }
		// for _, e := range validationErrors {
		// 	log.Printf("e: %v", e)
		// 	if e.Field() == "Username" {
		// 		switch e.Tag() {
		// 			case "required":
		// 				if e.Field() == "Username" {
		// 					errorMessage = "ユーザー名は必須です。"
		// 				}
		// 			case "max":
		// 				if e.Field() == "Username" {
		// 					errorMessage = "ユーザー名は255文字以内で入力してください。"
		// 				}
		// 		}
		// 	}
		// }
		utils.ErrorResponse[any](c, http.StatusBadRequest, errorMessage)
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
		message = "使用できます"
	} else {
		message = "このユーザーネームは使用されています"
	}
	utils.SuccessResponse[IsUniqueUsernameResponse](c, http.StatusOK, IsUniqueUsernameResponse{IsUnique: isUnique, Message: message}, message)
}

// @Summary ログイン済みの場合、ログインユーザーの情報を取得
// @Description ヘッダーのトークンからロ図イン済みのユーザーを取得する
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
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}
	userDetailResponse, err := userService.getUserByIdService(&GetUserByIdRequest{UserId: *loginUserId}, loginUserId)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}
	utils.SuccessResponse[UserDetailResponse](c, http.StatusOK, *userDetailResponse, "ユーザー情報の取得に成功しました。")
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
// @Router /api/user/update/{userId} [put]
func (h *UserHandler) UpdateUser(c *gin.Context, ik *imagekit.ImageKit) {
	userId, requestBody, err := userService.validateUpdateUserRequest(c)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}

	userResponse, err := userService.updateUserService(ik, userId, requestBody)
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
