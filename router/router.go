package router

import (
	"douyin-easy/mdw"
	"douyin-easy/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Router *gin.Engine

func init() {

	router := gin.Default()
	// 设置跨域
	router.Use(mdw.Cors())
	// 设置静态路由
	router.StaticFS("/public", http.Dir(utils.StaticRoot))
	// 此处配置路由

	initUser(router)
	initFeed(router)
	initComment(router)

	Router = router
}
