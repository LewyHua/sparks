package mysql

import (
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sparks/model"
)

func FindUserByName(username string) (user model.User, exist bool, err error) {
	user = model.User{}
	if err = DB.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, false, nil
		}
		// 处理其他查询错误
		zap.L().Error("Database err", zap.Error(err))
		return user, false, err
	}
	return user, true, nil
}

func FindUserByUserID(id uint) (user model.User, exist bool, err error) {
	user = model.User{}
	if err = DB.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, false, err
		}
		// 处理其他查询错误
		zap.L().Error("Database err", zap.Error(err))
		return user, false, err
	}
	return user, true, nil
}

func CreateUser(user *model.User) error {
	if err := DB.Create(user).Error; err != nil {
		zap.L().Error("create user failed", zap.Error(err))
		return err
	}
	return nil
}
