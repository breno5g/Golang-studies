package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greeting(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Olá, %s", name)
}

func HandlerGreeting(w http.ResponseWriter, r *http.Request) {
	Greeting(w, "Mundo!")
}

func main() {
	err := http.ListenAndServe(":5000", http.HandlerFunc(HandlerGreeting))

	if err != nil {
		fmt.Print(err)
	}
}
