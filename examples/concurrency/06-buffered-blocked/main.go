package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println("trying to send 4")

	// blocked channel full
	ch <- 4

	fmt.Println("never reaches here")
}
