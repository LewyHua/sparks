package video

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	eclient "go.etcd.io/etcd/client/v3"
	eresolver "go.etcd.io/etcd/client/v3/naming/resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sparks/config/constant"
	proto "sparks/grpc_gen/video"
	"sparks/utils"
	"strings"
)

var videoClient proto.VideoServiceClient
var MyEtcdURL = "http://localhost:2379"

const (
	maxFileSize = 50 * 1024 * 1024
	minFileSize = 1 * 1024 * 1024
)

func InitVideoClient() {
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
	etcdTarget := fmt.Sprintf("etcd:///%s", constant.VideoServiceName)

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

	videoClient = proto.NewVideoServiceClient(conn)
}

func Action(c *gin.Context) {
	userID, _ := utils.GetCurrentUserID(c)
	title := c.PostForm("title")
	file, err := c.FormFile("file")
	if err != nil || title == "" || file.Size == 0 {
		c.JSON(http.StatusOK, proto.PublishVideoResponse{
			StatusCode: utils.CodeInvalidParam,
			StatusMsg:  utils.MapErrMsg(utils.CodeInvalidParam),
		})
		return
	}
	// 检验文件类型
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !isValidFileType(ext) {
		zap.L().Error("file type error")
		c.JSON(http.StatusOK, proto.PublishVideoResponse{
			StatusCode: utils.CodeInvalidFileType,
			StatusMsg:  utils.MapErrMsg(utils.CodeInvalidFileType),
		})
		return
	}

	// 校验文件大小
	if file.Size > maxFileSize || file.Size < minFileSize {
		c.JSON(http.StatusOK, proto.PublishVideoResponse{
			StatusCode: utils.CodeInvalidFileSize,
			StatusMsg:  utils.MapErrMsg(utils.CodeInvalidFileSize),
		})
		return
	}

	// 打开上传的文件
	src, err := file.Open()
	defer src.Close()
	if err != nil {
		zap.L().Error("打开文件失败", zap.Error(err))
		c.JSON(http.StatusInternalServerError, proto.PublishVideoResponse{
			StatusCode: utils.CodeServerBusy,
			StatusMsg:  utils.MapErrMsg(utils.CodeServerBusy),
		})
		return
	}

	// 读取文件内容
	data, err := io.ReadAll(src)
	if err != nil {
		zap.L().Error("读取文件失败", zap.Error(err))
		c.JSON(http.StatusInternalServerError, proto.PublishVideoResponse{
			StatusCode: utils.CodeServerBusy,
			StatusMsg:  utils.MapErrMsg(utils.CodeServerBusy),
		})
		return
	}

	// 生成视频文件名
	videoFileName := strings.Replace(uuid.New().String(), "-", "", -1) + ".mp4"
	// 生成视频路径，用于本地存储
	videoPath := constant.FileLocalPath + videoFileName
	// 将文件写入本地，然后通过MQ异步处理视频的上传
	err = os.WriteFile(videoPath, data, constant.FileMode)
	if err != nil {
		zap.L().Error("写入文件失败", zap.Error(err))
		c.JSON(http.StatusInternalServerError, proto.PublishVideoResponse{
			StatusCode: utils.CodeServerBusy,
			StatusMsg:  utils.MapErrMsg(utils.CodeServerBusy),
		})
		return
	}

	req := &proto.PublishVideoRequest{
		UserId:    userID,
		Title:     title,
		VideoName: videoFileName,
	}

	resp, err := videoClient.PublishVideo(c, req)
	if err != nil {
		zap.L().Error("PublishVideo err.", zap.Error(err))
		c.JSON(http.StatusOK, proto.PublishVideoResponse{
			StatusCode: utils.CodeServerBusy,
			StatusMsg:  utils.MapErrMsg(utils.CodeServerBusy),
		})
		return
	}
	c.JSON(http.StatusOK, resp)

}
func List(c *gin.Context) {

}

func Feed(c *gin.Context) {

}

// 校验文件类型是否为视频类型
func isValidFileType(fileExt string) bool {
	validExts := []string{".mp4", ".avi", ".mov"}
	for _, ext := range validExts {
		if fileExt == ext {
			return true
		}
	}
	return false
}
