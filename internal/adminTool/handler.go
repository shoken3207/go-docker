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
// @Tags adminStadium
// @Secrity BearerAuth
// @Param keyword query string false "キーワード"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/admin/stadium/stadiums [get]
func (h *AdminToolHandler) GetStadiums(c *gin.Context) {
	keyword := c.DefaultQuery("keyword", "")
	log.Println("キーワード:", keyword)
	stadiums := []models.Stadium{}

	stadiums, err := adminToolService.getStadiumsService(keyword)

	log.Println("返却値：", stadiums)

	if err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}

	log.Println(stadiums)

	utils.SuccessResponse[any](c, http.StatusOK, stadiums, "スタジアムの検索に成功しました。")
}

// @Summary スタジアムid検索
// @Description idからスタジアム情報のレコードを取得して表示する。
// @Tags adminStadium
// @Secrity BearerAuth
// @Param id path uint true "スタジアムID"
// @Success 200 {object} utils.ApiResponse[Stadium]"成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/admin/stadium/idStadium/{id} [get]
func (h *AdminToolHandler) GetIdStadiums(c *gin.Context) {
	var request IdRequest

	if err := c.ShouldBindUri(&request); err != nil {
		log.Println("エラー:", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります。")
		return
	}

	stadium, err := adminToolService.StadiumGetIdService(request.Id)

	if err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}

	utils.SuccessResponse[any](c, http.StatusOK, stadium, "スタジアムの検索に成功しました。")
}

// @Summary スタジアム追加
// @Description リクエストからスタジアム情報を追加後、重複確認を行い登録する。
// @Tags adminStadium
// @Security BearerAuth
// @Param request body StadiumAddRequest true "スタジアム情報"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/admin/stadium/stadiumAdd [post]
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
// @Tags adminStadium
// @Security BearerAuth
// @Param id path uint true "スタジアムID"
// @Param request body StadiumUpdateRequest true "スタジアム情報"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/admin/stadium/update/{id} [put]
func (h *AdminToolHandler) StadiumUpdate(c *gin.Context) {
	var requestId IdRequest

	if err := c.ShouldBindUri(&requestId); err != nil {
		log.Println("エラー:", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります。")
		return
	}

	requestBody := StadiumUpdateRequest{}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}
	log.Println("重複検索が正常にリターンはされているよ:handler")
	if err := adminToolService.UpdateStadiumService(requestId.Id, &requestBody); err != nil {
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
// @Tags adminStadium
// @Security BearerAuth
// @Param id path uint true "スタジアムID"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/admin/stadium/delete/{id} [delete]
func (h *AdminToolHandler) DeleteStadium(c *gin.Context) {
	var request IdRequest

	if err := c.ShouldBindUri(&request); err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}

	if err := adminToolService.deleteStadiumService(request.Id); err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}

	utils.SuccessResponse[any](c, http.StatusOK, nil, "スタジアムが正常に削除されました。")
}

// @Summary スポーツ全件検索
// @Description スポーツ情報のレコードを全件取得して、一覧として表示する。
// @Tags adminSports
// @Secrity BearerAuth
// @Param keyword query string false "キーワード"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/admin/sports/sports [get]
func (h *AdminToolHandler) GetSports(c *gin.Context) {
	keyword := c.DefaultQuery("keyword", "")
	log.Println("キーワード:", keyword)
	sport := []models.Sport{}

	sport, err := adminToolService.getSportService(keyword)

	log.Println("返却値：", sport)

	if err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}

	log.Println(sport)

	utils.SuccessResponse[any](c, http.StatusOK, sport, "スポーツの検索に成功しました。")
}

// @Summary スポーツid検索
// @Description idからスポーツ情報のレコードを取得して表示する。
// @Tags adminSports
// @Secrity BearerAuth
// @Param id path uint true "スポーツID"
// @Success 200 {object} utils.ApiResponse[Sports]"成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/admin/sports/idSports/{id} [get]
func (h *AdminToolHandler) GetIdSports(c *gin.Context) {
	var request IdRequest

	if err := c.ShouldBindUri(&request); err != nil {
		log.Println("エラー:", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります。")
		return
	}

	sport, err := adminToolService.SportGetIdService(request.Id)

	if err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}

	utils.SuccessResponse[any](c, http.StatusOK, sport, "スポーツの検索に成功しました。")
}

