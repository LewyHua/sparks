package mysql

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"sparks/model"
)

// GetUserFavoriteCount 根据id查询用户点赞数
func GetUserFavoriteCount(userID int64) (int64, error) {
	var cnt int64
	err := DB.Model(&model.Favorite{}).Where("user_id = ?", userID).Count(&cnt).Error
	return cnt, err
}

// GetVideoFavoriteCountByVideoId 根据视频id查询点赞数
func GetVideoFavoriteCountByVideoId(videoID int64) (int64, error) {
	var cnt int64
	err := DB.Model(&model.Favorite{}).Where("video_id = ?", videoID).Count(&cnt).Error
	return cnt, err
}

// AddUserFavorite 添加喜欢关系
func AddUserFavorite(userID, videoID int64) bool {
	follow := model.Favorite{UserId: userID, VideoId: videoID}
	result := DB.Model(&model.Favorite{}).Create(&follow)
	return result.RowsAffected != 0
}

// DeleteUserFavorite 删除喜欢关系
func DeleteUserFavorite(userID, videoID int64) error {
	favorite := model.Favorite{UserId: userID, VideoId: videoID}
	result := DB.Delete(&model.Favorite{}, favorite)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Println("未找到喜欢关系", userID, videoID)
		return result.Error
	}
	return nil
}

func IsFavorite(userID, videoID int64) bool {
	var count int64
	DB.Model(&model.Favorite{}).
		Where("user_id = ? AND video_id = ?", userID, videoID).
		Count(&count)
	return count != 0
}

// GetFavoritesById 从数据库中获取点赞列表
func GetFavoritesById(userID int64) []uint {
	var videoList []uint
	DB.Model(&model.Favorite{}).
		Limit(30).
		Select("video_id").
		Where("user_id = ?", userID).
		Order("id desc").
		Find(&videoList)
	return videoList
}
