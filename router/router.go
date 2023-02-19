package router

import (
	"dousheng/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	/* @ShuCoding
	basic service：
		register、login、feed、getUserData
	 */
	basicAPI := r.Group("/douyin")
	//basicAPI.Use(middleware.JwtToken())
	basicAPI.POST("/user/register/", api.Register)
	basicAPI.GET("/user/", api.GetUserData)
	basicAPI.POST("/user/login/", api.Login)
	basicAPI.GET("/feed/", api.Feed)
	basicAPI.POST("/publish/action/",api.VideoPublish)
	err := r.Run(":8088")
	if err != nil {
		panic(err)
	}
}
