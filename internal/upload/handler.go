package upload

import (
	"go-docker/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imagekit-developer/imagekit-go"
)

type UploadHandler struct{}

var uploadService = NewUploadService()

// @Summary 画像をクラウドストレージ(imagekit)にアップロード
// @Description 画像をアップロードし、URLを返します。<br>プロフィール、スタジアム、遠征など、格納フォルダを指定してください。<br>画像は1枚から10枚アップロードできるが、Swagger UIでは1つしか選択できません。<br>ファイルの拡張子は、[".jpg", ".jpeg", ".png"]だけを受け付けています。ファイルサイズは最大5MBを上限としています。
// @Tags upload
// @Accept multipart/form-data
// @Param folder query string true "格納フォルダ"
// @Param images formData file true "画像ファイル"
// @Success 200 {object} utils.ApiResponse[UploadImagesResponse] "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 401 {object} utils.ErrorBasicResponse "認証エラー"
// @Failure 404 {object} utils.ErrorBasicResponse "not foundエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/upload/images [post]
func (h *UploadHandler) UploadImages(c *gin.Context, ik *imagekit.ImageKit) {
	query, files, err := uploadService.validateUploadImagesRequest(c)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}
	imageUrls, err := uploadService.UploadImagesService(ik, &query.Folder, files)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}
	utils.SuccessResponse[UploadImagesResponse](c, http.StatusOK, UploadImagesResponse{ImageUrls: *imageUrls}, "画像のアップロードに成功しました")
}

func NewUploadHandler() *UploadHandler {
	return &UploadHandler{}
}
