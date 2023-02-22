package dao

import (
	"dousheng/dao/dal"
	"dousheng/dao/model"
	"time"
)

func MessageSave(message model.Message) error {
	err := dal.Message.Save(&message)
	if err != nil {
		return err
	}
	return nil
}

func GetMessageToUser(userId, toUserId int64, preTime time.Time) ([]model.Message, error) {
	var msg []model.Message
	err := dal.Message.Where(dal.Message.ToUserID.Eq(userId), dal.Message.FromUserID.Eq(toUserId), dal.Message.CreateTime.Gt(preTime)).
		Order(dal.Message.CreateTime).
		Scan(&msg)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func GetAllMsgToUser(userId, toUserId int64) ([]model.Message, error) {
	var msg []model.Message
	err := dal.Message.Where(dal.Message.ToUserID.Eq(userId), dal.Message.FromUserID.Eq(toUserId)).
		Or(dal.Message.ToUserID.Eq(toUserId), dal.Message.FromUserID.Eq(userId)).
		Order(dal.Message.CreateTime).
		Scan(&msg)
	if err != nil {
		return nil, err
	}
	return msg, nil
}
