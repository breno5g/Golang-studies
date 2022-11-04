package main

import "testing"

func compareStrings(t *testing.T, result, expected string) {
	t.Helper()

	if result != expected {
		t.Errorf("result '%s', expected '%s', dado '%s'", result, expected, "test")
	}
}

func TestSearch(t *testing.T) {
	// [string] is key type and the other string is value type
	// dictionary := map[string]string{"teste": "isso é apenas um teste"}

	// result := Search(dictionary, "teste")
	// expected := "isso é apenas um teste"

	dictionary := Dictionary{"teste": "isso é apenas um teste"}
	result := dictionary.Search("teste")
	expected := "isso é apenas um teste"

	compareStrings(t, result, expected)
}
