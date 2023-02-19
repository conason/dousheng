package service

import "tk/dao/model"

type FavService interface {
	FavAction(favorite model.Favorite) error

	GetFavListByUserId(userId int64) ([]model.Video, error)
}
