package models

import "time"
import "gorm.io/gorm"

type User struct {
	ID        int64 `gorm:"primarykey"`
	Username  string
	Password  string
	Salt      string
	Avatar    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Status    int
}

func (u *User) TableName() string {
	return "user"
}
