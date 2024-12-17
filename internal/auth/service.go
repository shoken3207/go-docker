package auth

import (
	"errors"
	"go-docker/internal/db"
	"go-docker/models"
	"go-docker/pkg/utils"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
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
		log.Println("パスワード比較エラー:", err)
		return utils.NewCustomError(http.StatusBadRequest, "パスワードが違います。")
	}

	afterPassHash, err := authService.generateHashedPass(request.AfterPassword)
	if err != nil {
		log.Println("パスワードハッシュ化エラー:", err)
		return utils.NewCustomError(http.StatusInternalServerError, "ハッシュ化に失敗しました。")
	}
	user.PassHash = *afterPassHash

	if err := db.DB.Model(&user).Updates(models.User{PassHash: *afterPassHash}).Error; err != nil {
		log.Println("パスワード更新エラー:", err)
		return utils.NewCustomError(http.StatusInternalServerError, "パスワードの更新に失敗しました。")
	}

	return nil
}

func NewAuthService() *AuthService {
	return &AuthService{}
}
