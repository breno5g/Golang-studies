package main

import (
	"fmt" // Módulo nativo do go
)

const espanhol = "espanhol"
const frances = "francês"
const prefixoOlaPortugues = "Olá, " // const padrão
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "

func prefixodeSaudacao(idioma string) (prefixo string) {
	// A definição da variavel de retorno está sendo feita em cima
	switch idioma {
	case frances:
		prefixo = prefixoOlaFrances
	case espanhol:
		prefixo = prefixoOlaEspanhol
	default:
		prefixo = prefixoOlaPortugues
	}
	return // Como já definimos o retorno acima, não é preciso aqui
}

func Ola(name string, idioma string) string {
	// Declaração de função com tipagem de parametro e retorno
	if name == "" {
		name = "Mundo"
	}

	return prefixodeSaudacao(idioma) + name // Concatenação basica
}

func main() { // Função principal que o go chama ao executar o código
	fmt.Println(Ola("Breno", "espanhol")) // Print do retorno da outra função
}
