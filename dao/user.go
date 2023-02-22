package dao

import (
	"dousheng/dao/dal"
	"dousheng/dao/model"
)

// SaveUser 用户注册
func SaveUser(user model.User) error {
	err := dal.User.Create(&user)
	if err != nil {
		return err
	}

	return nil
}

// GetUserIdByName 通过username查找用户并返回其id
func GetUserIdByName(username string) (int64, error) {
	var user int64
	err := dal.User.
		Select(dal.User.ID).
		Where(dal.User.Name.Eq(username)).Scan(&user)
	if err != nil {
		return 0, err
	}

	return user, nil
}

// CountUserId 通过用户id判断用户是否存在
func CountUserId(userId int64) (int64, error) {
	count, err := dal.User.
		Select(dal.User.ID).
		Where(dal.User.ID.Eq(userId)).
		Count()
	if err != nil {
		return 0, err
	}

	return count, nil
}

// GetUserData 通过传入的id从数据库获取用户信息
func GetUserData(userId int64) (model.User, error) {
	var user model.User
	err := dal.User.Where(dal.User.ID.Eq(userId)).Scan(&user)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

// GetPWDByName 用户登录
func GetPWDByName(username string) (string, error) {
	var pwd string
	err := dal.User.Select(dal.User.Password).
		Where(dal.User.Name.Eq(username)).
		Scan(&pwd)
	if err != nil {
		return "", err
	}

	return pwd, nil
}

func GetUserById(userId int64) (model.User, error) {
	var user model.User
	err := dal.User.Where(dal.User.ID.Eq(userId)).Scan(&user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func GetUserListByIds(ids []int64) ([]model.User, error) {
	var user []model.User
	err := dal.User.Where(dal.User.ID.In(ids...)).Scan(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
