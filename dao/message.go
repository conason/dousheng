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

func GetMessageToUser(userId int64, preTime time.Time) ([]model.Message, error) {
	var msg []model.Message
	err := dal.Message.Where(dal.Message.ToUserID.Eq(userId), dal.Message.CreateTime.Gt(preTime)).
		Order(dal.Message.CreateTime.Desc()).
		Scan(&msg)
	if err != nil {
		return nil, err
	}

	return msg, nil
}
