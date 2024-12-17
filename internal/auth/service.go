package auth

import (
	"errors"
	"go-docker/internal/db"
	"go-docker/models"
	"go-docker/pkg/constants"
	"go-docker/pkg/utils"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct{}

func (s *AuthService) findUserByEmail(email string) (*models.User, error) {
	user := models.User{}
	if err := db.DB.Select("id", "email", "pass_hash").Where("email = ?", email).First(&user).Error; err != nil {
		log.Printf("ユーザーデータ取得エラー: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewCustomError(http.StatusNotFound, "ユーザーデータが見つかりませんでした。")
		} else {
			return nil, utils.NewCustomError(http.StatusInternalServerError, "ユーザーデータ取得に失敗しました。")
		}
	}

	return &user, nil
}

func (s *AuthService) createUser(newUser *models.User) error {
	if err := db.DB.Create(&newUser).Error; err != nil {
		log.Printf("ユーザーデータ追加エラー: %v", err)
		return  utils.NewCustomError(http.StatusInternalServerError, "ユーザーデータ追加に失敗しました。")
	}

	return nil
}

func (s *AuthService) generateHashedPass(password string) (*string, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("パスワードハッシュ化エラー: %v", err)
		return nil, utils.NewCustomError(http.StatusInternalServerError, "パスワードのハッシュ化に失敗しました。")
	}
	stringPassHash := string(passHash)
	return &stringPassHash, nil
}

func (s *AuthService) comparePassword(requestPass string, savedPass string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(savedPass), []byte(requestPass)); err != nil {
		log.Printf("パスワード比較エラー: %v", err)
		return utils.NewCustomError(http.StatusUnauthorized, "パスワードが一致しませんでした。")
	}
	return nil
}

func (s *AuthService) generateJwtToken(req TokenRequest, addExp time.Duration) (*string, error) {
	claims := jwt.MapClaims{
		"jti": uuid.New().String(),
		"exp": time.Now().Add(addExp).Unix(),
	}
	if req.UserID != nil {
		claims["userId"] = *req.UserID
	} else if req.Email != nil {
		claims["email"] = *req.Email
	} else {
		return nil, utils.NewCustomError(http.StatusBadRequest, "JWTトークン生成に必要なキーがありません。")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Printf("SECRET_KEYが設定されていません。")
		return nil, utils.NewCustomError(http.StatusInternalServerError, "SECRET_KEYが設定されていません。")
	}
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Printf("jwtトークン生成エラー: %v", err)
		return nil, utils.NewCustomError(http.StatusInternalServerError, "jwtトークンの生成に失敗しました。")
	}
	return &tokenString, nil
}

func (s *AuthService) updatePass(userId *uint, request *UpdatePassRequestBody) error {
	user := models.User{}
	if err := db.DB.Where("id = ?", &userId).First(&user).Error; err != nil {
		log.Println("ユーザー検索エラー: %v", err)
		return utils.NewCustomError(http.StatusNotFound, "ユーザー検索に失敗しました。")
	}

	if err := authService.comparePassword(request.BeforePassword, user.PassHash); err != nil {
		return err
	}

	afterPassHash, err := authService.generateHashedPass(request.AfterPassword)
	if err != nil {
		return err
	}
	user.PassHash = *afterPassHash

	if err := db.DB.Model(&user).Updates(models.User{PassHash: *afterPassHash}).Error; err != nil {
		log.Println("パスワード更新エラー:", err)
		return utils.NewCustomError(http.StatusInternalServerError, "パスワードの更新に失敗しました。")
	}

	return nil
}

func (s *AuthService) emailVerificationService(request *EmailVerificationRequest) error {
	user, err := authService.findUserByEmail(request.Email)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok && customErr.Code != http.StatusNotFound {
			return err
		}
	}
	if user != nil {
		return utils.NewCustomError(http.StatusConflict, "登録済みのメールアドレスです。")
	}

	token, err := authService.generateJwtToken(TokenRequest{Email: &request.Email}, constants.EmailVerificationTokenExpDate)
	if err != nil {
		return err
	}

	subject := "[ビジターGO] ユーザー登録"
	registerUrl := constants.RegisteBaserUrl + "/?token=" + *token
	body := `
[ビジターGO] メールアドレス確認のご案内とユーザー登録

こんにちは、

この度はビジターGOにご登録いただき、ありがとうございます。

以下のリンクから、ユーザー登録を完了させてください。

確認リンク:
` + registerUrl + `

※ 上記のリンクは、発行から30分以内にご利用ください。期限が過ぎると、再度新しいリンクをリクエストする必要があります。

もし、ご不明点がございましたら、お気軽にお問い合わせください。

どうぞよろしくお願いいたします。

ビジターGOサポートチーム
`

	if err := utils.SendEmail(request.Email, subject, body); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) registerService(request *RegisterRequest) error {
	claims, err := utils.ParseJWTToken(request.Token)
	if err != nil {
		return err
	}
	email, ok := claims["email"].(string)
	if !ok {
		return utils.NewCustomError(http.StatusUnauthorized, "トークンデータが不正な値です。")
	}
	user, err := authService.findUserByEmail(email)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok && customErr.Code != http.StatusNotFound {
			return err
		}
	}

	if user != nil {
		return utils.NewCustomError(http.StatusConflict, "登録済みのメールアドレスです。")
	}
	passHash, err := authService.generateHashedPass(request.Password)
	if err != nil {
		return err
	}

	newUser := models.User{Name: request.Name, Email: email, PassHash: *passHash, Description: request.Description, ProfileImage: request.ProfileImage}

	if err := authService.createUser(&newUser); err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "内部エラーが発生しました。")
	}

	return nil
}

func (s *AuthService) loginService(request *LoginRequest) (*string, error) {
	user, err := authService.findUserByEmail(request.Email)
	if err != nil {
		return nil, err
	}
	if err = authService.comparePassword(request.Password, user.PassHash); err != nil {
		return nil, err
	}

	token, err := authService.generateJwtToken(TokenRequest{UserID: &user.ID}, constants.LoginTokenExpDate)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *AuthService) validateUpdatePassRequest(c *gin.Context) (*uint, *UpdatePassRequestBody, error) {
	var requestBody UpdatePassRequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Printf("リクエストエラー: %v", err)
		return nil, nil, utils.NewCustomError(http.StatusBadRequest, "リクエストに不備があります。")
	}
	var requestPath UpdateUserRequestPath
	if err := c.ShouldBindUri(&requestPath); err != nil {
		log.Printf("リクエストエラー: %v", err)
		return nil, nil, utils.NewCustomError(http.StatusBadRequest, "リクエストに不備があります。")
	}
	userId, err := utils.StringToUint(c.GetString("userId"))
	if err != nil {
		return nil, nil, err
	}
	if *userId != requestPath.UserId {
		return nil, nil, utils.NewCustomError(http.StatusUnauthorized, "自分のパスワードしか更新できません。")
	}

	return userId, &requestBody, nil
}

func NewAuthService() *AuthService {
	return &AuthService{}
}
