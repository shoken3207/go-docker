package utils

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"gopkg.in/gomail.v2"
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



func SendEmailDev(from string, to string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	smtpHost := "mailhog"
	smtpPort := 1025
	d := gomail.NewDialer(smtpHost, smtpPort, "", "")
	if err := d.DialAndSend(m); err != nil {
		log.Printf("メール送信エラー: %v", err)
		return err
	}
	return nil
}

func SendEmailProd(from string, to string, subject string, body string) error {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("SENDGRID_API_KEY environment variable is not set")
	}
	client := sendgrid.NewSendClient(apiKey)
	log.Printf(body)
	fromEmail := mail.NewEmail("ビジターGOサポートチーム", from)
	toEmail := mail.NewEmail("Recipient", to)
	message := mail.NewSingleEmail(fromEmail, subject, toEmail, body, "")
	response, err := client.Send(message)
	if  err != nil {
		log.Printf("メール送信エラー: %v", err)
		return err
	} else {
		log.Println(response.StatusCode)
		log.Println(response.Body)
		log.Println(response.Headers)
	}

	return nil
}

func SendEmail(to string, subject string, body string) error {
	from := os.Getenv("FROM_EMAIL")
	env := os.Getenv("ENV")
	log.Printf(env)
	log.Printf(from)
	var err error
	if(env == "prod") {
		err = SendEmailProd(from, to, subject, body)
	}else {
		err = SendEmailDev(from, to, subject, body)
	}
	if err != nil {
		return err
	}

	return nil
}