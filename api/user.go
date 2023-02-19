package api

import (
	"dousheng/model"
	"dousheng/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserMsg struct {
	model.Response
	User   model.User `json:"user"`
}
// Register 用户注册
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	md5Password := utils.Md5(password)
	code := model.Register(username, md5Password)
	if code == utils.SUCCESS {
		status := model.UserRegister{
			Response:model.Response{
				StatusCode: code,
				StatusMsg: utils.GetStatusMsg(utils.USER_SUCCESS_REGISTER),
			},
			UserId: model.FindUser(username),
			Token: utils.BuildToken(model.FindUser(username),username),
		}
		fmt.Println("注册成功")
		c.JSON(http.StatusOK, status)
	} else {
		status := model.UserRegister{
			Response: model.Response{
				StatusCode: utils.FAIL,
				StatusMsg: utils.GetStatusMsg(utils.USER_FAIL_REGISTER),
			},
		}
		fmt.Println("注册失败")
		fmt.Println(status)
		c.JSON(http.StatusOK, status)
	}
}
// GetUserData 获取用户信息
func GetUserData(c *gin.Context) {
	userId := c.Query("user_id")
	token:= c.Query("token")
	userid,_ := strconv.Atoi(userId)
	id := int32(userid)
	resId := utils.ParseToken(token)
	if resId==id{
		userMsg := UserMsg{
			Response:model.Response{
				StatusCode: utils.SUCCESS,
				StatusMsg: utils.GetStatusMsg(utils.SUCCESS),
			},
			User: model.GetUserData(int32(userid)),
		}
		c.JSON(http.StatusOK, userMsg)
	}else {
		userMsg := UserMsg{
			Response:model.Response{
				StatusCode: utils.FAIL,
				StatusMsg: utils.GetStatusMsg(utils.USER_NOT_EXIT),
			},
		}
		c.JSON(http.StatusOK, userMsg)
	}
}
//Login 用户登录
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	md5Password := utils.Md5(password)
	res := model.Login(username, md5Password)
	if res {
		user := model.UserRegister{
			Response: model.Response{
				StatusCode: utils.SUCCESS,
				StatusMsg: utils.GetStatusMsg(utils.USER_SUCCESS_LOGIN),
			},
			UserId: model.FindUser(username),
			Token: utils.BuildToken(model.FindUser(username),username),
		}
		c.JSON(http.StatusOK, user)
	} else {
		user := model.UserRegister{
			Response: model.Response{
				StatusCode: utils.FAIL,
				StatusMsg: utils.GetStatusMsg(utils.USER_PASSWORD_IS_NOT_CORRECT),
			},
			UserId: 0,
		}
		c.JSON(http.StatusOK, user)
	}
}
