package comment

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	eclient "go.etcd.io/etcd/client/v3"
	eresolver "go.etcd.io/etcd/client/v3/naming/resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
	"sparks/config/constant"
	proto "sparks/grpc_gen/comment"
)

// CommentServiceClient is the client API for CommentService service.
var commentClient proto.CommentServiceClient
var MyEtcdURL = "http://localhost:2379"

func InitCommentClient() {
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
	etcdTarget := fmt.Sprintf("etcd:///%s", constant.CommentServiceName)

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

	commentClient = proto.NewCommentServiceClient(conn)
}

func Action(c *gin.Context) {
	resp, err := commentClient.CommentAction(context.Background(), &proto.CommentActionRequest{})
	if err != nil {
		zap.L().Error("comment action failed", zap.Error(err))
		return
	}
	zap.L().Info("comment action", zap.Any("resp", resp))
}

func List(c *gin.Context) {

}
