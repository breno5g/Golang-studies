package inteiros

import (
	"fmt"
	"testing"
)

func TestAdicionador(t *testing.T) {
	var soma int = Adiciona(2, 2)
	var esperado int = 4

	if soma != esperado {
		t.Errorf("esperado %d, resultado %d", esperado, soma)
	}
}

func ExampleAdiciona() {
	soma := Adiciona(1, 5)
	fmt.Println(soma)
	// Output: 6
}
