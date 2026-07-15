package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("worker: preparing value")

		time.Sleep(2 * time.Second)

		fmt.Println("worker: sending")

		ch <- 42

		fmt.Println("worker: sent")
	}()

	fmt.Println("main: waiting")

	value := <-ch

	fmt.Println("main: received", value)
}
