package expedition

import (
	"go-docker/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imagekit-developer/imagekit-go"
)

type ExpeditionHandler struct{}

var expeditionService = NewExpeditionService()

// @Summary idから遠征記録詳細を取得
// @Description 遠征記録詳細情報を取得
// @Tags expedition
// @Security BearerAuth
// @Param expeditionId path uint true "expeditionId"
// @Success 200 {object} utils.ApiResponse[GetExpeditionDetailResponse] "成功"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 403 {object} utils.BasicResponse "認証エラー"
// @Failure 404 {object} utils.BasicResponse "not foundエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/expedition/{expeditionId} [get]
func (h *ExpeditionHandler) GetExpeditionDetail(c *gin.Context) {
	var request GetExpeditionDetailRequest
	if err := c.ShouldBindUri(&request); err != nil {
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります。")
	}
	expeditionDetail, err := expeditionService.GetExpeditionDetailService(&request)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}

	utils.SuccessResponse[GetExpeditionDetailResponse](c, http.StatusOK, *expeditionDetail, "遠征記録詳細取得に成功しました")
}

// @Summary 遠征記録を作成
// @Description 遠征、出費、試合、訪れた施設の情報を保存する。
// @Tags expedition
// @Security BearerAuth
// @Param request body CreateExpeditionRequest true "遠征記録作成リクエスト"
// @Success 200 {object} utils.BasicResponse "成功"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 403 {object} utils.BasicResponse "認証エラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/expedition/create [post]
func (h *ExpeditionHandler) CreateExpedition(c *gin.Context) {
	var request CreateExpeditionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります。")
	}
	if err := expeditionService.CreateExpeditionService(&request); err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}

	utils.SuccessResponse[any](c, http.StatusOK, nil, "遠征記録作成に成功しました")
}

// @Summary 遠征記録を更新
// @Description 遠征、出費、試合、訪れた施設の情報を更新する。<br>Payment, VisitedFacility, Game, GameScoreのdeleteにはidの配列ですが、ExpeditionImageのdeleteにはfileId(string)の配列をリクエストで渡してください
// @Tags expedition
// @Param request body UpdateExpeditionRequestBody true "遠征記録更新リクエスト"
// @Security BearerAuth
// @Success 200 {object} utils.BasicResponse "成功"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 403 {object} utils.BasicResponse "認証エラー"
// @Failure 404 {object} utils.BasicResponse "ユーザーが見つかりません"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/expedition/update/{id} [put]
func (h *ExpeditionHandler) UpdateExpedition(c *gin.Context, ik *imagekit.ImageKit) {
	expeditionId, requestBody, err := expeditionService.ValidateUpdateExpeditionRequest(c)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}

	if err := expeditionService.UpdateExpeditionService(expeditionId, requestBody, ik); err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}
	utils.SuccessResponse[any](c, http.StatusOK, nil, "遠征記録更新に成功しました")
}

// @Summary 遠征記録を削除
// @Description pathのidをもとに遠征記録を削除する。
// @Tags expedition
// @Security BearerAuth
// @Success 200 {object} utils.BasicResponse "アップロードした画像のURL"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 403 {object} utils.BasicResponse "ユーザーが見つかりません"
// @Failure 404 {object} utils.BasicResponse "ユーザーが見つかりません"
// @Failure 500 {object} utils.BasicResponse "ユーザーが見つかりません"
// @Router /api/expedition/delete/{id} [delete]
func (h *ExpeditionHandler) DeleteExpedition(c *gin.Context) {
}

func NewExpeditionHandler() *ExpeditionHandler {
	return &ExpeditionHandler{}
}
