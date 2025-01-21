package stadium

import (
	"go-docker/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StadiumHandler struct{}

// @Summary スタジアム情報、そのスタジアムの遠征記録、周辺施設を取得するAPI
// @Description 遠征記録、周辺施設は1ページ目だけ返し、2ページ目以降は別APIから返す
// @Tags stadium
// @Param stadiumId path integer true "stadiumId"
// @Success 200 {object} utils.ApiResponse[GetStadiumResponse] "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/stadium/get [get]
func (h *StadiumHandler) GetStadium(c *gin.Context) {
	utils.SuccessResponse[GetStadiumResponse](c, http.StatusOK, GetStadiumResponse{}, "aaa")
}

func NewStadiumHandler() *StadiumHandler {
	return &StadiumHandler{}
}