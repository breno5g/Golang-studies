package iteracao

import "testing"

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
