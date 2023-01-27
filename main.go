package main

import (
	"douyin-easy/config"
	docs "douyin-easy/docs"
	"douyin-easy/router"
	"fmt"
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	// 读取配置，启动程序
	app := config.GlobalConfig
	addr := fmt.Sprintf("%s:%d", app.Server.Host, app.Server.Port)

	docs.SwaggerInfo.BasePath = "/"
	swagger := ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))

	router.Router.GET("/swagger/*any", swagger)
	err := router.Router.Run(addr)
	if err != nil {
		log.Printf("Route binding failed,err: %s\n", err)
		log.Panic("Application stop")
	}
}
