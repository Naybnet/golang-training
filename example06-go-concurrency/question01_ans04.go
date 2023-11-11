package main

import (
	"fmt"
	"sync"
)

// add time sleep func
func main() {
	wg := sync.WaitGroup{}
	// registers the number of new tasks
	wg.Add(1)
	go func() {
		defer wg.Done() // Done decrements the number of pending tasks
		fmt.Println("Let's Go")
	}()
	// Blocking statement
	wg.Wait()
}
