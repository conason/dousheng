package handlers

import (
	"dousheng/config"
	"dousheng/dao"
	"dousheng/dao/model"
	"dousheng/service/serviceImpl"
	"dousheng/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
	"strconv"
)

//type Video struct {
//	ID            int64  `json:"id"`             // 视频id
//	User          User   `json:"author"`         // 视频作者
//	PlayURL       string `json:"play_url"`       // 视频URL
//	CoverURL      string `json:"cover_url"`      // 封面URL
//	FavoriteCount int32  `json:"favorite_count"` // 点赞总数
//	CommentCount  int32  `json:"comment_count"`  // 评论总数
//	IsFavorite    bool   `json:"is_favorite"`    // 是否点赞
//	Title         string `json:"title"`          // 视频标题
//}

type DouyinFeedResponse struct {
	StatusCode int32         `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg  string        `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg,omitempty"` // 返回状态描述
	VideoList  []model.Video `protobuf:"bytes,3,rep,name=video_list,json=videoList,proto3" json:"video_list,omitempty"`       // 视频列表
	NextTime   int64         `protobuf:"varint,4,opt,name=next_time,json=nextTime,proto3,oneof" json:"next_time,omitempty"`   // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

type DouyinPublishActionResponse struct {
	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg,omitempty"` // 返回状态描述
}

type DouyinPublishListResponse struct {
	StatusCode int32         `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg  string        `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg,omitempty"` // 返回状态描述
	VideoList  []model.Video `protobuf:"bytes,3,rep,name=video_list,json=videoList,proto3" json:"video_list,omitempty"`       // 用户发布的视频列表
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
	//log.Println("timeInt:%V\n", timeInt)

	// 时间戳转日期
	//timeNext, err := utils.TimestampToDate(timeInt)
	//utils.ResolveError(err)

	timestamp, err := utils.GetTimestamp()
	utils.ResolveError(err)
	//log.Println("timestamp:%V\n", timestamp)

	if timeInt > timestamp {
		timeInt /= 1000
	}

	redisVideos, err := utils.Redis.ZRevRange(utils.Ctx, config.VIDEOSKEY, timeInt, timestamp).Result()
	if err != nil {
		//log.Println(err)
		c.JSON(http.StatusOK, DouyinFeedResponse{
			StatusCode: -1,
			StatusMsg:  "server is busy please try again",
			VideoList:  nil,
			NextTime:   timestamp,
		})
		return
	}
	// redis 返回为空
	if len(redisVideos) <= 0 {
		//fmt.Println("yse")
		redisVideos, err = utils.Redis.ZRevRange(utils.Ctx, config.VIDEOSKEY, timestamp, 0).Result()
		if err != nil {
			//log.Println(err)
			c.JSON(http.StatusOK, DouyinFeedResponse{
				StatusCode: -1,
				StatusMsg:  "server is busy please try again",
				VideoList:  nil,
				NextTime:   timestamp,
			})
			return
		}
	}
	// video缓存为空的情况
	if len(redisVideos) <= 0 {
		c.JSON(http.StatusOK, DouyinFeedResponse{
			StatusCode: -1,
			StatusMsg:  "no video",
			VideoList:  nil,
			NextTime:   timestamp,
		})
		return
	}
	// 返回视频最早时间戳
	//if redisVideos == nil {
	//	return
	//}
	//print("redisVideo:", redisVideos)
	score, err := utils.Redis.ZScore(utils.Ctx, config.VIDEOSKEY, redisVideos[0]).Result()
	if err != nil {
		utils.ResolveError(err)
	}
	nextTime := int64(score)
	//fmt.Println(redisVideos)

	// 限制返回 video 数量
	var len = len(redisVideos)
	if len > config.N {
		len = config.N
	}
	videoData := make([]model.Video, len)
	for i := 0; i < len; i++ {
		err = json.Unmarshal([]byte(redisVideos[i]), &videoData[i])
		if err != nil {
			c.JSON(http.StatusOK, DouyinFeedResponse{
				StatusCode: -1,
				StatusMsg:  "feed video failed",
				VideoList:  nil,
				NextTime:   timestamp,
			})
			return
		}
	}

	//for i, video := range redisVideos {
	//	err = json.Unmarshal([]byte(video), &videoData[i])
	//	if err != nil {
	//		c.JSON(http.StatusOK, DouyinFeedResponse{
	//			StatusCode: -1,
	//			StatusMsg:  "feed video failed",
	//			VideoList:  nil,
	//			NextTime:   timestamp,
	//		})
	//		return
	//	}
	//}

	// 用户未登录
	if userId == 0 {
		c.JSON(http.StatusOK, DouyinFeedResponse{
			StatusCode: 0,
			StatusMsg:  "feed video",
			VideoList:  videoData,
			NextTime:   nextTime,
		})
		return
	}

	// 用户登录判断是否点赞\订阅
	videos, err := IfLoginTransVideos(videoData, userId)
	utils.ResolveError(err)

	// feed响应
	c.JSON(http.StatusOK, DouyinFeedResponse{
		StatusCode: 0,
		StatusMsg:  "feed success",
		VideoList:  videos,
		NextTime:   nextTime,
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
	//upload success add to cache
	err = addCache(userId)
	if err != nil {
		c.JSON(http.StatusOK, DouyinPublishActionResponse{
			StatusCode: -1,
			StatusMsg:  "upload cache failed",
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

	//len := len(videos)
	//user, _ := serviceImpl.GetUserById(userId)
	//videoList := make([]Video, len)
	//for i := 0; i < len; i++ {
	//	videoList[i] = Video{
	//		ID: videos[i].ID,
	//		User:          user,
	//		PlayURL:       videos[i].PlayURL,
	//		CoverURL:      videos[i].CoverURL,
	//		FavoriteCount: videos[i].FavoriteCount,
	//		CommentCount:  videos[i].CommentCount,
	//		Title:         videos[i].Title,
	//	}
	//}
	ctx.JSON(http.StatusOK, DouyinPublishListResponse{
		StatusCode: 0,
		StatusMsg:  "get publish_list successfully",
		VideoList:  videos,
	})
}

func IfLoginTransVideos(videoData []model.Video, userId int64) ([]model.Video, error) {
	for i := 0; i < len(videoData); i++ {
		//是否点赞
		isFav, err := dao.IsFav(userId, videoData[i].ID)
		if err != nil {
			utils.ResolveError(err)
		}
		videoData[i].IsFavorite = isFav
		//是否订阅
		isSub, err := dao.IsSub(userId, videoData[i].User.ID)
		if err != nil {
			return nil, err
		}
		videoData[i].User.IsFowllow = isSub
	}
	return videoData, nil
}

func TransVideos(videoData []model.Video) error {
	count := len(videoData)
	//video := make([]handlers.Video, count)
	//zset := make([]redis.Z, count)
	for i := 0; i < int(count); i++ {
		//完善视频作者信息
		videoUserId := videoData[i].UserID
		userData, err := serviceImpl.GetUserById(videoUserId)
		if err != nil {
			utils.ResolveError(err)
		}
		video := model.Video{
			User: userData,
		}
		// JSON
		videoJSON, err := json.Marshal(&video)
		if err != nil {
			return err
		}
		// 视频时间戳
		createTime := videoData[i].CreateDate.Unix()
		arg := redis.Z{
			Score:  float64(createTime),
			Member: videoJSON,
		}
		utils.Redis.ZAdd(utils.Ctx, config.VIDEOSKEY, &arg)
	}

	return nil

}

func addCache(userId int64) error {
	daoVideo, err := dao.GetVideoByUserId(userId)
	if err != nil {
		return err
	}
	user, err := dao.GetUserById(userId)
	if err != nil {
		return err
	}

	video := model.Video{
		ID:            daoVideo.ID,
		User:          user,
		PlayURL:       daoVideo.PlayURL,
		CoverURL:      daoVideo.CoverURL,
		FavoriteCount: daoVideo.FavoriteCount,
		CommentCount:  daoVideo.CommentCount,
		Title:         daoVideo.Title,
	}
	//解析为JSON
	videoJSON, err := json.Marshal(video)
	if err != nil {
		return err
	}
	score := float64(daoVideo.CreateDate.Unix())
	z := redis.Z{
		Score:  score,
		Member: videoJSON,
	}
	//加入redis
	utils.Redis.ZAdd(utils.Ctx, config.VIDEOSKEY, &z)
	return nil
}
