package model

import "time"

const TableNameUser = "tb_user"

// User mapped from table <tb_user>
type User struct {
	ID              int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                        // 用户id
	Name            string    `gorm:"column:name" json:"name"`                                                  // 用户名
	FollowCount     int32     `gorm:"column:follow_count" json:"follow_count"`                                  // 关注总数
	FollowerCount   int32     `gorm:"column:follower_count" json:"follower_count"`                              // 粉丝总数
	BackgroundImage string    `gorm:"column:background_image" json:"background_image"`                          // 用户个人页顶部大图URL
	Signature       string    `gorm:"column:signature" json:"signature"`                                        // 个人简介
	TotalFavorited  int32     `gorm:"column:total_favorited" json:"total_favorited"`                            // 获赞数量
	WorkCount       int32     `gorm:"column:work_count" json:"work_count"`                                      // 作品数量
	FavoriteCount   int32     `gorm:"column:favorite_count" json:"favorite_count"`                              // 点赞数量
	Password        string    `gorm:"column:password" json:"password"`                                          // 密码
	CreateTime      time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间
	IsFowllow       bool      `json:"is_follow"`                                                                // 是否关注
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
