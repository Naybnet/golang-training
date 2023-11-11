package main

import (
	"fmt"
)

// Close channel
func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("Let's Go")
		c <- true
		close(c)
	}()

	// Needs a closed channel otherwise deadlock
	for v := range c {
		fmt.Println(v)
	}
}
