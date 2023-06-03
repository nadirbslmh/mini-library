package model

import (
	"time"

	"gorm.io/gorm"
)

type Rent struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	UserID    int            `json:"user_id"`
	BookID    int            `json:"book_id"`
}
