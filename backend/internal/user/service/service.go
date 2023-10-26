package service

import (
	"context"
	"go.uber.org/zap"
	"sparks/dal/mysql"
	proto "sparks/grpc_gen/user"
	"sparks/model"
	"sparks/utils"
)

// implement CommentServiceServer

type UserServiceImpl struct {
	proto.UnimplementedUserServiceServer
}

func (u UserServiceImpl) Register(ctx context.Context, req *proto.UserRegisterRequest) (*proto.UserRegisterResponse, error) {
	username := req.Username
	password := req.Password
	encryptPassword, err := utils.EncryptPassword(password)
	if err != nil {
		zap.L().Error("encrypt password failed", zap.Error(err))
		return nil, err
	}
	err = mysql.CreateUser(&model.User{UserName: username, Password: encryptPassword})
	if err != nil {
		zap.L().Error("create user failed", zap.Error(err))
		return nil, err
	}
	return &proto.UserRegisterResponse{
		StatusCode: 200,
		StatusMsg:  "Created",
		UserId:     0,
		Token:      "",
	}, nil
}

func (u UserServiceImpl) Login(ctx context.Context, req *proto.UserLoginRequest) (*proto.UserLoginResponse, error) {
	return nil, nil
}

func (u UserServiceImpl) GetUserInfoById(ctx context.Context, req *proto.UserInfoByIdRequest) (*proto.UserInfoByIdResponse, error) {
	return nil, nil
}

func (u UserServiceImpl) GetUserInfoByName(ctx context.Context, req *proto.UserInfoByUsernameRequest) (*proto.UserInfoByUsernameResponse, error) {
	return nil, nil
}

func (u UserServiceImpl) CheckUserExists(ctx context.Context, req *proto.UserExistsRequest) (*proto.UserExistsResponse, error) {
	return nil, nil
}
