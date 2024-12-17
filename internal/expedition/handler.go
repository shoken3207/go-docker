package expedition

import (
	"github.com/gin-gonic/gin"
)

type ExpeditionHandler struct{}

// @Summary 遠征記録を作成
// @Description 遠征、出費、試合、訪れた施設の情報を保存する。
// @Tags expedition
// @Security BearerAuth
// @Param request body CreateExpeditionRequest true "遠征記録作成リクエスト"
// @Success 200 {object} utils.BasicResponse "アップロードした画像のURL"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 403 {object} utils.BasicResponse "ユーザーが見つかりません"
// @Failure 404 {object} utils.BasicResponse "ユーザーが見つかりません"
// @Failure 500 {object} utils.BasicResponse "ユーザーが見つかりません"
// @Router /api/expedition/create [post]
func (h *ExpeditionHandler) CreateExpedition(c *gin.Context) {

}

// @Summary 遠征記録を更新
// @Description 遠征、出費、試合、訪れた施設の情報を更新する。
// @Tags expedition
// @Security BearerAuth
// @Success 200 {object} utils.BasicResponse "アップロードした画像のURL"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 403 {object} utils.BasicResponse "ユーザーが見つかりません"
// @Failure 404 {object} utils.BasicResponse "ユーザーが見つかりません"
// @Failure 500 {object} utils.BasicResponse "ユーザーが見つかりません"
// @Router /api/expedition/update/{id} [put]
func (h *ExpeditionHandler) UpdateExpedition(c *gin.Context) {
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
