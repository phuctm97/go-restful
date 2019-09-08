package main

import (
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOOS
	fmt.Printf("Hello World from %s!\n", os)
}
