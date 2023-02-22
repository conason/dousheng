package service

import (
	"dousheng/dao/model"
	"mime/multipart"
	"time"
)

type VideoService interface {
	ParseVideo(videoData *multipart.FileHeader) ([]byte, error)

	PushVideoToMysql(userId int64, playUrl, coverUrl, title string) error

	GetVideo(data time.Time) ([]model.Video, int64)

	GetVideoById(videoId int64) (model.Video, error)

	GetVideosByUserId(userId int64) ([]model.Video, error)

	GetNewestVideos() ([]model.Video, int64, error)
}
