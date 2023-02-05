package router

import (
	"douyin-easy/api"
	"github.com/gin-gonic/gin"
)

func initComment(r *gin.Engine) {
	publish := r.Group("/douyin/comment/")
	publish.POST("/action/", api.CommentAction)
	publish.GET("/list/", api.CommentList)
}
