package dao

import (
	"dousheng/dao/dal"
	"dousheng/dao/model"
	"errors"
	"time"
)

// Register 用户注册
func Register(username, password string) (int64, error) {
	user := model.User{
		Name:       username,
		Password:   password,
		CreateTime: time.Now(),
	}

	err := dal.User.Create(&user)
	if err != nil {
		return -1, err
	}

	return 0, nil

	//findUser, err := FindUser(username)
	//if err != nil {
	//	return -1, err
	//}

	//if findUser == 0 {

	//db.Db.Model(&user{}).Create(map[string]interface{}{"username": username, "password": password})
	//fmt.Println("用户创建成功")
	//return utils.SUCCESS
	//}

	//return -1, errors.New("creat user fail")
	//fmt.Println("用户创建失败")
	//return utils.FAIL
}

// FindUser 通过username查找用户并返回其id
func FindUser(username string) (int64, error) {
	var user int64
	err := dal.User.
		Select(dal.User.ID).
		Where(dal.User.Name.Eq(username)).Scan(&user)
	if err != nil {
		return 0, err
	}
	if err != nil {
		return -1, err
	}
	if user == 0 {
		return 0, nil
	}

	return user, nil
}

// FindUserWithId 通过用户id判断用户是否存在
func FindUserWithId(userId int64) (bool, error) {
	count, err := dal.User.
		Select(dal.User.ID).
		Where(dal.User.ID.Eq(userId)).
		Count()
	if err != nil {
		return false, err
	}

	//user := user{}
	//db.Db.Where("id = ?",id).Find(&user).Limit(1)
	if count == 0 {
		return false, nil
	}

	return true, nil
}

// GetUserData 通过传入的id从数据库获取用户信息
func GetUserData(userId int64) (model.User, error) {
	var user model.User
	err := dal.User.Where(dal.User.ID.Eq(userId)).Scan(&user)
	if err != nil {
		return model.User{}, err
	}

	if user == (model.User{}) {
		return model.User{}, errors.New("no such user")
	}

	return user, nil

	//user := User{}
	//db.Db.Where("id=?", id).First(&user)
	//return user
}

// Login 用户登录
func Login(username, password string) (int64, bool, error) {

	user, err := dal.User.Where(dal.User.Name.Eq(username)).First()
	if err != nil {
		return 0, false, err
	}

	if user.Password != password {
		return -1, false, nil //errors.New("incorrect account password")
	}

	return user.ID, true, nil
	//else {

	//db.Db.Where("username=?", username).First(&userData)
	//if userData.Password == password {
	//	res = true
	//} else {
	//	res = false
	//}
	//}
	//return res
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
