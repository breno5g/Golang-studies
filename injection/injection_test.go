package main

import (
	"bytes"
	"testing"
)

func TestGreeting(t *testing.T) {
	buffer := bytes.Buffer{}
	Greeting(&buffer, "Chris")

	result := buffer.String()
	expected := "Olá, Chris"

	if result != expected {
		t.Errorf("result '%s', expected '%s'", result, expected)
	}
}
