package main

import "fmt"

func main() {

	letters := make(map[string]int)
	letters["a"] = 1
	letters["b"] = 2
	letters["c"] = 3
	letters["d"] = 4
	fmt.Println(letters)

	for key, value := range letters {
		fmt.Println(key, value)
	}

	delete(letters, "b")
	fmt.Println(letters)
}
