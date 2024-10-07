package main

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	arr := [3]int{}
	for i, value := range arr {
		fmt.Println("arr: i =", i, "value =", value)
	}

	arr2 := [3]int{1, 2}
	for i, value := range arr2 {
		fmt.Println("arr2: i =", i, "value =", value)
	}

	arr3 := arr2[1:3]
	for i, value := range arr3 {
		fmt.Println("arr3: i =", i, "value =", value)
	}

	arr4 := append(arr3, 2, 4, 3, 2)
	for i, value := range arr4 {
		fmt.Println("arr4: i =", i, "value =", value)
	}

	arr5 := append(arr4[:1], arr4[2:]...)
	for i, value := range arr5 {
		fmt.Println("arr5: i =", i, "value =", value)
	}
}
