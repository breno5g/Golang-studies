package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	t.Run("Default iteration value / 5", func(t *testing.T) {
		result := Repeat("a", 0)
		expected := "aaaaa"

		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	})

	t.Run("Iteration with parameter", func(t *testing.T) {
		result := Repeat("a", 10)
		expected := "aaaaaaaaaa"

		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	})
}

func ExampleRepeat() {
	result := Repeat("a", 0)
	fmt.Println(result)
	// Output: aaaaa
}

// go test -bench=.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}
