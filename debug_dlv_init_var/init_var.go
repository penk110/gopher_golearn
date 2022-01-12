package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := make([]int, 3)
	for i := 0; i < len(nums); i++ {
		// append
		v := rand.Intn(10)
		nums[i] = v
		fmt.Printf("index: %d, value: %d\n", i, v)
	}

	// 切片不同初始化方式
	l1 := []int{}
	fmt.Println(&l1, len(l1))
	var l2 []int
	fmt.Println(&l2, len(l2))

}
