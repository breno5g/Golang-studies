package main

import (
	"fmt" // Módulo nativo do go
)

const espanhol = "espanhol"
const frances = "francês"
const prefixoOlaPortugues = "Olá, " // const padrão
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "

func Ola(name string, idioma string) string {
	// Declaração de função com tipagem de parametro e retorno
	if name == "" {
		name = "Mundo"
	}

	prefixo := prefixoOlaPortugues

	switch idioma {
	case espanhol:
		prefixo = prefixoOlaEspanhol
	case frances:
		prefixo = prefixoOlaFrances
	}

	return prefixo + name // Concatenação basica
}

func main() { // Função principal que o go chama ao executar o código
	fmt.Println(Ola("Breno", "espanhol")) // Print do retorno da outra função
}
