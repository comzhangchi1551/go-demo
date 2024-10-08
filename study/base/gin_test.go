package main

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGin(t *testing.T) {
	// 创建一个默认的路由引擎
	r := gin.Default()

	// GET：请求方式；/hello：请求的路径
	r.GET("/hello", sayHello)

	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	_ = r.Run()
}

func sayHello(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Hello golang",
	})
}
