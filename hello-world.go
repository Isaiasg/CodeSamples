package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello from", runtime.GOOS)

	func() {
		fmt.Println("anonymous func")
	}()
}
