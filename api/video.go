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

type VideoListResponse struct {
	Response
	VideoList []vo.VideoVo `json:"video_list"`
}

type feedApi struct{}

var FeedApi feedApi

// Publish 上传视频
// @Summary 上传视频
// @Schemes
// @Description 上传视频
// @Tags 基础模块
// @Param title body string true "用户名"
// @Param token body string true "密码"
// @Param data formData file true "视频文件"
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Router /douyin/publish/action/ [post]
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

// Feed 获取视频流
// @Summary 获取视频流
// @Schemes
// @Description 获取视频流
// @Tags 基础模块
// @Param lastTime query int64 false "最后一个视频时间戳"
// @Accept json
// @Produce json
// @Success 200 {object} FeedResponse
// @Router /douyin/feed/ [get]
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

// List 获取投稿视频
// @Summary 获取投稿视频
// @Schemes
// @Description 获取投稿视频
// @Tags 基础模块
// @Param token query string true "用户token"
// @Accept json
// @Produce json
// @Success 200 {object} VideoListResponse
// @Router /douyin/publish/list/ [get]
func (fApi feedApi) List(c *gin.Context) {
	token := c.Query("token")

	videoVos, err := service.VideoService.List(token)
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "获取投稿视频成功",
		},
		VideoList: videoVos,
	})
}
