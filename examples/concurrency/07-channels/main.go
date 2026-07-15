package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "worker 1"
	}()

	go func() {
		time.Sleep(time.Second)
		ch2 <- "worker 2" // this should be executed first
	}()

	select {
	case msg := <-ch1:
		fmt.Println(msg)

	case msg := <-ch2:
		fmt.Println(msg)
	}
}
