package user

import (
	"errors"
	"go-docker/internal/db"
	"go-docker/models"
	"go-docker/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService struct{}

func (s *UserService) createUserResponse(user *models.User) *UserResponse {
	userResponse := UserResponse{
		Id:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		Name:         user.Name,
		Description:  user.Description,
		ProfileImage: user.ProfileImage,
	}
	return &userResponse
}

func (s *UserService) findUserById(userId uint) (*models.User, error) {
	user := models.User{}
	if err := db.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		log.Printf("ユーザーデータ取得エラー: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewCustomError(http.StatusNotFound, "ユーザーデータが見つかりませんでした。。")
		} else {
			return nil, utils.NewCustomError(http.StatusInternalServerError, "ユーザーデータ取得に失敗しました。")
		}
	}
	return &user, nil
}

func (s *UserService) updateUser(userId *uint, request *UpdateUserRequestBody) (*models.User, error) {
	user, err := s.findUserById(*userId)
	if err != nil {
		return nil, err
	}

	user.Name = request.Name
	user.Description = request.Description
	user.ProfileImage = request.ProfileImage
	user.FileId = request.FileId

	if err := db.DB.Model(user).Updates(models.User{Name: request.Name, Description: request.Description, ProfileImage: request.ProfileImage, FileId: request.FileId}).Error; err != nil {
		log.Printf("ユーザー更新エラー: %v", err)
		return nil, utils.NewCustomError(http.StatusInternalServerError, "ユーザーデータがの更新に失敗しました。")
	}

	return user, nil
}

func (s *UserService) getUserByIdService(request *GetUserByIdRequest) (*UserResponse, error) {
	user, err := userService.findUserById(request.UserId)
	if err != nil {
		return nil, err
	}
	userResponse := s.createUserResponse(user)

	return userResponse, nil
}

func (s *UserService) getUserByUsernameService(request *GetUserByUsernameRequest) (*UserResponse, error) {
	user, err := utils.FindUserByUsername(request.Username)
	if err != nil {
		return nil, err
	}
	userResponse := s.createUserResponse(user)

	return userResponse, nil
}

func (s *UserService) validateUpdateUserRequest(c *gin.Context) (*uint, *UpdateUserRequestBody, error) {
	var requestBody UpdateUserRequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Printf("リクエストエラー: %v", err)
		return nil, nil, utils.NewCustomError(http.StatusBadRequest, "リクエストに不備があります。")
	}
	var requestPath UpdateUserRequestPath
	if err := c.ShouldBindUri(&requestPath); err != nil {
		log.Printf("リクエストエラー: %v", err)
		return nil, nil, utils.NewCustomError(http.StatusBadRequest, "リクエストに不備があります。")
	}
	userId, err := utils.StringToUint(c.GetString("userId"))
	if err != nil {
		return nil, nil, err
	}
	if *userId != requestPath.UserId {
		return nil, nil, utils.NewCustomError(http.StatusUnauthorized, "自分のユーザー情報しか更新できません。")
	}

	return userId, &requestBody, nil
}

func (s *UserService) updateUserService(userId *uint, requestBody *UpdateUserRequestBody) (*UserResponse, error) {
	user, err := userService.updateUser(userId, requestBody)
	if err != nil {
		return nil, err
	}

	userResponse := UserResponse{
		Id:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		Name:         user.Name,
		Description:  user.Description,
		ProfileImage: user.ProfileImage,
		FileId:       user.FileId,
	}

	return &userResponse, nil
}

func NewUserService() *UserService {
	return &UserService{}
}
