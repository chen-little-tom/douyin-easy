package api

import (
	"douyin-easy/model"
	"douyin-easy/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ActionResponse struct {
	StatusCode int                            `json:"status_code"`
	StatusMsg  string                         `json:"status_msg"`
	Comment    *service.CommentActionResponse `json:"comment"`
}

type ListResponse struct {
	StatusCode  int                               `json:"status_code"`
	StatusMsg   string                            `json:"status_msg"`
	CommentList *[]*service.CommentActionResponse `json:"comment_list"`
}

func CommentAction(c *gin.Context) {
	token := c.Query("token")
	videoIdParam := c.Query("video_id")
	actionType := c.Query("action_type")

	videoId, _ := strconv.ParseUint(videoIdParam, 10, 64)

	//将token转化为userId作为参数fromUserId的值
	user, _ := service.UserService.GetLoginUser(token)
	toUserId, _ := model.VideoModel.GetAuthorIdByVideoId(videoId)

	if user.Id == 0 {
		response := ActionResponse{
			StatusCode: -1,
			StatusMsg:  "Not login",
			Comment:    nil,
		}
		c.JSON(http.StatusOK, response)
		return
	}

	if actionType == "1" {
		//发布评论
		commentText := c.Query("comment_text")
		comment, err := service.PostComment(0, toUserId, videoId, user.Id, commentText)
		if err != nil {
			response := ActionResponse{
				StatusCode: -1,
				StatusMsg:  "Post comment failed",
				Comment:    nil,
			}
			c.JSON(http.StatusOK, response)
			return
		}
		response := ActionResponse{
			StatusCode: 0,
			StatusMsg:  "Success",
			Comment:    comment,
		}
		c.JSON(http.StatusOK, response)
		return
	} else if actionType == "2" {
		//删除评论
		commentIdParam := c.Query("comment_id")
		commentId, _ := strconv.ParseUint(commentIdParam, 10, 64)
		comment, err := service.DeleteComment(commentId, videoId, 4, toUserId)
		if err != nil {
			response := ActionResponse{
				StatusCode: -1,
				StatusMsg:  "Delete comment failed",
				Comment:    nil,
			}
			c.JSON(http.StatusOK, response)
			return
		}
		response := ActionResponse{
			StatusCode: 0,
			StatusMsg:  "Success",
			Comment:    comment,
		}
		c.JSON(http.StatusOK, response)
		return
	} else {
		//错误操作
		response := ActionResponse{
			StatusCode: -1,
			StatusMsg:  "Wrong action",
			Comment:    nil,
		}
		c.JSON(http.StatusOK, response)
		return
	}

}

func CommentList(c *gin.Context) {
	videoIdParam := c.Query("video_id")
	videoId, _ := strconv.ParseUint(videoIdParam, 10, 64)
	commentList, err := service.GetCommentList(videoId)
	if err != nil {
		response := ListResponse{
			StatusCode:  -1,
			StatusMsg:   "获取评论列表失败",
			CommentList: nil,
		}
		c.JSON(http.StatusOK, response)
		return
	}
	response := ListResponse{
		StatusCode:  0,
		StatusMsg:   "Success",
		CommentList: commentList,
	}
	c.JSON(http.StatusOK, response)
	return
}
