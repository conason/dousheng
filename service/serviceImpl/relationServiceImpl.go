package serviceImpl

import (
	"dousheng/dao"
	"dousheng/dao/model"
	"time"
)

func SubAction(userId, upId, actionType int64) error {
	relation := model.Relation{
		FollowerID:  userId,
		FollowingID: upId,
		IsDeleted:   0,
		CreateTime:  time.Now(),
	}
	var count = int64(1)
	if actionType == 2 {
		count = -1
		relation.IsDeleted = 1
		//err := dao.RelationSave(relation)
		//if err != nil {
		//	return err
		//}
		//err = dao.AddSubCount(userId, -1)
		//if err != nil {
		//	return err
		//}
		//err = dao.AddFansCount(upId, -1)
		//if err != nil {
		//	return err
		//}
		//return nil
	}

	err := dao.RelationSave(relation)
	if err != nil {
		return err
	}
	err = dao.AddSubCount(userId, count)
	if err != nil {
		return err
	}
	err = dao.AddFansCount(upId, count)
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
