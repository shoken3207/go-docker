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
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 401 {object} utils.ErrorBasicResponse "認証エラー"
// @Failure 404 {object} utils.ErrorBasicResponse "not foundエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/expedition/{expeditionId} [get]
func (h *ExpeditionHandler) GetExpeditionDetail(c *gin.Context) {
	var requestPath GetExpeditionDetailRequestPath
	loginUserId, err, customErr := utils.ValidateRequest(c, &requestPath, nil, nil, true)
	if err != nil {
		if customErr, ok := customErr.(*utils.CustomError); ok {
			utils.HandleCustomError(c, customErr, err, requestPath)
			return
		}
	}
	expeditionDetail, err := expeditionService.GetExpeditionDetailService(&requestPath, loginUserId)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, utils.CreateSingleMessage(customErr.Error()))
			return
		}
	}

	utils.SuccessResponse[GetExpeditionDetailResponse](c, http.StatusOK, *expeditionDetail, utils.CreateSingleMessage("遠征記録詳細取得に成功しました"))
}

// @Summary 遠征記録を作成
// @Description 遠征、出費、試合、訪れた施設の情報を保存する。
// @Tags expedition
// @Security BearerAuth
// @Param request body CreateExpeditionRequestBody true "遠征記録作成リクエスト"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 401 {object} utils.ErrorBasicResponse "認証エラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/expedition/create [post]
func (h *ExpeditionHandler) CreateExpedition(c *gin.Context) {
	var requestBody CreateExpeditionRequestBody
	loginUserId, err, customErr := utils.ValidateRequest(c, nil, nil, &requestBody, true)
	if err != nil {
		if customErr, ok := customErr.(*utils.CustomError); ok {
			utils.HandleCustomError(c, customErr, err, requestBody)
			return
		}
	}
	if err := expeditionService.CreateExpeditionService(&requestBody, loginUserId); err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, utils.CreateSingleMessage(customErr.Error()))
			return
		}
	}

	utils.SuccessResponse[any](c, http.StatusOK, nil, utils.CreateSingleMessage("遠征記録作成に成功しました"))
}

// @Summary 遠征記録を更新
// @Description 遠征、出費、試合、訪れた施設の情報を更新する。<br>Payment, VisitedFacility, Game, GameScoreのdeleteにはidの配列ですが、ExpeditionImageのdeleteにはurlの配列をリクエストで渡してください
// @Tags expedition
// @Param request body UpdateExpeditionRequestBody true "遠征記録更新リクエスト"
// @Param expeditionId path int true "遠征記録ID"
// @Security BearerAuth
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 401 {object} utils.ErrorBasicResponse "認証エラー"
// @Failure 404 {object} utils.ErrorBasicResponse "ユーザーが見つかりません"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/expedition/update/{expeditionId} [put]
func (h *ExpeditionHandler) UpdateExpedition(c *gin.Context, ik *imagekit.ImageKit) {
	var requestBody UpdateExpeditionRequestBody
	var requestPath UpdateExpeditionRequestPath
	loginUserId, err, customErr := utils.ValidateRequest(c, &requestPath, nil, &requestBody, true)
	if err != nil {
		if customErr, ok := customErr.(*utils.CustomError); ok {
			utils.HandleCustomError(c, customErr, err, requestBody)
			return
		}
	}

	if err := expeditionService.UpdateExpeditionService(&requestPath.ExpeditionId, loginUserId, &requestBody, ik); err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, utils.CreateSingleMessage(customErr.Error()))
			return
		}
	}
	utils.SuccessResponse[any](c, http.StatusOK, nil, utils.CreateSingleMessage("遠征記録更新に成功しました"))
}

