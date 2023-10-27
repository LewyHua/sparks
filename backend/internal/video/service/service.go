package service

import (
	"context"
	"go.uber.org/zap"
	"os"
	"sparks/config/constant"
	"sparks/dal/mysql"
	proto "sparks/grpc_gen/video"
	"sparks/model"
	"sparks/utils"
	"time"
)

// implement CommentServiceServer

type VideoServiceImpl struct {
	proto.UnimplementedVideoServiceServer
}

func (v *VideoServiceImpl) VideoFeed(ctx context.Context, req *proto.VideoFeedRequest) (*proto.VideoFeedResponse, error) {
	return nil, nil
}
func (v *VideoServiceImpl) PublishVideo(ctx context.Context, req *proto.PublishVideoRequest) (*proto.PublishVideoResponse, error) {

	// 上传视频
	go func() {
		videoPath := constant.FileLocalPath + req.GetVideoName()
		defer func(name string) {
			err := os.Remove(name)
			if err != nil {
				zap.L().Error("删除本地视频失败", zap.Error(err))
			}
		}(videoPath)

		// 本地视频路径
		zap.L().Info("上传视频到Kodo", zap.String("videoPath", videoPath))
		// 视频存储到oss
		if err := utils.UploadToKoDo(req.GetVideoName()); err != nil {
			zap.L().Error("上传视频到Kodo失败", zap.Error(err))
			return
		}

		// 视频信息存储到MySQL
		video := model.Video{
			AuthorId:  req.GetUserId(),
			VideoUrl:  "s36dukovu.hn-bkt.clouddn.com/" + req.GetVideoName(),
			CoverUrl:  "s36dukovu.hn-bkt.clouddn.com/cover/" + req.GetVideoName(),
			Title:     req.GetTitle(),
			CreatedAt: time.Now().Unix(),
		}
		mysql.InsertVideo(&video)
		zap.L().Info("视频信息存储到MySQL", zap.Any("video", video))
	}()

	return &proto.PublishVideoResponse{
		StatusCode: utils.CodeSuccess,
		StatusMsg:  utils.MapErrMsg(utils.CodeSuccess),
	}, nil

}

func (v *VideoServiceImpl) GetPublishVideoList(ctx context.Context, req *proto.PublishVideoListRequest) (*proto.PublishVideoListResponse, error) {
	return nil, nil
}
func (v *VideoServiceImpl) GetWorkCount(ctx context.Context, req *proto.GetWorkCountRequest) (*proto.GetWorkCountResponse, error) {
	return nil, nil
}
