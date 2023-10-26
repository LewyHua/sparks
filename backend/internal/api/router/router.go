package router

import (
	"github.com/gin-gonic/gin"
	"sparks/internal/api/biz/comment"
	"sparks/internal/api/biz/favorite"
	"sparks/internal/api/biz/relation"
	"sparks/internal/api/biz/user"
	"sparks/internal/api/biz/video"
)

func InitRouter(r *gin.Engine) {
	r.Static("/public", "public/")
	apiRouter := r.Group("/sparks")

	// user
	userGroup := apiRouter.Group("/user")
	{
		userGroup.GET("/", user.Info)
		userGroup.POST("/register/", user.Register)
		userGroup.POST("/login/", user.Login)
	}

	// video
	videoGroup := apiRouter.Group("/video")
	{
		videoGroup.POST("/publish/", video.Action)
		videoGroup.GET("/list/", video.List)
	}

	// comment
	commentGroup := apiRouter.Group("/comment")
	{
		commentGroup.POST("/action/", comment.Action)
		commentGroup.GET("/list/", comment.List)
	}

	// favorite
	favoriteGroup := apiRouter.Group("/favorite")
	{
		favoriteGroup.POST("/action/", favorite.Action)
		favoriteGroup.GET("/list/", favorite.List)
	}

	// 社交接口
	relationGroup := apiRouter.Group("/relation")
	{
		relationGroup.POST("/action/", relation.Action)
		relationGroup.GET("/follow/list/", relation.FollowList)
		relationGroup.GET("/follower/list/", relation.FollowerList)
	}

}
