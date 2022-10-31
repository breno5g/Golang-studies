package arrays

import (
	"fmt"
	"testing"
)

func TestSumArray(t *testing.T) {

}

func TestArray(t *testing.T) {
	t.Run("Collection with any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		result := SumArray(numbers)
		expected := 15

		if result != expected {
			t.Errorf("result %d, expected %d, with %v", result, expected, numbers)
		}
	})
}

func ExampleSumArray() {
	numbers := []int{1, 2, 3, 4, 5}

	result := SumArray(numbers)
	fmt.Println(result)
	// Output: 15
}
