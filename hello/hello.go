package hello

import "fmt"

const prefixInEnglish = "Hello, "
const prefixInPortuguese = "Olá, "
const prefixInFrench = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := prefixInEnglish

	switch language {
	case "pt-br":
		prefix = prefixInPortuguese
	case "fr":
		prefix = prefixInFrench
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("Breno", "pt-br"))
}
