package dao

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"time"
	"tk/config"
	"tk/dao/dal"
	"tk/dao/model"
	"tk/utils"
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

// PushVideoToMysql 将投稿信息存入数据库
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

	err := dal.Video.Create(&video)
	if err != nil {
		return err
	}

	return nil
	//pushVideoToMysql := video{
	//	Id:       videoId,
	//	PlayUrl:  playUrl,
	//	CoverUrl: coverUrl,
	//	Title:    title,
	//}
	//db.Db.Create(pushVideoToMysql)
}

// GetVideo 按照time降序的方式查找config.N个视频信息
func GetVideo(timeStr string) ([]model.Video, int64) {
	var videos []model.Video
	timeNext, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		return nil, 0
	}
	err = dal.Video.Where(dal.Video.CreateDate.Lt(timeNext)).
		Order(dal.Video.CreateDate.Desc()).
		Limit(config.N).
		Scan(&videos)
	if err != nil {
		return nil, 0
	}
	//db.Db.Where("create_date < ?", time).Limit(config.N).Find(&video).Count(&count)
	//if count >= 5 {
	//	count = 5
	//}
	count := cap(videos)
	return videos, int64(count)
}

func GetVideoById(videoId int64) (model.Video, error) {
	video, err := dal.Video.Where(dal.Video.ID.Eq(videoId)).First()
	if err != nil {
		return model.Video{}, err
	}
	return *video, nil
}

func GetVideosByUserId(userId int64) ([]model.Video, error) {
	var video []model.Video
	err := dal.Video.Where(dal.Video.UserID.Eq(userId)).Scan(&video)
	//log.Panicln(err)
	if err != nil {
		return nil, err
	}
	return video, nil
}
