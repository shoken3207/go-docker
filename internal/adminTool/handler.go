package adminTool

import (
	"go-docker/models"
	"go-docker/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminToolHandler struct{}

var adminToolService = NewAdminToolService()

// @Summary スタジアム全件検索
// @Description スタジアム情報のレコードを全件取得して、一覧として表示する。
// @Tags stadium
// @Secrity BearerAuth
// @Param keyword query string false "キーワード"
// @Success 200 {object} utils.BasicResponse "成功"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/stadium/stadiums [get]
func (h *AdminToolHandler) GetStadiums(c *gin.Context) {
	keyword := c.DefaultQuery("keyword", "")
	log.Println("キーワード:", keyword)
	stadiums := []models.Stadium{}

	stadiums, err := adminToolService.getStadiumsService(keyword)

	if err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}

	log.Println(stadiums)

	utils.SuccessResponse[any](c, http.StatusOK, stadiums, "スタジアムの検索に成功しました。")
}

// @Summary スタジアム追加
// @Description リクエストからスタジアム情報を追加後、重複確認を行い登録する。
// @Tags stadium
// @Security BearerAuth
// @Param request body StadiumAddRequest true "スタジアム情報"
// @Success 200 {object} utils.BasicResponse "成功"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/stadium/stadiumAdd [post]
func (h *AdminToolHandler) StadiumAdd(c *gin.Context) {
	request := StadiumAddRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}

	if err := adminToolService.createStadiumService(&request); err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}

	utils.SuccessResponse[any](c, http.StatusOK, nil, "スタジアム登録に成功しました。")
}

// @Summary スタジアム更新
// @Description リクエストボディに更新対象のIDを指定してスタジアムを更新します
// @Tags stadium
// @Security BearerAuth
// @Param request body StadiumUpdateRequest true "スタジアム情報"
// @Success 200 {object} utils.BasicResponse "成功"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/stadium/update [put]
func (h *AdminToolHandler) StadiumUpdate(c *gin.Context) {
	request := StadiumUpdateRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}
	log.Println("重複検索が正常にリターンはされているよ:handler")
	if err := adminToolService.UpdateStadiumService(&request); err != nil {
		log.Printf("リクエストエラー: %v", err)
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}

	utils.SuccessResponse[any](c, http.StatusOK, nil, "スタジアムが正常に更新更新されました。")
}

// @Summary スタジアム削除
// @Description リクエストボディに削除対象のIDを指定してスタジアムを削除します
// @Tags stadium
// @Security BearerAuth
// @Param request body DeleteRequest true "スタジアムID"
// @Success 200 {object} utils.BasicResponse "成功"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/stadium/delete [delete]
func (h *AdminToolHandler) DeleteStadium(c *gin.Context) {
	request := DeleteRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}

	if err := adminToolService.deleteStadiumService(&request); err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}

	utils.SuccessResponse[any](c, http.StatusOK, nil, "スタジアムが正常に削除されました。")
}

// @Summary スポーツ全件検索
// @Description スポーツ情報のレコードを全件取得して、一覧として表示する。
// @Tags sports
// @Secrity BearerAuth
// @Param keyword query string false "キーワード"
// @Success 200 {object} utils.BasicResponse "成功"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/sports/sports [get]
func (h *AdminToolHandler) GetSports(c *gin.Context) {
	keyword := c.DefaultQuery("keyword", "")
	log.Println("キーワード:", keyword)
	sport := []models.Sport{}

	sport, err := adminToolService.getSportService(keyword)

	if err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}

	log.Println(sport)

	utils.SuccessResponse[any](c, http.StatusOK, sport, "スポーツの検索に成功しました。")
}

// @Summary スポーツの追加
// @Description リクエストからスポーツ情報を取得後、重複確認を行い登録する。
// @Tags sports
// @Security BearerAuth
// @Param request body SportsAddRequest true "スポーツ情報"
// @Success 200 {object} utils.BasicResponse "成功"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/sports/sportsAdd [post]
func (h *AdminToolHandler) SportsAdd(c *gin.Context) {
	request := SportsAddRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}

	if err := adminToolService.createSportService(&request); err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}

	utils.SuccessResponse[any](c, http.StatusOK, nil, "スポーツ情報の登録に成功しました。")
}

// @Summary スポーツ更新
// @Description リクエストボディに更新対象のIDを指定してスポーツ情報を更新します
// @Tags sports
// @Security BearerAuth
// @Param request body SportsUpdateRequest true "スポーツ情報"
// @Success 200 {object} utils.BasicResponse "成功"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/sports/update [put]
func (h *AdminToolHandler) SportsUpdate(c *gin.Context) {
	request := SportsUpdateRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}
	log.Println("重複検索が正常にリターンはされているよ:handler")
	if err := adminToolService.UpdateSportService(&request); err != nil {
		log.Printf("リクエストエラー: %v", err)
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}

	utils.SuccessResponse[any](c, http.StatusOK, nil, "スポーツが正常に更新されました。")
}

// @Summary スポーツ削除
// @Description リクエストボディに削除対象のIDを指定してスポーツ情報を削除します
// @Tags sports
// @Security BearerAuth
// @Param request body DeleteRequest true "スポーツ情報"
// @Success 200 {object} utils.BasicResponse "成功"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/sports/delete [delete]
func (h *AdminToolHandler) DeleteSports(c *gin.Context) {
	request := DeleteRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}

	if err := adminToolService.deleteSportService(&request); err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}

	utils.SuccessResponse[any](c, http.StatusOK, nil, "スポーツが正常に削除されました。")
}

