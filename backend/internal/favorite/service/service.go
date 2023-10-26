package service

import (
	"context"
	proto "sparks/grpc_gen/favorite"
)

type FavoriteServiceImpl struct {
	proto.UnimplementedFavoriteServiceServer
}

func (f *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *proto.FavoriteActionRequest) (*proto.FavoriteActionResponse, error) {
	return nil, nil
}
func (f *FavoriteServiceImpl) GetVideoFavoriteCount(ctx context.Context, req *proto.VideoFavoriteCountRequest) (*proto.VideoFavoriteCountResponse, error) {
	return nil, nil
}
func (f *FavoriteServiceImpl) GetUserFavoriteList(ctx context.Context, req *proto.UserFavoriteListRequest) (*proto.UserFavoriteListResponse, error) {
	return nil, nil
}
func (f *FavoriteServiceImpl) GetUserFavoriteCount(ctx context.Context, req *proto.UserFavoriteCountRequest) (*proto.UserFavoriteCountResponse, error) {
	return nil, nil
}
func (f *FavoriteServiceImpl) GetUserFavoritedCount(ctx context.Context, req *proto.UserFavoritedCountRequest) (*proto.UserFavoritedCountResponse, error) {
	return nil, nil
}
func (f *FavoriteServiceImpl) IsUserFavorite(ctx context.Context, req *proto.IsUserFavoriteRequest) (*proto.IsUserFavoriteResponse, error) {
	return nil, nil
}
