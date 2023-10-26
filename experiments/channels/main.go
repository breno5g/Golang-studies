package main

import "fmt"

func main() {
	// First test
	channel := make(chan string)

	go func() {
		channel <- "Hello from another thread"
	}()

	msg := <-channel
	fmt.Println(msg)
}
