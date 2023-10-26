package service

import (
	"context"
	proto "sparks/grpc_gen/video"
)

// implement CommentServiceServer

type VideoServiceImpl struct {
	proto.UnimplementedVideoServiceServer
}

func (v *VideoServiceImpl) VideoFeed(ctx context.Context, req *proto.VideoFeedRequest) (*proto.VideoFeedResponse, error) {
	return nil, nil
}
func (v *VideoServiceImpl) PublishVideo(ctx context.Context, req *proto.PublishVideoRequest) (*proto.PublishVideoResponse, error) {
	return nil, nil
}

func (v *VideoServiceImpl) GetPublishVideoList(ctx context.Context, req *proto.PublishVideoListRequest) (*proto.PublishVideoListResponse, error) {
	return nil, nil
}
func (v *VideoServiceImpl) GetWorkCount(ctx context.Context, req *proto.GetWorkCountRequest) (*proto.GetWorkCountResponse, error) {
	return nil, nil
}
