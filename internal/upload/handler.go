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
// @Security BearerAuth
// @Accept multipart/form-data
// @Produce json
// @Param folder query string true "格納フォルダ"
// @Param images formData file true "画像ファイル"
// @Success 200 {object} utils.ApiResponse[UploadImagesResponse] "成功"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 401 {object} utils.BasicResponse "認証エラー"
// @Failure 404 {object} utils.BasicResponse "not foundエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/upload/images [post]
func (h *UploadHandler) UploadImages(c *gin.Context, ik *imagekit.ImageKit) {
	form, err := c.MultipartForm()
	if err != nil {
		utils.ErrorResponse[any](c, http.StatusBadRequest, "multipart formのパースに失敗")
		return
	}


	files := form.File["images"]
	if len(files) == 0 {
		utils.ErrorResponse[any](c, http.StatusBadRequest, "ファイルが選択されていません")
		return
	} else if len(files) > 10 {
		utils.ErrorResponse[any](c, http.StatusBadRequest, "ファイルの上限選択数を超えています")
		return
	}

	folder := c.DefaultQuery("folder", "default")
	if folder == "" {
		utils.ErrorResponse[any](c, http.StatusBadRequest, "画像格納フォルダを指定してください")
		return
	}

	var uploadedURLs []string
	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			utils.ErrorResponse[any](c, http.StatusInternalServerError, "ファイルの開封に失敗")
			return
		}
		defer src.Close()

		if err := uploadService.validateFile(file); err != nil {
			utils.ErrorResponse[any](c, http.StatusBadRequest, err.Error())
			return
		}

		url, err := uploadService.uploadToImageKit(ik, folder, file.Filename, src)
		if err != nil {
			utils.ErrorResponse[any](c, http.StatusInternalServerError, err.Error())
			return
		}

		uploadedURLs = append(uploadedURLs, url)
	}
	utils.SuccessResponse(c, http.StatusOK, uploadedURLs, "画像のアップロードに成功しました")
}

func NewUploadHandler() *UploadHandler {
	return &UploadHandler{}
}