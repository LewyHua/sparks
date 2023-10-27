package utils

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"testing"
)

func TestUpload(t *testing.T) {
	accessKey = "sB5H5Jq5_RZmwXyxDeJzyR7SGtjaxERVEfNlD6wu"
	secretKey = "8xUBykY0QZPgVewWwFsJ9myXofN8ty39gzy3RfM6"
	domain := "s36dukovu.hn-bkt.clouddn.com"
	bucket = "sparksovo"

	localFile := "/Users/lewyhua/Desktop/aaa.mp4"
	key := "aaa.mp4"

	mac := qbox.NewMac(accessKey, secretKey)
	// 根据七牛云存储区域选择对应的存储区域
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan, // 选择华南区域，您可以根据实际情况选择其他区域
		UseHTTPS:      false,               // 是否使用 HTTPS
		UseCdnDomains: false,               // 是否使用 CDN 加速域名
	}
	// 构建表单上传的对象

	formUploader := storage.NewFormUploader(&cfg)
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	upToken := putPolicy.UploadToken(mac)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)
	publicAccessURL := storage.MakePublicURL(domain, key)
	fmt.Println(publicAccessURL)
	videoURL := "https://" + domain + storage.ZoneHuanan.RsHost + "/" + ret.Key
	fmt.Println(videoURL)
	fmt.Println(videoURL + "?vframe/png/offset/1/w/480/h/360")

}

//
//func TestGetCover(t *testing.T) {
//
//	// 构建一个 Mac 实例，用于生成上传凭证
//	mac := qbox.NewMac(accessKey, secretKey)
//
//	// 初始化配置
//	cfg := storage.Config{
//		Zone:          &storage.ZoneHuanan, // 选择华南区域，您可以根据实际情况选择其他区域
//		UseHTTPS:      false,               // 是否使用 HTTPS
//		UseCdnDomains: false,               // 是否使用 CDN 加速域名
//	}
//
//	// 创建 BucketManager 对象
//	bucketManager := storage.NewBucketManager(mac, &cfg)
//
//	// 替换为您的存储空间名
//	// 替换为要截取帧的视频文件的 key
//	videoKey := "your_video_key.mp4"
//
//	// 截帧处理操作
//	fops := "vframe/png/offset/1/w/480/h/360"
//	pfop := storage.Per{
//		Pipeline:  "your_pipeline", // 替换为您的数据处理队列名
//		Force:     false,
//		SaveKey:   "frame.png", // 保存的帧图片名称
//		Fops:      fops,
//		NotifyURL: "your_notify_url", // 接收处理结果的通知地址
//		Separator: "-",
//	}
//	persistentID, err := bucketManager.Pfop(bucket, videoKey, []string{fops}, []string{}, pfop)
//	if err != nil {
//		log.Fatal("Failed to request video frame extraction:", err)
//	}
//	fmt.Println("Persistent ID:", persistentID)
//}
