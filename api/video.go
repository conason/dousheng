package api

import (
	"dousheng/model"
	"dousheng/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// Feed 视频feed流
//func Feed(c *gin.Context) {
//	demoUser := model.User{
//		Id:            1,
//		Username:      "shuxin",
//		FollowCount:   12,
//		FollowerCount: 13,
//		IsFollow:      false,
//	}
//	demo := []model.Video{{
//		Id:            1,
//		Author:        demoUser,
//		PlayUrl:       "http://rq9lt9dry.bkt.clouddn.com/ceshi.mp4?e=1708242197&token=XuigBGSCJ7vpAtRtpu04NqLGLXpEROCaqgOxTZ0W:cScrjx0quzLcdSNPImvbENW9epk=",
//		CoverUrl:      "http://rp8cwyjwy.hn-bkt.clouddn.com/douyin.mp4?e=1675504751&token=XuigBGSCJ7vpAtRtpu04NqLGLXpEROCaqgOxTZ0W:A5FTVckQlut9mvix4y1tzmTUXMU=",
//		FavoriteCount: 8555,
//		CommentCount:  131311,
//		IsFavorite:    false,
//		Title:         "hello golang",
//	},
//	}
//	video := model.VideoLists{
//		Response: model.Response{
//			StatusCode: utils.SUCCESS,
//		},
//		VideoList: demo,
//		NextTime:  int32(time.Now().UnixNano()),
//	}
//	c.JSON(http.StatusOK, video)
//}

// VideoPublish 视频发布
func VideoPublish(c *gin.Context) {
	title := c.PostForm("title")
	token := c.PostForm("token")
	videoData, err := c.FormFile("data")
	utils.ResolveError(err)
	userId := utils.ParseToken(token)
	data := model.ParseVideo(videoData)
	key := fmt.Sprintf("%s.mp4", title)
	code := utils.PushVideo(key, data)
	if code == utils.SUCCESS {
		model.PushVideoToMysql(userId, utils.GetVideo(fmt.Sprintf("%s.mp4", title)), "", title)
		c.JSON(http.StatusOK, model.Response{
			StatusCode: utils.SUCCESS,
			StatusMsg:  utils.GetStatusMsg(utils.VIDEO_PUSH_SUCCESS),
		})
	} else {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: utils.FAIL,
			StatusMsg:  utils.GetStatusMsg(utils.VIDEO_PUSH_FAIL),
		})
	}

}

func Feed(c *gin.Context) {
	latestTime := c.Query("latest_time")
	t ,err := time.Parse("2006-01-02 15:04",latestTime)
	utils.ResolveError(err)
	videoData,count := model.GetVideo(t.String())
	videoAndAuthor := make([]model.Video,count)
	for i := 0; i < int(count); i++ {
		user := model.GetUserData(videoData[i].Id)
		videoAndAuthor[i].Id=videoData[i].VideoId
		videoAndAuthor[i].PlayUrl=videoData[i].PlayUrl
		videoAndAuthor[i].CoverUrl=videoData[i].CoverUrl
		videoAndAuthor[i].CommentCount=videoData[i].CommentCount
		videoAndAuthor[i].FavoriteCount=videoData[i].FavoriteCount
		videoAndAuthor[i].IsFavorite=videoData[i].IsFavorite
		videoAndAuthor[i].Title=videoData[i].Title
		videoAndAuthor[i].Author=user
	}
	timeNow, err := strconv.ParseInt(time.Now().Format("2006-01-02 15:04"), 10, 32)
	utils.ResolveError(err)
	videoLists := model.VideoLists{
		Response:model.Response{
			StatusCode: utils.SUCCESS,
			StatusMsg: utils.GetStatusMsg(utils.VIDEO_GET_SUCCESS),
		},
		NextTime:int32(timeNow) ,
		VideoList: videoAndAuthor,
	}
	c.JSON(http.StatusOK,videoLists)
}
