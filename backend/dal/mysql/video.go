package mysql

import (
	"sparks/model"
)

// FindVideoByVideoId 根据视频ID查找视频
func FindVideoByVideoId(videoID int64) (model.Video, bool) {
	video := model.Video{}
	return video, DB.Where("id = ?", videoID).First(&video).RowsAffected != 0
}

// FindVideosByAuthorId 返回查询到的列表及是否出错
// 若未找到，返回空列表
func FindVideosByAuthorId(authorID int64) ([]model.Video, bool) {
	videos := make([]model.Video, 0)
	return videos, DB.Where(" author_id = ?", authorID).Find(&videos).RowsAffected != 0
}

// FindWorkCountsByAuthorId 返回查询到的列表
func FindWorkCountsByAuthorId(authorID int64) int64 {
	var count int64
	DB.Model(&model.Video{}).Where("author_id = ?", authorID).Count(&count)
	return count
}

// InsertVideo 插入视频
func InsertVideo(video *model.Video) bool {
	result := DB.Model(model.Video{}).Create(video)
	return result.RowsAffected != 0
}

// GetAllVideos 获取所有视频
func GetAllVideos(latestTime string) []model.Video {
	videos := make([]model.Video, 0, 30)
	DB.Model(&model.Video{}).Where("created_at < ?", latestTime).Order("created_at DESC").Find(&videos)
	return videos
}
