package serviceImpl

import (
	"dousheng/dao"
	"dousheng/dao/model"
	"time"
)

func FavAction(userId int64, videoId int64, actionType int32) error {
	favorite := model.Favorite{
		UserID:     userId,
		VideoID:    videoId,
		IsDeleted:  0,
		CreateTime: time.Now(),
	}
	var count = int64(1)
	if actionType == 2 {
		count = -1
		favorite.IsDeleted = 1
		//favorite = model.Favorite{
		//	UserID:     userId,
		//	VideoID:    videoId,
		//	IsDeleted:  1,
		//	CreateTime: time.Now(),
		//}
	}

	//事务
	//更新、插入fav表，以及更新video表
	//fav表操作
	err := dao.Fav(favorite)
	if err != nil {
		return err
	}
	err = dao.AddFavCount(userId, count)
	if err != nil {
		return err
	}
	err = dao.AddVideoFavCount(videoId, count)
	if err != nil {
		return err
	}
	//video表
	//更新完成

	return err
}

func GetFavListByUserId(userId int64) ([]int64, error) {
	favList, err := dao.FavList(userId)
	if err != nil {
		return nil, err
	}
	return favList, nil

}
