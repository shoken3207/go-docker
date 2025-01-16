package models

type User struct {
	BaseModel
	Username        string           `json:"username" gorm:"size:255;not null;unique"`
	Email           string           `json:"email" gorm:"size:100;not null;unique"`
	PassHash        string           `json:"passHash" gorm:"not null;column:pass_hash"`
	Name            string           `json:"name" gorm:"size:100;not null"`
	Description     *string          `json:"description" gorm:"type:text"`
	ProfileImage    *string          `json:"profileImage" gorm:"column:profile_image"`
	FileId          *string          `json:"fileId" gorm:"column:file_id;unique"`
	FavoriteTeams   []FavoriteTeam   `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
	ExpeditionLikes []ExpeditionLike `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
	Expeditions     []Expedition     `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
}

func (u *User) GetDescription() string {
	if u.Description != nil {
		return *u.Description
	}
	return ""
}
func (u *User) SetDescription(description string) {
	if description == "" {
		u.Description = nil
	} else {
		u.Description = &description
	}
}

func (u *User) GetFileId() string {
	if u.FileId != nil {
		return *u.FileId
	}
	return ""
}
func (u *User) SetFileId(fileId string) {
	if fileId == "" {
		u.FileId = nil
	} else {
		u.FileId = &fileId
	}
}

func (u *User) GetProfileImage() string {
	if u.ProfileImage != nil {
		return *u.ProfileImage
	}
	return ""
}
func (u *User) SetProfileImage(profileImage string) {
	if profileImage == "" {
		u.ProfileImage = nil
	} else {
		u.ProfileImage = &profileImage
	}
}