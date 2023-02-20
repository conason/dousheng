package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	ptime "time"
	"tk/dao/model"
	"tk/service/serviceImpl"
)

type DouyinMessageActionResponse struct {
	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg,omitempty"` // 返回状态描述
}

type DouyinMessageChatResponse struct {
	StatusCode  int32           `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg   string          `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg,omitempty"` // 返回状态描述
	MessageList []model.Message `protobuf:"bytes,3,rep,name=message_list,json=messageList,proto3" json:"message_list,omitempty"` // 消息列表
}

func Send(ctx *gin.Context) {
	token := ctx.Query("token")
	toUserIdStr := ctx.Query("to_user_id")
	//没啥用的参数
	//actionStr := ctx.Query("action_type")
	content := ctx.Query("content")
	userId, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusOK, DouyinMessageActionResponse{
			StatusCode: -1,
			StatusMsg:  "unknown user",
		})
		return
	}

	if token == "" {
		ctx.JSON(http.StatusOK, DouyinMessageActionResponse{
			StatusCode: -1,
			StatusMsg:  "invalid token",
		})
		return
	}

	toUserId, err := strconv.ParseInt(toUserIdStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinMessageActionResponse{
			StatusCode: -1,
			StatusMsg:  "invalid toUserId",
		})
		return
	}

	if content == "" {
		ctx.JSON(http.StatusOK, DouyinMessageActionResponse{
			StatusCode: -1,
			StatusMsg:  "content can not be blank",
		})
		return
	}

	err = serviceImpl.SendMsg(model.Message{
		ToUserID:   toUserId,
		FromUserID: userId.(int64),
		Content:    content,
		CreateTime: ptime.Now(),
	})
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinMessageActionResponse{
			StatusCode: -1,
			StatusMsg:  "send message failed",
		})
		return
	}

	ctx.JSON(http.StatusOK, DouyinMessageActionResponse{
		StatusCode: 0,
		StatusMsg:  "sean message successfully",
	})

}

func Receive(ctx *gin.Context) {
	token := ctx.Query("token")
	toUserIdStr := ctx.Query("to_user_id")
	preMsgTime := ctx.Query("pre_msg_time")

	if token == "" {
		ctx.JSON(http.StatusOK, DouyinMessageChatResponse{
			StatusCode:  -1,
			StatusMsg:   "invalid token",
			MessageList: nil,
		})
		return
	}

	toUserId, err := strconv.ParseInt(toUserIdStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinMessageChatResponse{
			StatusCode:  -1,
			StatusMsg:   "invalid toUserId",
			MessageList: nil,
		})
		return
	}

	preTimeInt, err := strconv.ParseInt(preMsgTime, 10, +64)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinMessageChatResponse{
			StatusCode:  -1,
			StatusMsg:   "invalid pre_msg_time",
			MessageList: nil,
		})
		return
	}
	timeTemplate := "2006-01-02 15:04:05"
	unix := ptime.Unix(preTimeInt, 0)
	preTime := unix.Format(timeTemplate)
	ptime, err := ptime.Parse(timeTemplate, preTime)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinMessageChatResponse{
			StatusCode:  -1,
			StatusMsg:   "pre_msg_time format conversion failed",
			MessageList: nil,
		})
		return
	}

	messages, err := serviceImpl.ReceiveMsg(toUserId, ptime)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinMessageChatResponse{
			StatusCode:  -1,
			StatusMsg:   "receive message failed",
			MessageList: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, DouyinMessageChatResponse{
		StatusCode:  0,
		StatusMsg:   "receive message successfully",
		MessageList: messages,
	})
}
