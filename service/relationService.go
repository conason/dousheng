package service

import "tk/dao/model"

type RelationService interface {
	SubAction(userId, upId int64) error

	SubList(userId int64) ([]model.User, error)

	FansList(userId int64) ([]model.User, error)
}
