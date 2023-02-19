package model

import (
	"dousheng/db"
	"dousheng/utils"
	"fmt"
	"time"
)

type UserRegister struct {
	Response
	UserId int32  `json:"user_id"`
	Token  string `json:"token,omitempty"`
}
// 用于使用gorm从数据库中读取，字段需与数据库字段保持一致
type user struct {
	Id       int32
	Username string
	FollowCount int32
	FollowerCount int32
	IsFollow bool
	Password string
	CreateTime time.Time
}

// Register 用户注册
func Register(username, password string) int32 {
	if FindUser(username) == 0 {
		db.Db.Model(&user{}).Create(map[string]interface{}{"username": username, "password": password})
		fmt.Println("用户创建成功")
		return utils.SUCCESS
	}
	fmt.Println("用户创建失败")
	return utils.FAIL
}

// FindUser 通过username查找用户并返回其id
func FindUser(username string) int32 {
	user := user{}
	db.Db.Where("username=?", username).First(&user)
	return user.Id
}

// FindUserWithId 通过用户id判断用户是否存在
func FindUserWithId(id int32) bool {
	user := user{}
	db.Db.Where("id = ?",id).Find(&user).Limit(1)
	if user.Username==""{
		return false
	}
	return true
}
//GetUserData 通过传入的id从数据库获取用户信息
func GetUserData(id int32) User {
	user := User{}
	db.Db.Where("id=?", id).First(&user)
	return user
}
//Login 用户登录
func Login(username, password string) bool {
	userData := user{}
	var res bool
	if FindUser(username) == 0 {
		res = false
	} else {
		db.Db.Where("username=?", username).First(&userData)
		if userData.Password == password {
			res = true
		} else {
			res = false
		}
	}
	return res
}
