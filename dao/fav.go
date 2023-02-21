package dao

import (
	"dousheng/dao/dal"
	"dousheng/dao/model"
)

func Fav(favorite model.Favorite) error {
	err := dal.Favorite.Where(dal.Favorite.VideoID.Eq(favorite.VideoID), dal.Favorite.UserID.Eq(favorite.UserID)).Save(&favorite)
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
