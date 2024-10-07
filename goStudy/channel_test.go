package main

import (
	"fmt"
	"testing"
)

func TestChan(t *testing.T) {
	// 在 TestGin go 中，定义一个有三个缓冲空间的 channel。
	chan1 := make(chan int)
	chan2 := make(chan int)

	for i := 0; i < 5; i++ {
		<-chan1
	}

	chan2 <- 0

	go func() {
		for {
			select {
			case v := <-chan1:
				fmt.Println("v =", v)
			case <-chan2:

			}
		}
	}()
}
