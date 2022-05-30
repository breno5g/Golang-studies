package main

import "fmt" // Módulo nativo do go

const prefixoOlaPortugues = "Olá, " // const padrão

func Ola(name string) string {
	// Declaração de função com tipagem de parametro e retorno
	if name == "" {
		name = "Mundo"
	}
	return prefixoOlaPortugues + name // Concatenação basica
}

func main() { // Função principal que o go chama ao executar o código
	fmt.Println(Ola("Breno")) // Print do retorno da outra função
}
