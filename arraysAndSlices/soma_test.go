package main

import "testing"

func TestSoma(t *testing.T) {

	t.Run("coleção de 5 números", func(t *testing.T) {
		// Array feião mas saudavel
		// numeros := [5]int{1, 2, 3, 4, 5}
		numeros := []int{1, 2, 3, 4, 5}

		resultado := Soma(numeros)
		esperado := 15

		if resultado != esperado {
			t.Errorf("resultado %d, want %d, dado %v", resultado, esperado, numeros)
		}
	})

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
