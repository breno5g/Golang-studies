package main

import (
	"reflect"
	"testing"
)

func TestSoma(t *testing.T) {

	// t.Run("coleção de 5 números", func(t *testing.T) {
	// 	// Array feião mas saudavel e com numero maximo de 5
	// 	numeros := [5]int{1, 2, 3, 4, 5}

	// 	resultado := Soma(numeros)
	// 	esperado := 15

	// 	if resultado != esperado {
	// 		t.Errorf("resultado %d, want %d, dado %v", resultado, esperado, numeros)
	// 	}
	// })

	// go test -cover // cobertura do código
	t.Run("coleção de qualquer tamanho", func(t *testing.T) {
		// se tu não passar um valor no array, ele vira um slice... que é um array grandão
		numeros := []int{1, 2, 3}

		resultado := Soma(numeros)
		esperado := 6

		if resultado != esperado {
			t.Errorf("resultado %d, esperado %d, dado %v", resultado, esperado, numeros)
		}
	})

}

func TestSomaTudo(t *testing.T) {

	resultado := SomaTudo([]int{1, 2}, []int{0, 9})
	esperado := []int{3, 9}

	// O == não consegue fazer deep equal, então dá erro
	// if resultado != esperado {
	// 	t.Errorf("resultado %v esperado %v", resultado, esperado)
	// }

	// reflect.DeepEqual não valida os tipos, então ele compila o código mesmo assim
	if !reflect.DeepEqual(resultado, esperado) {
		t.Errorf("recebido %v esperado %v", resultado, esperado)
	}
}
