package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-demo/entity"
	"net/http"
)

// Login 定义一个结构体，用来接受入参
type Login struct {
	Username string
	Password string
}

// 使用传统方法，一个一个解析入参
func sayHello(c *gin.Context) {

	// 1. get请求，从uri中获取入参，并且设定默认值
	username := c.DefaultQuery("username", "小王子")
	fmt.Println(username)

	// 1.1 从uri中获取入参，没有默认值
	username1 := c.Query("username")
	fmt.Println(username1)

	// 2. 从json中获取请求体
	b, _ := c.GetRawData()
	// 2.1 定义map或结构体，用来承接json入参
	var m map[string]interface{}
	// 2.2 反序列化
	_ = json.Unmarshal(b, &m)

	// 3. 从path中获取入参
	usernameFromPath := c.Param("username")
	fmt.Println("usernameFromPath =", usernameFromPath)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"obj": map[string]any{
			"param":            username,
			"json":             &m,
			"usernameFromPath": usernameFromPath,
		},
	})
}

// 使用bind，自动将入参解析并放到绑定对象中去。
//  1. 如果使用get请求，则只使用 Form 绑定引擎（query）。
//  2. 如果使用post请求，首先检查 content-type 是否为 JSON 或 XML，然后再使用 Form（form-data）。
func sayHello2(context *gin.Context) {
	bodyJson := Login{}
	// Bind 和 ShouldBind 的区别在于，Bind 如果出错，会在 header 中设定 code 为 400，而 ShouldBind 并不会。
	//err := context.Bind(&bodyJson)
	err := context.ShouldBind(&bodyJson)

	if err == nil {
		context.JSON(http.StatusOK, entity.Success(bodyJson))
	} else {
		fmt.Printf("context.ShouldBind fail! err = %v\n", err)
	}
}

// 解析入参为 file 的
func upload(context *gin.Context) {
	file, err := context.FormFile("file1")
	if err != nil {
		context.JSON(http.StatusInternalServerError, entity.Fail("upload file fail!"))
		return
	}

	fmt.Println(file.Filename)

}

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()

	// GET：请求方式；/hello：请求的路径
	//r.GET("/hello/:username", sayHello)
	r.POST("/hello2", sayHello2)

	r.POST("/upload", upload)

	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	_ = r.Run(":8080")
}
