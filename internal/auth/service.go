package auth

import (
	"errors"
	"fmt"
	"go-docker/internal/db"
	"go-docker/models"
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

func (s *AuthService) createUser(newUser *models.User) error {
	if err := db.DB.Create(&newUser).Error; err != nil {
		log.Printf("Error: %v", err)
		return err
	}

	return nil
}

func (s *AuthService) createEmailVerification(newEmailVerification *models.EmailVerification) error {
	if err := db.DB.Create(&newEmailVerification).Error; err != nil {
		log.Printf("Error: %v", err)
		return err
	}

	return nil
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

func (s *AuthService) comparePassword(requestPass string, savedPass string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(savedPass), []byte(requestPass)); err != nil {
		log.Printf("パスワード比較エラー: %v", err)
		return err
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
		return nil, fmt.Errorf("JWTトークン生成に必要なキーがありません。")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
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

func ParseJWTToken(tokenStr string) (jwt.MapClaims, error) {
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return nil, fmt.Errorf("SECRET_KEYが設定されていません。")
	}
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
		return []byte(secretKey), nil
	})
	if err != nil || !token.Valid {
		log.Printf("jwtトークンパースエラー %v", err)
		return nil, fmt.Errorf("トークンが不正な値です。")
	}
	return claims, nil
}

func NewAuthService() *AuthService {
	return &AuthService{}
}
