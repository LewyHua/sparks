package handler

import (
	"context"
	proto "sparks/grpc_gen/comment"
)

// implement CommentServiceServer

type CommentServiceImpl struct {
	proto.UnimplementedCommentServiceServer
}

func (c *CommentServiceImpl) CommentAction(context.Context, *proto.CommentActionRequest) (*proto.CommentActionResponse, error) {
	return &proto.CommentActionResponse{
		StatusCode: 200,
		StatusMsg:  "nice",
	}, nil
}

func (c *CommentServiceImpl) GetCommentList(context.Context, *proto.CommentListRequest) (*proto.CommentListResponse, error) {
	return nil, nil
}

func (c *CommentServiceImpl) GetCommentCount(context.Context, *proto.CommentCountRequest) (*proto.CommentCountResponse, error) {
	return nil, nil
}
