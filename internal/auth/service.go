package auth

import (
	"errors"
	"go-docker/internal/db"
	"go-docker/models"
	"go-docker/pkg/constants"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

func (s *AuthService) findUserByEmail(email string) (*models.User, error) {
	user := models.User{}
	if err := db.DB.Select("id", "email", "pass_hash").Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *AuthService) createUser(newUser *models.User) ( error) {
	if err := db.DB.Create(&newUser).Error; err != nil {
		log.Printf("Error: %v", err)
		return  err
	}

	return  nil
}

func (s *AuthService) createEmailVerification(newEmailVerification *models.EmailVerification) ( error) {
	if err := db.DB.Create(&newEmailVerification).Error; err != nil {
		log.Printf("Error: %v", err)
		return  err
	}

	return  nil
}

func (s *AuthService) generateHashedPass(password string) (*string, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	stringPassHash := string(passHash)
	return &stringPassHash, nil
}

func (s *AuthService) comparePassword(requestPass string, savedPass string) (error) {
	if err := bcrypt.CompareHashAndPassword([]byte(savedPass), []byte(requestPass)); err != nil {
		log.Printf("パスワード比較エラー: %v", err)
		return err
	}
	return nil
}

func (s *AuthService) generateJwtToken(userId uint) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"jti": uuid.New().String(),
		"exp": time.Now().Add(constants.JwtTokenExpDate).Unix(),
	})
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Printf("SECRET_KEYが設定されていません。")
		return nil, errors.New("SECRET_KEYが設定されていません。")
	}
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Printf("jwtトークン生成エラー: %v", err)
		return nil, err
	}
	return &tokenString, nil
}

func NewAuthService() *AuthService {
	return &AuthService{}
}