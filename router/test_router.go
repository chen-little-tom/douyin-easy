package router

import (
	"douyin-easy/api"

	"github.com/gin-gonic/gin"
)

func initTest(r *gin.Engine) {
	test := r.Group("/test")

	test.GET("/demo", api.TestApi.Demo)
}
