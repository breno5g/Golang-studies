package main

import (
	"fmt"
	"testing"
)

func TestPerimetro(t *testing.T) {
	retangulo := Retangulo{10.0, 10.0}
	resultado := Perimetro(retangulo)
	esperado := 40.0

	if resultado != esperado {
		t.Errorf("Resultado %.2f, esperado %.2f", resultado, esperado)
	}
}

// func TestArea(t *testing.T) {
// 	t.Run("retângulos", func(t *testing.T) {
// 		retangulo := Retangulo{12.0, 6.0}
// 		resultado := retangulo.Area()
// 		esperado := 72.0

// 		if resultado != esperado {
// 			t.Errorf("resultado %.2f, esperado %.2f", resultado, esperado)
// 		}
// 	})

// 	t.Run("círculos", func(t *testing.T) {
// 		circulo := Circulo{10}
// 		resultado := circulo.Area()
// 		esperado := 314.1592653589793

// 		if resultado != esperado {
// 			t.Errorf("resultado %.2f, esperado %.2f", resultado, esperado)
// 		}
// 	})
// }

type Forma interface {
	Area() float64
}

// func TestArea(t *testing.T) {
// 	verificaArea := func(t *testing.T, forma Forma, esperado float64) {
// 		t.Helper()
// 		resultado := forma.Area()

// 		if resultado != esperado {
// 			t.Errorf("resultado %.2f, esperado %.2f", resultado, esperado)
// 		}
// 	}

// 	t.Run("retângulos", func(t *testing.T) {
// 		retangulo := Retangulo{12.0, 6.0}
// 		verificaArea(t, retangulo, 72.0)
// 	})

// 	t.Run("círculos", func(t *testing.T) {
// 		circulo := Circulo{10}
// 		verificaArea(t, circulo, 314.1592653589793)
// 	})
// }

// TDT - table driven test

// func TestArea(t *testing.T) {
// 	testesArea := []struct { // Isso é um array de struct
// 		// Primeiro é definido o formato da struct
// 		forma    Forma
// 		esperado float64
// 	}{
// 		// Depois os valores do array
// 		{Retangulo{12, 6}, 72.0},
// 		// Está encapsulado em um objeto... mas ainda não captei o porque
// 		// Okay, tu passa entre {} porque o formato da struct é um objeto, então ele faz um auto assign
// 		// Primeiro parametro é a forma e o segundo o valor esperado
// 		{Circulo{10}, 314.1592653589793},
// 	}

// 	for _, tt := range testesArea {
// 		resultado := tt.forma.Area()

// 		if resultado != tt.esperado {
// 			t.Errorf("resultado %.2f, esperado %.2f", resultado, tt.esperado)
// 		}
// 	}
// }

func TestArea(t *testing.T) {
	testesArea := []struct {
		forma    Forma
		esperado float64
	}{
		{Retangulo{Largura: 12, Altura: 6}, 72.0},
		{Circulo{Raio: 10}, 314.1592653589793},
		{Triangulo{Base: 12, Altura: 6}, 36.0},
	}

	for _, tt := range testesArea {
		resultado := tt.forma.Area()
		if resultado != tt.esperado {
			t.Errorf("resultado %.2f, esperado %.2f", resultado, tt.esperado)
			fmt.Printf("Forma testada: %#v\n ", tt.forma)
		}
	}
}
