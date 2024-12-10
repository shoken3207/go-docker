package uploadimage

import (
	"github.com/gin-gonic/gin"
)

type UploadHandler struct{}

// @Summary 画像をクラウドストレージにアップロード
// @Description 画像をアップロードし、URLを返します。
// @Tags upload
// @Param file formData file true "画像ファイル"
// @Router /api/upload/images [post]
func (h *UploadHandler) UploadImages(c *gin.Context) {
}

func NewUploadHandler() *UploadHandler {
	return &UploadHandler{}
}