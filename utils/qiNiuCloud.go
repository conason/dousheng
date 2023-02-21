package utils

import (
	"bytes"
	"context"
	"dousheng/config"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"time"
)

func PushVideo(key string, data []byte) int32 {
	putPolicy := storage.PutPolicy{
		Scope: config.VideoBucket,
	}
	mac := qbox.NewMac(config.AccessKey, config.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.ZoneHuadongZheJiang2
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	dataLen := int64(len(data))
	err := formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret.Key, ret.Hash)
	if ret.Hash != "" {
		return SUCCESS
	}
	return FAIL
}

func GetVideo(key string) string {
	domain := config.Domain
	mac := qbox.NewMac(config.AccessKey, config.SecretKey)
	deadline := time.Now().Add(time.Second * 3600 * 24 * 365).Unix() //1年有效期
	privateAccessURL := storage.MakePrivateURL(mac, domain, key, deadline)
	return privateAccessURL
}

func DeleteVideo() {
	bucket := config.VideoBucket
	key := "github-x.jpg"
	accessKey := config.AccessKey
	secretKey := config.SecretKey
	mac := qbox.NewMac(accessKey, secretKey)
	cfg := storage.Config{}
	bucketManager := storage.NewBucketManager(mac, &cfg)
	err := bucketManager.Delete(bucket, key)
	if err != nil {
		fmt.Println(err)
		return
	}
}
