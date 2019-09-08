package main

import (
	"runtime"
	"fmt"
)

func main() {
	fmt.Printf("Hello World from %s!\n", runtime.GOOS)
}