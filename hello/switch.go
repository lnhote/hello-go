package main

import (
	"fmt"
	"runtime"
)

func main() {
	switch env := runtime.GOOS; env {
	case "darwin":
		fmt.Printf("OS X, env = %s", env)
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%v", env)
	}
}
