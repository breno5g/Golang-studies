package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	result := Repeat("a")
	expected := "aaaaa"

	if result != expected {
		t.Errorf("result '%s', expected '%s'", result, expected)
	}
}

func ExampleRepeat() {
	result := Repeat("a")
	fmt.Println(result)
	// Output: aaaaa
}
