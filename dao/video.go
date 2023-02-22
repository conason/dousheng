package dao

import (
	"dousheng/config"
	"dousheng/dao/dal"
	"dousheng/dao/model"
	"time"
)

// SaveVideo 将投稿信息存入数据库
func SaveVideo(video model.Video) error {
	err := dal.Video.Create(&video)
	if err != nil {
		return err
	}

	return nil
}

// GetVideoByTime 按照time降序的方式查找config.N个视频信息
func GetVideoByTime(data time.Time) ([]model.Video, error) {
	var videos []model.Video
	err := dal.Video.Where(dal.Video.CreateDate.Gt(data)).
		Order(dal.Video.CreateDate.Desc()).
		Limit(config.N).
		Scan(&videos)
	if err != nil {
		return nil, err
	}
	return videos, nil
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
	err := dal.Video.Where(dal.Video.UserID.Eq(userId)).
		Order(dal.Video.UpdateDate).
		Scan(&video)
	if err != nil {
		return nil, err
	}
	return video, nil
}

func GetNewestVideos() ([]model.Video, error) {
	var videos []model.Video
	err := dal.Video.Order(dal.Video.UpdateDate.Desc()).
		Limit(config.N).
		Scan(&videos)
	if err != nil {
		return nil, err
	}

	return videos, nil
}

func AddVideoFavCount(videoId, num int64) error {
	_, err := dal.Video.Where(dal.Video.ID.Eq(videoId)).UpdateSimple(dal.Video.FavoriteCount.Add(int32(num)))
	if err != nil {
		return err
	}
	return nil
}

func AddCommentCount(videoId, num int64) error {
	_, err := dal.Video.Where(dal.Video.ID.Eq(videoId)).UpdateSimple(dal.Video.CommentCount.Add(int32(num)))
	if err != nil {
		return err
	}
	return nil
}
