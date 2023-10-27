package mysql

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"sparks/model"
)

// AddFollow 添加关注关系
func AddFollow(userID, followID int64) error {
	follow := model.Relation{UserId: userID, FollowId: followID}
	result := DB.Model(&model.Relation{}).Create(&follow)
	// 判断是否创建成功
	if result.Error != nil {
		log.Println("创建 follow 失败:", result.Error)
		return result.Error
	} else {
		log.Println("成功创建 follow")
		return nil
	}
}

// DeleteFollowById 删除关注关系
func DeleteFollowById(userID, followID int64) error {
	follow := model.Relation{UserId: userID, FollowId: followID}
	result := DB.Delete(&model.Relation{}, follow)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Println("未找到 Follow", userID, followID)
		return result.Error
	}
	return nil
}

// GetFollowCnt 关注数
func GetFollowCnt(userID int64) (int64, error) {
	var cnt int64
	err := DB.Model(&model.Relation{}).Where("user_id = ?", userID).Count(&cnt).Error
	// 返回评论数和是否查询成功
	return cnt, err
}

// GetFollowerCnt 粉丝数
func GetFollowerCnt(userID int64) (int64, error) {
	var cnt int64
	err := DB.Model(&model.Relation{}).Where("follow_id = ?", userID).Count(&cnt).Error
	// 返回评论数和是否查询成功
	return cnt, err
}

// IsFollowing 是否关注
func IsFollowing(userA int64, userB int64) bool {
	var count int64
	DB.Model(&model.Relation{}).
		Where("user_id = ? AND follow_id = ?", userA, userB).
		Count(&count)
	return count > 0
}

// GetFollowList 获取关注列表
func GetFollowList(userID int64) ([]int64, error) {
	var result []int64
	err := DB.Model(&model.Relation{}).Select("follow_id").Where("user_id = ?", userID).Scan(&result).Error
	return result, err
}

// GetFollowerList 获取粉丝列表
func GetFollowerList(userID int64) ([]int64, error) {
	var result []int64
	err := DB.Model(&model.Relation{}).Select("user_id").Where("follow_id = ?", userID).Scan(&result).Error
	return result, err
}