// @Summary 遠征記録を削除
// @Description 遠征記録とそれに関連する全てのデータ（画像、いいね、支払い、試合情報など）を削除する
// @Tags expedition
// @Security BearerAuth
// @Param expeditionId path int true "遠征記録ID"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 401 {object} utils.ErrorBasicResponse "認証エラー"
// @Failure 404 {object} utils.ErrorBasicResponse "遠征記録が見つかりません"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/expedition/delete/{expeditionId} [delete]
func (h *ExpeditionHandler) DeleteExpedition(c *gin.Context, ik *imagekit.ImageKit) {
	var requestPath DeleteExpeditionRequestPath
	loginUserId, err, customErr := utils.ValidateRequest(c, &requestPath, nil, nil, true)
	if err != nil {
		if customErr, ok := customErr.(*utils.CustomError); ok {
			utils.HandleCustomError(c, customErr, err, requestPath)
			return
		}
	}

	if err := expeditionService.DeleteExpeditionService(&requestPath.ExpeditionId, loginUserId, ik); err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, utils.CreateSingleMessage(customErr.Error()))
			return
		}
	}

	utils.SuccessResponse[any](c, http.StatusOK, nil, utils.CreateSingleMessage("遠征記録を削除しました"))
}

// @Summary 遠征記録にいいね、いいね解除を行う
// @Description ユーザーが遠征記録にいいね済みならいいねを付ける。いいねしていなかったらいいねする
// @Tags expedition
// @Security BearerAuth
// @Param expeditionId path int true "遠征記録ID"
// @Success 200 {object} utils.ApiResponse[LikeExpeditionResponse] "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 401 {object} utils.ErrorBasicResponse "認証エラー"
// @Failure 404 {object} utils.ErrorBasicResponse "遠征記録が見つかりません"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/expedition/like/{expeditionId} [post]
func (h *ExpeditionHandler) LikeExpedition(c *gin.Context) {
	var requestPath LikeExpeditionRequestPath
	loginUserId, err, customErr := utils.ValidateRequest(c, &requestPath, nil, nil, true)
	if err != nil {
		if customErr, ok := customErr.(*utils.CustomError); ok {
			utils.HandleCustomError(c, customErr, err, requestPath)
			return
		}
	}

	likesCount, isLiked, message, err := expeditionService.ExpeditionLikeService(loginUserId, &requestPath.ExpeditionId);
	if  err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, utils.CreateSingleMessage(customErr.Error()))
			return
		}
	}

	utils.SuccessResponse[LikeExpeditionResponse](c, http.StatusOK, LikeExpeditionResponse{LikesCount: *likesCount, IsLiked: *isLiked}, utils.CreateSingleMessage(*message))
}

// @Summary 遠征記録一覧を取得
// @Description ページネーション付きで遠征記録一覧を取得します<br>teamIdとsportIdを指定すると、そのチーム、スポーツの遠征記録一覧を取得します。指定しなければ全ての遠征記録一覧を取得します<br>stadiumIdを入力したらそのstadiumIdに関する遠征記録を取得します。
// @Tags expedition
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param page query int true "ページ番号" minimum(1)
// @Param sportId query int false "スポーツID"
// @Param teamId query int false "チームID"
// @Param stadiumId query int false "スタジアムID"
// @Success 200 {object} utils.ApiResponse[[]ExpeditionListResponse] "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 401 {object} utils.ErrorBasicResponse "認証エラー"
// @Failure 404 {object} utils.ApiResponse[[]ExpeditionListResponse] "遠征記録が見つかりません"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/expedition/list [get]
func (h *ExpeditionHandler) GetExpeditionList(c *gin.Context) {
	var requestQuery GetExpeditionListRequestQuery
	loginUserId, err, customErr := utils.ValidateRequest(c, nil, &requestQuery, nil, true)
	if err != nil {
		if customErr, ok := customErr.(*utils.CustomError); ok {
			utils.HandleCustomError(c, customErr, err, requestQuery)
			return
		}
	}

	expeditions, err := expeditionService.GetExpeditionListService(&requestQuery, loginUserId)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, utils.CreateSingleMessage(customErr.Error()))
			return
		}
	}

	if len(expeditions) == 0 {
		var message string
		if requestQuery.Page == 1 {
			message = "いいねした遠征記録が見つかりません"
		} else {
			message = "最後のページです"
		}
		utils.SuccessResponse[[]ExpeditionListResponse](c, http.StatusNotFound, expeditions, utils.CreateSingleMessage(message))
		return
	}

	utils.SuccessResponse[[]ExpeditionListResponse](c, http.StatusOK, expeditions, utils.CreateSingleMessage("遠征記録一覧を取得しました"))
}

