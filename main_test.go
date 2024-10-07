package main

import (
	"fmt"
	"go-demo/entity"
	"testing"
)

func TestGo(t *testing.T) {
	user := entity.LocalBook{Username: "zhang", Gender: "M", Age: 12}

	fmt.Println(user)
	fmt.Printf("user =%v\n", user)
}
