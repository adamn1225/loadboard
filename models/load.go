package models

import (
	"time"

	"gorm.io/gorm"
)

type Load struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Weight      string         `json:"weight"`
	Dimensions  string         `json:"dimensions"`
	Pickup      string         `json:"pickup"`
	Dropoff     string         `json:"dropoff"`
	Date        time.Time      `json:"date"`
	CreatedBy   uint           `json:"created_by"` // FK to User
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Status      string         `json:"status" gorm:"default:open"` // open, claimed, delivered

}
