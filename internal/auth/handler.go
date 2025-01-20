package auth

import (
	"go-docker/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct{}

var authService = NewAuthService()

// @Summary メールアドレスの本人確認
// @Description リクエストからメールアドレス取得後、tokenTypeに応じてチェックし、メールアドレス宛にtokenを含めた画面URLをメールで送信
// @Tags auth
// @Param email query string true "メールアドレス"
// @Param tokenType query string true "トークンタイプ register or reset"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/admin/auth/emailVerification [get]
func (h *AuthHandler) EmailVerification(c *gin.Context) {
	request := EmailVerificationRequest{}
	if err := c.ShouldBindQuery(&request); err != nil {
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります。")
		return
	}

	if err := authService.emailVerificationService(&request); err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}

	utils.SuccessResponse[any](c, http.StatusOK, nil, "入力されたメールアドレス宛に本登録用URLを送信しました。")
}

// @Summary ユーザー登録
// @Description メールアドレス確認後にリクエスト内容をユーザーテーブルに保存
// @Tags auth
// @Param request body auth.RegisterRequest true "ユーザー情報"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/admin/auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var request RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります。")
		return
	}

	if err := authService.registerService(&request); err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}
	utils.SuccessResponse[any](c, http.StatusOK, nil, "ユーザー登録に成功しました。")
}

// @Summary ログイン
// @Description メールアドレスとパスワードが合致したら、jwtトークンをクライアントに返却
// @Tags auth
// @Param request body auth.LoginRequest true "ログイン情報"
// @Success 200 {object} utils.ApiResponse[LoginResponse] "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 404 {object} utils.ErrorBasicResponse "not foundエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/admin/auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります。")
		return
	}

	token, err := authService.loginService(&request)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}
	utils.SuccessResponse[LoginResponse](c, http.StatusOK, LoginResponse{Token: *token}, "ログインに成功しました。")
}

// @Summary ログアウト状態からパスワードを変更
// @Description メール内リンクで本人確認後、トークンと新しいパスワードをリクエストで取得し、パスワードを更新する
// @Tags auth
// @Param request body ResetPassRequest true "tokenと新しいパスワード"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 404 {object} utils.ErrorBasicResponse "not foundエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/admin/auth/resetPass [put]
func (h *AuthHandler) ResetPass(c *gin.Context) {
	var request ResetPassRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります。")
	}

	if err := authService.resetPassService(&request); err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}
	utils.SuccessResponse[any](c, http.StatusOK, nil, "パスワードのリセットに成功しました。")
}

// @Summary ログイン状態からパスワードを変更
// @Description 現在のパスワードと新しいパスワードをリクエストで取得し、現在のパスワードが合致したら、新しいパスワードに更新する
// @Tags auth
// @Security BearerAuth
// @param userId path uint true "userId"
// @Param request body UpdatePassRequestBody true "メールアドレス"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 401 {object} utils.ErrorBasicResponse "認証エラー"
// @Failure 404 {object} utils.ErrorBasicResponse "not foundエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/admin/auth/updatePass/{userId} [put]
func (h *AuthHandler) UpdatePass(c *gin.Context) {
	userId, requestBody, err := authService.validateUpdatePassRequest(c)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}

	if err := authService.updatePass(userId, requestBody); err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}

	utils.SuccessResponse[any](c, http.StatusOK, nil, "パスワードの更新に成功しました。")
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}
