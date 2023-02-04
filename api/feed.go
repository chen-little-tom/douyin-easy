package api

import (
	"douyin-easy/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	StatusCode int64   `json:"status_code"` // 状态码
	StatusMsg  *string `json:"status_msg"`  // 状态描述
}

type feedApi struct{}

var FeedApi feedApi

func (fApi feedApi) Publish(c *gin.Context) {
	err := service.VideoService.Add(c)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  " uploaded successfully",
	})
}
