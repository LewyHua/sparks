package models

import (
	"gorm.io/gorm"
	"time"
)

type Video struct {
	ID        int64 `gorm:"primarykey"`
	AuthorID  int64
	Title     string
	FilePath  string
	CoverPath string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (v *Video) TableName() string {
	return "video"
}
