package handlers

import (
	"dousheng/dao"
	"dousheng/dao/model"
	"dousheng/service/serviceImpl"
	"dousheng/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Video struct {
	ID            int64  `json:"id"`             // 视频id
	User          User   `json:"author"`         // 视频作者
	PlayURL       string `json:"play_url"`       // 视频URL
	CoverURL      string `json:"cover_url"`      // 封面URL
	FavoriteCount int32  `json:"favorite_count"` // 点赞总数
	CommentCount  int32  `json:"comment_count"`  // 评论总数
	IsFavorite    bool   `json:"is_favorite"`    // 是否点赞
	Title         string `json:"title"`          // 视频标题
}

type DouyinFeedResponse struct {
	StatusCode int32   `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg  string  `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg,omitempty"` // 返回状态描述
	VideoList  []Video `protobuf:"bytes,3,rep,name=video_list,json=videoList,proto3" json:"video_list,omitempty"`       // 视频列表
	NextTime   int64   `protobuf:"varint,4,opt,name=next_time,json=nextTime,proto3,oneof" json:"next_time,omitempty"`   // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

type DouyinPublishActionResponse struct {
	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg,omitempty"` // 返回状态描述
}

type DouyinPublishListResponse struct {
	StatusCode int32   `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg  string  `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg,omitempty"` // 返回状态描述
	VideoList  []Video `protobuf:"bytes,3,rep,name=video_list,json=videoList,proto3" json:"video_list,omitempty"`       // 用户发布的视频列表
}

func Feed(c *gin.Context) {
	token := c.Query("token")
	userId := int64(0)
	if token != "" {
		userId = utils.ParseToken(token)
	}
	latestTime := c.Query("latest_time")
	timeInt, err := strconv.ParseInt(latestTime, 10, 64)
	if err != nil {
		utils.ResolveError(err)
	}

	// 时间戳转日期
	timeNext, err := utils.TimestampToDate(timeInt)
	utils.ResolveError(err)

	timestamp, err := utils.GetTimestamp()
	utils.ResolveError(err)
	//fmt.Printf("time:%v\n", timestamp)

	videoData, err := serviceImpl.GetVideo(timeNext)
	if err != nil || len(videoData) == 0 {
		videoData, err := serviceImpl.GetNewestVideos()
		if err != nil {
			c.JSON(http.StatusOK, DouyinFeedResponse{
				StatusCode: -1,
				StatusMsg:  "server is busy please try again",
				VideoList:  nil,
				NextTime:   timestamp,
			})
		}

		videos, err := transformVideos(videoData, userId)
		utils.ResolveError(err)
		//fmt.Printf("nextTime:%v\n", videoData[0].UpdateDate.Unix())
		c.JSON(http.StatusOK, DouyinFeedResponse{
			StatusCode: -1,
			StatusMsg:  "feed video",
			VideoList:  videos,
			NextTime:   videoData[0].UpdateDate.Unix(),
		})
		return
	}

	videos, err := transformVideos(videoData, userId)
	utils.ResolveError(err)

	// feed响应
	//fmt.Printf("nextTime:%v\n", videoData[0].UpdateDate.Unix())
	c.JSON(http.StatusOK, DouyinFeedResponse{
		StatusCode: 0,
		StatusMsg:  "feed video",
		VideoList:  videos,
		NextTime:   videoData[0].UpdateDate.Unix(),
	})

}

// VideoPublish 视频发布
func VideoPublish(c *gin.Context) {
	title := c.PostForm("title")
	token := c.PostForm("token")
	videoData, err := c.FormFile("data")
	if err != nil {
		utils.ResolveError(err)
	}

	if token == "" || title == "" || videoData == nil {
		c.JSON(http.StatusOK, DouyinPublishActionResponse{
			StatusCode: -1,
			StatusMsg:  "invalid videoPublish request",
		})
		return
	}
	//token解析
	userId := utils.ParseToken(token)

	//upload video
	err = serviceImpl.Upload(videoData, title, userId)
	if err != nil {
		c.JSON(http.StatusOK, DouyinPublishActionResponse{
			StatusCode: -1,
			StatusMsg:  "video upload failed",
		})
		return
	}

	c.JSON(http.StatusOK, DouyinPublishActionResponse{
		StatusCode: 0,
		StatusMsg:  "publish_action successfully",
	})

}

func PublishList(ctx *gin.Context) {
	userStr := ctx.Query("user_id")
	userId, err := strconv.ParseInt(userStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinPublishListResponse{
			StatusCode: -1,
			StatusMsg:  "invalid request",
			VideoList:  nil,
		})
	}

	videos, err := serviceImpl.GetVideosByUserId(userId)
	if err != nil {
		return
	}
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinPublishListResponse{
			StatusCode: -1,
			StatusMsg:  "server Error",
			VideoList:  nil,
		})
	}

	len := len(videos)
	//user, _ := serviceImpl.GetUserById(userId)
	videoList := make([]Video, len)
	for i := 0; i < int(len); i++ {
		videoList[i] = Video{
			ID: videos[i].ID,
			//User:          user,
			PlayURL:       videos[i].PlayURL,
			CoverURL:      videos[i].CoverURL,
			FavoriteCount: videos[i].FavoriteCount,
			CommentCount:  videos[i].CommentCount,
			Title:         videos[i].Title,
		}
	}
	ctx.JSON(http.StatusOK, DouyinPublishListResponse{
		StatusCode: 0,
		StatusMsg:  "get publish_list successfully",
		VideoList:  videoList,
	})
}

func transformVideos(videoData []model.Video, userId int64) ([]Video, error) {
	count := len(videoData)
	video := make([]Video, count)
	for i := 0; i < int(count); i++ {
		videoUserId := videoData[i].UserID
		userData, err := serviceImpl.GetUserById(videoUserId)
		if err != nil {
			utils.ResolveError(err)
		}
		//是否点赞
		isFav, err := dao.IsFav(userId, videoData[i].ID)
		if err != nil {
			utils.ResolveError(err)
		}
		isSub, err := dao.IsSub(userId, videoUserId)
		if err != nil {
			return nil, err
		}
		user := User{
			ID:              userData.ID,
			Name:            userData.Name,
			FollowCount:     userData.FollowCount,
			FollowerCount:   userData.FollowerCount,
			BackgroundImage: userData.BackgroundImage,
			Signature:       userData.Signature,
			TotalFavorited:  userData.TotalFavorited,
			WorkCount:       userData.WorkCount,
			FavoriteCount:   userData.FavoriteCount,
			IsFowllow:       isSub,
		}
		video[i] = Video{
			ID:            videoData[i].ID,
			User:          user,
			PlayURL:       videoData[i].PlayURL,
			CoverURL:      videoData[i].CoverURL,
			FavoriteCount: videoData[i].FavoriteCount,
			CommentCount:  videoData[i].CommentCount,
			IsFavorite:    isFav,
			Title:         videoData[i].Title,
		}
	}
	return video, nil
}
