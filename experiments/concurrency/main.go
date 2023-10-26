package main

import (
	"fmt"
	"net/http"
)

var number uint64 = 0

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		number++
		w.Write([]byte(fmt.Sprintf("Hello, world! You are visitor number %d", number)))
	})

	http.ListenAndServe(":3001", nil)
}
