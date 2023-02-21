package serviceImpl

import (
	"time"
	"tk/dao"
	"tk/dao/model"
)

func SubAction(userId, upId, actionType int64) error {
	relation := model.Relation{
		FollowerID:  userId,
		FollowingID: upId,
		Isdeleted:   0,
		CreateTime:  time.Now(),
	}

	if actionType == 2 {
		relation.Isdeleted = 1
		err := dao.RelationSave(relation)
		if err != nil {
			return err
		}
		return nil
	}

	err := dao.RelationSave(relation)
	if err != nil {
		return err
	}
	return nil
}

func SubList(userId int64) ([]model.User, error) {
	upIds, err := dao.GetSubListByUserId(userId)
	if err != nil {
		return nil, err
	}

	users, err := dao.GetUserListByIds(upIds)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func FansList(userId int64) ([]model.User, error) {
	upIds, err := dao.GetFansListByUserId(userId)
	if err != nil {
		return nil, err
	}

	users, err := dao.GetUserListByIds(upIds)
	if err != nil {
		return nil, err
	}

	return users, nil
}
