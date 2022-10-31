package hello

import "fmt"

const prefixInEnglish = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return prefixInEnglish + name
}

func main() {
	fmt.Println(Hello("Breno"))
}
