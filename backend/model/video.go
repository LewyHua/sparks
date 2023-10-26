package model

import (
	"gorm.io/gorm"
)

type Video struct {
	ID        int64  `gorm:"primaryKey"`
	AuthorId  int64  `gorm:"index"`
	VideoUrl  string `gorm:"not null"`
	CoverUrl  string `gorm:"not null"`
	Title     string
	CreatedAt int64 `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt
}

func (*Video) TableName() string {
	return "video"
}

type VideoResponse struct {
	ID            int64        `json:"id"`
	Author        UserResponse `json:"author"`
	PlayUrl       string       `json:"play_url"`
	CoverUrl      string       `json:"cover_url"`
	FavoriteCount int64        `json:"favorite_count"` //点赞数
	CommentCount  int64        `json:"comment_count"`  //评论数
	IsFavorite    bool         `json:"is_favorite"`    //是否点赞
	Title         string       `json:"title"`          //视频标题
}

// VideoListResponse 用户所有投稿过的视频
type VideoListResponse struct {
	Response
	VideoResponse []VideoResponse `json:"video_list"`
}

// FeedListResponse 投稿时间倒序的视频列表
type FeedListResponse struct {
	Response
	NextTime      int64           `json:"next_time"`
	VideoResponse []VideoResponse `json:"video_list"`
}
