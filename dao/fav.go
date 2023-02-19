package dao

import (
	"tk/dao/dal"
	"tk/dao/model"
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
	favorites, err := dal.Favorite.Select(dal.Favorite.VideoID).
		Where(dal.Favorite.UserID.Eq(userId), dal.Favorite.IsDeleted.Eq(0)).
		Find()
	if err != nil {
		return nil, err
	}

	for i, val := range favorites {
		userIds[i] = val.UserID
	}

	return userIds, err
}
