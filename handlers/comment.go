package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"tk/dao/model"
	"tk/service/serviceImpl"
)

type DouyinCommentActionResponse struct {
	StatusCode int32         `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string        `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`     // 返回状态描述
	Comment    model.Comment `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`                          // 评论成功返回评论内容，不需要重新拉取整个列表
}

type DouyinCommentListResponse struct {
	StatusCode  int32           `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg   string          `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg,omitempty"` // 返回状态描述
	CommentList []model.Comment `protobuf:"bytes,3,rep,name=comment_list,json=commentList,proto3" json:"comment_list,omitempty"` // 评论列表
}

func CommentAction(ctx *gin.Context) {
	userId, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusOK, DouyinCommentActionResponse{
			StatusCode: -1,
			StatusMsg:  "user not logged in",
			Comment:    model.Comment{},
		})
	}
	userid, err := strconv.ParseInt(userId.(string), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinCommentActionResponse{
			StatusCode: -1,
			StatusMsg:  "invalid userId",
			Comment:    model.Comment{},
		})
	}

	//videoId解析
	videoId := ctx.Query("video_id")
	videoid, err := strconv.ParseInt(videoId, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinCommentActionResponse{
			StatusCode: -1,
			StatusMsg:  "invalid videoId",
			Comment:    model.Comment{},
		})
	}

	//text解析
	text := ctx.Query("comment_text")
	if text == "" {
		ctx.JSON(http.StatusOK, DouyinCommentActionResponse{
			StatusCode: -1,
			StatusMsg:  "invalid commentText",
			Comment:    model.Comment{},
		})
	}

	//actionType解析
	act := ctx.Query("action_type ")
	actionType, err := strconv.ParseInt(act, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinCommentActionResponse{
			StatusCode: -1,
			StatusMsg:  "invalid actionType ",
			Comment:    model.Comment{},
		})
	}

	comment := model.Comment{
		UserID:     userid,
		VideoID:    videoid,
		Content:    text,
		IsDeleted:  int32(actionType),
		CreateTime: time.Now(),
	}

	err = serviceImpl.Comment(comment)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinCommentActionResponse{
			StatusCode: -1,
			StatusMsg:  err.Error(),
			Comment:    model.Comment{},
		})
	}

	ctx.JSON(http.StatusOK, DouyinCommentActionResponse{
		StatusCode: 0,
		StatusMsg:  "comment success",
		Comment:    comment,
	})
}

func UserCommentList(ctx *gin.Context) {
	userId, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusOK, DouyinCommentListResponse{
			StatusCode:  -1,
			StatusMsg:   "user not logged in",
			CommentList: nil,
		})
		return
	}
	userid, err := strconv.ParseInt(userId.(string), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinCommentListResponse{
			StatusCode:  -1,
			StatusMsg:   "invalid userId",
			CommentList: nil,
		})
		return
	}

	userCommentList, err := serviceImpl.GetUserCommentList(userid)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinCommentListResponse{
			StatusCode:  -1,
			StatusMsg:   err.Error(),
			CommentList: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, DouyinCommentListResponse{
		StatusCode:  0,
		StatusMsg:   "get user_comment_list success",
		CommentList: userCommentList,
	})
}

func VideoCommentList(ctx *gin.Context) {
	//videoId解析
	videoId := ctx.Query("video_id")
	videoid, err := strconv.ParseInt(videoId, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinCommentListResponse{
			StatusCode:  -1,
			StatusMsg:   "invalid videoId",
			CommentList: nil,
		})
		return
	}
	commentList, err := serviceImpl.GetVideoCommentList(videoid)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinCommentListResponse{
			StatusCode:  -1,
			StatusMsg:   err.Error(),
			CommentList: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, DouyinCommentListResponse{
		StatusCode:  0,
		StatusMsg:   "get video_comment_list success",
		CommentList: commentList,
	})
}
