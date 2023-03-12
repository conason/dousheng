package model

import "time"

const TableNameComment = "tb_comment"

// Comment mapped from table <tb_comment>
type Comment struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                                    // 评论id
	UserID     int64     `gorm:"column:user_id" json:"user_id"`                                                        // 用户id
	VideoID    int64     `gorm:"column:video_id" json:"video_id"`                                                      // 视频id
	Content    string    `gorm:"column:content" json:"content"`                                                        // 评论
	IsDeleted  int32     `gorm:"column:is_deleted" json:"is_deleted"`                                                  // 删除评论
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"`             // 创建时间
	User       User      `gorm:"-" protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`                      // 评论用户信息
	CreateDate string    `protobuf:"bytes,4,opt,name=create_date,json=createDate,proto3" json:"create_date,omitempty"` // 评论发布日期，格式 mm-dd
}

// TableName Comment's table name
func (*Comment) TableName() string {
	return TableNameComment
}
