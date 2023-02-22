package serviceImpl

import (
	"dousheng/dao"
	"dousheng/dao/model"
	"dousheng/utils"
)

func SendMsg(message model.Message) error {
	err := dao.MessageSave(message)
	if err != nil {
		return err
	}
	return nil
}

func ReceiveMsg(userId, toUserId, preTimeInt int64) ([]model.Message, error) {
	if preTimeInt == 0 {
		messages, err := dao.GetAllMsgToUser(userId, toUserId)
		if err != nil {
			return nil, err
		}
		return messages, nil
	}

	date, err := utils.TimestampToDate(preTimeInt)
	if err != nil {
		return nil, err
	}

	messages, err := dao.GetMessageToUser(userId, toUserId, date)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
