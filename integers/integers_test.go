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

func TestDivide(t *testing.T) {
	result := Divide(10, 5)
	expected := 2

	if result != expected {
		t.Errorf("result '%d', expected '%d'", result, expected)
	}
}

func ExampleDivide() {
	result := Divide(10, 5)
	fmt.Println(result)
	// Output: 2
}

func TestRest(t *testing.T) {
	result := Rest(7, 2)
	expected := 1

	if result != expected {
		t.Errorf("result '%d', expected '%d'", result, expected)
	}
}

func ExampleRest() {
	result := Rest(7, 2)
	fmt.Println(result)
	// Output: 1
}
