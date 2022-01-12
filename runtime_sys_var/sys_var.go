package main

import (
	"fmt"
	"runtime"
)

func main() {

	// amd64 darwin /Users/zhang/sdk/go1.16 go1.16
	fmt.Println(runtime.GOARCH, runtime.GOOS, runtime.GOROOT(), runtime.Version())
}
