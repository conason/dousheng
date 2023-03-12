package dao

import (
	"dousheng/config"
	"dousheng/dao/model"
	"gorm.io/gorm"
	"time"
)

// SaveVideo 将投稿信息存入数据库
func SaveVideo(video model.Video) error {
	err := db.Model(&model.Video{}).Create(&video).Error

	return err
	//err := dal.Video.Create(&video)
	//if err != nil {
	//	return err
	//}
	//
	//return nil
}

// GetVideoByTime 按照time降序的方式查找config.N个视频信息
func GetVideoByTime(data time.Time) ([]model.Video, error) {
	var videos []model.Video

	err := db.Model(&model.Video{}).Where("create_date > ?", data).Order("create_date desc").Limit(config.N).Scan(&videos).Error

	//err := dal.Video.Where(dal.Video.CreateDate.Gt(data)).
	//	Order(dal.Video.CreateDate.Desc()).
	//	Limit(config.N).
	//	Scan(&videos)
	//if err != nil {
	//	return nil, err
	//}
	return videos, err
}

func GetVideoById(videoId int64) (model.Video, error) {
	var video model.Video
	err := db.Model(&model.Video{}).Where("video_id = ?", videoId).Find(&video).Error

	return video, err

	//video, err := dal.Video.Where(dal.Video.ID.Eq(videoId)).First()
	//if err != nil {
	//	return model.Video{}, err
	//}
	//return *video, nil
}

func GetVideosByUserId(userId int64) ([]model.Video, error) {
	var videos []model.Video
	err := db.Model(&model.Video{}).Where("user_id = ?", userId).Order("create_date desc").Scan(&videos).Error
	return videos, err

	//var video []model.Video
	//err := dal.Video.Where(dal.Video.UserID.Eq(userId)).
	//	Order(dal.Video.UpdateDate).
	//	Scan(&video)
	//if err != nil {
	//	return nil, err
	//}
	//return video, nil
}

func GetNewestVideos() ([]model.Video, error) {
	var videos []model.Video
	err := db.Model(&model.Video{}).Order("create_date desc").Limit(config.N).Scan(&videos).Error

	return videos, err
	//err := dal.Video.Order(dal.Video.UpdateDate.Desc()).
	//	Limit(config.N).
	//	Scan(&videos)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return videos, nil
}

func AddVideoFavCount(videoId, num int64) error {
	err := db.Model(&model.Video{}).Where("video_id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
	return err
	//_, err := dal.Video.Where(dal.Video.ID.Eq(videoId)).UpdateSimple(dal.Video.FavoriteCount.Add(int32(num)))
	//if err != nil {
	//	return err
	//}
	//return nil
}

func AddCommentCount(videoId, num int64) error {
	err := db.Model(&model.Video{}).Where("video_id = ?", videoId).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error
	return err

	//_, err := dal.Video.Where(dal.Video.ID.Eq(videoId)).UpdateSimple(dal.Video.CommentCount.Add(int32(num)))
	//if err != nil {
	//	return err
	//}
	//return nil
}

// GetAllVideos func may repeat -> GetNewestVideos
func GetAllVideos() ([]model.Video, error) {
	var videos []model.Video
	err := db.Model(&model.Video{}).Order("create_date").Limit(config.N).Scan(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}

// GetVideoByUserId func may repeat -> GetVideosByUserId
func GetVideoByUserId(userId int64) (model.Video, error) {
	var video model.Video
	err := db.Model(&model.Video{}).Where("user_id = ?", userId).Order("create_date desc").Limit(config.N).Scan(&video).Error
	return video, err
	//err := dal.Video.Where(dal.Video.UserID.Eq(userId)).
	//	Order(dal.Video.CreateDate.Desc()).
	//	Limit(1).
	//	Scan(&video)
	//if err != nil {
	//	return model.Video{}, err
	//}
	//return video, nil
}
