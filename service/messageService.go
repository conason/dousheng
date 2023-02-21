package service

import (
	"dousheng/dao/model"
	"time"
)

type MsgService interface {
	SendMsg(message model.Message) error

	ReceiveMsg(userId int64, pertime time.Time) ([]model.Message, error)
}
