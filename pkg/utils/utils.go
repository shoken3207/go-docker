package utils

import (
	"context"
	"errors"
	"fmt"
	"go-docker/internal/db"
	"go-docker/models"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/imagekit-developer/imagekit-go"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

func NewCustomError(code int, message string) error {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

func (e *CustomError) Error() string {
	return e.Message
}

func BoolPtr(b bool) *bool {
	return &b
}

func StringToUint(s string) (*uint, error) {
	u, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		log.Printf("変換エラー: %v", err)
		return nil, NewCustomError(http.StatusInternalServerError, "stringからuintへの変換エラー")
	}
	parseValue := uint(u)
	return &parseValue, nil
}

func SuccessResponse[T any](c *gin.Context, statusCode int, data T, messages []string) {
	c.JSON(statusCode, ApiResponse[T]{
		Success: true,
		Data:    data,
		Messages: messages,
	})
}

func ErrorResponse[T any](c *gin.Context, statusCode int, messages []string) {
	c.JSON(statusCode, ApiResponse[T]{
		Success: false,
		Messages: messages,
	})
}

func CreateSingleMessage (message string) []string {
	return []string{message}
}

func ParseJWTToken(tokenStr string) (jwt.MapClaims, error) {
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return nil, NewCustomError(http.StatusInternalServerError, "SECRET_KEYが設定されていません。")
	}
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
		return []byte(secretKey), nil
	})
	if err != nil || !token.Valid {
		log.Printf("jwtトークンパースエラー %v", err)
		return nil, NewCustomError(http.StatusUnauthorized, "トークンが不正な値です。")
	}
	return claims, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			ErrorResponse[any](c, http.StatusUnauthorized, CreateSingleMessage("jwtトークンがありません。"))
			c.Abort()
			return
		}
		claims, err := ParseJWTToken(tokenStr)
		if err != nil {
			if customErr, ok := err.(*CustomError); ok {
				ErrorResponse[any](c, customErr.Code, CreateSingleMessage(customErr.Error()))
				c.Abort()
				return
			}
		}

		userId, ok := claims["userId"].(float64)
		if !ok {
			ErrorResponse[any](c, http.StatusUnauthorized, CreateSingleMessage("トークンデータが不正な値です。"))
			c.Abort()
			return
		}
		userIdUint := uint(userId)
		c.Set("userId", fmt.Sprintf("%d", userIdUint))
		c.Next()
	}
}

func DeleteUploadImage(ik *imagekit.ImageKit, fileId *string) error {
	ctx := context.Background()
	_, err := ik.Media.DeleteFile(ctx, *fileId)
	if err != nil {
		log.Printf("画像削除エラー: %v", err)
		return NewCustomError(http.StatusInternalServerError, "アップロード画像の削除に失敗しました。")
	}

	return nil
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
		return NewCustomError(http.StatusInternalServerError, "メール送信に失敗しました。")
	}
	log.Printf("success")
	return nil
}

