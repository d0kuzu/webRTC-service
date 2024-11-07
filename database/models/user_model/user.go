package user_model

import (
	"aisale/database/models/company_model"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint                  `json:"id" gorm:"primaryKey;autoIncrement"`
	Username  string                `json:"username" gorm:"unique;not null"`
	Password  string                `json:"password" gorm:"not null"`
	Email     string                `json:"email" gorm:"unique;not null"`
	Aser      string                `json:"aser"`
	CompanyID uint                  `json:"company_id"`
	Company   company_model.Company `json:"-" gorm:"foreignKey:CompanyID"` // JsonIgnore equivalent
	Role      string                `json:"role" gorm:"not null"`
	CreatedAt time.Time             `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time             `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt        `json:"-" gorm:"index"`
}
