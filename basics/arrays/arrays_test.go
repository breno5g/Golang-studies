package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

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

func TestSumMultiplesArrays(t *testing.T) {
	result := SumMultiplesArrays([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	// DeepEqual comparison because compare two arrays is not possible
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result %v expected %v", result, expected)
	}
}

func TestSumMultiplesArraysRest(t *testing.T) {
	logger := func(t *testing.T, result, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("result %v expected %v", result, expected)
		}
	}

	t.Run("with two arrays full of values", func(t *testing.T) {
		result := SumMultiplesArraysRest([]int{1, 2, 3}, []int{0, 9})
		expected := []int{5, 9}

		logger(t, result, expected)
	})

	t.Run("with empty arrays", func(t *testing.T) {
		result := SumMultiplesArraysRest([]int{}, []int{})
		expected := []int{0, 0}

		logger(t, result, expected)
	})
}
