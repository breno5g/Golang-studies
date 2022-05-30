package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Breno")
	esperado := "Olá, Breno"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
