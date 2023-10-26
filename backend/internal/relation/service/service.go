package service

import (
	"context"
	proto "sparks/grpc_gen/relation"
)

// implement CommentServiceServer

type RelationServiceImpl struct {
	proto.UnimplementedRelationServiceServer
}

func (r *RelationServiceImpl) RelationAction(ctx context.Context, req *proto.RelationActionRequest) (*proto.RelationActionResponse, error) {
	return nil, nil
}
func (r *RelationServiceImpl) GetFollowList(ctx context.Context, req *proto.FollowListRequest) (*proto.FollowListResponse, error) {
	return nil, nil
}
func (r *RelationServiceImpl) GetFollowerList(ctx context.Context, req *proto.FollowerListRequest) (*proto.FollowerListResponse, error) {
	return nil, nil
}
func (r *RelationServiceImpl) GetFollowListCount(ctx context.Context, req *proto.FollowListCountRequest) (*proto.FollowListCountResponse, error) {
	return nil, nil
}
func (r *RelationServiceImpl) GetFollowerListCount(ctx context.Context, req *proto.FollowerListCountRequest) (*proto.FollowerListCountResponse, error) {
	return nil, nil
}
func (r *RelationServiceImpl) IsFollowing(ctx context.Context, req *proto.IsFollowingRequest) (*proto.IsFollowingResponse, error) {
	return nil, nil
}
