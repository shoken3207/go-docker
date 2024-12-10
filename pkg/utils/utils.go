package utils

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func SuccessResponse[T any](c *gin.Context, statusCode int, data T, message string) {
	c.JSON(statusCode, ApiResponse[T]{
		Success: true,
		Data: data,
		Message: message,
	})
}

func ErrorResponse[T any](c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, ApiResponse[T]{
		Success: false,
		Message: message,
	})
}


func GenerateRandomToken(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			ErrorResponse[any](c, http.StatusUnauthorized, "jwtトークンがありません。")
			c.Abort()
			return
		}
		secretKey := os.Getenv("SECRET_KEY")
		if secretKey == "" {
			ErrorResponse[any](c, http.StatusUnauthorized, "SECRET_KEYが設定されていません。")
			c.Abort()
			return
		}
		log.Print(tokenStr, secretKey)
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
			return []byte(secretKey), nil
		})
		if err != nil || !token.Valid {
			ErrorResponse[any](c, http.StatusUnauthorized, "トークンが不正な値です。")
			c.Abort()
			return
		}
		userID, ok := claims["userId"].(float64)
		if !ok {
			ErrorResponse[any](c, http.StatusUnauthorized, "トークンデータが不正な値です。")
			c.Abort()
			return
		}
		userIDUint := uint(userID)
		c.Set("userID", userIDUint)
		c.Next()
	}
}
