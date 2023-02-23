package dao

import (
	"dousheng/dao/dal"
	model "dousheng/dao/model"
)

func Fav(favorite model.Favorite) error {
	var fav model.Favorite
	err := dal.Favorite.Where(dal.Favorite.VideoID.Eq(favorite.VideoID), dal.Favorite.UserID.Eq(favorite.UserID)).Scan(&fav)
	if err != nil {
		return err
	}

	if fav != (model.Favorite{}) {
		_, err := dal.Favorite.Where(dal.Favorite.VideoID.Eq(favorite.VideoID), dal.Favorite.UserID.Eq(favorite.UserID)).Update(dal.Favorite.IsDeleted, favorite.IsDeleted)
		if err != nil {
			return err
		}
	}

	err = dal.Favorite.Where(dal.Favorite.VideoID.Eq(favorite.VideoID), dal.Favorite.UserID.Eq(favorite.UserID)).Save(&favorite)
	if err != nil {
		return err
	}
	return err
}

func FavList(userId int64) ([]int64, error) {
	var userIds []int64
	err := dal.Favorite.Select(dal.Favorite.VideoID).
		Where(dal.Favorite.UserID.Eq(userId), dal.Favorite.IsDeleted.Eq(0)).
		Scan(&userIds)
	if err != nil {
		return nil, err
	}

	return userIds, err
}

func IsFav(userId, videoId int64) (bool, error) {
	count, err := dal.Favorite.Where(dal.Favorite.UserID.Eq(userId), dal.Favorite.VideoID.Eq(videoId), dal.Favorite.IsDeleted.Eq(0)).Count()
	if err != nil {
		return false, err
	}

	if count <= 0 {
		return false, nil
	}

	return true, nil
}
