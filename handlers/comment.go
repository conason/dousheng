package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"tk/dao"
	"tk/dao/model"
	"tk/service/serviceImpl"
)

type Comment struct {
	Id         int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                  // 视频评论id
	User       model.User `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`                               // 评论用户信息
	Content    string     `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`                         // 评论内容
	CreateDate string     `protobuf:"bytes,4,opt,name=create_date,json=createDate,proto3" json:"create_date,omitempty"` // 评论发布日期，格式 mm-dd
}

type DouyinCommentActionResponse struct {
	StatusCode int32   `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string  `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`     // 返回状态描述
	Comment    Comment `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`                          // 评论成功返回评论内容，不需要重新拉取整个列表
}

type DouyinCommentListResponse struct {
	StatusCode  int32     `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg   string    `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg,omitempty"` // 返回状态描述
	CommentList []Comment `protobuf:"bytes,3,rep,name=comment_list,json=commentList,proto3" json:"comment_list,omitempty"` // 评论列表
}

func CommentAction(ctx *gin.Context) {
	userId, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusOK, DouyinCommentActionResponse{
			StatusCode: -1,
			StatusMsg:  "user not logged in",
			Comment:    Comment{},
		})
		return
	}
	userid, err := strconv.ParseInt(userId.(string), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinCommentActionResponse{
			StatusCode: -1,
			StatusMsg:  "invalid userId",
			Comment:    Comment{},
		})
		return
	}

	//videoId解析
	videoId := ctx.Query("video_id")
	videoid, err := strconv.ParseInt(videoId, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinCommentActionResponse{
			StatusCode: -1,
			StatusMsg:  "invalid videoId",
			Comment:    Comment{},
		})
		return
	}

	//text解析
	text := ctx.Query("comment_text")
	if text == "" {
		ctx.JSON(http.StatusOK, DouyinCommentActionResponse{
			StatusCode: -1,
			StatusMsg:  "invalid commentText",
			Comment:    Comment{},
		})
		return
	}

	//actionType解析
	act := ctx.Query("action_type ")
	actionType, err := strconv.ParseInt(act, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinCommentActionResponse{
			StatusCode: -1,
			StatusMsg:  "invalid actionType ",
			Comment:    Comment{},
		})
		return
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
			Comment:    Comment{},
		})
		return
	}

	user, err := dao.GetUserById(userid)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinCommentActionResponse{
			StatusCode: -1,
			StatusMsg:  err.Error(),
			Comment:    Comment{},
		})
		return
	}

	c := Comment{
		User:       user,
		Content:    text,
		CreateDate: string(comment.CreateTime.Unix()),
	}

	ctx.JSON(http.StatusOK, DouyinCommentActionResponse{
		StatusCode: 0,
		StatusMsg:  "comment success",
		Comment:    c,
	})
}

func UserCommentList(ctx *gin.Context) {
	userId, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusOK, DouyinCommentListResponse{
			StatusCode:  -1,
			StatusMsg:   "user not logged in",
			CommentList: []Comment{},
		})
		return
	}
	userid, err := strconv.ParseInt(userId.(string), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinCommentListResponse{
			StatusCode:  -1,
			StatusMsg:   "invalid userId",
			CommentList: []Comment{},
		})
		return
	}

	userCommentList, err := serviceImpl.GetUserCommentList(userid)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinCommentListResponse{
			StatusCode:  -1,
			StatusMsg:   err.Error(),
			CommentList: []Comment{},
		})
		return
	}

	user, err := dao.GetUserById(userid)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinCommentListResponse{
			StatusCode:  -1,
			StatusMsg:   err.Error(),
			CommentList: []Comment{},
		})
		return
	}

	len := len(userCommentList)
	var comments = make([]Comment, len)
	for i := 0; i < len; i++ {
		comments[i] = Comment{
			Id:         userCommentList[i].CommentID,
			User:       user,
			Content:    userCommentList[i].Content,
			CreateDate: string(userCommentList[i].CreateTime.Unix()),
		}
	}
	ctx.JSON(http.StatusOK, DouyinCommentListResponse{
		StatusCode:  0,
		StatusMsg:   "get user_comment_list success",
		CommentList: comments,
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

	len := len(commentList)
	var comments = make([]Comment, len)
	for i := 0; i < len; i++ {
		user, err := dao.GetUserById(commentList[i].UserID)
		if err != nil {
			ctx.JSON(http.StatusOK, DouyinCommentListResponse{
				StatusCode:  -1,
				StatusMsg:   err.Error(),
				CommentList: nil,
			})
			return
		}
		comments[i] = Comment{
			Id:         commentList[i].CommentID,
			User:       user,
			Content:    commentList[i].Content,
			CreateDate: string(commentList[i].CreateTime.Unix()),
		}
	}

	ctx.JSON(http.StatusOK, DouyinCommentListResponse{
		StatusCode:  0,
		StatusMsg:   "get video_comment_list success",
		CommentList: comments,
	})
}
