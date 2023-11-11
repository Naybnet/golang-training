package main

import (
	"fmt"
)

func print(c chan bool) {
	fmt.Println("Let's Go")
	c <- true

}

// create new channel
func main() {
	c := make(chan bool)
	go print(c)
	<-c
}
