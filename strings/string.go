package main

import (
	"fmt"
	"strings"
)

func StringContains() {
	// Compara se possui exatamente os caracteres naquela sequencia
	fmt.Println(strings.Contains("seafood", "foo")) // true
	fmt.Println(strings.Contains("seafood", "bar")) // false
	fmt.Println(strings.Contains("seafood", ""))    // true
	fmt.Println(strings.Contains("", ""))           // true
}

func StrintContainsAny() {
	// Compara se possui qualquer um dos caracteres
	fmt.Println(strings.ContainsAny("team", "i"))     // false
	fmt.Println(strings.ContainsAny("fail", "ui"))    // true
	fmt.Println(strings.ContainsAny("ure", "ui"))     // true
	fmt.Println(strings.ContainsAny("failure", "ui")) // true
	fmt.Println(strings.ContainsAny("foo", ""))       // false
	fmt.Println(strings.ContainsAny("", ""))          // false
}

func main() {
	// StringContains()
	StrintContainsAny()
}
