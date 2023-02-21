package serviceImpl

import (
	"dousheng/dao"
	"dousheng/dao/model"
)

func Comment(comment model.Comment) error {
	err := dao.SaveComment(comment)
	if err != nil {
		return err
	}
	return err
}

func GetVideoCommentList(videoId int64) ([]model.Comment, error) {
	comments, err := dao.GetCommentByVideoId(videoId)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func GetUserCommentList(userId int64) ([]model.Comment, error) {
	comments, err := dao.GetCommentByUserId(userId)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
