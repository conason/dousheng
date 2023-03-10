package dao

import (
	"dousheng/dao/model"
)

func RelationSave(relation model.Relation) error {
	var count int64
	err := db.Model(&relation).
		Where("follower_id = ? AND following_id = ?", relation.FollowerID, relation.FollowingID).
		Count(&count).
		Error
	if err != nil {
		return err
	}
	if count > 0 {
		err = db.Model(&relation).UpdateColumn("is_deleted", relation.IsDeleted).Error
		return err
	}
	err = db.Model(&model.Relation{}).Create(&relation).Error
	return err
	//var exits model.Relation
	//err := dal.Relation.
	//	Where(dal.Relation.FollowerID.Eq(relation.FollowerID), dal.Relation.FollowingID.Eq(relation.FollowingID)).
	//	Scan(&exits)
	//if err != nil {
	//	return err
	//}
	//
	//if exits != (model.Relation{}) {
	//	_, err := dal.Relation.
	//		Where(dal.Relation.FollowerID.Eq(relation.FollowerID), dal.Relation.FollowingID.Eq(relation.FollowingID)).
	//		Update(dal.Relation.Isdeleted, relation.Isdeleted)
	//	if err != nil {
	//		return err
	//	}
	//}
	//
	//err = dal.Relation.
	//	Where(dal.Relation.FollowerID.Eq(relation.FollowerID), dal.Relation.FollowingID.Eq(relation.FollowingID)).
	//	Save(&relation)
	//if err != nil {
	//	return err
	//}
	//
	//return nil
}

func GetSubListByUserId(userId int64) ([]int64, error) {
	var subList []int64
	err := db.Model(&model.Relation{}).
		Select("following_id").
		Where("follower_id = ? AND is_deleted = ?", userId, 0).
		Order("create_time desc").
		Scan(&subList).Error
	return subList, err
	//err := dal.Relation.Select(dal.Relation.FollowingID).
	//	Where(dal.Relation.FollowerID.Eq(userId), dal.Relation.Isdeleted.Eq(0)).
	//	Order(dal.Relation.CreateTime.Desc()).
	//	Scan(&subList)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return subList, nil
}

func GetFansListByUserId(userId int64) ([]int64, error) {
	var fansList []int64
	err := db.Model(&model.Relation{}).
		Select("follower_id").
		Where("following_id = ? AND is_deleted = ?", userId, 0).
		Order("create_time desc").
		Scan(&fansList).Error
	return fansList, err
	//err := dal.Relation.Select(dal.Relation.FollowerID).
	//	Where(dal.Relation.FollowingID.Eq(userId), dal.Relation.Isdeleted.Eq(0)).
	//	Order(dal.Relation.CreateTime.Desc()).
	//	Scan(&fansList)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return fansList, nil
}

func IsSub(userId, videoUserId int64) (bool, error) {
	var count int64
	err := db.Model(&model.Relation{}).Where("follower_id = ? AND following_id = ? AND is_deleted = ?", userId, videoUserId, 0).Count(&count).Error

	//count, err := dal.Relation.Where(dal.Relation.FollowerID.Eq(userId), dal.Relation.FollowingID.Eq(videoUserId), dal.Relation.Isdeleted.Eq(0)).
	//	Count()
	if err != nil {
		return false, err
	}
	if count <= 0 {
		return false, nil
	}
	return true, nil
}
