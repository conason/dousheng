package serviceImpl

import (
	"dousheng/dao"
	"dousheng/dao/model"
	"errors"
	"time"
)

func Register(username, password string) (int64, error) {
	userId, err := dao.GetUserIdByName(username)
	if err != nil {
		return 0, err
	}
	//用户已存在
	if userId != 0 {
		return -1, nil
	}

	user := model.User{
		Name:       username,
		Password:   password,
		CreateTime: time.Now(),
	}
	err = dao.SaveUser(user)
	if err != nil {
		return 0, err
	}

	userId, err = dao.GetUserIdByName(username)
	return userId, nil
}

func Login(username, password string) (int64, bool, error) {
	userId, err := dao.GetUserIdByName(username)
	if err != nil || userId == 0 {
		return 0, false, err
	}

	pwd, err := dao.GetPWDByName(username)
	if err != nil {
		return 0, false, err
	}

	if pwd != password {
		return 0, false, err
	}

	return userId, true, nil
}

func FindUser(username string) (int64, error) {
	userId, err := dao.GetUserIdByName(username)
	if err != nil {
		return 0, err
	}
	if userId == 0 {
		return 0, nil
	}
	return userId, nil
}

func FindUserWithId(userId int64) (int64, error) {
	count, err := dao.CountUserId(userId)
	if err != nil {
		return 0, err
	}
	if count == 0 {
		return 0, nil
	}
	return count, nil
}

func GetUserData(userId int64) (model.User, error) {
	user, err := dao.GetUserData(userId)
	if err != nil {
		return model.User{}, err
	}
	if user == (model.User{}) {
		return model.User{}, errors.New("no such user")
	}
	return user, nil
}

func GetUserById(userId int64) (model.User, error) {
	user, err := dao.GetUserById(userId)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func GetUserListByIds(ids []int64) ([]model.User, error) {
	users, err := dao.GetUserListByIds(ids)
	if err != nil {
		return nil, err
	}
	return users, nil
}
