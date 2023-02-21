package handlers

import (
	"dousheng/dao"
	"dousheng/service/serviceImpl"
	"dousheng/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DouyinFavoriteActionResponse struct {
	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg,omitempty"` // 返回状态描述
}

type DouyinFavoriteListResponse struct {
	StatusCode int32   `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg  string  `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg,omitempty"` // 返回状态描述
	VideoList  []Video `protobuf:"bytes,3,rep,name=video_list,json=videoList,proto3" json:"video_list,omitempty"`       // 用户点赞视频列表
}

func FavAction(ctx *gin.Context) {
	//userId解析
	//userIdStr := ctx.Query("user_id")
	//userid, err := strconv.ParseInt(userIdStr, 10, 64)
	//if err != nil {
	//	ctx.JSON(http.StatusOK, DouyinFavoriteActionResponse{
	//		StatusCode: -1,
	//		StatusMsg:  "invalid userId",
	//	})
	//	return
	//}

	token := ctx.Query("token")
	userId := utils.ParseToken(token)

	//videoId解析
	videoId := ctx.Query("video_id")
	videoid, err := strconv.ParseInt(videoId, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinFavoriteActionResponse{
			StatusCode: -1,
			StatusMsg:  "invalid video_id",
		})
		return
	}

	//actionType解析
	actType := ctx.Query("action_type")
	actionType, err := strconv.ParseInt(actType, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinFavoriteActionResponse{
			StatusCode: -1,
			StatusMsg:  "invalid action_type",
		})
		return
	}

	err = serviceImpl.FavAction(userId, videoid, int32(actionType))
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinFavoriteActionResponse{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, DouyinFavoriteActionResponse{
		StatusCode: 0,
		StatusMsg:  "favAction success",
	})
}

func FavList(ctx *gin.Context) {
	//鉴权

	//userId解析
	userId := ctx.Query("user_id")
	userid, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinFavoriteListResponse{
			StatusCode: -1,
			StatusMsg:  "invalid userId",
			VideoList:  nil,
		})
	}

	favList, err := serviceImpl.GetFavListByUserId(userid)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinFavoriteListResponse{
			StatusCode: -1,
			StatusMsg:  err.Error(),
			VideoList:  nil,
		})
	}
	//call video模块 resp-> []video
	len := len(favList)
	videos := make([]Video, len)
	user, err := dao.GetUserById(userid)
	if err != nil {
		utils.ResolveError(err)
	}
	for i := 0; i < len; i++ {
		video, err := dao.GetVideoById(favList[i])
		if err != nil {
			utils.ResolveError(err)
		}
		videos[i] = Video{
			ID:            video.ID,
			User:          user,
			PlayURL:       video.PlayURL,
			CoverURL:      video.CoverURL,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			Title:         video.Title,
			CreateDate:    video.CreateDate,
			UpdateDate:    video.UpdateDate,
		}
	}

	ctx.JSON(http.StatusOK, DouyinFavoriteListResponse{
		StatusCode: 0,
		StatusMsg:  "get faList success",
		VideoList:  videos,
	})
}
