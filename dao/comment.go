package dao

import (
	"dousheng/dao/dal"
	"dousheng/dao/model"
)

func SaveComment(comment model.Comment) error {
	err := dal.Comment.Save(&comment)
	if err != nil {
		return err
	}
	return err
}

func GetCommentByUserId(userId int64) ([]model.Comment, error) {
	var comments []model.Comment
	err := dal.Comment.Where(dal.Comment.UserID.Eq(userId), dal.Comment.IsDeleted.Eq(0)).
		Order(dal.Comment.CreateTime.Desc()).
		Scan(&comments)
	if err != nil {
		return nil, err
	}
	//var comments []model.Comment
	//for i, comment := range find {
	//	comments[i] = *comment
	//}

	return comments, err
}

func GetCommentByVideoId(videoId int64) ([]model.Comment, error) {
	var comments []model.Comment
	err := dal.Comment.Where(dal.Comment.VideoID.Eq(videoId), dal.Comment.IsDeleted.Eq(0)).
		Order(dal.Comment.CreateTime.Desc()).
		Scan(&comments)
	if err != nil {
		return nil, err
	}
	//var comments []model.Comment
	//for i, comment := range find {
	//	comments[i] = *comment
	//}
	return comments, err
}
