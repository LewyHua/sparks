package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"sparks/config"
	"sparks/dal/mysql"
	"sparks/internal/api/biz/comment"
	"sparks/internal/api/router"
	"sparks/logger"
	"sparks/middlewares/redis"
)

func main() {
	// 加载配置
	if err := config.Init(); err != nil {
		zap.L().Error("Load config failed, err:%v\n", zap.Error(err))
		return
	}
	// 加载日志
	if err := logger.Init(config.Conf.LogConfig, config.Conf.Mode); err != nil {
		zap.L().Error("Init logger failed, err:%v\n", zap.Error(err))
		return
	}

	// 初始化数据库
	if err := mysql.Init(config.Conf); err != nil {
		zap.L().Error("Init redis failed, err:%v\n", zap.Error(err))
		return
	}

	// 初始化中间件: redis
	if err := redis.Init(config.Conf); err != nil {
		zap.L().Error("Init redis failed, err:%v\n", zap.Error(err))
		return
	}

	// 初始化 grpc 客户端
	comment.InitializeCommentClient()

	r := gin.Default()
	router.InitRouter(r)
	r.Run(":8080")
}
