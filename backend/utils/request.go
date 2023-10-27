package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const ContextUserIDKey = "user_id"
const TokenValid = "tokenValid"

var ErrorUserNotLogin = errors.New("用户未登录")

func GetCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(ContextUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	return
}
