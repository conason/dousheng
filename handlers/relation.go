package handlers

import (
	"dousheng/dao/model"
	"dousheng/service/serviceImpl"
	"dousheng/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//type FriendUser struct {
//	user    model.User `json:"user"`
//	message string     `json:"message"`
//	msgType int64      `json:"msgType"`
//}

type DouyinRelationActionResponse struct {
	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg,omitempty"` // 返回状态描述
}

type DouyinRelationFollowListResponse struct {
	StatusCode int32        `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg  string       `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg,omitempty"` // 返回状态描述
	UserList   []model.User `protobuf:"bytes,3,rep,name=user_list,json=userList,proto3" json:"user_list,omitempty"`          // 用户信息列表
}

type DouyinRelationFollowerListResponse struct {
	StatusCode int32        `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg  string       `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg,omitempty"` // 返回状态描述
	UserList   []model.User `protobuf:"bytes,3,rep,name=user_list,json=userList,proto3" json:"user_list,omitempty"`          // 用户列表
}

type DouyinRelationFriendListResponse struct {
	StatusCode int32        `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg  string       `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg,omitempty"` // 返回状态描述
	UserList   []model.User `protobuf:"bytes,3,rep,name=user_list,json=userList,proto3" json:"user_list,omitempty"`          // 用户列表
}

func Relation(ctx *gin.Context) {
	token := ctx.Query("token")
	toUserIdStr := ctx.Query("to_user_id")
	actionStr := ctx.Query("action_type")

	if token == "" {
		ctx.JSON(http.StatusOK, DouyinRelationActionResponse{
			StatusCode: -1,
			StatusMsg:  "invalid token",
		})
		return
	}
	//token解析
	userId := utils.ParseToken(token)

	toUserId, err := strconv.ParseInt(toUserIdStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinRelationActionResponse{
			StatusCode: -1,
			StatusMsg:  "invalid toUserId",
		})
		return
	}

	actionType, err := strconv.ParseInt(actionStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinRelationActionResponse{
			StatusCode: -1,
			StatusMsg:  "invalid actionType",
		})
		return
	}

	err = serviceImpl.SubAction(userId, toUserId, actionType)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinRelationActionResponse{
			StatusCode: -1,
			StatusMsg:  "sub operation failed",
		})
		return
	}

	ctx.JSON(http.StatusOK, DouyinRelationActionResponse{
		StatusCode: 0,
		StatusMsg:  "sub Successfully",
	})

}

func FollowList(ctx *gin.Context) {
	userIdStr := ctx.Query("user_id")
	token := ctx.Query("token")

	if token == "" {
		ctx.JSON(http.StatusOK, DouyinRelationFollowListResponse{
			StatusCode: -1,
			StatusMsg:  "invalid token",
			UserList:   nil,
		})
		return
	}

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinRelationFollowListResponse{
			StatusCode: -1,
			StatusMsg:  "invalid userId",
			UserList:   nil,
		})
		return
	}

	subList, err := serviceImpl.SubList(userId)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinRelationFollowListResponse{
			StatusCode: -1,
			StatusMsg:  "get subList failed",
			UserList:   nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, DouyinRelationFollowListResponse{
		StatusCode: 0,
		StatusMsg:  "get subList successfully",
		UserList:   subList,
	})
}

func FollowerList(ctx *gin.Context) {
	userIdStr := ctx.Query("user_id")
	token := ctx.Query("token")

	if token == "" {
		ctx.JSON(http.StatusOK, DouyinRelationFollowerListResponse{
			StatusCode: -1,
			StatusMsg:  "invalid token",
			UserList:   nil,
		})
		return
	}

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinRelationFollowerListResponse{
			StatusCode: -1,
			StatusMsg:  "invalid userId",
			UserList:   nil,
		})
		return
	}

	fansList, err := serviceImpl.FansList(userId)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinRelationFollowerListResponse{
			StatusCode: -1,
			StatusMsg:  "get fansList failed",
			UserList:   nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, DouyinRelationFollowerListResponse{
		StatusCode: 0,
		StatusMsg:  "get fansList successfully",
		UserList:   fansList,
	})

}

func FriendList(ctx *gin.Context) {
	userIdStr := ctx.Query("user_id")
	token := ctx.Query("token")

	if token == "" {
		ctx.JSON(http.StatusOK, DouyinRelationFriendListResponse{
			StatusCode: -1,
			StatusMsg:  "invalid token",
			UserList:   nil,
		})
		return
	}

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinRelationFriendListResponse{
			StatusCode: -1,
			StatusMsg:  "invalid userId",
			UserList:   nil,
		})
		return
	}

	fansList, err := serviceImpl.FansList(userId)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinRelationFriendListResponse{
			StatusCode: -1,
			StatusMsg:  "get friendList failed",
			UserList:   nil,
		})
		return
	}
	fmt.Println(fansList)
	ctx.JSON(http.StatusOK, DouyinRelationFollowerListResponse{
		StatusCode: 0,
		StatusMsg:  "get friendList successfully",
		UserList:   fansList,
	})
}
