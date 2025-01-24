package utils

type CustomError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Bad Request"`
}

type ApiResponse[T any] struct {
	Success  bool     `json:"success" example:"true"`
	Data     T        `json:"data"`
	Messages []string `json:"messages" example:"成功しました！！"`
}

type SuccessBasicResponse struct {
	Success  bool     `json:"success" example:"true"`
	Messages []string `json:"messages" example:"成功しました！！"`
}
type ErrorBasicResponse struct {
	Success  bool     `json:"success" example:"false"`
	Messages []string `json:"messages" example:"エラーメッセージ"`
}

type FieldDetail struct {
	FieldName string `json:"field"`
	Min       *int   `json:"min"`
	Max       *int   `json:"max"`
}