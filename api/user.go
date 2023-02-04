package api

import (
	"douyin-easy/service"
	"douyin-easy/vo"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type userApi struct{}

var UserApi userApi

type UserLoginResponse struct {
	Response
	UserId uint64 `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type UserDetailResponse struct {
	Response
	User vo.UserVo `json:"user,omitempty"`
}

// Register 用户注册
// @Summary 用户注册
// @Schemes
// @Description 注册用户
// @Tags 基础模块
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Accept json
// @Produce json
// @Success 200 {object} UserLoginResponse
// @Router /douyin/user/register/ [post]
func (uApi userApi) Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	userId, token, err := service.UserService.UserRegister(username, password)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{400, err.Error()},
		})
		return
	}
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0, StatusMsg: "注册成功"},
		UserId:   userId,
		Token:    token,
	})
}

// Login 用户登录
// @Summary 用户登录
// @Schemes
// @Description 用户登录
// @Tags 基础模块
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Accept json
// @Produce json
// @Success 200 {object} UserLoginResponse
// @Router /douyin/user/login/ [post]
func (uApi userApi) Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	userId, token, err := service.UserService.UserLogin(username, password)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{400, err.Error()},
		})
		return
	}
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0, StatusMsg: "登录成功"},
		UserId:   userId,
		Token:    token,
	})
}

// Detail 用户详情
// @Summary 用户详情
// @Schemes
// @Description 用户详情
// @Tags 基础模块
// @Param user_id query int64 true "用户id"
// @Param token query string true "token"
// @Accept json
// @Produce json
// @Success 200 {object} UserDetailResponse
// @Router /douyin/user/ [get]
func (uApi userApi) Detail(c *gin.Context) {
	userId := c.Query("user_id")
	token := c.Query("token")
	uId, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		log.Printf("用户id转化失败,err->%s\n", err)
		c.JSON(http.StatusOK, UserDetailResponse{Response: Response{
			StatusCode: 400,
			StatusMsg:  "获取信息失败",
		}})
	}
	user, err := service.UserService.Detail(uId, token)
	if err != nil {
		c.JSON(http.StatusOK, UserDetailResponse{Response: Response{
			StatusCode: 400,
			StatusMsg:  err.Error(),
		}})
		return
	}
	c.JSON(http.StatusOK, UserDetailResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "获取用户信息成功",
		},
		User: user,
	})
}
