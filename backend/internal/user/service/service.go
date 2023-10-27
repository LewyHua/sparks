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

func (u UserServiceImpl) Register(ctx context.Context, req *proto.UserRegisterRequest) (resp *proto.UserRegisterResponse, err error) {
	username := req.Username
	password := req.Password
	// 加密密码
	encryptPassword, err := utils.EncryptPassword(password)
	if err != nil {
		zap.L().Error("encrypt password failed", zap.Error(err))
		return &proto.UserRegisterResponse{
			StatusCode: utils.CodeServerBusy,
			StatusMsg:  utils.MapErrMsg(utils.CodeServerBusy),
		}, nil
	}
	// 数据入库
	userModel := model.User{
		Username:  username,
		Password:  encryptPassword,
		Signature: "这个人很懒，什么都没有留下",
	}
	err = mysql.CreateUser(&userModel)
	if err != nil {
		zap.L().Error("create user failed", zap.Error(err))
		return &proto.UserRegisterResponse{
			StatusCode: utils.CodeUsernameAlreadyExists,
			StatusMsg:  utils.MapErrMsg(utils.CodeUsernameAlreadyExists),
		}, nil
	}

	// 将token存入redis

	// 用户名存入Bloom Filter
	go utils.AddToUserBloom(username)

	// 返回响应
	return &proto.UserRegisterResponse{
		StatusCode: utils.CodeSuccess,
		StatusMsg:  utils.MapErrMsg(utils.CodeSuccess),
		UserId:     userModel.ID,
		Token:      "fake_token",
	}, nil
}

func (u UserServiceImpl) Login(ctx context.Context, req *proto.UserLoginRequest) (*proto.UserLoginResponse, error) {
	// 布隆过滤器检查用户名是否存在
	exist := utils.TestUserBloom(req.Username)
	if !exist {
		zap.L().Info("Check user exists info:", zap.Bool("exist", exist))
		return &proto.UserLoginResponse{
			StatusCode: utils.CodeUsernameNotFound,
			StatusMsg:  utils.MapErrMsg(utils.CodeUsernameNotFound),
		}, nil
	}

	// 用户名存在，检查密码是否正确
	userModel, _, err := mysql.FindUserByName(req.Username)
	if err != nil {
		zap.L().Error("find user by name failed", zap.Error(err))
		return &proto.UserLoginResponse{
			StatusCode: utils.CodeUsernameNotFound,
			StatusMsg:  utils.MapErrMsg(utils.CodeUsernameNotFound),
		}, nil
	}

	// 检查密码是否正确
	err = utils.ComparePassword(userModel.Password, req.Password)
	if err != nil {
		return &proto.UserLoginResponse{
			StatusCode: utils.CodeWrongLoginCredentials,
			StatusMsg:  utils.MapErrMsg(utils.CodeWrongLoginCredentials),
		}, nil
	}

	// 生成token
	token := utils.GenerateToken(userModel.ID, userModel.Username)
	if err != nil {
		return &proto.UserLoginResponse{
			StatusCode: utils.CodeServerBusy,
			StatusMsg:  utils.MapErrMsg(utils.CodeServerBusy),
		}, nil
	}

	// TODO 将token存入redis

	// 返回响应
	return &proto.UserLoginResponse{
		StatusCode: utils.CodeSuccess,
		StatusMsg:  utils.MapErrMsg(utils.CodeSuccess),
		UserId:     userModel.ID,
		Token:      token,
	}, nil
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
