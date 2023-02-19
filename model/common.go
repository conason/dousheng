package model

type Response struct {
	StatusCode int32  `json:"status_code"`          //状态码
	StatusMsg  string `json:"status_msg,omitempty"` //状态信息描述（空时不返回）
}

type Video struct {
	Id            int32  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int32  `json:"favorite_count,omitempty"`
	CommentCount  int32  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	Title         string `json:"title,omitempty"`
}

type Comment struct {
	Id         int32  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type User struct {
	Id            int32  `json:"id,omitempty" gorm:"id"`                        //用户id
	Username      string `json:"name,omitempty" gorm:"name"`                //用户名称
	FollowCount   int32  `json:"follow_count,omitempty" gorm:"followcount"`     //关注总数
	FollowerCount int32  `json:"follower_count,omitempty" gorm:"followercount"` //粉丝总数
	IsFollow      bool   `json:"is_follow,omitempty" gorm:"isfollow"`           //true--已关注，false--未关注
}

type Message struct {
	Id         int32  `json:"id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}

type MessageSendEvent struct {
	UserId     int32  `json:"user_id,omitempty"`
	ToUserId   int32  `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId int32  `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}
