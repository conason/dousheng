package serviceImpl

import (
	"dousheng/dao"
	"dousheng/dao/model"
	"time"
)

func SendMsg(message model.Message) error {
	err := dao.MessageSave(message)
	if err != nil {
		return err
	}
	return nil
}

func ReceiveMsg(userId int64, preTime time.Time) ([]model.Message, error) {
	messages, err := dao.GetMessageToUser(userId, preTime)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
