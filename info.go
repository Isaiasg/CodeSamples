package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Getenv("USERNAME")
	fmt.Println(name)
}