func SendEmailProd(from string, to string, subject string, body string) error {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	if apiKey == "" {
		return NewCustomError(http.StatusInternalServerError, "SENDGRIDのAPI_KEYが設定されていません。")
	}
	client := sendgrid.NewSendClient(apiKey)
	fromEmail := mail.NewEmail("ビジターGOサポートチーム", from)
	toEmail := mail.NewEmail("Recipient", to)
	message := mail.NewSingleEmail(fromEmail, subject, toEmail, body, "")
	response, err := client.Send(message)
	if err != nil {
		log.Printf("メール送信エラー: %v", err)
		return NewCustomError(http.StatusInternalServerError, "メール送信に失敗しました。")
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

func FindUserByUsername(username string) (*models.User, error) {
	user := models.User{}
	if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		log.Printf("ユーザーデータ取得エラー: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, NewCustomError(http.StatusNotFound, "ユーザーデータが見つかりませんでした。")
		} else {
			return nil, NewCustomError(http.StatusInternalServerError, "ユーザーデータ取得に失敗しました。")
		}
	}

	return &user, nil
}

func CheckUsernameUnique(username string) (bool, error) {
	user, err := FindUserByUsername(username)
	if err != nil {
		if customErr, ok := err.(*CustomError); ok && customErr.Code != http.StatusNotFound {
			return false, err
		}
	}

	if user != nil {
		return false, nil
	}

	return true, nil
}

func ValidateAndPersistImages(tx *gorm.DB, imageUrls []string) ([]models.TempImage, error) {
	if len(imageUrls) == 0 {
		return []models.TempImage{}, nil
	}

	var tempImages []models.TempImage
	if err := tx.Where("image IN ?", imageUrls).Find(&tempImages).Error; err != nil {
		return []models.TempImage{}, NewCustomError(http.StatusInternalServerError, "画像情報の取得に失敗しました")
	}

	tempImageMap := make(map[string]models.TempImage)
	for _, img := range tempImages {
		tempImageMap[img.Image] = img
	}

	for _, imageUrl := range imageUrls {
		if _, exists := tempImageMap[imageUrl]; !exists {
			return nil, NewCustomError(http.StatusBadRequest, "無効な画像URLが含まれています")
		}
	}

	if err := tx.Where("image IN ?", imageUrls).Delete(&models.TempImage{}).Error; err != nil {
		return nil, NewCustomError(http.StatusInternalServerError, "一時画像の削除に失敗しました")
	}

	return tempImages, nil
}

func CreateFavoriteTeams(tx *gorm.DB, favoriteTeams *[]models.FavoriteTeam) error {
	if err := tx.Create(&favoriteTeams).Error; err != nil {
		log.Printf("お気に入りチーム追加エラー: %v", err)
		return NewCustomError(http.StatusInternalServerError, "お気に入りチーム追加に失敗しました。")
	}
	return nil
}

func GetFieldDetail(fieldName string, structName any) (FieldDetail, error) {
	val := reflect.ValueOf(structName)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	fieldDetail := FieldDetail{
		FieldName: fieldName,
		Min:    nil,
		Max:    nil,
	}
	field, found := val.Type().FieldByName(fieldName)
	if !found {
		log.Printf("フィールド %s が見つかりません", fieldName)
		return fieldDetail, NewCustomError(http.StatusInternalServerError, "フィールド詳細取得に失敗しました。")
	}


	fieldTag := field.Tag.Get("field")
	if fieldTag != "" {
		fieldDetail.FieldName = fieldTag
	}

	bindingTag := field.Tag.Get("binding")
	if bindingTag != "" {
		tagParts := strings.Split(bindingTag, ",")
		for _, part := range tagParts {
			if strings.HasPrefix(part, "min=") {
				if minValue, err := strconv.Atoi(strings.TrimPrefix(part, "min=")); err == nil {
					fieldDetail.Min = &minValue
				}
			}
			if strings.HasPrefix(part, "max=") {
				if maxValue, err := strconv.Atoi(strings.TrimPrefix(part, "max=")); err == nil {
					fieldDetail.Max = &maxValue
				}
			}
		}
	}

	return fieldDetail, nil
}


func GenerateRequestErrorMessages(err error, structName any) (*[]string, error) {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var errorMessages []string

		for _, ve := range validationErrors {
			fieldDetail, err := GetFieldDetail(ve.Field(), structName)
			if err != nil {
				return nil, err
			}

			var errorMessage string
			switch ve.Tag() {
				case "required":
					errorMessage = fmt.Sprintf("%sは必須項目です。", fieldDetail.FieldName)
				case "min":
					if fieldDetail.Min == nil {
						errorMessage = fmt.Sprintf("%sは下限文字数を下回っています。", fieldDetail.FieldName)
						} else {
							errorMessage = fmt.Sprintf("%sは%d文字以上で入力してください。", fieldDetail.FieldName, *fieldDetail.Min)
						}
					case "max":
						if fieldDetail.Max == nil {
							errorMessage = fmt.Sprintf("%sは上限文字数を上回っています。", fieldDetail.FieldName)
						} else {
							errorMessage = fmt.Sprintf("%sは%d文字以下で入力してください。", fieldDetail.FieldName, *fieldDetail.Max)
						}
				default:
					errorMessage = fmt.Sprintf("%sは不正な値です。", fieldDetail.FieldName)
			}
			errorMessages = append(errorMessages, errorMessage)
		}
		return &errorMessages, nil
	}

	return &[]string{"リクエストに不備があります。"}, nil
}

func HandleCustomError(c *gin.Context, customErr *CustomError, err error, request any) {
	log.Println(customErr.Code)
	switch customErr.Code {
	case http.StatusBadRequest:
		errorMessages, genErr := GenerateRequestErrorMessages(err, request)
		if genErr != nil {
			ErrorResponse[any](c, http.StatusInternalServerError, CreateSingleMessage(genErr.Error()))
		} else {
			ErrorResponse[any](c, http.StatusBadRequest, *errorMessages)
		}
		return
	default:
		ErrorResponse[any](c, customErr.Code, CreateSingleMessage(customErr.Error()))
	}
}

func ValidateRequest(c *gin.Context, requestPath any, requestQuery any, requestBody any, isProtected bool) (*uint, error, error) {
	if requestPath != nil {
		log.Printf("Path parameters: %v", c.Params)
		if err := c.ShouldBindUri(requestPath); err != nil {
			log.Printf("リクエストエラー: %v", err)
				return nil, err, NewCustomError(http.StatusBadRequest, "リクエストに不備があります。")
			}
	}
	if requestQuery != nil {
		log.Printf("Query parameters: %v", c.Request.URL.Query())
		if err := c.ShouldBindQuery(requestQuery); err != nil {
			log.Printf("リクエストエラー: %v", err)
			return nil, err, NewCustomError(http.StatusBadRequest, "リクエストに不備があります。")
		}
	}
	if requestBody != nil {
		log.Printf("Request body: %v", c.Request.Body)
		if err := c.ShouldBindJSON(requestBody); err != nil {
			log.Printf("リクエストエラー: %v", err)
			return nil, err, NewCustomError(http.StatusBadRequest, "リクエストに不備があります。")
		}
	}
	if isProtected {
		loginUserId, err := StringToUint(c.GetString("userId"))
		if err != nil {
			return nil, err, NewCustomError(http.StatusInternalServerError, err.Error())
		}
		return loginUserId, nil, nil
	}
	return nil, nil, nil
}

