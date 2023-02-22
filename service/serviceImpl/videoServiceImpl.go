package serviceImpl

import (
	"dousheng/dao"
	"dousheng/dao/model"
	"dousheng/utils"
	"io"
	"io/ioutil"
	"mime/multipart"
	"time"
)

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

	err := dao.SaveVideo(video)
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

func GetNewestVideos() ([]model.Video, int64, error) {
	videos, err := dao.GetNewestVideos()
	if err != nil {
		return nil, 0, err
	}
	return videos, int64(len(videos)), nil
}
