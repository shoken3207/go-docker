package auth

import (
	"errors"
	"go-docker/models"
	"go-docker/pkg/constants"
	"go-docker/pkg/utils"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthHandler struct{}

var authService = NewAuthService()

// @Summary メールアドレスの本人確認
// @Description リクエストからメールアドレス取得後、ユーザー登録されていないか確認し、メールアドレス宛に本登録URLをメールで送信
// @Tags auth
// @Param email path string true "メールアドレス"
// @Success 200 {object} utils.BasicResponse "成功"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/auth/emailVerified/{email} [get]
func (h *AuthHandler) EmailVerification(c *gin.Context) {
	request := EmailVerificationRequest{}
	if err := c.ShouldBindUri(&request); err != nil {
		utils.ErrorResponse[interface{}](c, http.StatusBadRequest, "リクエストに不備があります。")
		return
	}
	log.Printf(request.Email)
	user, err := authService.findUserByEmail(request.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ErrorResponse[interface{}](c, http.StatusInternalServerError, "内部エラーが発生しました。")
		return
	}
	log.Print("called")
	if user != nil {
		utils.ErrorResponse[interface{}](c, http.StatusConflict, "登録済みのメールアドレスです。")
		return
	}

	token := utils.GenerateRandomToken(32)
	EmailVerificationData := models.EmailVerification{
		Email: request.Email,
		Token: token,
		ExpiresAt: time.Now().Add(constants.EmailVerificationExpDate),
	}
	if err := authService.createEmailVerification(&EmailVerificationData); err != nil {
		utils.ErrorResponse[any](c, http.StatusInternalServerError, "内部エラーが発生しました。")
		return
	}

	// registerUrl := constants.RegisteBaserUrl + "/?token=" + token
	// authService.SendEmail(registerUrl)
	if err := authService.SendEmail(request.Email,"タイトル","本文"); err != nil {
		log.Printf("メール送信に失敗しました: %v\n", err)
		utils.ErrorResponse[any](c, http.StatusInternalServerError, "メール送信に失敗しました")
		return
	}
	utils.SuccessResponse[any](c, http.StatusOK, nil, "入力されたメールアドレス宛に本登録用URLを送信しました。")
}


// @Summary ユーザー登録
// @Description メールアドレス確認後にリクエスト内容をユーザーテーブルに保存
// @Tags auth
// @Param request body auth.RegisterRequest true "ユーザー情報"
// @Success 200 {object} utils.BasicResponse "成功"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var request RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c,http.StatusBadRequest, "リクエストに不備があります。")
		return
	}


	user, err := authService.findUserByEmail(request.Email)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("データベースエラー: %v", err)
			utils.ErrorResponse[any](c,http.StatusInternalServerError, "内部エラーが発生しました。")
			return;
		}
	}


	if user != nil {
		utils.ErrorResponse[any](c,http.StatusConflict, "登録済みのメールアドレスです。")
		return
	}
	passHash, err := authService.generateHashedPass(request.Password)
	if err != nil {
		utils.ErrorResponse[any](c,http.StatusInternalServerError,"内部エラーが発生しました。")
		return
	}

	newUser := models.User{Name: request.Name, Email: request.Email, PassHash: *passHash, Description: request.Description, ProfileImage: request.ProfileImage}

	if err := authService.createUser(&newUser); err != nil {
		utils.ErrorResponse[any](c,http.StatusInternalServerError,"内部エラーが発生しました。")
		return
	}
	utils.SuccessResponse[any](c, http.StatusOK, nil, "ユーザー登録に成功しました。")
}


// @Summary ログイン
// @Description メールアドレスとパスワードが合致したら、jwtトークンをCookieに保存
// @Tags auth
// @Param request body auth.LoginRequest true "ログイン情報"
// @Success 200 {object} utils.ApiResponse[LoginResponse] "成功"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 404 {object} utils.BasicResponse "not foundエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/auth/login [post]
func (h *AuthHandler)Login(c *gin.Context){
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[interface{}](c,http.StatusBadRequest, "リクエストに不備があります。")
		return;
	}

	user, err := authService.findUserByEmail(request.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.ErrorResponse[interface{}](c,http.StatusNotFound, "認証に失敗しました。")
		} else {
			utils.ErrorResponse[interface{}](c,http.StatusInternalServerError, "内部エラーが発生しました。")
		}
		return;
	}
	if  err = authService.comparePassword(request.Password, user.PassHash); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			utils.ErrorResponse[interface{}](c,http.StatusNotFound, "認証に失敗しました。")
		}else {
			utils.ErrorResponse[interface{}](c,http.StatusInternalServerError, "内部エラーが発生しました。")
		}
		return;
	}

	token, err := authService.generateJwtToken(user.ID)
	if err != nil {
		utils.ErrorResponse[interface{}](c,http.StatusInternalServerError, "内部エラーが発生しました。")
		return;
	}

	domain := os.Getenv("COOKIE_DOMAIN")
	if domain == "" {
		domain = "localhost"
	}

	utils.SuccessResponse[LoginResponse](c, http.StatusOK, LoginResponse{Token: *token}, "ログインに成功しました。")
}


// @Summary ログアウト状態からパスワードを変更
// @Description メール内リンクで本人確認後、トークンと新しいパスワードをリクエストで取得し、
// @Tags auth
// @Param request body auth.EmailVerificationRequest true "メールアドレス"
// @Success 200 {object} utils.BasicResponse "成功"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/auth/resetPass [put]
func (h *AuthHandler) ResetPass(c *gin.Context) {
}


// @Summary ログイン状態からパスワードを変更
// @Description 現在のパスワードと新しいパスワードをリクエストで取得し、現在のパスワードが合致したら、新しいパスワードに更新する
// @Tags auth
// @Security BearerAuth
// @Param request body auth.UpdatePassRequest true "メールアドレス"
// @Success 200 {object} utils.BasicResponse "成功"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 404 {object} utils.BasicResponse "リクエストエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/auth/updatePass [put]
func (h *AuthHandler) UpdatePass(c *gin.Context) {

}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}