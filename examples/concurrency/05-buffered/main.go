package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	fmt.Println("sending 1")
	ch <- 1

	fmt.Println("sending 2")
	ch <- 2

	fmt.Println("sending 3")
	ch <- 3

	fmt.Println("all values sent")
}
