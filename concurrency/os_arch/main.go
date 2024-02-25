package main

import (
	"fmt"
	"runtime"
)

//race condition
func main() {

	fmt.Println(runtime.GOOS , " " ,runtime.GOARCH)
	
	
}