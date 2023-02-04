package router

import (
	"douyin-easy/api"
	"github.com/gin-gonic/gin"
)

func initFeed(r *gin.Engine) {
	publish := r.Group("/douyin/publish")
	publish.POST("/action/", api.FeedApi.Publish)
}
