package mw

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sparks/middlewares/redis"
	"sparks/utils"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

// Auth 鉴权中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		// 解析token
		claims, err := utils.ParseToken(token)
		if err != nil {
			// token有误（含len(token)=0的情况），阻止后面函数执行
			c.Abort()
			c.JSON(http.StatusUnauthorized, Response{
				StatusCode: utils.CodeInvalidToken,
				StatusMsg:  utils.MapErrMsg(utils.CodeInvalidParam),
			})
			return
		}
		// 查看token是否在redis中, 若在，给token续期, 若不在，则阻止后面函数执行
		exist := redis.TokenIsExisted(claims.ID)
		if !exist {
			// token有误，阻止后面函数执行
			c.Abort()
			c.JSON(http.StatusOK, Response{
				StatusCode: utils.CodeInvalidToken,
				StatusMsg:  utils.MapErrMsg(utils.CodeInvalidParam),
			})
			return
		}
		go func(id int64) {
			// 给token续期，并标记位活跃
			redis.SetToken(id, token)
		}(claims.ID)

		c.Set(utils.ContextUserIDKey, claims.ID)
		c.Next()
	}
}

// AuthWithoutLogin 未登录情况，若携带token,解析用户id放入context;如果没有携带，则将用户id默认为0
func AuthWithoutLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.PostForm("token")
		var userId int64
		var tokenValid bool
		claims, err := utils.ParseToken(token)
		if err != nil {
			// token有误或token为空，未登录
			tokenValid = false
			userId = 0
		} else {
			// 查看token是否在redis中, 若在，则返回用户id, 并且给token续期, 若不在，则将userID设为0
			exist := redis.TokenIsExisted(claims.ID)
			if !exist {
				// token有误，设置userId为0,tokenValid为false
				userId = 0
				tokenValid = false
			} else {
				userId = claims.ID
				go func(id int64) {
					// 给token续期
					redis.SetToken(id, token)
				}(userId)
				tokenValid = true
			}
		}
		c.Set(utils.TokenValid, tokenValid)
		c.Set(utils.ContextUserIDKey, userId)
		c.Next()
	}
}

// AuthBody 若token在请求体里，解析token
func AuthBody() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.PostForm("token")
		// 解析token
		claims, err := utils.ParseToken(token)
		if err != nil {
			// token有误（含len(token)=0的情况），阻止后面函数执行
			c.Abort()
			c.JSON(http.StatusUnauthorized, Response{
				StatusCode: utils.CodeInvalidToken,
				StatusMsg:  utils.MapErrMsg(utils.CodeInvalidParam),
			})
			return
		}
		// 查看token是否在redis中, 若在，给token续期, 若不在，则阻止后面函数执行
		exist := redis.TokenIsExisted(claims.ID)
		if !exist {
			// token有误，阻止后面函数执行
			c.Abort()
			c.JSON(http.StatusOK, Response{
				StatusCode: utils.CodeInvalidToken,
				StatusMsg:  utils.MapErrMsg(utils.CodeInvalidParam),
			})
			return
		}
		go func(id int64) {
			// 给token续期
			redis.SetToken(id, token)
		}(claims.ID)

		c.Set(utils.ContextUserIDKey, claims.ID)
		c.Next()
	}
}
