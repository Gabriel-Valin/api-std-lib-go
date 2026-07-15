package main

import (
	"fmt"
	"sync"
)

func worker(
	wg *sync.WaitGroup,
	ch chan int,
) {
	defer wg.Done()

	ch <- 42
}

func main() {
	var wg sync.WaitGroup

	ch := make(chan int)

	wg.Add(1)

	go worker(&wg, ch)

	value := <-ch

	fmt.Println(value)

	wg.Wait()
}
