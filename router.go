package main

import (
	"github.com/gin-gonic/gin"
	"tk/handlers"
)

func InitRouter(r *gin.Engine) {
	douying := r.Group("/douyin")

	//feed
	douying.GET("/feed/", handlers.Feed)

	//user
	user := douying.Group("/user")
	user.POST("/login/", handlers.Login)
	user.POST("/register/", handlers.Register)
	user.GET("/", handlers.GetUserData)

	//publish
	publish := douying.Group("/publish")
	publish.POST("/action/", handlers.VideoPublish)
	publish.POST("/list/", handlers.PublishList)

	//fav
	fav := douying.Group("/favorite")
	fav.POST("/action/", handlers.FavAction)
	fav.GET("/list/", handlers.FavList)

	//comment
	comment := douying.Group("/comment")
	comment.POST("/action/", handlers.CommentAction)
	comment.GET("/list/", handlers.VideoCommentList)

	//relation
	//relation := douying.Group("/relation")
	//relation.POST("/action/", handlers.Relation)
	//relation.GET("/follow/list/", handlers.FollowList)
	//relation.GET("/follower/list/", handlers.FollowerList)
	//relation.GET("/friend/list/", handlers.FriendList)

	//message
	//msg := douying.Group("/message")
	//msg.GET("/chat/", handlers.Chat)
	//msg.POST("/action/", handlers.ChatMsg)

}
