package main

import "fmt"

func main() {
	europe := make([]string, 1)
	fmt.Println("Length", len(europe), cap(europe))
	europe[0] = "Italy"
	fmt.Println(europe[0])

	europe = append(europe, "Spain")
	fmt.Println(europe[1], cap(europe))

	america := []string{"Mexico"}
	world := append(europe, america...)
	fmt.Println(world, cap(world))
}
