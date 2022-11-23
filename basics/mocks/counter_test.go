package main

import (
	"bytes"
	"testing"
)

func TestCounter(t *testing.T) {
	buffer := &bytes.Buffer{}

	Counter(buffer)

	result := buffer.String()
	expected := "3"
	if result != expected {
		t.Errorf("result '%s', expected '%s'", result, expected)
	}
}
