package upload

// リクエスト
type UploadImagesRequestQuery struct {
	Folder string `form:"folder" binding:"required"`
}

// レスポンス
type UploadImagesResponse struct {
	Urls []string `json:"urls"`
}
