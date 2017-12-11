package main

import "fmt"

func main() {
	name := "lolo"
	fmt.Println(name)
	changeName(&name)
	fmt.Println(name)
}

func changeName(name *string) string {
	*name = "Eloi"
	return *name
}
