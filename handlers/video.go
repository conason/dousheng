package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"tk/dao"
	"tk/dao/model"
	"tk/utils"
)

type Video struct {
	ID            int64      `gorm:"column:video_id;primaryKey;autoIncrement:true" json:"id"`                  // 视频id
	User          model.User `gorm:"column:user_id" json:"author"`                                             // 视频作者
	PlayURL       string     `gorm:"column:play_url" json:"play_url"`                                          // 视频URL
	CoverURL      string     `gorm:"column:cover_url" json:"cover_url"`                                        // 封面URL
	FavoriteCount int32      `gorm:"column:favorite_count" json:"favorite_count"`                              // 点赞总数
	CommentCount  int32      `gorm:"column:comment_count" json:"comment_count"`                                // 评论总数
	Title         string     `gorm:"column:title" json:"title"`                                                // 视频标题
	CreateDate    time.Time  `gorm:"column:create_date;not null;default:CURRENT_TIMESTAMP" json:"create_date"` // 创建时间
	UpdateDate    time.Time  `gorm:"column:update_date;not null;default:CURRENT_TIMESTAMP" json:"update_date"` // 更新时间
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
	latestTime := c.Query("latest_time")
	timeInt, err := strconv.ParseInt(latestTime, 10, 64)
	if err != nil {
		utils.ResolveError(err)
	}
	timeNext := time.Unix(timeInt, 0).Format("2006-01-02 15:04:05")
	utils.ResolveError(err)

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		utils.ResolveError(err)
	}
	timeNow := time.Now().Format("2006-01-02 15:04:05")
	nowTime, err := time.ParseInLocation("2006-01-02 15:04:05", timeNow, loc)
	utils.ResolveError(err)

	videoData, count := dao.GetVideo(timeNext)
	if err != nil {
		return
	}
	if count == 0 {
		c.JSON(http.StatusOK, DouyinFeedResponse{
			StatusCode: -1,
			StatusMsg:  "Reach bottom of video list",
			VideoList:  nil,
			NextTime:   nowTime.Unix(),
		})
	}

	video := make([]Video, count)
	for i := 0; i < int(count); i++ {
		userId := videoData[i].UserID
		user, err := dao.GetUserData(userId)
		if err != nil {
			utils.ResolveError(err)
		}
		video[i] = Video{
			ID:            videoData[i].ID,
			User:          user,
			PlayURL:       videoData[i].PlayURL,
			CoverURL:      videoData[i].CoverURL,
			FavoriteCount: videoData[i].FavoriteCount,
			CommentCount:  videoData[i].CommentCount,
			Title:         videoData[i].Title,
			CreateDate:    videoData[i].CreateDate,
			UpdateDate:    videoData[i].UpdateDate,
		}
	}

	//feed响应
	c.JSON(http.StatusOK, DouyinFeedResponse{
		StatusCode: 0,
		StatusMsg:  "feed video",
		VideoList:  video,
		NextTime:   nowTime.Unix(),
	})

	//videoAndAuthor := make([]model.Video, count)
	//for i := 0; i < int(count); i++ {
	//	user := model.GetUserData(videoData[i].Id)
	//	videoAndAuthor[i].Id = videoData[i].VideoId
	//	videoAndAuthor[i].PlayUrl = videoData[i].PlayUrl
	//	videoAndAuthor[i].CoverUrl = videoData[i].CoverUrl
	//	videoAndAuthor[i].CommentCount = videoData[i].CommentCount
	//	videoAndAuthor[i].FavoriteCount = videoData[i].FavoriteCount
	//	videoAndAuthor[i].IsFavorite = videoData[i].IsFavorite
	//	videoAndAuthor[i].Title = videoData[i].Title
	//	videoAndAuthor[i].Author = user
	//}
	//timeNow, err := strconv.ParseInt(time.Now().Format("2006-01-02 15:04"), 10, 64)
	//utils.ResolveError(err)
	//videoLists := model.VideoLists{
	//	Response: model.Response{
	//		StatusCode: utils.SUCCESS,
	//		StatusMsg:  utils.GetStatusMsg(utils.VIDEO_GET_SUCCESS),
	//	},
	//	NextTime:  int32(timeNow),
	//	VideoList: videoAndAuthor,
	//}
	//c.JSON(http.StatusOK, videoLists)
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

	userId := utils.ParseToken(token)

	parseVideo, err := dao.ParseVideo(videoData)
	if err != nil {
		c.JSON(http.StatusOK, DouyinPublishActionResponse{
			StatusCode: -1,
			StatusMsg:  "server is busy please try again",
		})
		return
	}

	key := fmt.Sprintf("%s.mp4", title)
	fmt.Printf("%s\n", title)
	code := utils.PushVideo(key, parseVideo)
	if code != 0 {
		c.JSON(http.StatusOK, DouyinPublishActionResponse{
			StatusCode: -1,
			StatusMsg:  "server is busy please try again",
		})
		return
	}

	playURL := utils.GetVideo(key)
	err = dao.PushVideoToMysql(userId, playURL, "", title)
	if err != nil {
		c.JSON(http.StatusOK, DouyinPublishActionResponse{
			StatusCode: -1,
			StatusMsg:  "server is busy please try again",
		})
		return
	}

	c.JSON(http.StatusOK, DouyinPublishActionResponse{
		StatusCode: 0,
		StatusMsg:  "publish_action successfully",
	})

	//userId := utils.ParseToken(token)
	//data := model.ParseVideo(videoData)
	//key := fmt.Sprintf("%s.mp4", title)
	//code := utils.PushVideo(key, data)
	//if code == utils.SUCCESS {
	//	model.PushVideoToMysql(userId, utils.GetVideo(fmt.Sprintf("%s.mp4", title)), "", title)
	//	c.JSON(http.StatusOK, model.Response{
	//		StatusCode: utils.SUCCESS,
	//		StatusMsg:  utils.GetStatusMsg(utils.VIDEO_PUSH_SUCCESS),
	//	})
	//} else {
	//	c.JSON(http.StatusOK, model.Response{
	//		StatusCode: utils.FAIL,
	//		StatusMsg:  utils.GetStatusMsg(utils.VIDEO_PUSH_FAIL),
	//	})
	//}

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
	videos, err := dao.GetVideosByUserId(userId)
	if err != nil {
		ctx.JSON(http.StatusOK, DouyinPublishListResponse{
			StatusCode: -1,
			StatusMsg:  "server Error",
			VideoList:  nil,
		})
	}

	len := len(videos)
	user, err := dao.GetUserData(userId)
	videoList := make([]Video, len)
	for i := 0; i < int(len); i++ {
		videoList[i] = Video{
			ID:            videos[i].ID,
			User:          user,
			PlayURL:       videos[i].PlayURL,
			CoverURL:      videos[i].CoverURL,
			FavoriteCount: videos[i].FavoriteCount,
			CommentCount:  videos[i].CommentCount,
			Title:         videos[i].Title,
			CreateDate:    videos[i].CreateDate,
			UpdateDate:    videos[i].UpdateDate,
		}
	}
	ctx.JSON(http.StatusOK, DouyinPublishListResponse{
		StatusCode: 0,
		StatusMsg:  "get publish_list successfully",
		VideoList:  videoList,
	})
}
