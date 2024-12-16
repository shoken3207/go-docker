package upload

// リクエスト
type UploadImagesRequestQuery struct {
	Folder string `json:"folder" binding:"required"`
}

// レスポンス
type UploadImagesResponse struct {
	Urls []string `json:"urls"`
}
