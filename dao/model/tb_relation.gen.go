// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameRelation = "tb_relation"

// Relation mapped from table <tb_relation>
type Relation struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                        // 关注id
	FollowerID  int64     `gorm:"column:follower_id" json:"follower_id"`                                    // 粉丝id
	FollowingID int64     `gorm:"column:following_id" json:"following_id"`                                  // 博主id
	Isdeleted   int32     `gorm:"column:isdeleted" json:"isdeleted"`                                        // 取消关注
	CreateTime  time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间
}

// TableName Relation's table name
func (*Relation) TableName() string {
	return TableNameRelation
}
