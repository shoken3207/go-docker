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

// @Summary idから遠征記録詳細を取得
// @Description 遠征記録詳細情報を取得
// @Tags expedition
// @Security BearerAuth
// @Param expeditionId path uint true "expeditionId"
// @Success 200 {object} utils.ApiResponse[GetExpeditionDetailResponse] "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 403 {object} utils.ErrorBasicResponse "認証エラー"
// @Failure 404 {object} utils.ErrorBasicResponse "not foundエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/admin/expedition/{expeditionId} [get]
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
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 403 {object} utils.ErrorBasicResponse "認証エラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/admin/expedition/create [post]
func (h *ExpeditionHandler) CreateExpedition(c *gin.Context) {
	var request CreateExpeditionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("リクエストエラー: %v", err)
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
// @Description 遠征、出費、試合、訪れた施設の情報を更新する。<br>Payment, VisitedFacility, Game, GameScoreのdeleteにはidの配列ですが、ExpeditionImageのdeleteにはurlの配列をリクエストで渡してください
// @Tags expedition
// @Param request body UpdateExpeditionRequestBody true "遠征記録更新リクエスト"
// @Param expeditionId path int true "遠征記録ID"
// @Security BearerAuth
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 403 {object} utils.ErrorBasicResponse "認証エラー"
// @Failure 404 {object} utils.ErrorBasicResponse "ユーザーが見つかりません"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/admin/expedition/update/{expeditionId} [put]
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
// @Router /api/admin/expedition/delete/{expeditionId} [delete]
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

	if err := expeditionService.DeleteExpeditionService(&requestPath.ExpeditionId, userId, ik); err != nil {
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
// @Router /api/admin/expedition/like/{expeditionId} [post]
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
// @Router /api/admin/expedition/unlike/{expeditionId} [delete]
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

// @Summary 遠征記録一覧を取得
// @Description ページネーション付きで遠征記録一覧を取得します<br>teamIdとsportIdを指定すると、そのチーム、スポーツの遠征記録一覧を取得します。指定しなければ全ての遠征記録一覧を取得します
// @Tags expedition
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param page query int true "ページ番号" minimum(1)
// @Param sportId query int false "スポーツID"
// @Param teamId query int false "チームID"
// @Success 200 {object} utils.ApiResponse[[]ExpeditionListResponse] "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 403 {object} utils.ErrorBasicResponse "認証エラー"
// @Failure 404 {object} utils.ErrorBasicResponse "遠征記録が見つかりません"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/admin/expedition/list [get]
func (h *ExpeditionHandler) GetExpeditionList(c *gin.Context) {
	var req ExpeditionListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Printf("リクエストパラメータが不正です: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストパラメータが不正です")
		return
	}

	expeditions, err := expeditionService.GetExpeditionList(&req)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}

	utils.SuccessResponse(c, http.StatusOK, expeditions, "遠征記録一覧を取得しました")
}

func NewExpeditionHandler() *ExpeditionHandler {
	return &ExpeditionHandler{}
}
