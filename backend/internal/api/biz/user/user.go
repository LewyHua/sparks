package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	eclient "go.etcd.io/etcd/client/v3"
	eresolver "go.etcd.io/etcd/client/v3/naming/resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"sparks/config/constant"
	proto "sparks/grpc_gen/user"
)

var userClient proto.UserServiceClient
var MyEtcdURL = "http://localhost:2379"

func InitUserClient() {
	// 创建 etcd 客户端
	etcdClient, err := eclient.NewFromURL(MyEtcdURL)
	if err != nil {
		fmt.Println("etcd client err:", err)
		zap.L().Error("etcd client failed", zap.Error(err))
		return
	}

	// 创建 etcd 实现的 grpc 服务注册发现模块 resolver
	etcdResolverBuilder, err := eresolver.NewBuilder(etcdClient)
	if err != nil {
		fmt.Println("resolver builder err:", err)
		zap.L().Error("resolver builder failed", zap.Error(err))
		return
	}

	// 拼接服务名称，需要固定定义 etcd:/// 作为前缀
	etcdTarget := fmt.Sprintf("etcd:///%s", constant.UserServiceName)

	// 创建 grpc 连接代理
	conn, err := grpc.Dial(
		// 服务名称
		etcdTarget,
		// 注入 etcd resolver
		grpc.WithResolvers(etcdResolverBuilder),
		// 声明使用的负载均衡策略为 roundrobin
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		fmt.Println("grpc dial err:", err)
		zap.L().Error("grpc dial failed", zap.Error(err))
		return
	}

	userClient = proto.NewUserServiceClient(conn)
}

func Info(c *gin.Context) {

}

func Register(c *gin.Context) {
	// 获取用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")

	// TODO 检查用户名和密码是否合法

	// 合法则注册
	resp, err := userClient.Register(c, &proto.UserRegisterRequest{
		Username: username,
		Password: password,
	})

	// 其实这里可以不用判断 err，因为即使 err != nil，resp 也不会为 nil
	// 但是防御性编程，还是判断一下
	// 而且如果 err != nil，resp 也是 nil，所以这里可以直接返回
	if err != nil {
		zap.L().Error("register failed", zap.Error(err))
		c.JSON(http.StatusOK, resp)
		return
	}

	// 注册成功，返回响应
	c.JSON(http.StatusOK, resp)
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	resp, err := userClient.Login(c, &proto.UserLoginRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		zap.L().Error("login failed", zap.Error(err))
		c.JSON(http.StatusOK, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}
