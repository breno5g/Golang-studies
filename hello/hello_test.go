package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	verifyMessage := func(t *testing.T, result, expected string) {
		// Defines a helper function
		// because of that, the test error will appear in the origin function
		t.Helper()
		if result != expected {
			t.Errorf("result '%s', expected  '%s'", result, expected)
		}
	}

	t.Run("Say hello with parameter", func(t *testing.T) {
		result := Hello("Breno", "")
		expected := "Hello, Breno"

		// t param is required because helper function doesn't have access to to t
		verifyMessage(t, result, expected)
	})

	t.Run("Return default response with no parameter", func(t *testing.T) {
		result := Hello("", "")
		expected := "Hello, World"

		verifyMessage(t, result, expected)
	})

	t.Run("Say hello in portuguese", func(t *testing.T) {
		result := Hello("Breno", "pt-br")
		expect := "Olá, Breno"

		verifyMessage(t, result, expect)
	})

	t.Run("Say hello in french", func(t *testing.T) {
		result := Hello("Breno", "fr")
		expect := "Bonjour, Breno"

		verifyMessage(t, result, expect)
	})
}
