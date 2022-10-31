package hello

import "fmt"

const prefixInEnglish = "Hello, "
const prefixInPortuguese = "Olá, "
const prefixInFrench = "Bonjour, "

// This is a private function, that's why it starts with a lowercase letter.
func selectPrefix(language string) (prefix string) {
	// prefix string create a variavel with start ""
	switch language {
	case "pt-br":
		prefix = prefixInPortuguese
	case "fr":
		prefix = prefixInFrench
	default:
		prefix = prefixInEnglish
	}
	// empty return will return the return typed variable
	return
}

// This is a private function, that's why it starts with a lowercase letter.
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return selectPrefix(language) + name
}

func main() {
	fmt.Println(Hello("Breno", "pt-br"))
}
