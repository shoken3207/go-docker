package user

import (
	"go-docker/internal/db"
	"go-docker/models"
)

type UserService struct{}

func (s *UserService) findUserById(userId uint) (*models.User, error) {
	user := models.User{}
	if err := db.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func NewUserService() *UserService {
	return &UserService{}
}
