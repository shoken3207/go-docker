package upload

// リクエスト
type UploadImagesRequestQuery struct {
	Folder string `form:"folder" binding:"required"`
}

// レスポンス
type UploadImagesResponse struct {
	Images []UploadToImageKitResponse `json:"images"`
}

type UploadToImageKitResponse struct {
	Url    string `json:"url"`
	FileId string `json:"fileId"`
}
