package integers

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	result := Sum(10, 10)
	expected := 20

	if result != expected {
		t.Errorf("result '%d', expected '%d'", result, expected)
	}
}

func ExampleSum() {
	result := Sum(10, 10)
	fmt.Println(result)
	// Output: 20
}
