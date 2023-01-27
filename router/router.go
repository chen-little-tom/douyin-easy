package router

import "github.com/gin-gonic/gin"

var Router *gin.Engine

func init() {

	router := gin.Default()
	// 此处配置路由
	initTest(router)

	Router = router
}