// @Summary スポーツの追加
// @Description リクエストからスポーツ情報を取得後、重複確認を行い登録する。
// @Tags adminSports
// @Security BearerAuth
// @Param request body SportsAddRequest true "スポーツ情報"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/admin/sports/sportsAdd [post]
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
// @Tags adminSports
// @Security BearerAuth
// @Param id path uint true "スポーツID"
// @Param request body SportsUpdateRequest true "スポーツ情報"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/admin/sports/update/{id} [put]
func (h *AdminToolHandler) SportsUpdate(c *gin.Context) {
	var requestId IdRequest

	if err := c.ShouldBindUri(&requestId); err != nil {
		log.Println("エラー:", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります。")
		return
	}

	requestBody := StadiumUpdateRequest{}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}
	log.Println("重複検索が正常にリターンはされているよ:handler")
	if err := adminToolService.UpdateStadiumService(requestId.Id, &requestBody); err != nil {
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
// @Tags adminSports
// @Security BearerAuth
// @Param id path uint true "スポーツID"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/admin/sports/delete/{id} [delete]
func (h *AdminToolHandler) DeleteSports(c *gin.Context) {
	var request IdRequest

	if err := c.ShouldBindUri(&request); err != nil {
		log.Printf("エラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}

	if err := adminToolService.deleteSportService(request.Id); err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}

	utils.SuccessResponse[any](c, http.StatusOK, nil, "スポーツが正常に削除されました。")
}

// @Summary リーグ全件検索
// @Description リーグ情報のレコードを全件取得して、一覧として表示する。
// @Tags adminLeague
// @Secrity BearerAuth
// @Param keyword query string false "キーワード"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/admin/league/leagues [get]
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

	log.Println("返却値：", league)

	log.Println(league)

	utils.SuccessResponse[any](c, http.StatusOK, league, "リーグの検索に成功しました。")
}

// @Summary リーグid検索
// @Description idからリーグ情報のレコードを取得して表示する。
// @Tags adminLeague
// @Secrity BearerAuth
// @Param id path uint true "リーグID"
// @Success 200 {object} utils.ApiResponse[League]"成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/admin/league/idLeague/{id} [get]
func (h *AdminToolHandler) GetIdLeague(c *gin.Context) {
	var request IdRequest

	if err := c.ShouldBindUri(&request); err != nil {
		log.Println("エラー:", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります。")
		return
	}

	league, err := adminToolService.LeagueGetIdService(request.Id)

	if err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}

	utils.SuccessResponse[any](c, http.StatusOK, league, "リーグの検索に成功しました。")
}

// @Summary リーグの追加
// @Description リクエストからリーグ情報を取得後、重複確認を行い登録する。
// @Tags adminLeague
// @Security BearerAuth
// @Param request body LeagueAddRequest true "リーグ情報"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/admin/league/leagueAdd [post]
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
// @Tags adminLeague
// @Security BearerAuth
// @Param id path uint true "リーグID"
// @Param request body LeagueUpdateRequest true "リーグ情報"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/admin/league/update/{id} [put]
func (h *AdminToolHandler) LeagueUpdate(c *gin.Context) {
	var requestId IdRequest

	if err := c.ShouldBindUri(&requestId); err != nil {
		log.Println("エラー:", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります。")
		return
	}

	requestBody := LeagueUpdateRequest{}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}
	log.Println("重複検索が正常にリターンはされているよ:handler")
	if err := adminToolService.UpdateLeagueService(requestId.Id, &requestBody); err != nil {
		log.Printf("リクエストエラー: %v", err)
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}

	utils.SuccessResponse[any](c, http.StatusOK, nil, "リーグが正常に更新されました。")
}

// @Summary リーグ削除
// @Description リクエストボディに削除対象のIDを指定してリーグ情報を削除します
// @Tags adminLeague
// @Security BearerAuth
// @Param id path uint true "リーグID"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/admin/league/delete/{id} [delete]
func (h *AdminToolHandler) DeleteLeague(c *gin.Context) {
	var request IdRequest

	if err := c.ShouldBindUri(&request); err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}

	if err := adminToolService.deleteLeagueService(request.Id); err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}

	utils.SuccessResponse[any](c, http.StatusOK, nil, "リーグが正常に削除されました。")
}

// @Summary チーム全件検索
// @Description チーム情報のレコードを全件取得して、一覧として表示する。
// @Tags adminTeam
// @Secrity BearerAuth
// @Param keyword query string false "キーワード"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/admin/team/teams [get]
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

// @Summary チームid検索
// @Description idからチーム情報のレコードを取得して表示する。
// @Tags adminTeam
// @Secrity BearerAuth
// @Param id path uint true "チームID"
// @Success 200 {object} utils.ApiResponse[Team]"成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/admin/team/idTeam/{id} [get]
func (h *AdminToolHandler) GetIdTeam(c *gin.Context) {
	var request IdRequest

	if err := c.ShouldBindUri(&request); err != nil {
		log.Println("エラー:", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります。")
		return
	}

	team, err := adminToolService.TeamGetIdService(request.Id)

	if err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}
	log.Println("チーム名:", team.Name)

	utils.SuccessResponse[any](c, http.StatusOK, team, "チームの検索に成功しました。")
}

// @Summary チームの追加
// @Description リクエストからチーム情報を取得後、重複確認を行い登録する。
// @Tags adminTeam
// @Security BearerAuth
// @Param request body TeamAddRequest true "チーム情報"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/admin/team/teamAdd [post]
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
// @Tags adminTeam
// @Security BearerAuth
// @Param id path uint true "チームID"
// @Param request body TeamUpdateRequest true "チーム情報"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/admin/team/update/{id} [put]
func (h *AdminToolHandler) TeamUpdate(c *gin.Context) {
	var requestId IdRequest

	if err := c.ShouldBindUri(&requestId); err != nil {
		log.Println("リクエストエラー(URI):", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります。")
		return
	}

	requestBody := TeamUpdateRequest{}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Printf("リクエストエラー(ボディ): %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}
	if err := adminToolService.UpdateTeamService(requestId.Id, &requestBody); err != nil {
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
// @Tags adminTeam
// @Security BearerAuth
// @Param id path uint true "チームID"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/admin/team/delete/{id} [delete]
func (h *AdminToolHandler) DeleteTeam(c *gin.Context) {
	var request IdRequest

	if err := c.ShouldBindUri(&request); err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}

	if err := adminToolService.deleteTeamService(request.Id); err != nil {
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
