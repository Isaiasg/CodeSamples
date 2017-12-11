package main

import (
	"fmt"
	"reflect"
)

var (
	name   string
	module float64
)

func main() {
	lastName := "Lolo"
	ptr := &lastName
	fmt.Println("Name is set to", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Module is set to", module, "and is of type", reflect.TypeOf(module))
	fmt.Println(lastName)
	fmt.Println(ptr, *ptr)
}
