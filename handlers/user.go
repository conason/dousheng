package handlers

import (
	"dousheng/dao/model"
	"dousheng/service/serviceImpl"
	"dousheng/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//type User struct {
//	ID              int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // 用户id
//	Name            string `gorm:"column:name" json:"name"`                           // 用户名
//	FollowCount     int32  `gorm:"column:follow_count" json:"follow_count"`           // 关注总数
//	FollowerCount   int32  `gorm:"column:follower_count" json:"follower_count"`       // 粉丝总数
//	BackgroundImage string `gorm:"column:background_image" json:"background_image"`   // 用户个人页顶部大图URL
//	Signature       string `gorm:"column:signature" json:"signature"`                 // 个人简介
//	TotalFavorited  int32  `gorm:"column:total_favorited" json:"total_favorited"`     // 获赞数量
//	WorkCount       int32  `gorm:"column:work_count" json:"work_count"`               // 作品数量
//	FavoriteCount   int32  `gorm:"column:favorite_count" json:"favorite_count"`       // 点赞数量
//	IsFowllow       bool   `json:"is_follow"`                                         // 是否关注
//}

type DouyinUserRegisterResponse struct {
	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg,omitempty"` // 返回状态描述
	UserId     int64  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`               // 用户id
	Token      string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`                                // 用户鉴权token
}

type DouyinUserLoginResponse struct {
	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg,omitempty"` // 返回状态描述
	UserId     int64  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`               // 用户id
	Token      string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`                                // 用户鉴权token
}

type DouyinUserResponse struct {
	StatusCode int32      `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg  string     `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg,omitempty"` // 返回状态描述
	User       model.User `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`                                  // 用户信息
}

// Register 用户注册
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	md5Password := utils.Md5(password)

	if username == "" || password == "" {
		c.JSON(http.StatusOK, DouyinUserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  "bad register request",
		})
		return
	}

	userId, err := serviceImpl.Register(username, md5Password)
	if err != nil {
		c.JSON(http.StatusOK, DouyinUserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  "registration failed",
		})
		return
	}
	if userId == -1 {
		c.JSON(http.StatusOK, DouyinUserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  "user already exists",
		})
		return
	}
	//注册成功
	c.JSON(http.StatusOK, DouyinUserRegisterResponse{
		StatusCode: 0,
		StatusMsg:  "registration success",
		UserId:     userId,
		Token:      utils.BuildToken(userId, username),
	})
}

// GetUserData 获取用户信息
func GetUserData(c *gin.Context) {
	userId := c.Query("user_id")
	token := c.Query("token")

	if userId == "" || token == "" {
		c.JSON(http.StatusOK, DouyinUserResponse{
			StatusCode: -1,
			StatusMsg:  "invalid userId or token",
		})
		return
	}

	userid, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, DouyinUserResponse{
			StatusCode: -1,
			StatusMsg:  "invalid userId",
		})
		return
	}

	resId := utils.ParseToken(token)
	if resId != userid {
		c.JSON(http.StatusOK, DouyinUserResponse{
			StatusCode: -1,
			StatusMsg:  "invalid token",
		})
		return
	}

	user, err := serviceImpl.GetUserById(userid)
	//获取用户数据失败
	if err != nil {
		c.JSON(http.StatusOK, DouyinUserResponse{
			StatusCode: -1,
			StatusMsg:  "Failed to get user information",
		})
		return
	}
	//成功获取用户信息
	c.JSON(http.StatusOK, DouyinUserResponse{
		StatusCode: 0,
		StatusMsg:  "Successfully obtained user information",
		User:       user,
	})

}

// Login 用户登录
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	md5Password := utils.Md5(password)

	if username == "" || password == "" {
		c.JSON(http.StatusOK, DouyinUserLoginResponse{
			StatusCode: -1,
			StatusMsg:  "invalid username or password",
		})
		return
	}

	userId, succ, err := serviceImpl.Login(username, md5Password)
	if err != nil || !succ {
		c.JSON(http.StatusOK, DouyinUserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  "registration failed",
		})
		return
	}
	//生成token
	token := utils.BuildToken(userId, username)

	//登录成功
	c.JSON(http.StatusOK, DouyinUserLoginResponse{
		StatusCode: 0,
		StatusMsg:  "successfully login",
		UserId:     userId,
		Token:      token,
	})
}
