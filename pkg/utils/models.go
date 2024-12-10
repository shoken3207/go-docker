package utils

type ApiResponse[T any] struct {
	Success bool   `json:"success"`
	Data    T      `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

type BasicResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}