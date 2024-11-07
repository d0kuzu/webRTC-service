package company_model

import (
	"aisale/database/models/user_model"
	"gorm.io/gorm"
	"time"
)

type Company struct {
	ID        uint              `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string            `json:"name" gorm:"not null"`
	Users     []user_model.User `json:"-" gorm:"foreignKey:CompanyID"` // One-to-many relationship
	CreatedAt time.Time         `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time         `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt    `json:"-" gorm:"index"`
}
