package handlers

import (
	"dousheng/dao"
	"dousheng/dao/model"
	"dousheng/service/serviceImpl"
	"dousheng/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
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

	//时间戳转日期
	timeNext, err := utils.TimestampToDate(timeInt)
	utils.ResolveError(err)

	timestamp, err := utils.GetTimestamp()
	utils.ResolveError(err)

	videoData, err := serviceImpl.GetVideo(timeNext)
	if err != nil || len(videoData) == 0 {
		videoData, err = dao.GetNewestVideos()
		if err != nil {
			c.JSON(http.StatusOK, DouyinFeedResponse{
				StatusCode: -1,
				StatusMsg:  "server is busy please try again",
				VideoList:  nil,
				NextTime:   timestamp,
			})
		}

		videos, err := transformVideos(videoData)
		utils.ResolveError(err)

		c.JSON(http.StatusOK, DouyinFeedResponse{
			StatusCode: -1,
			StatusMsg:  "feed video",
			VideoList:  videos,
			NextTime:   videos[0].UpdateDate.Unix(),
		})
		return
	}

	videos, err := transformVideos(videoData)
	utils.ResolveError(err)

	//feed响应
	c.JSON(http.StatusOK, DouyinFeedResponse{
		StatusCode: 0,
		StatusMsg:  "feed video",
		VideoList:  videos,
		NextTime:   videos[0].UpdateDate.Unix(),
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
	//将*multipart.FileHeader类型转化为[]byte
	parseVideo, err := serviceImpl.ParseVideo(videoData)
	if err != nil {
		c.JSON(http.StatusOK, DouyinPublishActionResponse{
			StatusCode: -1,
			StatusMsg:  "server is busy please try again",
		})
		return
	}
	//视频文件名
	videoName := fmt.Sprintf("%s.mp4", title)
	//封面文件名
	//coverName := strings.Replace(videoName, ".mp4", ".jpeg", 1)
	//fmt.Printf("%s\n", title)
	//视频上传
	code := utils.PushVideo(videoName, parseVideo)
	if code != 0 {
		c.JSON(http.StatusOK, DouyinPublishActionResponse{
			StatusCode: -1,
			StatusMsg:  "server is busy please try again",
		})
		return
	}
	//获取playURL
	playURL := utils.GetVideo(videoName)
	//上传至数据库
	err = serviceImpl.PushVideoToMysql(userId, playURL, "", title)
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
		}
	}
	ctx.JSON(http.StatusOK, DouyinPublishListResponse{
		StatusCode: 0,
		StatusMsg:  "get publish_list successfully",
		VideoList:  videoList,
	})
}

func transformVideos(videoData []model.Video) ([]Video, error) {
	count := len(videoData)
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
		}
	}
	return video, nil
}
