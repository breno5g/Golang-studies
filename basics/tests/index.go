package main

import "net/http"

func main() {
	res, err := http.Get("http://www.google.com")
}
