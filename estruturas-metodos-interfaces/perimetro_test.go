package perimetro

import "testing"

func TestPerimetro(t *testing.T) {
	retangulo := Retangulo{10.0, 10.0}
	resultado := Perimetro(retangulo)
	esperado := 40.0

	if resultado != esperado {
		t.Errorf("Resultado %.2f, esperado %.2f", resultado, esperado)
	}
}

func TestArea(t *testing.T) {
	retangulo := Retangulo{10.0, 10.0}
	resultado := Area(retangulo)
	esperado := 100.0

	if resultado != esperado {
		t.Errorf("Resultado %.2f, esperado %.2f", resultado, esperado)
	}
}
