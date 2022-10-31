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

// go test -bench=.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
