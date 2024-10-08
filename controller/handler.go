package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go-demo/dbope"
	"go-demo/entity"
	"go-demo/model"
)

// SayHello 使用传统方法，一个一个解析入参
func SayHello(c *gin.Context) {

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

	resultMap := map[string]any{
		"Param":            username,
		"Json":             m,
		"UsernameFromPath": usernameFromPath,
	}
	c.JSON(http.StatusOK, entity.SuccessBaseResult(resultMap))
}

// Bind 使用bind，自动将入参解析并放到绑定对象中去。
//  1. 如果使用get请求，则只使用 Form 绑定引擎（query）。
//  2. 如果使用post请求，首先检查 content-type 是否为 JSON 或 XML，然后再使用 Form（form-data）。
func Bind(context *gin.Context) {
	bodyJson := entity.Login{}
	// Bind 和 ShouldBind 的区别在于，Bind 如果出错，会在 header 中设定 code 为 400，而 ShouldBind 并不会。
	// err := context.Bind(&bodyJson)
	err := context.ShouldBind(&bodyJson)

	if err == nil {
		context.JSON(http.StatusOK, entity.SuccessBaseResult(bodyJson))
	} else {
		fmt.Printf("context.ShouldBind fail! err = %v\n", err)
	}
}

// UploadFile  解析入参为 file 的
func UploadFile(context *gin.Context) {
	file, err := context.FormFile("file1")
	if err != nil {
		context.JSON(http.StatusInternalServerError, entity.FailBaseResult("UploadFile file fail!"))
		return
	}

	fmt.Println(file.Filename)

	dst := fmt.Sprintf("./tmp/%s", file.Filename)

	// 上传文件到指定的目录
	_ = context.SaveUploadedFile(file, dst)

	context.JSON(http.StatusOK, entity.SuccessBaseResult(nil))
}

// UploadFiles 测试多文件上传
func UploadFiles(context *gin.Context) {
	form, err := context.MultipartForm()
	if err != nil {
		fmt.Printf("UploadFiles fail! err = %v\n", err)
	}

	headers := form.File["file"]
	for _, file := range headers {

		dst := fmt.Sprintf("./tmp/%s", file.Filename)

		// 上传文件到指定的目录
		_ = context.SaveUploadedFile(file, dst)
	}

	context.JSON(http.StatusOK, entity.SuccessBaseResult(nil))
}

func SearchUser(context *gin.Context) {
	age := context.Query("age")

	var users model.TempUser
	dbope.Db.Raw("select * from temp_user where age = ?", age).Find(&users)

	fmt.Println(users)

	context.JSON(http.StatusOK, entity.SuccessBaseResult(users))
}
