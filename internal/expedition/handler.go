package expedition

import (
	"go-docker/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imagekit-developer/imagekit-go"
)

type ExpeditionHandler struct{}

var expeditionService = NewExpeditionService()

// @Summary 遠征記録を作成
// @Description 遠征、出費、試合、訪れた施設の情報を保存する。fileIdはimagekitのfileIdで、ユニークな値です。
// @Tags expedition
// @Security BearerAuth
// @Param request body CreateExpeditionRequest true "遠征記録作成リクエスト"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 403 {object} utils.ErrorBasicResponse "認証エラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/expedition/create [post]
func (h *ExpeditionHandler) CreateExpedition(c *gin.Context) {
	var request CreateExpeditionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("リクエストに不備があります。: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります。")
		return
	}
	userId, err := utils.StringToUint(c.GetString("userId"))
	if err != nil {
		utils.ErrorResponse[any](c, http.StatusBadRequest, err.Error())
		return
	}
	if err := expeditionService.CreateExpeditionService(&request, userId); err != nil {
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
// @Param expeditionId path int true "遠征記録ID"
// @Security BearerAuth
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 403 {object} utils.ErrorBasicResponse "認証エラー"
// @Failure 404 {object} utils.ErrorBasicResponse "ユーザーが見つかりません"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/expedition/update/{expeditionId} [put]
func (h *ExpeditionHandler) UpdateExpedition(c *gin.Context, ik *imagekit.ImageKit) {
	expeditionId, requestBody, err := expeditionService.ValidateUpdateExpeditionRequest(c)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}

	userId, err := utils.StringToUint(c.GetString("userId"))
	if err != nil {
		utils.ErrorResponse[any](c, http.StatusBadRequest, err.Error())
		return
	}

	if err := expeditionService.UpdateExpeditionService(expeditionId, userId, requestBody, ik); err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}
	utils.SuccessResponse[any](c, http.StatusOK, nil, "遠征記録更新に成功しました")
}

// @Summary 遠征記録を削除
// @Description 遠征記録とそれに関連する全てのデータ（画像、いいね、支払い、試合情報など）を削除する
// @Tags expedition
// @Security BearerAuth
// @Param expeditionId path int true "遠征記録ID"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 403 {object} utils.ErrorBasicResponse "認証エラー"
// @Failure 404 {object} utils.ErrorBasicResponse "遠征記録が見つかりません"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/expedition/delete/{expeditionId} [delete]
func (h *ExpeditionHandler) DeleteExpedition(c *gin.Context, ik *imagekit.ImageKit) {
	var requestPath DeleteExpeditionRequestPath
	if err := c.ShouldBindUri(&requestPath); err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}

	userId, err := utils.StringToUint(c.GetString("userId"))
	if err != nil {
		utils.ErrorResponse[any](c, http.StatusBadRequest, err.Error())
		return
	}

	if err := expeditionService.DeleteExpedition(&requestPath.ExpeditionId, userId, ik); err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}

	utils.SuccessResponse[any](c, http.StatusOK, nil, "遠征記録を削除しました")
}

// @Summary 遠征記録にいいねする
// @Description ユーザーが遠征記録にいいねを付ける
// @Tags expedition
// @Security BearerAuth
// @Param expeditionId path int true "遠征記録ID"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 403 {object} utils.ErrorBasicResponse "認証エラー"
// @Failure 404 {object} utils.ErrorBasicResponse "遠征記録が見つかりません"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/expedition/like/{expeditionId} [post]
func (h *ExpeditionHandler) LikeExpedition(c *gin.Context) {
	var requestPath LikeExpeditionRequestPath
    if err := c.ShouldBindUri(&requestPath); err != nil {
		log.Printf("リクエストエラー: %v", err)
        utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
        return
	}

	userId, err := utils.StringToUint(c.GetString("userId"))
	if err != nil {
		utils.ErrorResponse[any](c, http.StatusBadRequest, err.Error())
		return
	}

	if err := expeditionService.CreateExpeditionLike(userId, &requestPath.ExpeditionId); err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}

	utils.SuccessResponse[any](c, http.StatusOK, nil, "いいねしました")
}

// @Summary 遠征記録のいいねを外す
// @Description ユーザーが遠征記録のいいねを外す
// @Tags expedition
// @Security BearerAuth
// @Param expeditionId path int true "遠征記録ID"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 403 {object} utils.ErrorBasicResponse "認証エラー"
// @Failure 404 {object} utils.ErrorBasicResponse "いいねが見つかりません"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/expedition/unlike/{expeditionId} [delete]
func (h *ExpeditionHandler) UnlikeExpedition(c *gin.Context) {
	var requestPath UnlikeExpeditionRequestPath
    if err := c.ShouldBindUri(&requestPath); err != nil {
		log.Printf("リクエストエラー: %v", err)
        utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
        return
    }

	userId, err := utils.StringToUint(c.GetString("userId"))
	if err != nil {
		utils.ErrorResponse[any](c, http.StatusBadRequest, err.Error())
		return
	}
	if err := expeditionService.DeleteExpeditionLike(userId, &requestPath.ExpeditionId); err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}

	utils.SuccessResponse[any](c, http.StatusOK, nil, "いいねを外しました")
}

func NewExpeditionHandler() *ExpeditionHandler {
	return &ExpeditionHandler{}
}
