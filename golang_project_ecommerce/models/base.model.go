// Package models ...
package models

import (
	"time"

	"gorm.io/gorm"
)

// Base Modal struct
type BaseModel struct {
	ID        uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt *time.Time      `gorm:"type:time" json:"created_at"`
	UpdatedAt *time.Time      `gorm:"type:time" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"type:time" json:"deleted_at"`
}
