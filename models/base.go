package models

import "gorm.io/gorm"

type BaseModel struct {
    gorm.Model
}

func (BaseModel) BeforeDelete(tx *gorm.DB) error {
    tx.Statement.Unscoped = true
    return nil
}