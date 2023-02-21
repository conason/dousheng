package dao

import (
	"dousheng/dao/dal"
	"dousheng/dao/model"
)

func RelationSave(relation model.Relation) error {
	err := dal.Relation.
		Where(dal.Relation.FollowerID.Eq(relation.FollowerID), dal.Relation.FollowingID.Eq(relation.FollowingID)).
		Save(&relation)
	if err != nil {
		return err
	}
	return nil
}

func GetSubListByUserId(userId int64) ([]int64, error) {
	var subList []int64
	err := dal.Relation.Select(dal.Relation.FollowingID).
		Where(dal.Relation.FollowerID.Eq(userId), dal.Relation.Isdeleted.Eq(0)).
		Order(dal.Relation.CreateTime.Desc()).
		Scan(&subList)
	if err != nil {
		return nil, err
	}

	return subList, nil
}

func GetFansListByUserId(userId int64) ([]int64, error) {
	var fansList []int64
	err := dal.Relation.Select(dal.Relation.FollowerID).
		Where(dal.Relation.FollowingID.Eq(userId), dal.Relation.Isdeleted.Eq(0)).
		Order(dal.Relation.CreateTime.Desc()).
		Scan(&fansList)
	if err != nil {
		return nil, err
	}

	return fansList, nil
}
