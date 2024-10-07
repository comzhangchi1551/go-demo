package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"
)

func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func newTask(count int) {
	goID := GoID()
	for i := range count {
		fmt.Println("new Goroutine, i =", i, ", goID =", goID)
		time.Sleep(time.Second)
	}
}

func TestGoroutine(t *testing.T) {
	// 1. 通过 go 关键字，调用定义好的方法，实现异步
	go newTask(5)

	// 2. 使用 go 关键字 + 匿名方法 + 尾巴的"()" 完成异步方法的定义，以及调用
	go func(count int) {
		goID := GoID()
		for i := range count {
			fmt.Println("new Goroutine, i =", i, ", goID =", goID)
			time.Sleep(time.Second)
		}
	}(10)

	time.Sleep(3 * time.Second)
	newTask(3)
}
