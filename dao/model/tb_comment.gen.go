// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

//
//import (
//	"time"
//)

//const TableNameComment = "tb_comment"
//
//// Comment mapped from table <tb_comment>
//type Comment struct {
//	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                        // 评论id
//	UserID     int64     `gorm:"column:user_id" json:"user_id"`                                            // 用户id
//	VideoID    int64     `gorm:"column:video_id" json:"video_id"`                                          // 视频id
//	Content    string    `gorm:"column:content" json:"content"`                                            // 评论
//	IsDeleted  int32     `gorm:"column:is_deleted" json:"is_deleted"`                                      // 删除评论
//	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间
//}
//
//// TableName Comment's table name
//func (*Comment) TableName() string {
//	return TableNameComment
//}
