package router

import (
	"douyin-easy/api"
	"github.com/gin-gonic/gin"
)

func initFeed(r *gin.Engine) {
	r.GET("/douyin/feed", api.FeedApi.Feed)
	publish := r.Group("/douyin/publish")
	publish.POST("/action/", api.FeedApi.Publish)
	publish.GET("/list/", api.FeedApi.List)
}