// @Summary リーグ全件検索
// @Description リーグ情報のレコードを全件取得して、一覧として表示する。
// @Tags league
// @Secrity BearerAuth
// @Param keyword query string false "キーワード"
// @Success 200 {object} utils.BasicResponse "成功"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/league/leagues [get]
func (h *AdminToolHandler) GetLeagues(c *gin.Context) {
	keyword := c.DefaultQuery("keyword", "")
	log.Println("キーワード:", keyword)
	league := []models.League{}

	league, err := adminToolService.getLeagueService(keyword)

	if err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}

	log.Println(league)

	utils.SuccessResponse[any](c, http.StatusOK, league, "リーグの検索に成功しました。")
}

// @Summary リーグの追加
// @Description リクエストからリーグ情報を取得後、重複確認を行い登録する。
// @Tags league
// @Security BearerAuth
// @Param request body LeagueAddRequest true "リーグ情報"
// @Success 200 {object} utils.BasicResponse "成功"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/league/leagueAdd [post]
func (h *AdminToolHandler) LeagueAdd(c *gin.Context) {
	request := LeagueAddRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}

	if err := adminToolService.createLeagueService(&request); err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}

	utils.SuccessResponse[any](c, http.StatusOK, nil, "リーグ情報の登録に成功しました。")
}

// @Summary リーグ更新
// @Description リクエストボディに更新対象のIDを指定してリーグ情報を更新します
// @Tags league
// @Security BearerAuth
// @Param request body LeagueUpdateRequest true "リーグ情報"
// @Success 200 {object} utils.BasicResponse "成功"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/league/update [put]
func (h *AdminToolHandler) LeagueUpdate(c *gin.Context) {
	request := LeagueUpdateRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}
	log.Println("重複検索が正常にリターンはされているよ:handler")
	if err := adminToolService.UpdateLeagueService(&request); err != nil {
		log.Printf("リクエストエラー: %v", err)
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}

	utils.SuccessResponse[any](c, http.StatusOK, nil, "スポーツが正常に更新されました。")
}

// @Summary リーグ削除
// @Description リクエストボディに削除対象のIDを指定してリーグ情報を削除します
// @Tags league
// @Security BearerAuth
// @Param request body DeleteRequest true "リーグID"
// @Success 200 {object} utils.BasicResponse "成功"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/league/delete [delete]
func (h *AdminToolHandler) DeleteLeague(c *gin.Context) {
	request := DeleteRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}

	if err := adminToolService.deleteLeagueService(&request); err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}

	utils.SuccessResponse[any](c, http.StatusOK, nil, "リーグが正常に削除されました。")
}

// @Summary チーム全件検索
// @Description チーム情報のレコードを全件取得して、一覧として表示する。
// @Tags team
// @Secrity BearerAuth
// @Param keyword query string false "キーワード"
// @Success 200 {object} utils.BasicResponse "成功"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/team/teams [get]
func (h *AdminToolHandler) GetTeams(c *gin.Context) {
	keyword := c.DefaultQuery("keyword", "")
	log.Println("キーワード:", keyword)
	team := []models.Team{}

	team, err := adminToolService.getTeamService(keyword)

	if err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}

	log.Println(team)

	utils.SuccessResponse[any](c, http.StatusOK, team, "チームの検索に成功しました。")
}

// @Summary チームの追加
// @Description リクエストからチーム情報を取得後、重複確認を行い登録する。
// @Tags team
// @Security BearerAuth
// @Param request body TeamAddRequest true "チーム情報"
// @Success 200 {object} utils.BasicResponse "成功"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/team/teamAdd [post]
func (h *AdminToolHandler) TeamAdd(c *gin.Context) {
	request := TeamAddRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}

	if err := adminToolService.createTeamService(&request); err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}

	utils.SuccessResponse[any](c, http.StatusOK, nil, "チーム情報の登録に成功しました。")
}

// @Summary チーム情報更新
// @Description リクエストボディに更新対象のIDを指定してチーム情報を更新します
// @Tags team
// @Security BearerAuth
// @Param request body TeamUpdateRequest true "チーム情報"
// @Success 200 {object} utils.BasicResponse "成功"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/team/update [put]
func (h *AdminToolHandler) TeamUpdate(c *gin.Context) {
	request := TeamUpdateRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}
	log.Println("重複検索が正常にリターンはされているよ:handler")
	if err := adminToolService.UpdateTeamService(&request); err != nil {
		log.Printf("リクエストエラー: %v", err)
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}

	utils.SuccessResponse[any](c, http.StatusOK, nil, "チーム情報が正常に更新されました。")
}

// @Summary チーム削除
// @Description リクエストボディに削除対象のIDを指定してチーム情報を削除します
// @Tags team
// @Security BearerAuth
// @Param request body DeleteRequest true "チームID"
// @Success 200 {object} utils.BasicResponse "成功"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/team/delete [delete]
func (h *AdminToolHandler) DeleteTeam(c *gin.Context) {
	request := DeleteRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}

	if err := adminToolService.deleteTeamService(&request); err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}

	utils.SuccessResponse[any](c, http.StatusOK, nil, "チームが正常に削除されました。")
}

func NewAdminToolHandler() *AdminToolHandler {
	return &AdminToolHandler{}
}