// @Summary ユーザーが投稿した遠征記録一覧を取得
// @Description リクエストのuserIdからページネーション付きで遠征記録一覧を取得します<br>ログインユーザーの場合はisPublicがfalse（プライベート）な投稿も取得し、そうじゃなければisPublicがtrue（パブリック）な投稿だけ取得します。
// @Tags expedition
// @Security BearerAuth
// @Param page query int true "ページ番号" minimum(1)
// @Param userId query int true "ユーザID"
// @Success 200 {object} utils.ApiResponse[[]ExpeditionListResponse] "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 401 {object} utils.ErrorBasicResponse "認証エラー"
// @Failure 404 {object} utils.ApiResponse[[]ExpeditionListResponse] "遠征記録が見つかりません"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/expedition/list/user [get]
func (h *ExpeditionHandler) GetExpeditionListByUserId(c *gin.Context) {
	var requestQuery GetExpeditionListByUserIdRequestQuery
	loginUserId, err, customErr := utils.ValidateRequest(c, nil, &requestQuery, nil, true)
	if err != nil {
		if customErr, ok := customErr.(*utils.CustomError); ok {
			utils.HandleCustomError(c, customErr, err, requestQuery)
			return
		}
	}

	expeditions, err := expeditionService.GetExpeditionListService(&GetExpeditionListRequestQuery{Page: requestQuery.Page, UserId: &requestQuery.UserId}, loginUserId)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, utils.CreateSingleMessage(customErr.Error()))
			return
		}
	}

	if len(expeditions) == 0 {
		var message string
		if requestQuery.Page == 1 {
			message = "いいねした遠征記録が見つかりません"
		} else {
			message = "最後のページです"
		}
		utils.SuccessResponse[[]ExpeditionListResponse](c, http.StatusNotFound, expeditions, utils.CreateSingleMessage(message))
		return
	}
	utils.SuccessResponse[[]ExpeditionListResponse](c, http.StatusOK, expeditions, utils.CreateSingleMessage("遠征記録一覧を取得しました"))
}

// @Summary ユーザーがいいねした遠征記録一覧を取得
// @Description リクエストのuserIdからページネーション付きで遠征記録一覧を取得します<br>ログインユーザーの場合はisPublicがfalse（プライベート）な投稿も取得し、そうじゃなければisPublicがtrue（パブリック）な投稿だけ取得します。
// @Tags expedition
// @Security BearerAuth
// @Param page query int true "ページ番号" minimum(1)
// @Param userId query int true "ユーザID"
// @Success 200 {object} utils.ApiResponse[[]ExpeditionListResponse] "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 401 {object} utils.ErrorBasicResponse "認証エラー"
// @Failure 404 {object} utils.ApiResponse[[]ExpeditionListResponse] "遠征記録が見つかりません"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/expedition/list/user/likes [get]
func (h *ExpeditionHandler) GetLikedExpeditionListByUserId(c *gin.Context) {
	var requestQuery GetExpeditionListByUserIdRequestQuery
	loginUserId, err, customErr := utils.ValidateRequest(c, nil, &requestQuery, nil, true)
	if err != nil {
		if customErr, ok := customErr.(*utils.CustomError); ok {
			utils.HandleCustomError(c, customErr, err, requestQuery)
			return
		}
	}

	expeditions, err := expeditionService.GetLikedExpeditionListService(&GetExpeditionListRequestQuery{Page: requestQuery.Page, UserId: &requestQuery.UserId}, loginUserId)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, utils.CreateSingleMessage(customErr.Error()))
			return
		}
	}

	if len(expeditions) == 0 {
		var message string
		if requestQuery.Page == 1 {
			message = "いいねした遠征記録が見つかりません"
		} else {
			message = "最後のページです"
		}
		utils.SuccessResponse[[]ExpeditionListResponse](c, http.StatusNotFound, expeditions, utils.CreateSingleMessage(message))
		return
	}
	utils.SuccessResponse[[]ExpeditionListResponse](c, http.StatusOK, expeditions, utils.CreateSingleMessage("遠征記録一覧を取得しました"))
}

func NewExpeditionHandler() *ExpeditionHandler {
	return &ExpeditionHandler{}
}
