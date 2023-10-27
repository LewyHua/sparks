package mysql

import (
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
	"sparks/model"
)

// AddComment 添加评论
func AddComment(comment *model.Comment) (int64, error) {
	result := DB.Model(model.Comment{}).Create(comment)
	// 判断是否创建成功
	if result.Error != nil {
		zap.L().Error("创建 Comment 失败:", zap.Error(result.Error))
		return 0, result.Error
	} else {
		return comment.ID, nil
	}
}

// FindCommentsByVideoId 根据视频ID查找评论
func FindCommentsByVideoId(videoId uint) ([]model.Comment, error) {
	comments := make([]model.Comment, 0)
	result := DB.Where("video_id = ?", videoId).Order("created_at desc").Find(&comments)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	log.Println(comments)
	return comments, nil
}

// FindCommentById 根据评论ID查找评论
func FindCommentById(commentId uint) (model.Comment, error) {
	comment := model.Comment{}
	result := DB.Find(&comment, commentId)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return comment, result.Error
	}
	return comment, nil
}

// DeleteCommentById 删除评论
func DeleteCommentById(commentId uint) error {
	result := DB.Delete(&model.Comment{}, commentId)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Println("未找到 Comment")
		return result.Error
	}
	return nil
}

// GetCommentCount 获取评论数
func GetCommentCount(videoId uint) (int64, error) {
	var cnt int64
	err := DB.Model(&model.Comment{}).Where("video_id = ?", videoId).Count(&cnt).Error
	// 返回评论数和是否查询成功
	return cnt, err
}

// IsCommentBelongsToUser 判断评论是否属于用户
func IsCommentBelongsToUser(commentId *int64, userId int64) (bool, error) {
	var cnt int64
	err := DB.Model(&model.Comment{}).Where("id = ? and user_id = ?", commentId, userId).Count(&cnt).Error
	return cnt != 0, err
}
