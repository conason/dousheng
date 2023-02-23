package serviceImpl

import (
	"dousheng/dao"
	"dousheng/dao/model"
	"dousheng/utils"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"strings"
	"time"
)

// 视频上传
func Upload(videoData *multipart.FileHeader, title string, userId int64) error {
	//将*multipart.FileHeader类型转化为[]byte
	parseVideo, err := ParseVideo(videoData)
	if err != nil {
		log.Panicln(err)
		return err
	}
	//视频文件名
	videoName := fmt.Sprintf("%s.mp4", title)
	//封面文件名
	coverName := strings.Replace(videoName, ".mp4", "cover.jpeg", 1)
	//fmt.Printf("%s\n", title)
	//视频上传
	code := utils.PushVideo(videoName, parseVideo)
	if code != 0 {
		return errors.New("upload failed")
	}
	//获取视频地址
	playURL := utils.GetVideo(videoName)
	//截取封面
	parseCover, err := utils.ParseCover(playURL, 1)
	if err != nil {
		log.Panicln(err)
		//return err
	}

	//封面上传
	succ := utils.PushCover(coverName, parseCover)
	if succ != 0 {
		return errors.New("upload failed")
	}
	//获取封面地址
	coverURL := utils.GetCover(coverName)
	//上传至数据库
	err = PushVideoToMysql(userId, playURL, coverURL, title)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// ParseVideo 将*multipart.FileHeader类型转化为 []byte
func ParseVideo(videoData *multipart.FileHeader) ([]byte, error) {
	file, err := videoData.Open()
	utils.ResolveError(err)
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			utils.ResolveError(err)
		}
	}(file)
	data, err := ioutil.ReadAll(file)
	if err != nil && err == io.EOF {
		return nil, err
		//fmt.Println(err)

	}
	return data, nil
}

func PushVideoToMysql(userId int64, playUrl, coverUrl, title string) error {
	now := time.Now()
	video := model.Video{
		UserID:     userId,
		PlayURL:    playUrl,
		CoverURL:   coverUrl,
		Title:      title,
		CreateDate: now,
		UpdateDate: now,
	}
	//事务
	err := dao.SaveVideo(video)
	if err != nil {
		return err
	}
	err = dao.AddWorkCount(userId, 1)
	if err != nil {
		return err
	}

	return nil
}

func GetVideo(data time.Time) ([]model.Video, error) {
	videos, err := dao.GetVideoByTime(data)
	if err != nil {
		return nil, err
	}

	return videos, nil
}

func GetVideoById(videoId int64) (model.Video, error) {
	videos, err := dao.GetVideoById(videoId)
	if err != nil {
		return model.Video{}, err
	}
	return videos, nil
}

func GetVideosByUserId(userId int64) ([]model.Video, error) {
	videos, err := dao.GetVideosByUserId(userId)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func GetNewestVideos() ([]model.Video, error) {
	videos, err := dao.GetNewestVideos()
	if err != nil {
		return nil, err
	}
	return videos, nil
}
