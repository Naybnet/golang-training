package main

import (
	"fmt"
)

func print() {
	fmt.Println("Let's Go")

}

// How to print "Let's Go"?
func main() {
	go print()
	// fmt.Println("exit!!!!")
}
