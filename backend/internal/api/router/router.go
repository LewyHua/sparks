package router

import (
	"github.com/gin-gonic/gin"
	"sparks/internal/api/biz/comment"
	"sparks/internal/api/biz/favorite"
	"sparks/internal/api/biz/relation"
	"sparks/internal/api/biz/user"
	"sparks/internal/api/biz/video"
	"sparks/internal/api/mw"
)

func InitRouter(r *gin.Engine) {
	r.Static("/public", "public/")
	apiRouter := r.Group("/sparks")

	// user
	userGroup := apiRouter.Group("/user")
	{
		userGroup.GET("/", mw.AuthWithoutLogin(), user.Info)
		userGroup.POST("/register/", user.Register)
		userGroup.POST("/login/", user.Login)
	}

	// video
	videoGroup := apiRouter.Group("/video")
	{
		videoGroup.POST("/publish/", mw.AuthBody(), video.Action)
		videoGroup.GET("/list/", mw.AuthWithoutLogin(), video.List)
		videoGroup.GET("/feed/", mw.AuthWithoutLogin(), video.Feed)
	}

	// comment
	commentGroup := apiRouter.Group("/comment")
	{
		commentGroup.POST("/action/", mw.AuthBody(), comment.Action)
		commentGroup.GET("/list/", mw.AuthWithoutLogin(), comment.List)
	}

	// favorite
	favoriteGroup := apiRouter.Group("/favorite")
	{
		favoriteGroup.POST("/action/", mw.AuthBody(), favorite.Action)
		favoriteGroup.GET("/list/", mw.AuthWithoutLogin(), favorite.List)
	}

	// 社交接口
	relationGroup := apiRouter.Group("/relation")
	{
		relationGroup.POST("/action/", mw.AuthBody(), relation.Action)
		relationGroup.GET("/follow/list/", mw.AuthWithoutLogin(), relation.FollowList)
		relationGroup.GET("/follower/list/", mw.AuthWithoutLogin(), relation.FollowerList)
	}

}
