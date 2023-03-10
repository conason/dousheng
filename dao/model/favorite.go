package model

import "time"

const TableNameFavorite = "tb_favorite"

// Favorite mapped from table <tb_favorite>
type Favorite struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                        // 点赞id
	Username   string    `gorm:"column:username" json:"username"`                                          // 用户名
	UserID     int64     `gorm:"column:user_id" json:"user_id"`                                            // 用户id
	VideoID    int64     `gorm:"column:video_id" json:"video_id"`                                          // 视频id
	IsDeleted  int32     `gorm:"column:is_deleted" json:"is_deleted"`                                      // 取消点赞
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间
}

// TableName Favorite's table name
func (*Favorite) TableName() string {
	return TableNameFavorite
}
