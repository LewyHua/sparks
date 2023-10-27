package main

import (
	"context"
	"fmt"
	eclient "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"sparks/config"
	"sparks/config/constant"
	"sparks/dal/mysql"
	proto "sparks/grpc_gen/video"
	"sparks/internal/video/service"
	"sparks/logger"
	"sparks/middlewares/redis"
	"time"
)

const (
	// grpc 服务名
	MyService = constant.VideoServiceName
	// grpc 服务端口
	port = constant.VideoServicePort
	// etcd 端口
	MyEtcdURL = "http://localhost:2379"
)

func main() {
	// 加载配置
	if err := config.Init(); err != nil {
		zap.L().Error("Load config failed, err:%v\n", zap.Error(err))
		return
	}
	// 加载日志
	if err := logger.Init(config.Conf.LogConfig, config.Conf.Mode); err != nil {
		zap.L().Error("Init logger failed, err:%v\n", zap.Error(err))
		return
	}

	// 初始化数据库
	if err := mysql.Init(config.Conf); err != nil {
		zap.L().Error("Init redis failed, err:%v\n", zap.Error(err))
		return
	}

	// 初始化中间件: redis
	if err := redis.Init(config.Conf); err != nil {
		zap.L().Error("Init redis failed, err:%v\n", zap.Error(err))
		return
	}

	// 接收命令行指定的 grpc 服务端口
	addr := fmt.Sprintf("localhost:%s", port)

	// 创建 tcp 端口监听器
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("listen tcp failed, err: ", err)
		return
	}
	// 创建 grpc server
	server := grpc.NewServer()
	proto.RegisterVideoServiceServer(server, &service.VideoServiceImpl{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// 注册 grpc 服务节点到 etcd 中
	go registerEndPointToEtcd(ctx, addr)

	// 启动 grpc 服务
	if err := server.Serve(listener); err != nil {
		fmt.Println("start grpc server failed, err: ", err)
	}
}

func registerEndPointToEtcd(ctx context.Context, addr string) {
	// 创建 etcd 客户端
	etcdClient, _ := eclient.NewFromURL(MyEtcdURL)
	etcdManager, _ := endpoints.NewManager(etcdClient, MyService)

	// 创建一个租约，每隔 10s 需要向 etcd 汇报一次心跳，证明当前节点仍然存活
	var ttl int64 = 10
	lease, _ := etcdClient.Grant(ctx, ttl)

	// 添加注册节点到 etcd 中，并且携带上租约 id
	err := etcdManager.AddEndpoint(ctx, fmt.Sprintf("%s/%s", MyService, addr), endpoints.Endpoint{Addr: addr}, eclient.WithLease(lease.ID))
	if err != nil {
		fmt.Println("add endpoint to etcd failed, err: ", err)
		return
	}
	// 每隔 5 s进行一次延续租约的动作
	for {
		select {
		case <-time.After(5 * time.Second):
			// 续约操作
			resp, _ := etcdClient.KeepAliveOnce(ctx, lease.ID)
			fmt.Printf("keep alive resp: %+v\n", resp)
		case <-ctx.Done():
			return
		}
	}
}
