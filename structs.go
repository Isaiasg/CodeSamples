package main

import "fmt"

func main() {
	type car struct {
		Maker string
		Year  int
		Color string
	}

	truck := car{
		Maker: "Ford",
		Year:  2016,
		Color: "Green",
	}

	fmt.Println(truck)
	fmt.Println("Color:", truck.Color)
}
