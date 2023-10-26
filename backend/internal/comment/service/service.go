package service

import (
	"context"
	proto "sparks/grpc_gen/comment"
)

// implement CommentServiceServer

type CommentServiceImpl struct {
	proto.UnimplementedCommentServiceServer
}

func (c *CommentServiceImpl) CommentAction(ctx context.Context, req *proto.CommentActionRequest) (*proto.CommentActionResponse, error) {
	return &proto.CommentActionResponse{
		StatusCode: 200,
		StatusMsg:  "nice",
	}, nil
}

func (c *CommentServiceImpl) GetCommentList(ctx context.Context, req *proto.CommentListRequest) (*proto.CommentListResponse, error) {
	return nil, nil
}

func (c *CommentServiceImpl) GetCommentCount(ctx context.Context, req *proto.CommentCountRequest) (*proto.CommentCountResponse, error) {
	return nil, nil
}
