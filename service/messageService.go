package service

import (
	"time"
	"tk/dao/model"
)

type MsgService interface {
	SendMsg(message model.Message) error

	ReceiveMsg(userId int64, pertime time.Time) ([]model.Message, error)
}
