package main

import (
	"fmt"
	"time"
)

// try to run: go run -race
func main() {
	msg := "Let's Go"
	done := make(chan bool)
	go func() {
		// Print: "Let's Go"
		fmt.Println(msg)
		done <- true
	}()
	<-done
	msg = "Let's GoGoGo"
	time.Sleep(1 * time.Second)

}
