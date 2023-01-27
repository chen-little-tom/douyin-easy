package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type testApi struct{}

var TestApi testApi

// Demo 测试接口
// @Summary 测试
// @Schemes
// @Description 测试接口
// @Tags 测试模块
// @Success 200 {objects} string
// @Router /test/demo [get]
func (api testApi) Demo(c *gin.Context) {
	c.JSON(http.StatusOK, "hello world")
}
