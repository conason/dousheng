package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tk/dao"
	"tk/dao/model"
	"tk/utils"
)

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
	code, err := dao.Register(username, md5Password)
	if err != nil || code != 0 {
		c.JSON(http.StatusOK, DouyinUserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  "registration failed",
		})
	}

	findUser, err := dao.FindUser(username)
	if err != nil {
		c.JSON(http.StatusOK, DouyinUserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  "please login again",
		})
	}

	//注册成功
	c.JSON(http.StatusOK, DouyinUserRegisterResponse{
		StatusCode: 0,
		StatusMsg:  "registration success",
		UserId:     findUser,
		Token:      utils.BuildToken(findUser, username),
	})

	//code := model.Register(username, md5Password)
	//if code == utils.SUCCESS {
	//status := model.UserRegister{
	//	Response: model.Response{
	//		StatusCode: code,
	//		StatusMsg:  utils.GetStatusMsg(utils.USER_SUCCESS_REGISTER),
	//	},
	//	UserId: model.FindUser(username),
	//	Token:  utils.BuildToken(model.FindUser(username), username),
	//}
	//	fmt.Println("注册成功")
	//	c.JSON(http.StatusOK, status)
	//} else {
	//	status := model.UserRegister{
	//		Response: model.Response{
	//			StatusCode: utils.FAIL,
	//			StatusMsg: utils.GetStatusMsg(utils.USER_FAIL_REGISTER),
	//		},
	//	}
	//	fmt.Println("注册失败")
	//	fmt.Println(status)
	//	c.JSON(http.StatusOK, status)
	//}
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
	}

	userid, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, DouyinUserResponse{
			StatusCode: -1,
			StatusMsg:  "invalid userId",
		})
	}

	resId := utils.ParseToken(token)
	if resId != userid {
		c.JSON(http.StatusOK, DouyinUserResponse{
			StatusCode: -1,
			StatusMsg:  "invalid token",
		})
	}

	user, err := dao.GetUserData(userid)
	//获取用户数据失败
	if err != nil {
		c.JSON(http.StatusOK, DouyinUserResponse{
			StatusCode: -1,
			StatusMsg:  "Failed to get user information",
		})
	}
	//成功获取用户信息
	c.JSON(http.StatusOK, DouyinUserResponse{
		StatusCode: 0,
		StatusMsg:  "Successfully obtained user information",
		User:       user,
	})

	//userid, _ := strconv.Atoi(userId)
	//id := int32(userid)
	//resId := utils.ParseToken(token)
	//if resId == id {
	//	userMsg := UserMsg{
	//		Response: model.Response{
	//			StatusCode: utils.SUCCESS,
	//			StatusMsg:  utils.GetStatusMsg(utils.SUCCESS),
	//		},
	//		User: model.GetUserData(int32(userid)),
	//	}
	//	c.JSON(http.StatusOK, userMsg)
	//} else {
	//	userMsg := UserMsg{
	//		Response: model.Response{
	//			StatusCode: utils.FAIL,
	//			StatusMsg:  utils.GetStatusMsg(utils.USER_NOT_EXIT),
	//		},
	//	}
	//	c.JSON(http.StatusOK, userMsg)
	//}
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
	}

	login, err := dao.Login(username, md5Password)
	if err != nil || login {
		c.JSON(http.StatusOK, DouyinUserLoginResponse{
			StatusCode: -1,
			StatusMsg:  "Login failed, please try again",
		})
	}

	userId, err := dao.FindUser(username)
	if err != nil {
		c.JSON(http.StatusOK, DouyinUserLoginResponse{
			StatusCode: -1,
			StatusMsg:  "Login failed, please try again",
		})
	}

	token := utils.BuildToken(userId, username)

	//登录成功
	c.JSON(http.StatusOK, DouyinUserLoginResponse{
		StatusCode: 0,
		StatusMsg:  "successfully login",
		UserId:     userId,
		Token:      token,
	})

	//res := model.Login(username, md5Password)
	//if res {
	//	user := model.UserRegister{
	//		Response: model.Response{
	//			StatusCode: utils.SUCCESS,
	//			StatusMsg:  utils.GetStatusMsg(utils.USER_SUCCESS_LOGIN),
	//		},
	//		UserId: model.FindUser(username),
	//		Token:  utils.BuildToken(model.FindUser(username), username),
	//	}
	//	c.JSON(http.StatusOK, user)
	//} else {
	//	user := model.UserRegister{
	//		Response: model.Response{
	//			StatusCode: utils.FAIL,
	//			StatusMsg:  utils.GetStatusMsg(utils.USER_PASSWORD_IS_NOT_CORRECT),
	//		},
	//		UserId: 0,
	//	}
	//	c.JSON(http.StatusOK, user)
	//}
}
