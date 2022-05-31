package iteracao

import (
	"fmt"
	"testing"
)

func TestRepetir(t *testing.T) {
	t.Run("Testa se repete 5 vezes / default", func(t *testing.T) {
		var resultado string = Repetir("a", 0)
		var esperado string = "aaaaa"

		if resultado != esperado {
			t.Errorf("esperado %s, resultado %s", esperado, resultado)
		}
	})
	t.Run("Testa se repete a quantidade de vezes passada no parametro", func(t *testing.T) {
		var resultado string = Repetir("a", 10)
		var esperado string = "aaaaaaaaaa"

		if resultado != esperado {
			t.Errorf("esperado %s, resultado %s", esperado, resultado)
		}
	})
}

// Comando bonito pra rodar o benchmark
// go test -bench=.

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a", 0)
	}
}

func ExampleRepetir() {
	res := Repetir("a", 10)
	fmt.Println(res)
	// Output: aaaaaaaaaa
}
