package models

// 点赞表
type Favorite struct {
	ID      int64 `gorm:"primarykey"`
	userID  int64
	videoID int64
}

func (f *Favorite) TableName() string {
	return "favorite"
}
