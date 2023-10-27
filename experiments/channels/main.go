package main

import "fmt"

func main() {
	// First test
	// channel := make(chan string)

	// go func() {
	// channel <- "Hello from another thread"
	// }()

	// msg := <-channel
	// fmt.Println(msg)

	// Second test
	ch := make(chan string)
	go publish(ch)
	reader(ch)

	// Third test
	// ch := make(chan string)
	// go read(ch)
	// write("Hello from another thread", ch)
}

func reader(ch chan string) {
	for msg := range ch {
		fmt.Println(msg)
	}
}

func publish(ch chan string) {
	for i := 0; i < 10; i++ {
		ch <- fmt.Sprintf("Hello from another thread: %d", i)
	}

	close(ch)
}

// // receive data
// func read(ch <-chan string) {
// 	fmt.Println(<-ch)
// }

// // adding data
// func write(name string, ch chan<- string) {
// 	ch <- name
// }
