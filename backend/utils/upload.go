package utils

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	_ "github.com/qiniu/go-sdk/v7/storage"
	"go.uber.org/zap"
	"os"
	"sparks/config/constant"
)

var (
	accessKey = "sB5H5Jq5_RZmwXyxDeJzyR7SGtjaxERVEfNlD6wu"
	secretKey = "8xUBykY0QZPgVewWwFsJ9myXofN8ty39gzy3RfM6"
	bucket    = "sparksovo"
	domain    = "s36dukovu.hn-bkt.clouddn.com"
)

func CreateDirectoryIfNotExist() error {
	if _, err := os.Stat(constant.FileLocalPath); os.IsNotExist(err) {
		// 创建文件夹
		err = os.MkdirAll(constant.FileLocalPath, 0700)
		if err != nil {
			zap.Error(err)
			return err
		}
	}
	return nil
}

func UploadToKoDo(videoName string) error {
	accessKey = "sB5H5Jq5_RZmwXyxDeJzyR7SGtjaxERVEfNlD6wu"
	secretKey = "8xUBykY0QZPgVewWwFsJ9myXofN8ty39gzy3RfM6"
	domain = "s36dukovu.hn-bkt.clouddn.com"
	bucket = "sparksovo"

	videoPath := constant.FileLocalPath + videoName

	mac := qbox.NewMac(accessKey, secretKey)
	// 根据七牛云存储区域选择对应的存储区域
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan, // 选择华南区域，您可以根据实际情况选择其他区域
		UseHTTPS:      false,               // 是否使用 HTTPS
		UseCdnDomains: false,               // 是否使用 CDN 加速域名
	}

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
	err := formUploader.PutFile(context.Background(), &ret, upToken, videoName, videoPath, &putExtra)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(ret.Key, ret.Hash)
	publicAccessURL := storage.MakePublicURL(domain, videoName)
	zap.L().Info("UploadToKoDo", zap.String("publicAccessURL", publicAccessURL))
	return nil
}

//
//func GetVideoCover(videoName string) (string, error) {
//	// 生成图片 UUID
//	imgId := uuid.New().String()
//	// 修改文件名
//	imgName := strings.Replace(imgId, "-", "", -1) + ".jpg"
//	//调用oss 获取封面图
//	err := postSnapShot(videoName, imgName)
//	if err != nil {
//		return "", err
//	}
//	return imgName, nil
//}

//func postSnapShot(videoName string, imgName string) error {
//	c := getClient()
//	PostSnapshotOpt := &cos.PostSnapshotOptions{
//		Input: &cos.JobInput{
//			Object: videoName,
//		},
//		Time:   "1",
//		Width:  720,
//		Height: 1280,
//		Format: "jpg",
//		Output: &cos.JobOutput{
//			Region: "ap-nanjing",
//			Bucket: "tiktok-1319971229",
//			Object: imgName,
//		},
//	}
//	_, _, err := c.CI.PostSnapshot(context.Background(), PostSnapshotOpt)
//	if err != nil {
//		return err
//	}
//	return nil
//}
