package router

import (
	"douyin-easy/api"
	"github.com/gin-gonic/gin"
)

func initUser(r *gin.Engine) {
	user := r.Group("/douyin/user")
	user.GET("/", api.UserApi.Detail)
	user.POST("/login/", api.UserApi.Login)
	user.POST("/register/", api.UserApi.Register)
}
