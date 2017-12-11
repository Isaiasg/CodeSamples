package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 20; i++ {
		fmt.Println(i)
	}

	for i := 0; ; {
		i++
		if i > 10 {
			break
		}

		fmt.Println(i)
	}

	days := []string{"Monday", "Friday"}
	for index, day := range days {
		fmt.Println(index, day)
	}

getout:
	for i := 0; i < 10; {
		fmt.Println("Using labels", i)
		for {
			i++
			break getout
		}
	}

	fmt.Println("why?")
}
