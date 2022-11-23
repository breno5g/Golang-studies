package main

import (
	"testing"
)

func compareStrings(t *testing.T, result, expected string) {
	t.Helper()

	if result != expected {
		t.Errorf("result '%s', expected '%s', dado '%s'", result, expected, "test")
	}
}

func compareError(t *testing.T, result, expected error) {
	t.Helper()

	if result != expected {
		t.Errorf("result '%s', expected '%s'", result, expected)
	}
}

func TestSearch(t *testing.T) {
	t.Run("Word exist", func(t *testing.T) {
		// [string] is key type and the other string is value type
		// dictionary := map[string]string{"teste": "isso é apenas um teste"}

		// result := Search(dictionary, "teste")
		// expected := "isso é apenas um teste"

		dictionary := Dictionary{"teste": "isso é apenas um teste"}
		result, _ := dictionary.Search("teste")
		expected := "isso é apenas um teste"

		compareStrings(t, result, expected)
	})

	t.Run("Word non exist", func(t *testing.T) {
		dictionary := Dictionary{"teste": "isso é apenas um teste"}
		_, err := dictionary.Search("kjdaklsj")

		compareError(t, err, WordNotFoundErr)
	})
}

func TestAdd(t *testing.T) {
	t.Run("Add new word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("farofa", "é bão")

		expected := "é bão"

		result, err := dictionary.Search("farofa")

		compareError(t, err, nil)
		compareStrings(t, result, expected)
	})
}
