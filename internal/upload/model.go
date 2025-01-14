package upload

// リクエスト
type UploadImagesRequestQuery struct {
	Folder string `form:"folder" binding:"required" example:"expedition"`
}

// レスポンス
type UploadImagesResponse struct {
	Images []UploadToImageKitResponse `json:"images"`
}

type UploadToImageKitResponse struct {
	Url    string `json:"url" example:"https://ik.imagekit.io/your_imagekit_id/image.jpg"`
	FileId string `json:"fileId" example:"file_1234567890"`
}
