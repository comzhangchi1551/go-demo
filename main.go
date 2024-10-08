package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go-demo/controller"
	_ "go-demo/dbope"
)

func main() {
	r := gin.Default()

	r.GET("/hello", controller.M1, controller.SayHello)
	r.POST("/bind", controller.Bind)

	group1 := r.Group("/upload")
	{
		group1.POST("/", controller.UploadFile)
		group1.POST("/files", controller.UploadFiles)
	}

	r.GET("/search/user", controller.SearchUser)

	err := r.Run(":8080")

	if err != nil {
		fmt.Printf("Gin start fail! err = %v\n", err)
	}
}
