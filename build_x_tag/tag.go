package main

import "fmt"

var (
	SERVICE_INFO string
)

// 编译时注入

func main() {
	fmt.Println("SERVICE_INFO: " + SERVICE_INFO)
}
