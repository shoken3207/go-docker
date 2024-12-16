package utils

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"gopkg.in/gomail.v2"
)

func BoolPtr(b bool) *bool {
	return &b
}

func StringToUint(s string) (*uint, error) {
	u, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		log.Printf("変換エラー:", err)
		return nil, fmt.Errorf("stringからuintへの変換エラー")
	}
	parseValue := uint(u)
	return &parseValue, nil
}

func (e *CustomError) Error() string {
	return e.Message
}

func NewCustomError(code int, message string) error {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

func SuccessResponse[T any](c *gin.Context, statusCode int, data T, message string) {
	c.JSON(statusCode, ApiResponse[T]{
		Success: true,
		Data:    data,
		Message: message,
	})
}

func ErrorResponse[T any](c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, ApiResponse[T]{
		Success: false,
		Message: message,
	})
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

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			ErrorResponse[any](c, http.StatusUnauthorized, "jwtトークンがありません。")
			c.Abort()
			return
		}
		claims, err := ParseJWTToken(tokenStr)
		if err != nil {
			ErrorResponse[any](c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		userId, ok := claims["userId"].(float64)
		if !ok {
			ErrorResponse[any](c, http.StatusUnauthorized, "トークンデータが不正な値です。")
			c.Abort()
			return
		}
		userIdUint := uint(userId)
		c.Set("userId", fmt.Sprintf("%d", userIdUint))
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
	fromEmail := mail.NewEmail("ビジターGOサポートチーム", from)
	toEmail := mail.NewEmail("Recipient", to)
	message := mail.NewSingleEmail(fromEmail, subject, toEmail, body, "")
	response, err := client.Send(message)
	if err != nil {
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
	if env == "prod" {
		err = SendEmailProd(from, to, subject, body)
	} else {
		err = SendEmailDev(from, to, subject, body)
	}
	if err != nil {
		return err
	}

	return nil
}
