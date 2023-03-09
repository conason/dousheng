package model

import "time"

const TableNameVideo = "tb_video"

// Video mapped from table <tb_video>
type Video struct {
	ID            int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                        // 视频id
	UserID        int64     `gorm:"column:user_id" json:"user_id"`                                            // 用户作者id
	User          User      `json:"author"`                                                                   // 视频作者
	PlayURL       string    `gorm:"column:play_url" json:"play_url"`                                          // 视频URL
	CoverURL      string    `gorm:"column:cover_url" json:"cover_url"`                                        // 封面URL
	FavoriteCount int32     `gorm:"column:favorite_count" json:"favorite_count"`                              // 点赞总数
	CommentCount  int32     `gorm:"column:comment_count" json:"comment_count"`                                // 评论总数
	Title         string    `gorm:"column:title" json:"title"`                                                // 视频标题
	CreateDate    time.Time `gorm:"column:create_date;not null;default:CURRENT_TIMESTAMP" json:"create_date"` // 创建时间
	UpdateDate    time.Time `gorm:"column:update_date;not null;default:CURRENT_TIMESTAMP" json:"update_date"` // 更新时间
	IsFavorite    bool      `json:"is_favorite"`                                                              // 是否点赞
}

//type Video struct {
//	ID            int64  `json:"id"`             // 视频id
//	User          User   `json:"author"`         // 视频作者
//	PlayURL       string `json:"play_url"`       // 视频URL
//	CoverURL      string `json:"cover_url"`      // 封面URL
//	FavoriteCount int32  `json:"favorite_count"` // 点赞总数
//	CommentCount  int32  `json:"comment_count"`  // 评论总数
//	IsFavorite    bool   `json:"is_favorite"`    // 是否点赞
//	Title         string `json:"title"`          // 视频标题
//}

// TableName Video's table name
func (*Video) TableName() string {
	return TableNameVideo
}
