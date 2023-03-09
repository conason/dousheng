package dao

import (
	"dousheng/dao/model"
	"gorm.io/gorm"
)

// SaveUser 用户注册
func SaveUser(user *model.User) error {
	err := db.Create(&user).Error
	return err
	//err := dal.User.Create(&user)
	//if err != nil {
	//	return err
	//}
	//
	//return nil
}

// GetUserIdByName 通过username查找用户并返回其id
func GetUserIdByName(username string) (int64, error) {
	var userId int64

	err := db.Select("id").Where("name=?", username).Take(&userId).Error

	return userId, err
	//err := dal.User.
	//	Select(dal.User.ID).
	//	Where(dal.User.Name.Eq(username)).Scan(&user)
	//if err != nil {
	//	return 0, err
	//}
	//
	//return user, nil
}

// CountUserId 通过用户id判断用户是否存在
func CountUserId(userId int64) (int64, error) {
	var count int64

	err := db.Select("id").Where("id=?", userId).Count(&count).Error
	return count, err
	//count, err := dal.User.
	//	Select(dal.User.ID).
	//	Where(dal.User.ID.Eq(userId)).
	//	Count()
	//if err != nil {
	//	return 0, err
	//}
	//
	//return count, nil
}

// GetUserData 通过传入的id从数据库获取用户信息
func GetUserData(userId int64) (model.User, error) {
	var user model.User

	err := db.Where("id=?", userId).Scan(&user).Error

	return user, err
	//err := dal.User.Where(dal.User.ID.Eq(userId)).Scan(&user)
	//if err != nil {
	//	return model.User{}, err
	//}
	//
	//return user, nil
}

// GetPWDByName 用户登录
func GetPWDByName(username string) (string, error) {
	var pwd string

	err := db.Select("password").Where("name=?", username).Take(&pwd).Error

	//err := dal.User.Select(dal.User.Password).
	//	Where(dal.User.Name.Eq(username)).
	//	Scan(&pwd)
	//if err != nil {
	//	return "", err
	//}

	return pwd, err
}

func GetUserById(userId int64) (model.User, error) {
	var user model.User

	err := db.Where("id=?", userId).Scan(&user).Error

	return user, err
	//err := dal.User.Where(dal.User.ID.Eq(userId)).Scan(&user)
	//if err != nil {
	//	return model.User{}, err
	//}
	//return user, nil
}

func GetUserListByIds(ids []int64) ([]model.User, error) {
	var user []model.User

	err := db.Where("id IN ?", ids).Find(&user).Error

	return user, err
	//err := dal.User.Where(dal.User.ID.In(ids...)).Scan(&user)
	//if err != nil {
	//	return nil, err
	//}
	//return user, nil
}

func AddWorkCount(userId, num int64) error {
	err := db.Model(&model.User{}).Where("id = ?", userId).Update("work_count", gorm.Expr("work_count + ?", 1)).Error

	return err
	//_, err := dal.User.Where(dal.User.ID.Eq(userId)).UpdateSimple(dal.User.WorkCount.Add(int32(num)))
	//if err != nil {
	//	return err
	//}
	//return nil
}

func AddFavCount(userId, num int64) error {
	err := db.Model(&model.User{}).Where("id = ?", userId).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error

	return err

	//_, err := dal.User.Where(dal.User.ID.Eq(userId)).UpdateSimple(dal.User.FavoriteCount.Add(int32(num)))
	//if err != nil {
	//	return err
	//}
	//return nil
}

func AddSubCount(userId, num int64) error {
	err := db.Model(&model.User{}).Where("id = ?", userId).Update("follow_count", gorm.Expr("follow_count + ?", 1)).Error

	return err

	//_, err := dal.User.Where(dal.User.ID.Eq(userId)).UpdateSimple(dal.User.FollowCount.Add(int32(num)))
	//if err != nil {
	//	return err
	//}
	//return nil
}

func AddFansCount(userId, num int64) error {
	err := db.Model(&model.User{}).Where("id = ?", userId).Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error

	return err

	//_, err := dal.User.Where(dal.User.ID.Eq(userId)).UpdateSimple(dal.User.FollowerCount.Add(int32(num)))
	//if err != nil {
	//	return err
	//}
	//return nil
}
