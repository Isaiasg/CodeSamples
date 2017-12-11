package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	go func() {
		defer waitGroup.Done()
		time.Sleep(3 * time.Second)
		fmt.Println("x")
	}()

	waitGroup.Wait()
}
