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
// @Description リクエストからメールアドレス取得後、tokenTypeに応じてチェックし、メールアドレス宛にtokenを含めた画面URLをメールで送信<br>ユーザー登録、パスワードリセット時に使います。<br>
// @Tags auth
// @Param email query string true "メールアドレス"
// @Param tokenType query string true "トークンタイプ register or reset"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/auth/emailVerification [get]
func (h *AuthHandler) EmailVerification(c *gin.Context) {
	var requestQuery EmailVerificationRequestQuery
	_, err, customErr := utils.ValidateRequest(c, nil, &requestQuery, nil, false)
	log.Printf("err: %v", err)
	if err != nil {
		if customErr, ok := customErr.(*utils.CustomError); ok {
			utils.HandleCustomError(c, customErr, err, requestQuery)
			return
		}
	}
	log.Println(requestQuery)

	if err := authService.emailVerificationService(&requestQuery); err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, utils.CreateSingleMessage(customErr.Error()))
			return
		}
	}

	utils.SuccessResponse[any](c, http.StatusOK, nil, utils.CreateSingleMessage("入力されたメールアドレス宛に本登録用URLを送信しました。"))
}

// @Summary ユーザー登録
// @Description メール内リンクから遷移できる本登録用画面からリクエスト内容を取得し、ユーザーテーブルに保存
// @Tags auth
// @Param request body auth.RegisterRequestBody true "ユーザー情報"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var requestBody RegisterRequestBody
	_, err, customErr := utils.ValidateRequest(c, nil, nil, &requestBody, false)
	if err != nil {
		if customErr, ok := customErr.(*utils.CustomError); ok {
			utils.HandleCustomError(c, customErr, err, requestBody)
			return
		}
	}

	if err := authService.registerService(&requestBody); err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, utils.CreateSingleMessage(customErr.Error()))
			return
		}
	}
	utils.SuccessResponse[any](c, http.StatusOK, nil, utils.CreateSingleMessage("ユーザー登録に成功しました。"))
}

// @Summary ログイン
// @Description メールアドレスとパスワードが合致したら、jwtトークンをクライアントに返却
// @Tags auth
// @Param request body auth.LoginRequestBody true "ログイン情報"
// @Success 200 {object} utils.ApiResponse[LoginResponse] "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 404 {object} utils.ErrorBasicResponse "not foundエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var requestBody LoginRequestBody
	_, err, customErr := utils.ValidateRequest(c, nil, nil, &requestBody, false)
	if err != nil {
		if customErr, ok := customErr.(*utils.CustomError); ok {
			utils.HandleCustomError(c, customErr, err, requestBody)
			return
		}
	}

	token, err := authService.loginService(&requestBody)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, utils.CreateSingleMessage(customErr.Error()))
			return
		}
	}
	utils.SuccessResponse[LoginResponse](c, http.StatusOK, LoginResponse{Token: *token}, utils.CreateSingleMessage("ログインに成功しました。"))
}

// @Summary ログアウト状態からパスワードを変更
// @Description メール内リンクから遷移できるパスワードリセット画面から、トークンと新しいパスワードをリクエストで取得し、パスワードを更新する
// @Tags auth
// @Param request body ResetPassRequestBody true "tokenと新しいパスワード"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 404 {object} utils.ErrorBasicResponse "not foundエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/auth/resetPass [put]
func (h *AuthHandler) ResetPass(c *gin.Context) {
	var requestBody ResetPassRequestBody
	_, err, customErr := utils.ValidateRequest(c, nil, nil, &requestBody, false)
	if err != nil {
		if customErr, ok := customErr.(*utils.CustomError); ok {
			utils.HandleCustomError(c, customErr, err, requestBody)
			return
		}
	}

	if err := authService.resetPassService(&requestBody); err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, utils.CreateSingleMessage(customErr.Error()))
			return
		}
	}
	utils.SuccessResponse[any](c, http.StatusOK, nil, utils.CreateSingleMessage("パスワードのリセットに成功しました。"))
}

// @Summary ログイン状態からパスワードを変更
// @Description 現在のパスワードと新しいパスワードをリクエストで取得し、現在のパスワードが合致したら、新しいパスワードに更新する
// @Tags auth
// @Security BearerAuth
// @Param request body UpdatePassRequestBody true "現在のパスワードと新しいパスワード"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 401 {object} utils.ErrorBasicResponse "認証エラー"
// @Failure 404 {object} utils.ErrorBasicResponse "not foundエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/auth/updatePass [put]
func (h *AuthHandler) UpdatePass(c *gin.Context) {
	var requestBody UpdatePassRequestBody
	loginUserId, err, customErr := utils.ValidateRequest(c, nil, nil, &requestBody, true)
	if err != nil {
		if customErr, ok := customErr.(*utils.CustomError); ok {
			utils.HandleCustomError(c, customErr, err, requestBody)
			return
		}
	}

	if err := authService.updatePass(loginUserId, &requestBody); err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, utils.CreateSingleMessage(customErr.Error()))
			return
		}
	}

	utils.SuccessResponse[any](c, http.StatusOK, nil, utils.CreateSingleMessage("パスワードの更新に成功しました。"))
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}
