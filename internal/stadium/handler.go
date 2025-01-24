package stadium

import (
	"go-docker/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StadiumHandler struct{}

var stadiumService = NewStadiumService()

// @Summary スタジアム情報、そのスタジアムの遠征記録、周辺施設を取得するAPI
// @Description 遠征記録は1ページ目（15件）だけ返し、2ページ目以降は別APIから返す<br>周辺施設は20件を上限としてランキング形式で返す
// @Tags stadium
// @Security BearerAuth
// @Param stadiumId path integer true "stadiumId"
// @Success 200 {object} utils.ApiResponse[GetStadiumResponse] "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 401 {object} utils.ErrorBasicResponse "認証エラー"
// @Failure 404 {object} utils.ErrorBasicResponse "not foundエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/stadium/{stadiumId} [get]
func (h *StadiumHandler) GetStadium(c *gin.Context) {
	var requestPath GetStadiumRequestPath
	loginUserId, err, customErr := utils.ValidateRequest(c, &requestPath, nil, nil, true)
	if err != nil {
		if customErr, ok := customErr.(*utils.CustomError); ok {
			utils.HandleCustomError(c, customErr, err, requestPath)
			return
		}
	}

	stadiumResponse, err := stadiumService.GetStadiumService(loginUserId, &requestPath)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, utils.CreateSingleMessage(customErr.Error()))
			return
		}
	}
	utils.SuccessResponse[GetStadiumResponse](c, http.StatusOK, *stadiumResponse, utils.CreateSingleMessage("スタジアムの取得に成功しました。"))
}

func NewStadiumHandler() *StadiumHandler {
	return &StadiumHandler{}
}