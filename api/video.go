package api

import (
	"douyin-easy/service"
	"douyin-easy/vo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []vo.VideoVo `json:"video_list,omitempty"`
	NextTime  int64        `json:"next_time,omitempty"`
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

func (fApi feedApi) Feed(c *gin.Context) {
	lastTime := c.Query("lastTime")
	var last time.Time
	if lastTime != "" {
		t, err := strconv.ParseInt(lastTime, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 400,
				StatusMsg:  "lastTime时间格式有误",
			})
			return
		}
		last = time.Unix(t, 0)
	} else {
		last = time.Now()
	}
	token := c.Query("token")
	videoVos, nextTime, err := service.VideoService.Feed(last, token)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 400,
			StatusMsg:  "获取视频流失败",
		})
		return
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "获取视频流成功",
		},
		NextTime:  nextTime.Unix(),
		VideoList: videoVos,
	})
}
