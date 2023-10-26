package models

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID        int64 `gorm:"primarykey"`
	UserID    int64
	VideoID   int64
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (c *Comment) TableName() string {
	return "comment"
}
