package service

import "dousheng/dao/model"

type UserService interface {
	Register(username, password string) (int64, error)

	Login(username, password string) (int64, bool, error)

	FindUser(username string) (int64, error)

	FindUserWithId(userId int64) (int64, error)

	GetUserData(userId int64) (model.User, error)

	GetUserById(userId int64) (model.User, error)

	GetUserListByIds(ids []int64) ([]model.User, error)
}
