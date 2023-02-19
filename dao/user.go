package dao

import (
	"errors"
	"time"
	"tk/dao/dal"
	"tk/dao/model"
)

// Register 用户注册
func Register(username, password string) (int64, error) {
	user := model.User{
		Name:       username,
		Password:   password,
		CreateTime: time.Now(),
	}

	findUser, err := FindUser(username)
	if err != nil {
		return -1, err
	}

	if findUser == 0 {

		err := dal.User.Create(&user)
		if err != nil {
			return -1, err
		}

		return 0, nil
		//db.Db.Model(&user{}).Create(map[string]interface{}{"username": username, "password": password})
		//fmt.Println("用户创建成功")
		//return utils.SUCCESS
	}

	return -1, errors.New("creat user fail")
	//fmt.Println("用户创建失败")
	//return utils.FAIL
}

// FindUser 通过username查找用户并返回其id
func FindUser(username string) (int64, error) {
	user, err := dal.User.Where(dal.User.Name.Eq(username)).First()
	if err != nil {
		return -1, err
	}
	if user == nil {
		return 0, errors.New("no such user")
	}

	return user.UserID, nil
}

// FindUserWithId 通过用户id判断用户是否存在
func FindUserWithId(userId int64) (bool, error) {
	count, err := dal.User.
		Select(dal.User.UserID).
		Where(dal.User.UserID.Eq(userId)).
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
	user, err := dal.User.Where(dal.User.UserID.Eq(userId)).First()
	if err != nil {
		return model.User{}, err
	}

	if user == nil {
		return model.User{}, errors.New("no such user")
	}

	return *user, nil

	//user := User{}
	//db.Db.Where("id=?", id).First(&user)
	//return user
}

// Login 用户登录
func Login(username, password string) (bool, error) {
	//userData := user{}
	//var res bool
	//密码不能为空
	if password == "" {
		return false, errors.New("invalid password")
	}

	findUser, err := FindUser(username)
	if err != nil {
		return false, err
	}
	if findUser == 0 {
		return false, errors.New("no such user")
	}

	user, err := dal.User.Select(dal.User.Password).Where(dal.User.UserID.Eq(findUser)).First()
	if err != nil {
		return false, err
	}

	if user.Password != password {
		return false, errors.New("Incorrect account password")
	}

	return true, nil
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
	user, err := dal.User.Where(dal.User.UserID.Eq(userId)).First()
	if err != nil {
		return model.User{}, err
	}
	return *user, nil
}
