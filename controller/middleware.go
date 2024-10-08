package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go-demo/entity"
)

// 拦截器
func M1(context *gin.Context) {
	fmt.Printf("M1 start!\n")

	now := time.Now()

	i := rand.Intn(10)
	if i%2 == 1 {
		fmt.Println("do Next, i =", i)
		context.Next()
	} else {
		fmt.Println("do Abort, i =", i)
		context.Abort()

		context.JSON(http.StatusInternalServerError, entity.FailBaseResult("M1 Abort!"))
	}

	cost := time.Since(now)

	fmt.Printf("cost =%v\n", cost)

	fmt.Printf("M1 end!\n")
}
