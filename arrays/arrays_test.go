package arrays

import (
	"testing"
)

func TestSumArray(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	result := SumArray(numbers)
	expected := 15

	if result != expected {
		t.Errorf("result %d, expected %d, with %v", result, expected, numbers)
	}
}
