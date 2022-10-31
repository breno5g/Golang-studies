package hello

import "testing"

func TestHello(t *testing.T) {
	result := Hello("Breno")
	expected := "Hello, Breno"

	if result != expected {
		t.Errorf("result '%s', expected  '%s'", result, expected)
	}
}
