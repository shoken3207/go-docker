package utils

type CustomError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Bad Request"`
}

type ApiResponse[T any] struct {
	Success bool   `json:"success" example:"true"`
	Data    T      `json:"data"`
	Message string `json:"message,omitempty" example:"成功しました！！"`
}

type SuccessBasicResponse struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"成功しました！！"`
}
type ErrorBasicResponse struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"エラーメッセージ"`
}
