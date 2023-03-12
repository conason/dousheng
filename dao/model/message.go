package model

import "time"

const TableNameMessage = "tb_message"

// Message mapped from table <tb_message>
type Message struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`     // 消息id
	ToUserID   int64     `gorm:"column:to_user_id" json:"to_user_id"`                   // 接收方id
	FromUserID int64     `gorm:"column:from_user_id" json:"from_user_id"`               // 发送方id
	Content    string    `gorm:"column:content" json:"content"`                         // 消息内容
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP"` // 创建时间
	CreateData int64     `gorm:"-" json:"create_time"`                                  // 创建时间
}

// TableName Message's table name
func (*Message) TableName() string {
	return TableNameMessage
}
