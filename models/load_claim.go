package models

import (
	"time"

	"gorm.io/gorm"
)

type LoadClaim struct {
	ID        uint `gorm:"primaryKey"`
	LoadID    uint `json:"load_id"`
	UserID    uint `json:"user_id"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
