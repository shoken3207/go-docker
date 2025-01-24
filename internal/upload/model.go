package upload

// リクエスト
type UploadImagesRequestQuery struct {
	Folder string `form:"folder" binding:"required" example:"expedition" field:"フォルダ名"`
}

// レスポンス
type UploadImagesResponse struct {
	ImageUrls []string `json:"imageUrls"`
}

type UploadToImageKitResponse struct {
	Url    string `json:"url" example:"https://ik.imagekit.io/your_imagekit_id/image.jpg"`
	FileId string `json:"fileId" example:"file_1234567890"`
}
