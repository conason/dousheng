package model

import (
	"dousheng/config"
	"dousheng/db"
	"dousheng/utils"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
)

type VideoLists struct {
	Response
	NextTime  int32   `json:"next_time"`
	VideoList []Video `json:"video_list"`
}


type video struct {
	VideoId       int32
	Id            int32
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int32
	CommentCount  int32
	Title         string
	IsFavorite bool
}

// ParseVideo 将*multipart.FileHeader类型转化为 []byte
func ParseVideo(videoData *multipart.FileHeader) []byte {
	file, err := videoData.Open()
	utils.ResolveError(err)
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil && err == io.EOF {
		fmt.Println(err)
	}
	return data
}

// PushVideoToMysql 将投稿信息存入数据库
func PushVideoToMysql(id int32, playUrl, coverUrl, title string) {
	pushVideoToMysql := video{
		Id:       id,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
		Title:    title,
	}
	db.Db.Create(pushVideoToMysql)
}

//GetVideo 按照time降序的方式查找config.N个视频信息
func GetVideo(time string) ([]video,int32){
	var count int64
	var video []video
	db.Db.Where("create_date < ?", time).Limit(config.N).Find(&video).Count(&count)
	if count >=5{
		count=5
	}
	return video,int32(count)
}
