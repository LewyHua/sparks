package model

type Favorite struct {
	ID      int64 `gorm:"primaryKey"`
	UserId  int64 `gorm:"index;not null"`
	VideoId int64 `gorm:"index;not null"`
}

func (*Favorite) TableName() string {
	return "favorite"
}

type FavoriteAction struct {
	UserId     int64
	VideoId    int64
	ActionType int
}

type FavoriteListResponse struct {
	FavoriteRes   Response
	VideoResponse []VideoResponse `json:"video_list"`
}
