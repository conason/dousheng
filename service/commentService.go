package service

import "tk/dao/model"

type CommentService interface {
	Comment(comment model.Comment) error

	GetVideoCommentList(videoId int64) ([]model.Comment, error)

	GetUserCommentList(userId int64) ([]model.Comment, error)
}
