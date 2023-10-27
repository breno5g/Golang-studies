package main

import (
	"fmt"
	"time"
)

func worker(id int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d got %d\n", id, x)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)
	workersQtd := 10

	for i := 0; i < workersQtd; i++ {
		go worker(i, data)
	}

	for i := 0; i < 100; i++ {
		data <- i
	}
}
